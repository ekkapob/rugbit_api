package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/gorilla/securecookie"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"success":true}`))
}

func main() {
	fmt.Println(securecookie.GenerateRandomKey(64))
	fmt.Println(securecookie.GenerateRandomKey(32))
	r := mux.NewRouter()
	r.HandleFunc("/", Index)
	http.ListenAndServe(":8080", r)
}
