// Command delta is an unofficial, agent-friendly CLI for Delta Air Lines.
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
	fmt.Fprintf(os.Stderr, `delta — unofficial CLI for Delta Air Lines

USAGE:
  delta search [--json] --from ORIGIN --to DEST --depart DATE [--return DATE]
  delta read [--json] <id|url>
  delta session chrome [--wait] [--timeout 3m]
  delta session sync
  delta session doctor [--json]
  delta brands
  delta version | help
`)
}
