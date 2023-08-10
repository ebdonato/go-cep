package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type CEP struct {
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

const URL = "http://viacep.com.br/ws/%s/json/"

func main() {
	file, err := os.Create("ceps.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	for k, cep := range os.Args[1:] {
		fmt.Println(k+1, cep)

		url := fmt.Sprintf(URL, cep)

		req, err := http.Get(url)
		if err != nil {
			fmt.Println(err)
			continue
		}
		defer req.Body.Close()

		res, err := io.ReadAll(req.Body)
		if err != nil {
			fmt.Println(err)
			continue
		}

		var data CEP
		err = json.Unmarshal(res, &data)
		if err != nil {
			fmt.Println(err)
			continue
		}

		_, err = file.Write([]byte(fmt.Sprintf("CEP %s: %s | %s, %s | DDD: %s\n", data.Cep, data.Logradouro, data.Localidade, data.Uf, data.Ddd)))
		if err != nil {
			fmt.Println(err)
			continue
		}
	}

	fmt.Println("Pronto!")
}
