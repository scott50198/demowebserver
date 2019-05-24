package handler

import (
	"demowebserver/auth"
	"demowebserver/config"
	"demowebserver/model"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/sessions"
)

var store = sessions.NewCookieStore([]byte("goSessionId"))

func IndexHandler(w http.ResponseWriter, r *http.Request) {

	session, err := store.Get(r, "goSessionId")

	if err != nil {
		fmt.Println(err.Error())
	}

	if session.IsNew {
		http.Redirect(w, r, "/welcome", http.StatusFound)
		return
	}

	contents, _ := ioutil.ReadFile(config.FrontendRoot + "/html/main.html")
	w.Write(contents)
}

func WelcomeHandler(w http.ResponseWriter, r *http.Request) {
	contents, err := ioutil.ReadFile(config.FrontendRoot + "/html/welcome.html")
	if err != nil {
		fmt.Println(err.Error())
	}
	w.Write(contents)
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("login"))
}

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		contents, err := ioutil.ReadFile(config.FrontendRoot + "/html/register.html")
		if err != nil {
			fmt.Println(err.Error())
		}
		w.Write(contents)
	} else if r.Method == "POST" {

		r.ParseForm()

		info := model.UserInfo{
			Account:  r.Form.Get("account"),
			Password: r.Form.Get("password"),
			Name:     r.Form.Get("name"),
			Email:    r.Form.Get("email"),
		}

		err := auth.Register(info)

		if err != nil {
			fmt.Println(err.Error())
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			jj, _ := json.Marshal(struct {
				Status int    `json:"status"`
				Msg    string `json:"msg"`
			}{
				200,
				err.Error(),
			})

			w.Write(jj)
		} else {
			w.Write([]byte("ok"))
		}

	}
}
