// Command castillatermal is an unofficial, agent-friendly CLI for Castilla Termal.
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
	fmt.Fprintf(os.Stderr, `castillatermal — unofficial CLI for Castilla Termal

USAGE:
  castillatermal search [--json] [--limit N] <destination...>
  castillatermal read [--json] <id|url>
  castillatermal availability [--json] --check-in DATE --check-out DATE [--guests N] [--rooms N] <hotel-id>
  castillatermal session chrome [--wait] [--timeout 3m]
  castillatermal session sync
  castillatermal session doctor [--json]
  castillatermal brands
  castillatermal network doctor [--json]
  castillatermal version | help
`)
}
