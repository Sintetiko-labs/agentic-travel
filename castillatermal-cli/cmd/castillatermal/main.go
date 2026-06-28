// Command castillatermal is an unofficial, agent-friendly CLI for Castilla Termal.
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
	fmt.Fprintf(os.Stderr, `castillatermal — unofficial CLI for Castilla Termal

USAGE:
  castillatermal search [--json] [--limit N] <destination...>
  castillatermal read [--json] <id|url>
  castillatermal availability [--json] --check-in DATE --check-out DATE [--guests N] [--rooms N] <hotel-id>
  castillatermal brands
  castillatermal version | help
`)
}
