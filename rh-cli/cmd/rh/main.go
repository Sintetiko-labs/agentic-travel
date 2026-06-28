// Command rh is an unofficial, agent-friendly CLI for Hoteles RH.
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
	fmt.Fprintf(os.Stderr, `rh — unofficial CLI for Hoteles RH

USAGE:
  rh search [--json] [--limit N] <destination...>
  rh read [--json] <id|url>
  rh availability [--json] --check-in DATE --check-out DATE [--guests N] [--rooms N] <hotel-id>
  rh session chrome [--wait] [--timeout 3m]
  rh session sync
  rh session doctor [--json]
  rh brands
  rh version | help
`)
}
