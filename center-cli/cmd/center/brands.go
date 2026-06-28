package main

import (
	"fmt"

	"github.com/fbelchi/center-cli/internal/client"
)

func cmdBrands() {
	for _, b := range client.Brands {
		fmt.Println(b)
	}
}
