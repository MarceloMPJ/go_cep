package providers

import (
	"reflect"
	"testing"

	"github.com/MarceloMPJ/go_cep/providers"
)

type testCep struct {
	cep     string
	address providers.Address
}

var testCepTable = []testCep{
	testCep{
		"74835655",
		providers.Address{
			Cep:   "74835655",
			State: "GO",
			City:  "Goiânia",
		},
	},
}

func TestCorreiosFindByCep(t *testing.T) {
	for i := 0; i < len(testCepTable); i++ {
		chanAddress := make(chan providers.Address)

		cep := testCepTable[i].cep
		correios := providers.Correios{Cep: cep}

		go correios.FindByCep(chanAddress)

		expected := testCepTable[i].address
		actual := <-chanAddress

		if !reflect.DeepEqual(actual, expected) {
			t.Errorf("Expected is %v but is %v", expected, actual)
		}
	}
}
