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

	contents, err := ioutil.ReadFile(config.FrontendRoot + "/index.html")
	if err != nil {
		fmt.Println(err.Error())
	}
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

			setSession(account, w, r)

			resp := model.Response{
				StatusCode: http.StatusOK,
				Msg:        "登入成功",
				Contents: struct {
					Path string `json:"path"`
				}{
					"/",
				},
			}
			w.Header().Set("Content-Type", "application/json")
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

		resp := new(model.Response)
		if err != nil {
			resp.StatusCode = http.StatusConflict
			resp.Msg = err.Error()
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(resp)
		} else {

			setSession(info.Account, w, r)
			resp.StatusCode = http.StatusOK
			resp.Msg = "OK"
			resp.Contents = struct {
				Path string `json:"path"`
			}{
				"/",
			}

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(resp)
		}
	}
}

func setSession(account string, w http.ResponseWriter, r *http.Request) {
	session, err := store.Get(r, "goSessionId")

	if err != nil {
		fmt.Println(err.Error())
	}

	info, err := dbhelper.GetUserInfoFromAccount(account)

	if err != nil {
		fmt.Println(err.Error())
	}

	session.Values["Account"] = info.Account
	session.Values["Id"] = info.Id
	session.Values["Name"] = info.Name
	session.Values["Email"] = info.Email

	session.Save(r, w)

}
