package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func login(w http.ResponseWriter, r *http.Request) {
	var tmplt = template.Must(template.ParseFiles("login.html"))
	tmplt.Execute(w, nil)
}

func signup(w http.ResponseWriter, r *http.Request) {
	var tmplt = template.Must(template.ParseFiles("signup.html"))
	tmplt.Execute(w, nil)
}
func signup_user(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	var username = r.Form["new-username"]
	var password = r.Form["new-password"]
	fmt.Println(username, " ", password)
	var tmplt = template.Must(template.ParseFiles("index.html"))
	tmplt.Execute(w, nil)
}

func login_user(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	var username = r.Form["username"]
	var password = r.Form["password"]
	fmt.Println(username, " ", password)
	var tmplt = template.Must(template.ParseFiles("index.html"))
	tmplt.Execute(w, nil)
}

func main() {
	http.HandleFunc("/", login)
	http.HandleFunc("/signup", signup)
	http.HandleFunc("/login_user", login_user)
	http.HandleFunc("/signup_user", signup_user)
	http.ListenAndServe(":8000", nil)
}
