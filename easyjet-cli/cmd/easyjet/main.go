// Command easyjet is an unofficial, agent-friendly CLI for easyJet.
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
	fmt.Fprintf(os.Stderr, `easyjet — unofficial CLI for easyJet

USAGE:
  easyjet search [--json] --from ORIGIN --to DEST --depart DATE [--return DATE]
  easyjet read [--json] <id|url>
  easyjet brands
  easyjet session chrome [--wait] [--port N]
  easyjet session sync
  easyjet session doctor [--json]
  easyjet version | help
`)
}
