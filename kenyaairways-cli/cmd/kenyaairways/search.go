package main

import (
	"flag"
	"fmt"

	"github.com/fbelchi/kenyaairways-cli/internal/client"
)

func cmdSearch(args []string) error {
	fs := flag.NewFlagSet("search", flag.ExitOnError)
	cf := addCommon(fs)
	from := fs.String("from", "", "origin airport/city (IATA)")
	to := fs.String("to", "", "destination airport/city (IATA)")
	depart := fs.String("depart", "", "departure date (YYYY-MM-DD)")
	ret := fs.String("return", "", "return date (YYYY-MM-DD, optional)")
	limit := fs.Int("limit", 24, "max results")
	page := fs.Int("page", 1, "page number")
	_ = fs.Parse(reorderArgs(fs, args))
	if *from == "" || *to == "" || *depart == "" {
		return fmt.Errorf("usage: kenyaairways search [flags] (requires --from --to --depart)")
	}
	cl := client.New("")
	res, err := cl.Search(*from, *to, *depart, *ret, *page, *limit)
	if err != nil {
		return err
	}
	if cf.jsonOut {
		return emitJSON(res)
	}
	fmt.Printf("%s→%s depart=%s total=%d\n", res.Origin, res.Dest, res.Depart, res.Total)
	for _, f := range res.Flights {
		fmt.Printf("  [%s] %s %s→%s %s — %s\n", f.ID, f.FlightNumber, f.Origin, f.Destination, f.Depart, f.Price)
	}
	return nil
}
