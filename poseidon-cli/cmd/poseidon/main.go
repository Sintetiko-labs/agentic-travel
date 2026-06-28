// Command poseidon is an unofficial, agent-friendly CLI for Hoteles Poseidón.
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
	fmt.Fprintf(os.Stderr, `poseidon — unofficial CLI for Hoteles Poseidón

USAGE:
  poseidon search [--json] [--limit N] <destination...>
  poseidon read [--json] <id|url>
  poseidon availability [--json] --check-in DATE --check-out DATE [--guests N] [--rooms N] <hotel-id>
  poseidon session chrome [--wait] [--timeout 3m]
  poseidon session sync
  poseidon session doctor [--json]
  poseidon brands
  poseidon version | help
`)
}
