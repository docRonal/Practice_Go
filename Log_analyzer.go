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
XSSPatern := regexp.MustCompile(`(?i)<script`)
PathPatern := regexp.MustCompile(`\.\./`)

	for _, line :=range lines {
		if sqlPattern.MatchString(line){	
				fmt.Println("ALERT: SQL Injection detected:", line)
		}
		if XSSPatern.MatchString(line){
			fmt.Println("XSS atack detected", line)
		}
		if PathPatern.MatchString(line){
			fmt.Println("Path Traversal detected", line)
		}

	}
	
}

