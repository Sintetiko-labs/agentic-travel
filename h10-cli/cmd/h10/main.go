// Command h10 is an unofficial, agent-friendly CLI for H10.
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
	fmt.Fprintf(os.Stderr, `h10 — unofficial CLI for H10

USAGE:
  h10 search [--json] [--limit N] [--brand BRAND] <destination...>
  h10 read [--json] [--brand BRAND] <id|url>
  h10 availability [--json] [--brand BRAND] --check-in DATE --check-out DATE [--guests N] [--rooms N] <hotel-id>
  h10 brands
  h10 version | help
`)
}
