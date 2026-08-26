package main

import(
"fmt"
"time"
"os"
"crypto/sha256"
"encoding/hex"
)

func main() {

expectedHash :="ecb72a159b63efd0f17bf4720674d188979e2035dc14fefc650193af57c45369"


	for {
		data, err := os.ReadFile("test.txt")
		if err!=nil{
			fmt.Println("Couldnt open file")
			continue 
		}
		
		hash := sha256.Sum256(data)
		hashString :=hex.EncodeToString(hash[:])
		if hashString != expectedHash {
			fmt.Println("Different hashes")
			return 
		}else{
			fmt.Println("Same hashes")
		}
		time.Sleep(5*time.Second)
	}


}
