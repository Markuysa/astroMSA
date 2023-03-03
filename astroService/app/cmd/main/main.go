package main

import (
	"astroService/internal/config"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// checking the response
func main() {

	url := "https://sameer-kumar-aztro-v1.p.rapidapi.com/?sign=aquarius&day=today"

	req, _ := http.NewRequest("POST", url, nil)

	config, err := config.New()
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Add("X-RapidAPI-Key", config.ApiKey)
	req.Header.Add("X-RapidAPI-Host", config.ApiHost)

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println(string(body))

}
