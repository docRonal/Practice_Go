package main
import "fmt"

func main (){
	var password string
	var stopword string
	for {
	  fmt.Print("enter password")
	  fmt.Scan(&password)
	  if len(password) < 8 {
		fmt.Println("short password")
	  } else if len (password) > 20 {
		fmt.Println("long password")
	  } else{
		fmt.Println("goot password")
	  }
	  if (stopword == "stop"){
	  break
	  }
  }
}
