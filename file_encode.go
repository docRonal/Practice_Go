package main

import(
"fmt"
"os"
)

func main() {

key:="hacker"
keyBytes := []byte(key)


file, err:= os.ReadFile("test.txt")
	if err!=nil{
	fmt.Println("Couldnt open file", err)
	return 
	}
	for i:=0; i<len(file);i++{
	file[i] = file[i] ^ keyBytes[i % len(keyBytes)]
	}
	err = os.WriteFile("test.enc", file, 0644)
	if err!=nil{
	fmt.Println("Couldnt save file", err)
	return 
	}
	fmt.Println("All good")

}	
