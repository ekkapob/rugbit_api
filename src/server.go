package main

import (
	"authen"
	"database/sql"
	dblib "db"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

type Response struct {
	error string
	data  authen.User
}

func indexHandler(res http.ResponseWriter, req *http.Request) {
	// w.Header().Set("Content-Type", "application/json")
	// w.Write([]byte(`{"success":true}`))
	//fmt.Println("index: ", authen.GetCookieUserName(req))
}

func loginHandler(res http.ResponseWriter, req *http.Request, db *sql.DB) {
	username := req.FormValue("username")
	password := req.FormValue("password")
	// fmt.Println(username, ":", password)
	err, _ := authen.Authen(db, username, password)

	res.Header().Set("Content-Type", "application/json")

	if err != nil {
		res.WriteHeader(http.StatusUnauthorized)

		res.Write([]byte("Invalid username or password"))
	}

	res.Write([]byte("Invalid username or password"))
}

func logoutHandler(res http.ResponseWriter, req *http.Request) {
	authen.ClearCookie(res)
}

func dbHandler(fn func(http.ResponseWriter, *http.Request, *sql.DB), db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fn(w, r, db)
	}
}

func main() {
	db := dblib.GetDb()
	defer db.Close()

	// fmt.Println(db.GetUserPwd(db, "ekkapob@gmail.com"))
	// authen.GetUser("123", "thailand")

	// Test login
	// authen.Login(db, "ekkapob@gmail.com", "thailand")
	// authen.Login(db, "test@gmail.com", "hello")

	// Test Signup
	err := authen.Signup(db, "abc@gmail.com", "Icanfly")
	if err != nil {
	}

	// Test
	// fmt.Println(authen.GetUser(db, "ekkapob@gmail.com"))
	// fmt.Println(authen.GetUser(db, "test@gmail.com"))
	fmt.Println(authen.Authen(db, "ekkapob@gmail.com", "thailand"))
	fmt.Println(authen.Authen(db, "test@gmail.com", "hello"))

	r := mux.NewRouter()
	r.HandleFunc("/", indexHandler)
	r.HandleFunc("/login", dbHandler(loginHandler, db)).Methods("POST")
	r.HandleFunc("/logout", logoutHandler).Methods("POST")
	http.ListenAndServe(":8080", r)
}
