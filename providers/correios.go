package providers

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Correios is a struct used to query api of correios
type Correios struct {
	cep string
}

type addressCorreios struct {
	City  string `json:"cidade"`
	State string `json:"uf"`
	Cep   string `json:"cep"`
}

func (address addressCorreios) getCity() string {
	return address.City
}

func (address addressCorreios) getState() string {
	return address.State
}

func (address addressCorreios) getCep() string {
	return address.Cep
}

func (provider Correios) findByCep(chanAddressable chan Addressable) {
	url := webserviceCorreiosURL()
	resp := queryWebservice(url, provider.cep)
	body := convertResponseToByte(resp)
	address := addressCorreios{}
	json.Unmarshal(body, &address)
	chanAddressable <- address
}

func webserviceCorreiosURL() string {
	urlBase := "https://webmaniabr.com/api/1/cep/%s/"
	appKey := ""
	appSecret := ""
	queryKey := fmt.Sprintf("?app_key=%s&app_secret=%s", appKey, appSecret)
	urlArr := []string{urlBase, queryKey}
	return strings.Join(urlArr, "")
}
