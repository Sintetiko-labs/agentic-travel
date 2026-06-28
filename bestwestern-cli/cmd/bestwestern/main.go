// Command bestwestern is an unofficial, agent-friendly CLI for Best Western.
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
	fmt.Fprintf(os.Stderr, `bestwestern — unofficial CLI for Best Western

USAGE:
  bestwestern search [--json] [--limit N] [--brand BRAND] <destination...>
  bestwestern read [--json] [--brand BRAND] <id|url>
  bestwestern availability [--json] [--brand BRAND] --check-in DATE --check-out DATE [--guests N] [--rooms N] <hotel-id>
  bestwestern brands
  bestwestern version | help
`)
}
