package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/fbelchi/travelkit/network"
)

func cmdNetwork(args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("usage: %s network doctor [--json]", os.Args[0])
	}
	switch args[0] {
	case "doctor":
		return runNetworkDoctor(args[1:])
	default:
		return fmt.Errorf("unknown subcommand %q — use doctor", args[0])
	}
}

func runNetworkDoctor(args []string) error {
	fs := flag.NewFlagSet("network doctor", flag.ExitOnError)
	jsonOut := fs.Bool("json", false, "emit JSON to stdout")
	_ = fs.Parse(reorderArgs(fs, args))

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	res := network.Doctor(ctx)
	if *jsonOut {
		enc := json.NewEncoder(os.Stdout)
		enc.SetIndent("", "  ")
		enc.SetEscapeHTML(false)
		return enc.Encode(res)
	}
	network.PrintDoctor(res)
	if res.Status != network.DoctorOK && res.Status != network.DoctorProxySet {
		return fmt.Errorf("%s", res.Message)
	}
	return nil
}
