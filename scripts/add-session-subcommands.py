#!/usr/bin/env python3
"""Add session chrome/sync/doctor subcommands to all travel CLIs."""

from __future__ import annotations

import json
import re
from pathlib import Path

ROOT = Path(__file__).resolve().parent.parent

# Priority CLIs with custom probe URLs (others use base URL GET).
PROBE_URLS: dict[str, str] = {
    "melia": "https://www.melia.com/services/search/hotels/v2/search",
    "nh": "https://www.nh-hotels.com/nh/es/api/v1/hotels/search?query=Madrid&locale=es&page=1&size=1",
    "iberostar": "https://www.iberostar.com/api/graphql",
    "vueling": "https://tickets.vueling.com/ScheduleSelect.aspx?culture=es-ES",
    "easyjet": "https://www.easyjet.com/ejavailability/api/v5/availability/query?DepartureIata=MAD&ArrivalIata=PMI&MinDepartureDate=2026-07-05&MaxDepartureDate=2026-07-05&AdultSeats=1&ChildSeats=0&Infants=0&LanguageCode=ES&IncludePrices=true",
    "aireuropa": "https://dapi.aireuropa.com/api/channel-home/v1/redirect/flow/BOOKING/urldata",
    "iberiaexpress": "https://www.iberiaexpress.com/api/availability/v1/flights?market=ES&language=es&origin=MAD&destination=BCN&departureDate=2026-07-05&adults=1&operatingCarrier=I2",
}

EXTRA_URLS: dict[str, list[str]] = {
    "vueling": ["https://tickets.vueling.com", "https://www.vueling.com"],
    "aireuropa": ["https://dapi.aireuropa.com"],
}

PROBE_METHODS: dict[str, str] = {
    "melia": "POST",
    "iberostar": "POST",
    "aireuropa": "POST",
}


def load_groups() -> list[dict]:
    out = []
    for cli in sorted(ROOT.glob("*-cli")):
        slug = cli.name.replace("-cli", "")
        client_go = cli / "internal/client/client.go"
        if not client_go.exists():
            continue
        text = client_go.read_text()
        m = re.search(r'const BaseURL = "([^"]+)"', text)
        if not m:
            continue
        cat = "airline" if (cli / "cmd" / slug / "search.go").exists() else "hotel"
        out.append({"slug": slug, "url": m.group(1), "cat": cat})
    return out


def start_url(g: dict) -> str:
    url = g["url"].rstrip("/")
    if g["slug"] == "vueling":
        return "https://tickets.vueling.com/ScheduleSelect.aspx?culture=es-ES"
    if g["slug"] == "aireuropa":
        return url + "/es/es"
    if g["cat"] == "hotel":
        return url + "/es"
    return url + "/es"


def gen_session_go(slug: str, g: dict) -> str:
    mod = f"github.com/fbelchi/{slug}-cli/internal/client"
    start = start_url(g)
    extra = EXTRA_URLS.get(slug, [])
    extra_lines = ""
    if extra:
        parts = ",\n\t\t".join(f'"{u}"' for u in extra)
        extra_lines = f"\n\t\tExtraURLs:   []string{{\n\t\t\t{parts},\n\t\t}},"
    probe = PROBE_URLS.get(slug, g["url"].rstrip("/") + "/")
    probe_method = PROBE_METHODS.get(slug, "GET")
    return f'''package main

import (
\t"flag"
\t"fmt"
\t"os"
\t"time"

\t"{mod}"
\t"github.com/fbelchi/travelkit/session"
)

func cmdSession(args []string) error {{
\tif len(args) == 0 {{
\t\treturn fmt.Errorf("usage: {slug} session <chrome|sync|doctor>")
\t}}
\tswitch args[0] {{
\tcase "chrome":
\t\treturn runSessionChrome(args[1:], false)
\tcase "sync":
\t\treturn runSessionChrome(args[1:], true)
\tcase "doctor":
\t\treturn runSessionDoctor(args[1:])
\tdefault:
\t\treturn fmt.Errorf("unknown subcommand %q — use chrome, sync, or doctor", args[0])
\t}}
}}

func runSessionChrome(args []string, syncOnly bool) error {{
\tfs := flag.NewFlagSet("session chrome", flag.ExitOnError)
\tport := fs.Int("port", 9222, "Chrome remote debugging port")
\twait := fs.Bool("wait", true, "wait for WAF cookies (_abck+bm_sz, cf_clearance, or Incapsula)")
\tnoWait := fs.Bool("no-wait", false, "capture immediately without waiting")
\treplace := fs.Bool("replace", false, "quit Chrome and relaunch with dedicated profile")
\ttimeout := fs.Duration("timeout", 3*time.Minute, "max wait for WAF session cookies (headed Chrome required)")
\tcf := addCommon(fs)
\t_ = fs.Parse(reorderArgs(fs, args))

\tdoWait := *wait && !*noWait
\tif syncOnly {{
\t\tdoWait = false
\t}}

\tcl := client.New("")
\tstartURL := "{start}"
\tif !syncOnly {{
\t\tfmt.Fprintln(os.Stderr, "Capturing session from headed Chrome…")
\t\tfmt.Fprintf(os.Stderr, "  Save to: %s\\n", cl.CookiesFilePath())
\t\tfmt.Fprintf(os.Stderr, "  URL: %s\\n", startURL)
\t\tif doWait {{
\t\t\tfmt.Fprintln(os.Stderr, "  Waiting for _abck+bm_sz (or WAF equivalent) — browse the site if needed")
\t\t}}
\t}}

\tresult, err := session.CaptureChrome(session.ChromeOptions{{
\t\tEnvPrefix:   cl.EnvPrefix,
\t\tBaseURL:     client.BaseURL,
\t\tStartURL:    startURL,{extra_lines}
\t\tPort:        *port,
\t\tWait:        doWait,
\t\tWaitTimeout: *timeout,
\t\tReplace:     *replace,
\t\tSyncOnly:    syncOnly,
\t}})
\tif err != nil {{
\t\tif result.Cookie != "" {{
\t\t\tcl.ApplyCookieHeader(result.Cookie)
\t\t\t_ = cl.SavePersistedCookies()
\t\t}}
\t\treturn err
\t}}
\tcl.ApplyCookieHeader(result.Cookie)
\tif err := cl.SavePersistedCookies(); err != nil {{
\t\treturn err
\t}}
\tif cf.jsonOut {{
\t\treturn emitJSON(map[string]any{{
\t\t\t"path":      cl.CookiesFilePath(),
\t\t\t"ready":     result.Ready,
\t\t\t"has_abck":  result.HasAbck,
\t\t\t"has_bm_sz": result.HasBmSz,
\t\t}})
\t}}
\tfmt.Fprintln(os.Stderr, "Session saved →", cl.CookiesFilePath())
\tif result.Ready {{
\t\tfmt.Fprintln(os.Stderr, "WAF/session cookies OK")
\t}} else {{
\t\tfmt.Fprintln(os.Stderr, "Warning: WAF cookies incomplete — re-run: {slug} session chrome --wait --timeout 3m")
\t}}
\treturn nil
}}

func runSessionDoctor(args []string) error {{
\tfs := flag.NewFlagSet("session doctor", flag.ExitOnError)
\tcf := addCommon(fs)
\t_ = fs.Parse(reorderArgs(fs, args))

\tcl := client.New("")
\tres := session.Doctor(session.DoctorOptions{{
\t\tSlug:        "{slug}",
\t\tEnvPrefix:   cl.EnvPrefix,
\t\tBaseURL:     client.BaseURL,
\t\tCookie:      cl.Cookie,
\t\tProbeURL:    "{probe}",
\t\tProbeMethod: "{probe_method}",
\t}})
\tif cf.jsonOut {{
\t\treturn emitJSON(res)
\t}}
\tfmt.Fprintf(os.Stderr, "status: %s\\n", res.Status)
\tfmt.Fprintf(os.Stderr, "file:   %s (exists=%v)\\n", res.SessionFile, res.SessionFileExists)
\tif res.SessionAge != "" {{
\t\tfmt.Fprintf(os.Stderr, "age:    %s\\n", res.SessionAge)
\t}}
\tfmt.Fprintf(os.Stderr, "cookies: abck=%v bm_sz=%v cf=%v incap=%v\\n",
\t\tres.Cookies.HasAbck, res.Cookies.HasBmSz, res.Cookies.HasCF, res.Cookies.HasIncapsula)
\tif res.ProbeHTTPStatus > 0 {{
\t\tfmt.Fprintf(os.Stderr, "probe:  HTTP %d\\n", res.ProbeHTTPStatus)
\t}}
\tfmt.Fprintln(os.Stderr, res.Message)
\tif res.NextStep != "" {{
\t\tfmt.Fprintln(os.Stderr, "next:", res.NextStep)
\t}}
\tif res.Status != session.DoctorOK {{
\t\treturn fmt.Errorf("%s", res.Message)
\t}}
\treturn nil
}}
'''


def patch_main(slug: str, path: Path) -> bool:
    text = path.read_text()
    if 'case "session":' in text:
        changed = False
        if "session doctor" not in text and "session chrome" in text:
            text = text.replace(
                "  " + slug + " session sync\n",
                "  " + slug + " session sync\n  " + slug + " session doctor [--json]\n",
            )
            changed = True
        if "session doctor" not in text and "session sync" not in text:
            # add usage lines before version
            text = re.sub(
                r"(  " + re.escape(slug) + r" brands\n)",
                r"\1  " + slug + " session chrome [--wait] [--timeout 3m]\n  " + slug + " session sync\n  " + slug + " session doctor [--json]\n",
                text,
                count=1,
            )
            changed = True
        if changed:
            path.write_text(text)
        return changed
    # insert session case before brands
    text = text.replace(
        '\tcase "brands":',
        '\tcase "session":\n\t\terr = cmdSession(os.Args[2:])\n\tcase "brands":',
        1,
    )
    # usage block
    text = re.sub(
        r"(  " + re.escape(slug) + r" brands\n  " + re.escape(slug) + r" version)",
        "  " + slug + " session chrome [--wait] [--timeout 3m]\n  " + slug + " session sync\n  " + slug + " session doctor [--json]\n\\1",
        text,
        count=1,
    )
    path.write_text(text)
    return True


def main() -> None:
    groups = load_groups()
    by_slug = {g["slug"]: g for g in groups}
    added = 0
    for cli in sorted(ROOT.glob("*-cli")):
        slug = cli.name.replace("-cli", "")
        g = by_slug.get(slug)
        if not g:
            continue
        session_path = cli / "cmd" / slug / "session.go"
        session_path.write_text(gen_session_go(slug, g))
        main_path = cli / "cmd" / slug / "main.go"
        if main_path.exists():
            patch_main(slug, main_path)
        added += 1
    print(f"session subcommands written for {added} CLIs")


if __name__ == "__main__":
    main()
