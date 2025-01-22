package main

import "fmt"

func main() {
	proxies, err := getJsonData()
	if err != nil {
		fmt.Println("No working proxy found:", err)
		return
	}

	ip, port, checkError := checkFirstWorkingProxy(proxies)

	if checkError != nil {
		fmt.Println(checkError)
	}

	fmt.Printf("Рабочий прокси: %s:%d", ip, port)
	fmt.Println("")
}
