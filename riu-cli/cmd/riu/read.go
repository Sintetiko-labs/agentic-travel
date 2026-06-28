package main

import (
	"flag"
	"fmt"

	"github.com/fbelchi/riu-cli/internal/client"
)

func cmdRead(args []string) error {
	fs := flag.NewFlagSet("read", flag.ExitOnError)
	cf := addCommon(fs)
	_ = fs.Parse(reorderArgs(fs, args))
	if fs.NArg() != 1 {
		return fmt.Errorf("usage: riu read [flags] <id|url>")
	}
	cl := client.New("")
	pv, err := cl.Read(fs.Arg(0))
	if err != nil {
		return err
	}
	if cf.jsonOut {
		return emitJSON(pv)
	}
	fmt.Printf("[%s] %s\n", pv.ID, pv.Name)
	if pv.Price.Price != "" {
		fmt.Printf("  price: %s %s\n", pv.Price.Price, pv.Price.Currency)
	}
	return nil
}
