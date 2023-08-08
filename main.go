package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

type Cep struct {
	Cep         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	Uf          string `json:"uf"`
	Ibge        string `json:"ibge"`
	Gia         string `json:"gia"`
	Ddd         string `json:"ddd"`
	Siafi       string `json:"siafi"`
}
type Cep2 struct {
	Cep          string `json:"cep"`
	State        string `json:"state"`
	City         string `json:"city"`
	Neighborhood string `json:"neighborhood"`
	Street       string `json:"street"`
	Service      string `json:"service"`
}

func main() {
	c1 := make(chan Cep2)
	c2 := make(chan Cep)

	go func() {
		for {
			req, err := http.Get("https://brasilapi.com.br/api/cep/v1/06233030")
			if err != nil {
				fmt.Fprintf(os.Stderr, "Erro ao fazer requisição: %v\n", err)
			}
			defer req.Body.Close()
			res, err := ioutil.ReadAll(req.Body)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Erro ao ler resposta:%s\n", err)
			}
			var data Cep2
			err = json.Unmarshal([]byte(res), &data)
			if err != nil {
				fmt.Fprintf(os.Stderr, "%+v\n", err)
			}
			//println(time.After(time.Second))
			c1 <- data
		}
	}()

	go func() {
		for {
			req, err := http.Get("https://viacep.com.br/ws/06233030/json/")
			if err != nil {
				fmt.Fprintf(os.Stderr, "Erro ao fazer requisição: %v\n", err)
			}
			defer req.Body.Close()
			res, err := ioutil.ReadAll(req.Body)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Erro ao ler resposta:%s\n", err)
			}
			var data Cep
			err = json.Unmarshal([]byte(res), &data)
			if err != nil {
				fmt.Fprintf(os.Stderr, "%+v\n", err)
			}
			c2 <- data
		}
	}()
	for {
		select {
		case data := <-c1:
			fmt.Printf("\nAPICEP ", data.City, data.State)

		case data := <-c2:
			fmt.Printf("\nVIACEP ", data.Cep, data.Bairro, data.Complemento, data.Logradouro, data.Uf)

		case <-time.After(time.Second * 1):
			fmt.Println("timeout")
		}

		//https: //cdn.apicep.com/file/apicep/68537-000.json
		//	https: //viacep.com.br/ws/68537000/json/
	}
}
