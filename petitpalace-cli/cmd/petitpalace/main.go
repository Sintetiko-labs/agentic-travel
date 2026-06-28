// Command petitpalace is an unofficial, agent-friendly CLI for Petit Palace.
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
	fmt.Fprintf(os.Stderr, `petitpalace — unofficial CLI for Petit Palace

USAGE:
  petitpalace search [--json] [--limit N] <destination...>
  petitpalace read [--json] <id|url>
  petitpalace availability [--json] --check-in DATE --check-out DATE [--guests N] [--rooms N] <hotel-id>
  petitpalace brands
  petitpalace version | help
`)
}
