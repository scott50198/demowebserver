package handler

import (
	"io/ioutil"
	"net/http"
	"sexdealwebserver/config"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {

	contents, err := ioutil.ReadFile(config.FrontendRoot + "/html/index.html")

	if err != nil {
		panic(err)
	}

	w.Write(contents)
}
