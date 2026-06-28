// Command axel is an unofficial, agent-friendly CLI for Axel Hotels.
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
	case "availability":
		err = cmdAvailability(os.Args[2:])
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
	fmt.Fprintf(os.Stderr, `axel — unofficial CLI for Axel Hotels

USAGE:
  axel search [--json] [--limit N] <destination...>
  axel read [--json] <id|url>
  axel availability [--json] --check-in DATE --check-out DATE [--guests N] [--rooms N] <hotel-id>
  axel session chrome [--wait] [--timeout 3m]
  axel session sync
  axel session doctor [--json]
  axel brands
  axel network doctor [--json]
  axel version | help
`)
}
