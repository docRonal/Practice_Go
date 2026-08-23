package main

import(
"fmt"
 "encoding/hex" 
)

func main() {

payload := "http://my-c2-server.com/malware.exe"
payloadBytes := []byte(payload)
key:="secret" 
keyBytes := []byte(key)


	for i:=0; i<len(payloadBytes); i++{
	payloadBytes[i] = payloadBytes[i] ^ keyBytes[i % len(keyBytes)]
	}
fmt.Println("Encode ", hex.EncodeToString(payloadBytes))

	for i:=0; i<len(payloadBytes); i++{
	payloadBytes[i] = payloadBytes[i] ^ keyBytes[i % len(keyBytes)]
	}
fmt.Println("Encode ", string(payloadBytes))


}
