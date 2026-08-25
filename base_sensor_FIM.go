package main

import (
	"fmt"
	"os"
	"encoding/hex"
	"crypto/sha256"
)

func main() {

data, err := os.ReadFile("test.txt")
	if err!=nil{
		fmt.Println("Couldnt open file")
		return nil
	}
hash := sha256.Sum256(data)
expectedHash:="2eb742fa145d1a145a6b5ffebbb031273c47641b69ee0170da0da530f73057a1"
hashString := hex.EncodeToString(hash[:])

//fmt.Println ("file hash ",hashString)

	if hashString != expectedHash {
		fmt.Println("different hashes")
		return
	}else{
		fmt.Println("All good the file is default")
	}
}
