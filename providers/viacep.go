package providers

import (
	"encoding/json"
)

// ViaCep is a struct used to query api of viacep
type ViaCep struct {
	cep string
}

type addressViaCep struct {
	City  string `json:"localidade"`
	State string `json:"uf"`
	Cep   string `json:"cep"`
}

func (address addressViaCep) getCity() string {
	return address.City
}

func (address addressViaCep) getState() string {
	return address.State
}

func (address addressViaCep) getCep() string {
	return address.Cep
}

func (provider ViaCep) findByCep(chanAddressable chan Addressable) {
	resp := queryWebservice("https://viacep.com.br/ws/%s/json/", provider.cep)
	body := convertResponseToByte(resp)
	address := addressViaCep{}
	json.Unmarshal(body, &address)
	chanAddressable <- address
}
