package main

import (
	"fmt"

	"github.com/fbelchi/designhotels-cli/internal/client"
)

func cmdBrands() {
	for _, b := range client.Brands {
		fmt.Println(b)
	}
}
