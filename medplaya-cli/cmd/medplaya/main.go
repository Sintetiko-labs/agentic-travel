// Command medplaya is an unofficial, agent-friendly CLI for MedPlaya.
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
	fmt.Fprintf(os.Stderr, `medplaya — unofficial CLI for MedPlaya

USAGE:
  medplaya search [--json] [--limit N] <destination...>
  medplaya read [--json] <id|url>
  medplaya availability [--json] --check-in DATE --check-out DATE [--guests N] [--rooms N] <hotel-id>
  medplaya session chrome [--wait] [--timeout 3m]
  medplaya session sync
  medplaya session doctor [--json]
  medplaya brands
  medplaya network doctor [--json]
  medplaya version | help
`)
}
