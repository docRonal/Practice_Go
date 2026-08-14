package main
import (
"fmt" 
"unicode"
)

func main (){
	var password string
	var stopword string 
	for {
	fmt.Print("Enter paswwrod: ")
	fmt.Scan(&password)
	hasDigit := false
	hasUpper := false 
	for _, char := range password{
		if unicode.IsDigit(char){
		hasDigit = true  
		}
		if unicode.IsUpper(char){
		hasUpper = true
		}
	}
	if len(password) < 8 {
		fmt.Println("short password")
	} else if len (password) > 20 {
		fmt.Println("long password")
	} else if (!hasUpper){
		fmt.Println("bad password")
	}else if (!hasDigit){
		fmt.Println("bad password")
	}else{
	fmt.Println("good")
	}
	if (stopword == "stop"){
	break
	}
	
}

}
