package main

import (
	"authen"
	"github.com/gorilla/mux"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	// w.Header().Set("Content-Type", "application/json")
	// w.Write([]byte(`{"success":true}`))
	authen.SetCookie(w, map[string]string{"foo": "bar"})
}

func Login(w http.ResponseWriter, r *http.Request) {
	authen.GetCookie(r, "session")
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", Index)
	r.HandleFunc("/login", Login).Methods("POST")
	http.ListenAndServe(":8080", r)
}
