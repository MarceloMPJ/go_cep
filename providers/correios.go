package providers

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Correios is a struct used to query api of correios
type Correios struct {
	Cep string
}

type addressCorreios struct {
	City  string `json:"cidade"`
	State string `json:"uf"`
}

/*
FindByCep is a function which receive a cep (string) and return an Addressable
*/
func (provider Correios) FindByCep(chanAddress chan Address) {
	url := webserviceCorreiosURL()
	resp := queryWebservice(url, provider.Cep)
	body := convertResponseToByte(resp)
	address := addressCorreios{}
	json.Unmarshal(body, &address)
	chanAddress <- convertAddressCorreiosToAddress(address)
}

func webserviceCorreiosURL() string {
	urlBase := "https://webmaniabr.com/api/1/cep/%s/"
	appKey := "hhHkMEF8MOGnDFipswdjO093rb9z5rMK"
	appSecret := "e48albDu6WG1Z7xQlIJ4AFW3Yw1smwzbXNWBqOJrDiBjYjYs"
	queryKey := fmt.Sprintf("?app_key=%s&app_secret=%s", appKey, appSecret)
	urlArr := []string{urlBase, queryKey}
	return strings.Join(urlArr, "")
}

func convertAddressCorreiosToAddress(address addressCorreios) Address {
	return Address{
		City:  address.City,
		State: address.State,
	}
}
