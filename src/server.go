package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/gorilla/securecookie"
	"net/http"
)

var s = securecookie.New(securecookie.GenerateRandomKey(64), securecookie.GenerateRandomKey(32))

func Index(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"success":true}`))
}

func Login(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.FormValue("name"))
	fmt.Println(r.FormValue("password"))
}

func main() {
	fmt.Println(s)
	r := mux.NewRouter()
	r.HandleFunc("/", Index)
	r.HandleFunc("/login", Login).Methods("POST")

	http.ListenAndServe(":8080", r)
}
