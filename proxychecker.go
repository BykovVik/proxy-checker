package main

import (
	"fmt"
	"net/http"
	"net/url"
	"time"
)

func proxyChecker(ip string, port int) bool {

	proxyURL := fmt.Sprintf("http://%s:%d", ip, port)
	proxyURLParsed, err := url.Parse(proxyURL)

	if err != nil {
		fmt.Println(err)
	}

	client := &http.Client{
		Transport: &http.Transport{
			Proxy: http.ProxyURL(proxyURLParsed),
		},
		Timeout: 5 * time.Second,
	}
	resp, err := client.Get("http://httpbin.org/ip")
	fmt.Println(resp, err)
	if err != nil {
		return false
	}
	defer resp.Body.Close()

	return resp.StatusCode == http.StatusOK
}

func checkFirstWorkingProxy(proxies []Proxy) (string, int, error) {
	for _, proxy := range proxies {
		if proxyChecker(proxy.IP, proxy.PORT) {
			return proxy.IP, proxy.PORT, nil
		}
	}
	return "", 0, fmt.Errorf("no working proxy found")
}
