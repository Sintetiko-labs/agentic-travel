package main

import (
	"flag"
	"fmt"

	"github.com/fbelchi/riu-cli/internal/client"
)

func cmdAvailability(args []string) error {
	fs := flag.NewFlagSet("availability", flag.ExitOnError)
	cf := addCommon(fs)
	checkIn := fs.String("check-in", "", "check-in date (YYYY-MM-DD)")
	checkOut := fs.String("check-out", "", "check-out date (YYYY-MM-DD)")
	guests := fs.Int("guests", 2, "number of guests")
	rooms := fs.Int("rooms", 1, "number of rooms")
	_ = fs.Parse(reorderArgs(fs, args))
	if fs.NArg() != 1 {
		return fmt.Errorf("usage: riu availability [flags] <hotel-id>")
	}
	if *checkIn == "" || *checkOut == "" {
		return fmt.Errorf("--check-in and --check-out are required")
	}
	cl := client.New("")
	av, err := cl.Availability(fs.Arg(0), *checkIn, *checkOut, *guests, *rooms)
	if err != nil {
		return err
	}
	if cf.jsonOut {
		return emitJSON(av)
	}
	fmt.Printf("status=%s from=%s %s\n", av.Status, av.From, av.Currency)
	return nil
}
