package main

import (
	"fmt"

	"github.com/fbelchi/htop-cli/internal/client"
)

func cmdBrands() {
	for _, b := range client.Brands {
		fmt.Println(b)
	}
}
