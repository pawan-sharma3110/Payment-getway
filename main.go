package main

import (
	"fmt"
	"log"
	"net/http"
	"payment_getway/db"
)

func main() {
	database, err := db.DbIn()
	if err != nil {
		panic(err)
	}
	defer database.Close()

	log.Fatal(http.ListenAndServe(":8080", nil))
	fmt.Println("Server run on port :8080")
	http.HandleFunc("/register", registerUser)
	http.HandleFunc("/login", loginUser)

}
