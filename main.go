package main

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

func main() {
	c1:= make(chan Cep)
	c2:= make(chan Cep)

	go func(){
		for{
			req, err := http.Get("cdn.apicep.com/file/apicep/68537-000.json/")
			if err != nil {
				fmt.Fprintf(os.Stderr, "Erro ao fazer requisição: %v\n", err)
			}
			println("APICEP")
			c1<-req
		}
	}()

	go func(){
		for{
			req, err := http.Get("https://viacep.com.br/ws/68537000/json/")
			if err != nil {
				fmt.Fprintf(os.Stderr, "Erro ao fazer requisição: %v\n", err)
			}
			println("VIACEP")
			c1<-req
		}
	}
https: //cdn.apicep.com/file/apicep/68537-000.json
https: //viacep.com.br/ws/68537000/json/
}
