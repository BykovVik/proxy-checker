package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Proxy struct {
	IP   string `json:"ip_address"`
	PORT int    `json:"port"`
}

func getJsonData() ([]Proxy, error) {

	// Read proxy list
	var data, readError = os.ReadFile("proxies.json")

	if readError != nil {
		fmt.Println("JSON file upload error")
	}

	// json decode
	var proxyList []Proxy
	var err = json.Unmarshal(data, &proxyList)

	if err != nil {
		fmt.Println("JSON file upload error")
	}

	return proxyList, nil

}
