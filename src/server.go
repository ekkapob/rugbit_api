package main

import (
	"authen"
	"db"
	"fmt"
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

	fmt.Println(db.GetUserPwd(database, "ekkapob@gmail.com"))
	authen.GetUser("123", "thailand")

	r := mux.NewRouter()
	r.HandleFunc("/", indexHandler)
	r.HandleFunc("/login", loginHandler).Methods("POST")
	r.HandleFunc("/logout", logoutHandler).Methods("POST")
	http.ListenAndServe(":8080", r)
}
