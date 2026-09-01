package main

import (
"fmt"
"os"
"strings"
"regexp"
)

func main() {

data, err:= os.ReadFile("access.log")
	if err!=nil{
		fmt.Println("Couldnt open file")
		return 
	}
lines :=strings.Split(string(data), "\n")
sqlPattern := regexp.MustCompile(`(?i)union.*select`)

	for _, line :=range lines {
		if sqlPattern.MatchString(line){	
				fmt.Println("ALERT: SQL Injection detected:", line)
			
		}

	}
	
}
