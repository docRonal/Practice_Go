package main

import(
"encoding/json"
"fmt"
"net/http"
)

type User struct {
	ID int `json:"id"` 
	Name string `json:"name"`
	Role string `json:"role"`
}

func main(){

http.HandleFunc("/api/users", func(w http.ResponseWriter, r *http.Request){

	w.Header().Set("Content-Type", "application/json")
	
	users:=[]User{
		{ID:1, Name: "Admin", Role:"Admin"},
		{ID:2, Name: "Hacker", Role:"Superuser"},
		{ID:3, Name: "Guest", Role:"User"},
	}

	json.NewEncoder(w).Encode(users)
})

fmt.Println("Open server")

err:=http.ListenAndServe(":8082", nil)
	if err!=nil{
	fmt.Println("Problems", err)
	}
}
