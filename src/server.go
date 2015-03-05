package main

import (
	"authen"
	"db"
	//"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func indexHandler(res http.ResponseWriter, req *http.Request) {
	// w.Header().Set("Content-Type", "application/json")
	// w.Write([]byte(`{"success":true}`))
	//fmt.Println("index: ", authen.GetCookieUserName(req))
}

func loginHandler(res http.ResponseWriter, req *http.Request) {
	authen.SetCookie(res, "ek")
}

func logoutHandler(res http.ResponseWriter, req *http.Request) {
	authen.ClearCookie(res)
}

func main() {
	database := db.GetDb()
	defer database.Close()

	// fmt.Println(db.GetUserPwd(database, "ekkapob@gmail.com"))
	// authen.GetUser("123", "thailand")

	// Test login
	// authen.Login(database, "ekkapob@gmail.com", "thailand")
	// authen.Login(database, "test@gmail.com", "hello")

	// Test Signup
	err := authen.Signup(database, "abc@gmail.com", "Icanfly")
	if err != nil {

	}

	r := mux.NewRouter()
	r.HandleFunc("/", indexHandler)
	r.HandleFunc("/login", loginHandler).Methods("POST")
	r.HandleFunc("/logout", logoutHandler).Methods("POST")
	http.ListenAndServe(":8080", r)
}
