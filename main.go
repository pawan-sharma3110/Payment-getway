package main

import (
	"fmt"
	"log"
	"net/http"
	"payment_getway/db"
	"payment_getway/handler"
)

func main() {
	database, err := db.DbIn()
	if err != nil {
		panic(err)
	}
	defer database.Close()

	fmt.Println("Server run on port :8080")

	http.HandleFunc("/register", handler.RegisterUser)
	http.HandleFunc("/login", handler.LoginUser)

	log.Fatal(http.ListenAndServe(":8080", nil))

}
