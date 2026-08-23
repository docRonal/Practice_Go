package main

import(
"fmt"
)

func main() {

target := "hacker"

key:=byte(42) //random number from 0 to 255

data:= []byte(target)

	for i:=0; i<len(data); i++{
	data[i] = data[i] ^ key
	}
fmt.Println("Encode ", string(data))

	for i:=0; i<len(data); i++{
	data[i] = data[i] ^ key
	}
	
fmt.Println("Dencode ", string(data))
}
