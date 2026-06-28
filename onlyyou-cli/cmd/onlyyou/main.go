// Command onlyyou is an unofficial, agent-friendly CLI for Only YOU.
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
	fmt.Fprintf(os.Stderr, `onlyyou — unofficial CLI for Only YOU

USAGE:
  onlyyou search [--json] [--limit N] <destination...>
  onlyyou read [--json] <id|url>
  onlyyou availability [--json] --check-in DATE --check-out DATE [--guests N] [--rooms N] <hotel-id>
  onlyyou brands
  onlyyou version | help
`)
}
