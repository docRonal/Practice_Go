package main

import(
"fmt"
"io"
"net/http"
)

func main() {

client:=&http.Client{}
targetIP:="8.8.8.8"
url:="https://www.virustotal.com/api/v3/ip_addresses/" + targetIP
req, err:= http.NewRequest("GET", url, nil)

	if err!=nil{
		fmt.Println("cant connect")
		return
	}
	
	req.Header.Add("x-apikey", "API_KEY")
	resp, err:=client.Do(req)
	if err!=nil{
		fmt.Println("Connection Problem", err)
		return
	}
	defer resp.Body.Close()
	
	bodyBytes, _ := io.ReadAll(resp.Body)
	fmt.Println("The server answer: ", string(bodyBytes))
}


