package service

import (
	"graphql-intro/helper"
	"io/ioutil"
	"log"
	"net/http"
)

// FetchFromRajaongkir from api.rajaongkir.com
func FetchFromRajaongkir(data string) []byte {
	baseURL := "https://api.rajaongkir.com/starter"
	apiKey := helper.GetEnvVar("RAJAONGKIR_APIKEY")
	res, err := http.Get(baseURL + data + "?key=" + apiKey)
	if err != nil {
		log.Fatalln("Error -> ", err.Error())
	}

	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatalln("Error -> ", err.Error())
	}

	return resBody
}
