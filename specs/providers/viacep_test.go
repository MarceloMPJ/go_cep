package providers

import (
	"reflect"
	"testing"

	"github.com/MarceloMPJ/go_cep/providers"
)

func TestViacepFindByCep(t *testing.T) {
	for i := 0; i < len(testAddressTable); i++ {
		chanAddress := make(chan providers.Address)

		cep := testAddressTable[i].cep
		correios := providers.ViaCep{Cep: cep}

		go correios.FindByCep(chanAddress)

		expected := testAddressTable[i].address
		actual := <-chanAddress

		if !reflect.DeepEqual(actual, expected) {
			t.Errorf("Expected is %v but is %v", expected, actual)
		}
	}
}
