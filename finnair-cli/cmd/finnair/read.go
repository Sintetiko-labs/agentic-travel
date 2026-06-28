package main

import (
	"flag"
	"fmt"

	"github.com/fbelchi/finnair-cli/internal/client"
)

func cmdRead(args []string) error {
	fs := flag.NewFlagSet("read", flag.ExitOnError)
	cf := addCommon(fs)
	_ = fs.Parse(reorderArgs(fs, args))
	if fs.NArg() != 1 {
		return fmt.Errorf("usage: finnair read [flags] <id|url>")
	}
	cl := client.New("")
	fv, err := cl.Read(fs.Arg(0))
	if err != nil {
		return err
	}
	if cf.jsonOut {
		return emitJSON(fv)
	}
	fmt.Printf("[%s] %s %s→%s %s\n", fv.ID, fv.FlightNumber, fv.Origin, fv.Destination, fv.Depart)
	if fv.Price.Price != "" {
		fmt.Printf("  price: %s %s\n", fv.Price.Price, fv.Price.Currency)
	}
	return nil
}
