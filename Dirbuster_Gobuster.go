package main

import (
"fmt"
"net/http"
)

func main() {

baseURL := "http://scanme.nmap.org"

paths := []string{"/admin", "/login", "/images","/robots.txt", "/backup.zip"}

	for _, path:= range paths {
	targetUrl := baseURL + path
	resp, err := http.Get(targetUrl)
	
	if err!=nil{
	fmt.Println("cant connect to web-site")
	continue
	}
	if resp.StatusCode == 200 {
	fmt.Println("web-site is real with way: " + targetUrl)
	}
	defer resp.Body.Close()
	
	}

}
