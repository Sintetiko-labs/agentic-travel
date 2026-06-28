package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/fbelchi/iberostar-cli/internal/client"
	"github.com/fbelchi/travelkit/session"
)

func cmdSession(args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("usage: iberostar session <chrome|sync>")
	}
	switch args[0] {
	case "chrome":
		return runSessionChrome(args[1:], false)
	case "sync":
		return runSessionChrome(args[1:], true)
	default:
		return fmt.Errorf("unknown subcommand %q — use chrome or sync", args[0])
	}
}

func runSessionChrome(args []string, syncOnly bool) error {
	fs := flag.NewFlagSet("session chrome", flag.ExitOnError)
	port := fs.Int("port", 9222, "Chrome remote debugging port")
	wait := fs.Bool("wait", true, "wait for WAF/session cookies (_abck, bm_sz, cf_clearance)")
	noWait := fs.Bool("no-wait", false, "capture immediately without waiting")
	replace := fs.Bool("replace", false, "quit Chrome and relaunch with dedicated profile")
	timeout := fs.Duration("timeout", 180*time.Second, "max wait for session cookies")
	cf := addCommon(fs)
	_ = fs.Parse(reorderArgs(fs, args))

	doWait := *wait && !*noWait
	if syncOnly {
		doWait = false
	}

	cl := client.New("")
	startURL := "https://www.iberostar.com/es"
	if !syncOnly {
		fmt.Fprintln(os.Stderr, "Capturing session from Chrome…")
		fmt.Fprintf(os.Stderr, "  Save to: %s\n", cl.CookiesFilePath())
		fmt.Fprintf(os.Stderr, "  URL: %s\n", startURL)
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
		fmt.Fprintln(os.Stderr, "Warning: WAF cookies not detected — browse the site in Chrome, then re-run session sync")
	}
	return nil
}
