// Command aerolineas is an unofficial, agent-friendly CLI for Aerolíneas Argentinas.
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
	fmt.Fprintf(os.Stderr, `aerolineas — unofficial CLI for Aerolíneas Argentinas

USAGE:
  aerolineas search [--json] --from ORIGIN --to DEST --depart DATE [--return DATE]
  aerolineas read [--json] <id|url>
  aerolineas session chrome [--wait] [--timeout 3m]
  aerolineas session sync
  aerolineas session doctor [--json]
  aerolineas brands
  aerolineas network doctor [--json]
  aerolineas version | help
`)
}
