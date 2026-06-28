// Command nobu is an unofficial, agent-friendly CLI for Nobu Hotels.
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
	case "availability":
		err = cmdAvailability(os.Args[2:])
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
	fmt.Fprintf(os.Stderr, `nobu — unofficial CLI for Nobu Hotels

USAGE:
  nobu search [--json] [--limit N] <destination...>
  nobu read [--json] <id|url>
  nobu availability [--json] --check-in DATE --check-out DATE [--guests N] [--rooms N] <hotel-id>
  nobu session chrome [--wait] [--timeout 3m]
  nobu session sync
  nobu session doctor [--json]
  nobu brands
  nobu version | help
`)
}
