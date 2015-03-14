package main

import (
	"authen"
	"database/sql"
	dblib "db"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"model"
	"net/http"
	"strings"
)

type Response struct {
	Success bool   `json:"success"`
	Error   string `json:"error,omitempty"`
	UserResponse
	PostResponse
}

type UserResponse struct {
	User *model.User `json:"user,omitempty"`
}

type PostResponse struct {
	Post *model.Post `json:"post,omitempty"`
}

func indexHandler(res http.ResponseWriter, req *http.Request) {
	// w.Header().Set("Content-Type", "application/json")
	// w.Write([]byte(`{"success":true}`))
	//fmt.Println("index: ", authen.GetCookieUserName(req))
}

func signupHandler(res http.ResponseWriter, req *http.Request, db *sql.DB) {

	user := model.User{
		Username:  strings.TrimSpace(req.FormValue("username")),
		Password:  strings.TrimSpace(req.FormValue("password")),
		Firstname: strings.TrimSpace(req.FormValue("firstname")),
		Lastname:  strings.TrimSpace(req.FormValue("lastname")),
	}

	response := &Response{}
	if len(user.Username) == 0 || len(user.Password) == 0 {
		res.WriteHeader(http.StatusInternalServerError)
		response.Error = "username and password are required"
		resData, _ := json.Marshal(response)
		res.Write(resData)
		return
	}

	err := authen.Signup(db, user)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		response.Error = err.Error()
		resData, _ := json.Marshal(response)
		res.Write(resData)
		return
	}

	response.Success = true
	resData, _ := json.Marshal(response)
	res.Write(resData)
}

func loginHandler(res http.ResponseWriter, req *http.Request, db *sql.DB) {
	username := req.FormValue("username")
	password := req.FormValue("password")
	err, user := authen.Authen(db, username, password)

	response := &Response{}
	if err != nil {
		res.WriteHeader(http.StatusUnauthorized)
		response.Error = "Invalid username or password"
	} else {
		response.User = &user
		fmt.Println(user)
		authen.SetCookie(res, user.Username)
	}
	fmt.Println(response.User)
	resData, _ := json.Marshal(response)
	fmt.Println(string(resData))
	res.Write(resData)
}

func logoutHandler(res http.ResponseWriter, req *http.Request) {
	authen.ClearCookie(res)
}

func dbHandler(fn func(http.ResponseWriter, *http.Request, *sql.DB), db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		fn(w, r, db)
	}
}

func postHandler(res http.ResponseWriter, req *http.Request) {

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
	//err := authen.Signup(db, "abc@gmail.com", "Icanfly")
	// if err != nil {
	// }

	// Test
	// fmt.Println(authen.GetUser(db, "ekkapob@gmail.com"))
	// fmt.Println(authen.GetUser(db, "test@gmail.com"))
	// fmt.Println(authen.Authen(db, "ekkapob@gmail.com", "thailand"))
	// fmt.Println(authen.Authen(db, "test@gmail.com", "hello"))

	r := mux.NewRouter()
	r.HandleFunc("/", indexHandler)
	r.HandleFunc("/signup", dbHandler(signupHandler, db)).Methods("POST")
	r.HandleFunc("/login", dbHandler(loginHandler, db)).Methods("POST")
	r.HandleFunc("/logout", logoutHandler).Methods("POST")

	r.HandleFunc("/post/{id}", postHandler)

	http.ListenAndServe(":8080", r)
}
