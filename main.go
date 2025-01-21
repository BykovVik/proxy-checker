package main

import "fmt"

func main() {
	proxies, err := getJsonData()
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, proxy := range proxies {
		fmt.Printf("IP: %s, PORT: %d\n", proxy.IP, proxy.PORT)
	}
}
