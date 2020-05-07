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
	FindByCep(chan Address)
}

/*
Address is a struct which is possible get the attributes of a address
(city, state, cep)
*/
type Address struct {
	City  string
	State string
}

/*
FindByCep is a function which receive a cep (string) and return an Addressable
*/
func FindByCep(cep string) Address {
	providers := []Provider{
		ViaCep{Cep: cep},
		Correios{Cep: cep},
	}
	return consultProviders(providers)
}

func consultProviders(providers []Provider) Address {
	chanAddress := make(chan Address)
	for i := 0; i < len(providers); i++ {
		go providers[i].FindByCep(chanAddress)
	}
	return <-chanAddress
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
