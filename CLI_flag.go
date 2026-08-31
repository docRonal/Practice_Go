package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"net/http"
)

type VTResponse struct {
	Data struct {
		Attributes struct {
			Stats struct {
				Malicious int `json:"malicious"`
				Harmless  int `json:"harmless"`
			} `json:"last_analysis_stats"`
		} `json:"attributes"`
	} `json:"data"`
}

func main() {
	hashFlag := flag.String("hash", "", "24d004a104d4d54034dbcffc2a4b19a11f39008a575aa614ea04703480b1022c")
	flag.Parse()

	if *hashFlag == "" {
		fmt.Println("provide a hash using the -hash flag.")
		return
	}

	targetHash := *hashFlag
	url := "https://www.virustotal.com/api/v3/files/" + targetHash

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("Can't create request:", err)
		return
	}

	req.Header.Add("x-apikey", "API_KEY")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Connection problem:", err)
		return
	}
	defer resp.Body.Close()

	bodyBytes, _ := io.ReadAll(resp.Body)
	if resp.StatusCode != 200 {
		fmt.Println("API error:", resp.StatusCode)
		fmt.Println("Raw response data:", string(bodyBytes))
		return
	}

	var result VTResponse
	err = json.Unmarshal(bodyBytes, &result)
	if err != nil {
		fmt.Println("Can't parse JSON:", err)
		return
	}

	fmt.Println("Target Hash:", targetHash)
	fmt.Println("Malicious detections:", result.Data.Attributes.Stats.Malicious)
	fmt.Println("Harmless detections:", result.Data.Attributes.Stats.Harmless)
}
