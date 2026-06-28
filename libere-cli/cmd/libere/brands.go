package main

import (
	"fmt"

	"github.com/fbelchi/libere-cli/internal/client"
)

func cmdBrands() {
	for _, b := range client.Brands {
		fmt.Println(b)
	}
}
