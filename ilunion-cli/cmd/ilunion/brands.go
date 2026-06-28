package main

import (
	"fmt"

	"github.com/fbelchi/ilunion-cli/internal/client"
)

func cmdBrands() {
	for _, b := range client.Brands {
		fmt.Println(b)
	}
}
