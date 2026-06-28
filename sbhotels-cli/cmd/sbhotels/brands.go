package main

import (
	"fmt"

	"github.com/fbelchi/sbhotels-cli/internal/client"
)

func cmdBrands() {
	for _, b := range client.Brands {
		fmt.Println(b)
	}
}
