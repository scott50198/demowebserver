package main

import (
	"demowebserver/dbhelper"
	"demowebserver/handler"
	"log"
	"net/http"
)

func main() {

	err := dbhelper.OpenDB()
	if err != nil {
		panic(err)
	}

	http.HandleFunc("/", handler.IndexHandler)
	http.HandleFunc("/login", handler.LoginHandler)
	http.HandleFunc("/register", handler.RegisterHandler)
	http.HandleFunc("/welcome", handler.WelcomeHandler)
	http.Handle("/fronted/", http.StripPrefix("/fronted/", http.FileServer(http.Dir("fronted"))))

	log.Fatal(http.ListenAndServe(":80", nil))

}
