package main

import (
	"fmt"

	"github.com/fbelchi/swiftair-cli/internal/client"
)

func cmdBrands() {
	for _, b := range client.Brands {
		fmt.Println(b)
	}
}
