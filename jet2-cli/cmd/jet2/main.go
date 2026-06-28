// Command jet2 is an unofficial, agent-friendly CLI for Jet2.
package main

import (
	"fmt"
	"os"

	"github.com/fbelchi/travelkit/network"
)

var version = "dev"

func main() {
	os.Args = network.PreprocessArgs(os.Args)
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
	case "network":
		err = cmdNetwork(os.Args[2:])
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
	fmt.Fprintf(os.Stderr, `jet2 — unofficial CLI for Jet2

USAGE:
  jet2 search [--json] --from ORIGIN --to DEST --depart DATE [--return DATE]
  jet2 read [--json] <id|url>
  jet2 session chrome [--wait] [--timeout 3m]
  jet2 session sync
  jet2 session doctor [--json]
  jet2 brands
  jet2 network doctor [--json]
  jet2 version | help
`)
}
