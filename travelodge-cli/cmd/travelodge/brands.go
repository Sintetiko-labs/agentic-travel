package main

import (
	"fmt"

	"github.com/fbelchi/travelodge-cli/internal/client"
)

func cmdBrands() {
	for _, b := range client.Brands {
		fmt.Println(b)
	}
}
