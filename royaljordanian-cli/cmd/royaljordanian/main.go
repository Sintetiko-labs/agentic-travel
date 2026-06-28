// Command royaljordanian is an unofficial, agent-friendly CLI for Royal Jordanian.
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
	fmt.Fprintf(os.Stderr, `royaljordanian — unofficial CLI for Royal Jordanian

USAGE:
  royaljordanian search [--json] --from ORIGIN --to DEST --depart DATE [--return DATE]
  royaljordanian read [--json] <id|url>
  royaljordanian session chrome [--wait] [--timeout 3m]
  royaljordanian session sync
  royaljordanian session doctor [--json]
  royaljordanian brands
  royaljordanian version | help
`)
}
