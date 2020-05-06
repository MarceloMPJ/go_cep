package main

import (
	"fmt"

	"github.com/MarceloMPJ/go_cep/providers"
)

func main() {
	fmt.Println(providers.FindByCep("74835655"))
}
