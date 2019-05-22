package main

import (
	"demowebserver/handler"
	"log"
	"net/http"
)

type ContactDetails struct {
	Email   string
	Subject string
	Message string
}

func main() {

	http.HandleFunc("/", handler.IndexHandler)
	http.HandleFunc("/login", handler.LoginHandler)
	http.HandleFunc("/register", handler.RegisterHandler)
	http.HandleFunc("/welcome", handler.WelcomeHandler)
	http.Handle("/fronted/", http.StripPrefix("/fronted/", http.FileServer(http.Dir("fronted"))))

	log.Fatal(http.ListenAndServe(":8080", nil))

}
