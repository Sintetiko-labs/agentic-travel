package main

import (
	"flag"
	"fmt"
	"strings"

	"github.com/fbelchi/alegria-cli/internal/client"
)

func cmdSearch(args []string) error {
	fs := flag.NewFlagSet("search", flag.ExitOnError)
	cf := addCommon(fs)
	limit := fs.Int("limit", 24, "max results")
	page := fs.Int("page", 1, "page number")
	_ = fs.Parse(reorderArgs(fs, args))
	if fs.NArg() == 0 {
		return fmt.Errorf("usage: alegria search [flags] <destination...>")
	}
	cl := client.New("")
	query := strings.Join(fs.Args(), " ")
	res, err := cl.Search(query, *page, *limit)
	if err != nil {
		return err
	}
	if cf.jsonOut {
		return emitJSON(res)
	}
	fmt.Printf("query=%q total=%d page=%d source=%s\n", res.Query, res.Total, res.Page, res.Source)
	for _, h := range res.Hotels {
		fmt.Printf("  [%s] %s — %s\n", h.ID, h.Name, h.Price)
	}
	return nil
}
