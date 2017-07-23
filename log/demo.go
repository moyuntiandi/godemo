package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	logger "git.oschina.net/janpoem/go-logger"
)

func init() {
	logger.SetRollingDaily("/opt/var/log", "mm.log")
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
		for i := 0; i < 100; i++ {
			logger.Log("username", r.Form["username"])
			logger.Log("password", r.Form["password"])
		}
		t, _ := template.ParseFiles("login.gtpl")
		t.Execute(w, nil)
	}
}
func main() {

	http.HandleFunc("/", defaultIndex)
	http.HandleFunc("/login", login)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

}
