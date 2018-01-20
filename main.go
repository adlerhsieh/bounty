package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"net/http"
)

var username string = ""
var password string = ""
var cst string = ""
var xst string = ""
var base_url string = ""
var api_key string = "859611b2fc1eaee629198189391ced734af866a9"

func init() {
	setConfig()
}

func setConfig() {
	data, err := ioutil.ReadFile("./config.yml")
	if err != nil {
		panic(err)
	}

	c := make(map[string]interface{})

	err = yaml.Unmarshal([]byte(data), &c)
	if err != nil {
		panic(err)
	}

	username = c["identifier"].(string)
	password = c["password"].(string)
}

func refreshTokens() {
	client := http.Client{}

	bodyString := map[string]string{"identifier": username, "password": password}
	reqBody, err := json.Marshal(bodyString)
	if err != nil {
		panic(err)
	}

	req, err := http.NewRequest("POST", "https://api.ig.com/gateway/deal/session", bytes.NewBuffer(reqBody))
	if err != nil {
		panic(err)
	}

	req.Header.Set("X-IG-API-KEY", api_key)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json; charset=UTF-8")

	res, err := client.Do(req)

	if err != nil {
		panic(err)
	}

	cst = res.Header.Get("CST")
	xst = res.Header.Get("X-SECURITY-TOKEN")
}

func main() {
	refreshTokens()

	client := http.Client{}

	req, err := http.NewRequest("GET", "https://api.ig.com/gateway/deal/watchlists", nil)
	if err != nil {
		panic(err)
	}

	req.Header.Set("X-SECURITY-TOKEN", xst)
	req.Header.Set("CST", cst)
	req.Header.Set("X-IG-API-KEY", api_key)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json; charset=UTF-8")

	res, err := client.Do(req)

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(body))
}
