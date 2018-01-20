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
var api_key string = "859611b2fc1eaee629198189391ced734af866a9"

func init() {
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

func main() {
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

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(body))
}
