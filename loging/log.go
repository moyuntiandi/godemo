package loging

import (
	"fmt"
	"html/template"
	"net/http"

	logger "git.oschina.net/janpoem/go-logger"
)

func init() {
	logger.SetRollingDaily("D://logs", "mm.log")
	logger.SetLevel(logger.LOG)
}

func defaultIndex(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "Hello World")
}

func login(w http.ResponseWriter, r *http.Request) {

	r.ParseForm()

	if r.Method == "GET" {
		t, _ := template.ParseFiles("login.gtpl")
		t.Execute(w, nil)
	} else {
		for i := 0; i < 10; i++ {
			logger.Log("username", r.Form["username"])
			logger.Log("password", r.Form["password"])
		}
	}
}
