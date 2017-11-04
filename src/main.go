package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
)

func main() {
	var email, user, empresa, token string
	email = os.Getenv("EMAIL")
	user = os.Getenv("USER")
	empresa = os.Getenv("EMPRESA")
	token = os.Getenv("TOKEN")

	// Generated by curl-to-Go: https://mholt.github.io/curl-to-go

	req, err := http.NewRequest("GET", "https://"+empresa+".acelerato.com/v2/apontamentos/usuarios/"+user, nil)
	if err != nil {
		// handle err
	}
	req.SetBasicAuth(email, token)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		// handle err
	}
	defer resp.Body.Close()

	var parsed map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&parsed); err != nil {
		log.Println(err)
	}
	//fmt.Println(parsed["content"])
	apontamentos := parsed["content"].([]interface{})
	var total float64
	var minutosTotal float64
	for _, apontamento := range apontamentos {
		mes := "11/2017"
		dataDoLancamento := apontamento.(map[string]interface{})
		quantidadeFormatada := apontamento.(map[string]interface{})
		if mes == dataDoLancamento["dataDoLancamento"].(string)[3:10] {
			//fmt.Println(apontamento, index, dataDoLancamento["dataDoLancamento"])
			horasToStr := quantidadeFormatada["quantidadeFormatada"].(string)[0:2]
			minutosToStr := quantidadeFormatada["quantidadeFormatada"].(string)[3:5]
			minutos, _ := strconv.ParseFloat(minutosToStr, 64)
			horas, _ := strconv.ParseFloat(horasToStr, 64)
			minutosTotal += minutos
			total += horas
		}
	}
	fmt.Printf("Você bilhetou um total de %.2f horas este mês", total+(minutosTotal/60))
}
