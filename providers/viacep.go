package providers

import (
	"encoding/json"
)

// ViaCep is a struct used to query api of viacep
type ViaCep struct {
	Cep string
}

type addressViaCep struct {
	City  string `json:"localidade"`
	State string `json:"uf"`
}

/*
FindByCep is a function which receive a cep (string) and return an Addressable
*/
func (provider ViaCep) FindByCep(chanAddress chan Address) {
	resp := queryWebservice("https://viacep.com.br/ws/%s/json/", provider.Cep)
	body := convertResponseToByte(resp)
	address := addressViaCep{}
	json.Unmarshal(body, &address)
	chanAddress <- converAddressViaCeptToAddress(address)
}

func converAddressViaCeptToAddress(address addressViaCep) Address {
	return Address{
		City:  address.City,
		State: address.State,
	}
}
