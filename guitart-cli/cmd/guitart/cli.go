package main

import (
	"encoding/json"
	"flag"
	"os"
	"strings"
)

type common struct {
	jsonOut bool
}

func addCommon(fs *flag.FlagSet) *common {
	c := &common{}
	fs.BoolVar(&c.jsonOut, "json", false, "emit JSON to stdout")
	return c
}

func reorderArgs(fs *flag.FlagSet, args []string) []string {
	var flags, positional []string
	for i := 0; i < len(args); i++ {
		a := args[i]
		if a == "--" {
			positional = append(positional, args[i+1:]...)
			break
		}
		if len(a) > 1 && a[0] == '-' {
			flags = append(flags, a)
			name := strings.TrimLeft(a, "-")
			if strings.IndexByte(name, '=') >= 0 {
				continue
			}
			if f := fs.Lookup(name); f != nil {
				if _, ok := f.Value.(interface{ IsBoolFlag() bool }); !ok && i+1 < len(args) {
					flags = append(flags, args[i+1])
					i++
				}
			}
			continue
		}
		positional = append(positional, a)
	}
	out := append(flags, "--")
	return append(out, positional...)
}

func emitJSON(v any) error {
	enc := json.NewEncoder(os.Stdout)
	enc.SetIndent("", "  ")
	enc.SetEscapeHTML(false)
	return enc.Encode(v)
}
