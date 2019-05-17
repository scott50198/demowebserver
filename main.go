package main

import (
	"log"
	"net/http"
	"sexdealwebserver/handler"
)

type ContactDetails struct {
	Email   string
	Subject string
	Message string
}

func main() {
	// tmpl := template.Must(template.ParseFiles(config.FrontendRoot + "/html/forms.html"))

	// http.HandleFunc("/index", func(w http.ResponseWriter, r *http.Request) {
	// 	if r.Method != http.MethodPost {
	// 		tmpl.Execute(w,
	// 			struct {
	// 				Success bool
	// 				Qitle   string
	// 			}{false, "123"})
	// 		fmt.Println("1")
	// 		return
	// 	}

	// 	details := ContactDetails{
	// 		Email:   r.FormValue("email"),
	// 		Subject: r.FormValue("subject"),
	// 		Message: r.FormValue("message"),
	// 	}

	// 	// do something with details
	// 	_ = details
	// 	fmt.Printf("%v", details)

	// 	tmpl.Execute(w, struct{ Success bool }{true})
	// })

	http.HandleFunc("/", handler.IndexHandler)
	http.Handle("/fronted/", http.StripPrefix("/fronted/", http.FileServer(http.Dir("fronted"))))

	log.Fatal(http.ListenAndServe(":8080", nil))

}
