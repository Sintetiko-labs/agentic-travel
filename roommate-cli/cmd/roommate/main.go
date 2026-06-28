// Command roommate is an unofficial, agent-friendly CLI for Room Mate.
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
	fmt.Fprintf(os.Stderr, `roommate — unofficial CLI for Room Mate

USAGE:
  roommate search [--json] [--limit N] <destination...>
  roommate read [--json] <id|url>
  roommate availability [--json] --check-in DATE --check-out DATE [--guests N] [--rooms N] <hotel-id>
  roommate session chrome [--wait] [--timeout 3m]
  roommate session sync
  roommate session doctor [--json]
  roommate brands
  roommate version | help
`)
}
