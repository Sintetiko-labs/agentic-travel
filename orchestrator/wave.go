// Command wave runs the Madrid→London parallel search wave via mac-build-cli.sh.
package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
	"sync"
	"time"
)

type sourceSpec struct {
	Name string
	Path string
	MS   int
	OK   bool
}

type waveMeta struct {
	Route         string `json:"route"`
	RyanairRoute  string `json:"ryanair_route"`
	VuelingTarget string `json:"vueling_target"`
	Depart        string `json:"depart"`
	CheckIn       string `json:"check_in"`
	CheckOut      string `json:"check_out"`
}

func msNow() int64 { return time.Now().UnixMilli() }

func runTimed(ctx context.Context, tmp, name string, args ...string) sourceSpec {
	out := filepath.Join(tmp, name+".json")
	errPath := filepath.Join(tmp, name+".err")
	t0 := msNow()
	cmd := exec.CommandContext(ctx, args[0], args[1:]...)
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	runErr := cmd.Run()
	_ = os.WriteFile(out, stdout.Bytes(), 0o644)
	_ = os.WriteFile(errPath, stderr.Bytes(), 0o644)
	ok := runErr == nil
	ms := int(msNow() - t0)
	return sourceSpec{Name: name, Path: out, MS: ms, OK: ok}
}

func runBuild(ctx context.Context, build, slug string, slugArgs ...string) *exec.Cmd {
	all := append([]string{build, slug}, slugArgs...)
	return exec.CommandContext(ctx, all[0], all[1:]...)
}

func flightTotal(raw []byte) int {
	var m map[string]any
	if json.Unmarshal(raw, &m) != nil {
		return 0
	}
	if t, ok := m["total"].(float64); ok {
		return int(t)
	}
	return 0
}

func runVueling(ctx context.Context, build, tmp, from, depart string) []sourceSpec {
	directOut := filepath.Join(tmp, "vueling_direct.json")
	t0 := msNow()
	cmd := runBuild(ctx, build, "vueling", "search", "--json", "--from", from, "--to", "LGW", "--depart", depart)
	out, err := cmd.Output()
	if err == nil && flightTotal(out) > 0 {
		_ = os.WriteFile(directOut, out, 0o644)
		return []sourceSpec{{Name: "vueling", Path: directOut, MS: int(msNow() - t0), OK: true}}
	}
	leg1 := filepath.Join(tmp, "vueling_leg1.json")
	leg2 := filepath.Join(tmp, "vueling_leg2.json")
	tLeg := msNow()
	var b1, b2 []byte
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		c := runBuild(ctx, build, "vueling", "search", "--json", "--from", from, "--to", "BCN", "--depart", depart)
		b1, _ = c.Output()
	}()
	go func() {
		defer wg.Done()
		c := runBuild(ctx, build, "vueling", "search", "--json", "--from", "BCN", "--to", "LGW", "--depart", depart)
		b2, _ = c.Output()
	}()
	wg.Wait()
	_ = os.WriteFile(leg1, b1, 0o644)
	_ = os.WriteFile(leg2, b2, 0o644)
	ms1 := int(msNow() - tLeg)
	return []sourceSpec{
		{Name: "vueling_leg1", Path: leg1, MS: ms1, OK: true},
		{Name: "vueling_leg2", Path: leg2, MS: 0, OK: true},
	}
}

func env(k, def string) string {
	if v := strings.TrimSpace(os.Getenv(k)); v != "" {
		return v
	}
	return def
}

func main() {
	root := env("AGENTIC_TRAVEL_ROOT", "")
	if root == "" {
		wd, _ := os.Getwd()
		root = filepath.Clean(filepath.Join(wd, ".."))
	}
	build := filepath.Join(root, "scripts", "mac-build-cli.sh")
	merge := filepath.Join(root, "scripts", "wave-merge.py")
	from := env("WAVE_FROM", "MAD")
	depart := env("WAVE_DEPART", "2026-07-15")
	checkIn := env("WAVE_CHECK_IN", "2026-07-15")
	checkOut := env("WAVE_CHECK_OUT", "2026-07-18")
	outPath := env("WAVE_OUT", filepath.Join(root, "wave-result-go.json"))

	tmp, err := os.MkdirTemp("", "wave-go-*")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	defer os.RemoveAll(tmp)

	ctx := context.Background()
	wallT0 := msNow()
	var mu sync.Mutex
	var specs []sourceSpec
	var wg sync.WaitGroup

	add := func(s sourceSpec) {
		mu.Lock()
		specs = append(specs, s)
		mu.Unlock()
	}

	if tok := strings.TrimSpace(os.Getenv("DUFFEL_ACCESS_TOKEN")); tok != "" {
		wg.Add(1)
		go func() {
			defer wg.Done()
			add(runTimed(ctx, tmp, "duffel", "node", filepath.Join(root, "mcp", "call-search-flights.mjs"), "--from", from, "--to", "STN", "--depart", depart))
		}()
	}

	wg.Add(1)
	go func() {
		defer wg.Done()
		add(runTimed(ctx, tmp, "ryanair", build, "ryanair", "search", "--json", "--from", from, "--to", "STN", "--depart", depart))
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		for _, s := range runVueling(ctx, build, tmp, from, depart) {
			add(s)
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		add(runTimed(ctx, tmp, "travelodge", build, "travelodge", "search", "--json", "London"))
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		add(runTimed(ctx, tmp, "hilton", build, "hilton", "search", "--json", "London"))
	}()

	wg.Wait()
	wallMS := int(msNow() - wallT0)

	meta, _ := json.Marshal(waveMeta{
		Route:         from + "->London",
		RyanairRoute:  from + "->STN",
		VuelingTarget: from + "->LGW",
		Depart:        depart,
		CheckIn:       checkIn,
		CheckOut:      checkOut,
	})

	mergeArgs := []string{merge, "--out", outPath, "--wall-ms", strconv.Itoa(wallMS), "--meta", string(meta)}
	for _, s := range specs {
		ok := "1"
		if !s.OK {
			ok = "0"
		}
		mergeArgs = append(mergeArgs, fmt.Sprintf("%s:%s:%d:%s", s.Name, s.Path, s.MS, ok))
	}

	cmd := exec.Command("python3", mergeArgs...)
	cmd.Dir = root
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		os.Exit(1)
	}
}
