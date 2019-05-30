package handler

import (
	"demowebserver/auth"
	"demowebserver/config"
	"demowebserver/dbhelper"
	"demowebserver/model"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/sessions"
)

var store = sessions.NewCookieStore([]byte("goSessionId"))

func IndexHandler(w http.ResponseWriter, r *http.Request) {

	// session, err := store.Get(r, "goSessionId")

	// if err != nil {
	// 	fmt.Println(err.Error())
	// }

	// if session.IsNew {
	// 	http.Redirect(w, r, "/welcome", http.StatusFound)
	// 	return
	// }

	// contents, _ := ioutil.ReadFile(config.FrontendRoot + "/html/main.html")
	// w.Write(contents)
	w.Write([]byte("hello world"))
}

func WelcomeHandler(w http.ResponseWriter, r *http.Request) {
	contents, err := ioutil.ReadFile(config.FrontendRoot + "/html/welcome.html")
	if err != nil {
		fmt.Println(err.Error())
	}
	w.Write(contents)
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		contents, err := ioutil.ReadFile(config.FrontendRoot + "/html/login.html")
		if err != nil {
			fmt.Println(err.Error())
		}
		w.Write(contents)
	} else if r.Method == "POST" {
		r.ParseForm()

		account := r.Form.Get("account")
		password := r.Form.Get("password")

		if !dbhelper.CheckAccountAndPasswordValidate(account, password) {
			resp := model.Response{
				StatusCode: http.StatusConflict,
				Msg:        "帳號或密碼錯誤，請重新輸入",
			}
			json.NewEncoder(w).Encode(resp)
		} else {
			resp := model.Response{
				StatusCode: http.StatusOK,
				Msg:        "登入成功",
			}
			json.NewEncoder(w).Encode(resp)
		}
	}
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
			resp := model.Response{
				StatusCode: http.StatusConflict,
				Msg:        err.Error(),
			}
			json.NewEncoder(w).Encode(resp)
		} else {

			setSession(info, w, r)

			// resp := model.Response{
			// 	StatusCode: http.StatusOK,
			// 	Msg:        "OK",
			// }
			// json.NewEncoder(w).Encode(resp)
			http.Redirect(w, r, "/", http.StatusFound)
		}
	}
}

func setSession(info model.UserInfo, w http.ResponseWriter, r *http.Request) {
	session, err := store.Get(r, "goSessionId")
	w.Header().Set("Content-Type", "text/html")

	if err != nil {
		fmt.Println(err.Error())
	}

	id, err := dbhelper.GetUserIdFromAccount(info.Account)

	if err != nil {
		fmt.Println(err.Error())
	}

	info.Id = id

	session.Values["Account"] = info.Account
	session.Values["Id"] = info.Id
	session.Values["Name"] = info.Name
	session.Values["Email"] = info.Email

	session.Save(r, w)

}
