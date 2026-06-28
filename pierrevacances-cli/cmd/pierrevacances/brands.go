package main

import (
	"fmt"

	"github.com/fbelchi/pierrevacances-cli/internal/client"
)

func cmdBrands() {
	for _, b := range client.Brands {
		fmt.Println(b)
	}
}
