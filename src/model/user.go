package model

type User struct {
	Id        int    `json:"id"`
	Password  string `json:"-"`
	Username  string `json:"username"`
	Firstname string `json:"firstname,omitempty"`
	Lastname  string `json:"lastname,omitempty"`
}
