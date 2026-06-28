// Command designhotels is an unofficial, agent-friendly CLI for Design Hotels.
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
	fmt.Fprintf(os.Stderr, `designhotels — unofficial CLI for Design Hotels

USAGE:
  designhotels search [--json] [--limit N] <destination...>
  designhotels read [--json] <id|url>
  designhotels availability [--json] --check-in DATE --check-out DATE [--guests N] [--rooms N] <hotel-id>
  designhotels session chrome [--wait] [--timeout 3m]
  designhotels session sync
  designhotels session doctor [--json]
  designhotels brands
  designhotels network doctor [--json]
  designhotels version | help
`)
}
