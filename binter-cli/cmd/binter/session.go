package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/fbelchi/binter-cli/internal/client"
	"github.com/fbelchi/travelkit/session"
)

func cmdSession(args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("usage: binter session <chrome|sync|doctor>")
	}
	switch args[0] {
	case "chrome":
		return runSessionChrome(args[1:], false)
	case "sync":
		return runSessionChrome(args[1:], true)
	case "doctor":
		return runSessionDoctor(args[1:])
	default:
		return fmt.Errorf("unknown subcommand %q — use chrome, sync, or doctor", args[0])
	}
}

func runSessionChrome(args []string, syncOnly bool) error {
	fs := flag.NewFlagSet("session chrome", flag.ExitOnError)
	port := fs.Int("port", 9222, "Chrome remote debugging port")
	wait := fs.Bool("wait", true, "wait for WAF cookies (_abck+bm_sz, cf_clearance, or Incapsula)")
	noWait := fs.Bool("no-wait", false, "capture immediately without waiting")
	replace := fs.Bool("replace", false, "quit Chrome and relaunch with dedicated profile")
	timeout := fs.Duration("timeout", 3*time.Minute, "max wait for WAF session cookies (headed Chrome required)")
	cf := addCommon(fs)
	_ = fs.Parse(reorderArgs(fs, args))

	doWait := *wait && !*noWait
	if syncOnly {
		doWait = false
	}

	cl := client.New("")
	startURL := "https://www.bintercanarias.com/es"
	if !syncOnly {
		fmt.Fprintln(os.Stderr, "Capturing session from headed Chrome…")
		fmt.Fprintf(os.Stderr, "  Save to: %s\n", cl.CookiesFilePath())
		fmt.Fprintf(os.Stderr, "  URL: %s\n", startURL)
		if doWait {
			fmt.Fprintln(os.Stderr, "  Waiting for _abck+bm_sz (or WAF equivalent) — browse the site if needed")
		}
	}

	result, err := session.CaptureChrome(session.ChromeOptions{
		EnvPrefix:   cl.EnvPrefix,
		BaseURL:     client.BaseURL,
		StartURL:    startURL,
		Port:        *port,
		Wait:        doWait,
		WaitTimeout: *timeout,
		Replace:     *replace,
		SyncOnly:    syncOnly,
	})
	if err != nil {
		if result.Cookie != "" {
			cl.ApplyCookieHeader(result.Cookie)
			_ = cl.SavePersistedCookies()
		}
		return err
	}
	cl.ApplyCookieHeader(result.Cookie)
	if err := cl.SavePersistedCookies(); err != nil {
		return err
	}
	if cf.jsonOut {
		return emitJSON(map[string]any{
			"path":      cl.CookiesFilePath(),
			"ready":     result.Ready,
			"has_abck":  result.HasAbck,
			"has_bm_sz": result.HasBmSz,
		})
	}
	fmt.Fprintln(os.Stderr, "Session saved →", cl.CookiesFilePath())
	if result.Ready {
		fmt.Fprintln(os.Stderr, "WAF/session cookies OK")
	} else {
		fmt.Fprintln(os.Stderr, "Warning: WAF cookies incomplete — re-run: binter session chrome --wait --timeout 3m")
	}
	return nil
}

func runSessionDoctor(args []string) error {
	fs := flag.NewFlagSet("session doctor", flag.ExitOnError)
	cf := addCommon(fs)
	_ = fs.Parse(reorderArgs(fs, args))

	cl := client.New("")
	res := session.Doctor(session.DoctorOptions{
		Slug:        "binter",
		EnvPrefix:   cl.EnvPrefix,
		BaseURL:     client.BaseURL,
		Cookie:      cl.Cookie,
		ProbeURL:    "https://www.bintercanarias.com/",
		ProbeMethod: "GET",
	})
	if cf.jsonOut {
		return emitJSON(res)
	}
	fmt.Fprintf(os.Stderr, "status: %s\n", res.Status)
	fmt.Fprintf(os.Stderr, "file:   %s (exists=%v)\n", res.SessionFile, res.SessionFileExists)
	if res.SessionAge != "" {
		fmt.Fprintf(os.Stderr, "age:    %s\n", res.SessionAge)
	}
	fmt.Fprintf(os.Stderr, "cookies: abck=%v bm_sz=%v cf=%v incap=%v\n",
		res.Cookies.HasAbck, res.Cookies.HasBmSz, res.Cookies.HasCF, res.Cookies.HasIncapsula)
	if res.ProbeHTTPStatus > 0 {
		fmt.Fprintf(os.Stderr, "probe:  HTTP %d\n", res.ProbeHTTPStatus)
	}
	fmt.Fprintln(os.Stderr, res.Message)
	if res.NextStep != "" {
		fmt.Fprintln(os.Stderr, "next:", res.NextStep)
	}
	if res.Status != session.DoctorOK {
		return fmt.Errorf("%s", res.Message)
	}
	return nil
}
