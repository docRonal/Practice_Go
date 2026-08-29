package main

import(
"fmt"
"io"
"net/http"
"encoding/json"
)

type VTResposnse struct{
	Data struct{
		Atributes struct{
			Stats struct{
				Malicious int `json:"malicious"`
				Harmless  int `json:"harmless"`
			} `json:"last_analysis_stats"`
		} `json:"attributes"`
	} `json:"data"`
}

func main() {

client:=&http.Client{}
targetHash:="24d004a104d4d54034dbcffc2a4b19a11f39008a575aa614ea04703480b1022c" // hash of virus WannaCry
url:="https://www.virustotal.com/api/v3/files/" + targetHash
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
	if resp.StatusCode != 200 {
		fmt.Println ("API error", resp.StatusCode)
		fmt.Println ("Unsorted data", string(bodyBytes))
		return 
	}
	
	var result VTResposnse
	
	err = json.Unmarshal(bodyBytes, &result)
	if err!=nil{
		fmt.Println("Cant pars JSON file", err)
		return
	}
	
	

	fmt.Println("The server danger part: ", result.Data.Atributes.Stats.Malicious)
	fmt.Println("The server save part ", result.Data.Atributes.Stats.Harmless)
	
}



