// Command limehome is an unofficial, agent-friendly CLI for Limehome.
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
	fmt.Fprintf(os.Stderr, `limehome — unofficial CLI for Limehome

USAGE:
  limehome search [--json] [--limit N] <destination...>
  limehome read [--json] <id|url>
  limehome availability [--json] --check-in DATE --check-out DATE [--guests N] [--rooms N] <hotel-id>
  limehome session chrome [--wait] [--timeout 3m]
  limehome session sync
  limehome session doctor [--json]
  limehome brands
  limehome version | help
`)
}
