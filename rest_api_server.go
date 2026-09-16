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
type Loot struct {
	BotID string `json:"bot_id"`
	Data  string `json:"data"`
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

http.HandleFunc ("/api/loot/", func(w http.ResponseWriter, r *http.Request){

	if r.Method!=http.MethodPost{
	http.Error(w, "Only Post Method", http.StatusMethodNotAllowed)
	return
	}
	
	var l Loot
	
	err := json.NewDecoder(r.Body).Decode(&l)
		if err != nil {
			http.Error(w, "Ошибка чтения JSON", http.StatusBadRequest)
			return
		}

		fmt.Printf("Bot '%s' Send data: %s\n", l.BotID, l.Data)

		fmt.Fprintf(w, "Data already on the server")


})
fmt.Println("Open server")

err:=http.ListenAndServe(":8082", nil)
	if err!=nil{
	fmt.Println("Problems", err)
	}
}
