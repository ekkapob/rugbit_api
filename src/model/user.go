package model

type User struct {
	Id        int    `json:"id"`
	Password  string `json:"-"`
	Username  string `json:"username"`
	Firstname string `json:"firstname,omitempty"`
	Lastname  string `json:"lastname,omitempty"`
}

type Post struct {
	Id          int      `json:"id"`
	Title       string   `json:"title"`
	Description string   `json:"description"`
	CreatedDate string   `json:"createdDate"`
	UpdatedDate string   `json:"updatedDate"`
	Skills      []string `json:"skills,omitempty"`
	Budget      string   `json:"budget,omitempty"`
	Owner       string   `json:"owner"`
}
