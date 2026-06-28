// Command plusultra is an unofficial, agent-friendly CLI for Plus Ultra.
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
	fmt.Fprintf(os.Stderr, `plusultra — unofficial CLI for Plus Ultra

USAGE:
  plusultra search [--json] --from ORIGIN --to DEST --depart DATE [--return DATE]
  plusultra read [--json] <id|url>
  plusultra session chrome [--wait] [--timeout 3m]
  plusultra session sync
  plusultra session doctor [--json]
  plusultra brands
  plusultra version | help
`)
}
