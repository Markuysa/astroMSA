package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

// checking the response
func main() {

	url := "https://sameer-kumar-aztro-v1.p.rapidapi.com/?sign=aquarius&day=today"

	req, _ := http.NewRequest("POST", url, nil)

	req.Header.Add("X-RapidAPI-Key", "54208fc179msh16e7cc7ea2939dcp1eb840jsnde68a4a6c957")
	req.Header.Add("X-RapidAPI-Host", "sameer-kumar-aztro-v1.p.rapidapi.com")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println(string(body))
	
}
