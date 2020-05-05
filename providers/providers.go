package providers

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

/*
Provider is the resource used to find Address
*/
type Provider interface {
	findByCep(chan Addressable)
}

/*
Addressable is a interface which is possible get the attributes of a address
(city, state, cep)
*/
type Addressable interface {
	getCity() string
	getState() string
	getCep() string
}

/*
FindByCep is a function which receive a cep (string) and return an Addressable
*/
func FindByCep(cep string) Addressable {
	chanAddressable := make(chan Addressable)

	viaCep := ViaCep{cep: cep}
	correios := Correios{cep: cep}

	go viaCep.findByCep(chanAddressable)
	go correios.findByCep(chanAddressable)

	return <-chanAddressable
}

func queryWebservice(urlProvider string, cep string) *http.Response {
	url := fmt.Sprintf(urlProvider, cep)
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("Error on query API: %v\n", err)
		os.Exit(1)
	}
	return resp
}

func convertResponseToByte(resp *http.Response) []byte {
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error on convert response to byte: %v\n", err)
		os.Exit(1)
	}
	return body
}
