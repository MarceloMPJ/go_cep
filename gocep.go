package gocep

import "github.com/MarceloMPJ/go_cep/providers"

/*
FindByCep is a function which receive a string and return datas of address
quering providers (ViaCep and Correios)
*/
func FindByCep(string cep) Address {
	return providers.FindByCep("74835655")
}
