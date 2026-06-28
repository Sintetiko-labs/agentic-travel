// Command tui is an unofficial, agent-friendly CLI for TUI.
package main

import (
	"fmt"
	"os"
)

var version = "dev"

func main() {
	if len(os.Args) < 2 {
		usage()
		os.Exit(2)
	}
	var err error
	switch os.Args[1] {
	case "search":
		err = cmdSearch(os.Args[2:])
	case "read":
		err = cmdRead(os.Args[2:])
	case "session":
		err = cmdSession(os.Args[2:])
	case "brands":
		cmdBrands()
	case "version", "--version", "-v":
		fmt.Println(version)
	case "help", "-h", "--help":
		usage()
	default:
		fmt.Fprintf(os.Stderr, "unknown command %q\n\n", os.Args[1])
		usage()
		os.Exit(2)
	}
	if err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}
}

func usage() {
	fmt.Fprintf(os.Stderr, `tui — unofficial CLI for TUI

USAGE:
  tui search [--json] [--brand BRAND] --from ORIGIN --to DEST --depart DATE [--return DATE]
  tui read [--json] [--brand BRAND] <id|url>
  tui session chrome [--wait] [--timeout 3m]
  tui session sync
  tui session doctor [--json]
  tui brands
  tui version | help
`)
}
