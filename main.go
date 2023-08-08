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
		for i := 0; i < 1; i++ {
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
			c1 <- data
			i++

		}
	}()

	go func() {
		for i := 0; i < 1; i++ {
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
			i++
		}
	}()
	for i := 0; i < 3; i++ {
		select {
		case data := <-c1:
			fmt.Printf("\nbrasilapi ", data.Cep, data.City, data.State, data.Street, data.Neighborhood)

		case data := <-c2:
			fmt.Printf("\nVIACEP ", data.Cep, data.Localidade, data.Uf, data.Logradouro, data.Bairro)

		case <-time.After(time.Second * 1):
			fmt.Println("timeout")
		}
		i++

	}
}
