package providers

import (
	"github.com/MarceloMPJ/go_cep/providers"
)

type testAddress struct {
	cep     string
	address providers.Address
}

var testAddressTable = []testAddress{
	testAddress{
		"74835655",
		providers.Address{
			State: "GO",
			City:  "Goiânia",
		},
	},
}
