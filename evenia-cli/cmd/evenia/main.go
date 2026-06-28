// Command evenia is an unofficial, agent-friendly CLI for Evenia.
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
	fmt.Fprintf(os.Stderr, `evenia — unofficial CLI for Evenia

USAGE:
  evenia search [--json] [--limit N] <destination...>
  evenia read [--json] <id|url>
  evenia availability [--json] --check-in DATE --check-out DATE [--guests N] [--rooms N] <hotel-id>
  evenia session chrome [--wait] [--timeout 3m]
  evenia session sync
  evenia session doctor [--json]
  evenia brands
  evenia version | help
`)
}
