package main

import (
	"fmt"

	"github.com/fbelchi/vietnamairlines-cli/internal/client"
)

func cmdBrands() {
	for _, b := range client.Brands {
		fmt.Println(b)
	}
}
