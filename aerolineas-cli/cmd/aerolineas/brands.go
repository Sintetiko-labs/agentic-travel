package main

import (
	"fmt"

	"github.com/fbelchi/aerolineas-cli/internal/client"
)

func cmdBrands() {
	for _, b := range client.Brands {
		fmt.Println(b)
	}
}
