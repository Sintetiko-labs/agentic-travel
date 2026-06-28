package main

import (
	"fmt"

	"github.com/fbelchi/25hours-cli/internal/client"
)

func cmdBrands() {
	for _, b := range client.Brands {
		fmt.Println(b)
	}
}
