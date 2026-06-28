package main

import (
	"fmt"

	"github.com/fbelchi/airalgerie-cli/internal/client"
)

func cmdBrands() {
	for _, b := range client.Brands {
		fmt.Println(b)
	}
}
