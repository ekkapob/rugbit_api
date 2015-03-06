package authen

import (
	"database/sql"
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

// func Login(db *sql.DB, username, password string) (userId int) {
// 	var (
// 		id      int
// 		hashPwd string
// 	)
// 	_ = db.QueryRow("select id, password from users where username = $1", username).Scan(&id, &hashPwd)
// 	if id != 0 {
// 		err := bcrypt.CompareHashAndPassword([]byte(hashPwd), []byte(password))
// 		if err == nil {
// 			userId = id
// 		}
// 	}
// 	return userId
// }

func Authen(db *sql.DB, username, password string) (err error, user User) {

	err, hashPassword := getHashPassword(db, username)
	if err == nil {
		err = bcrypt.CompareHashAndPassword([]byte(hashPassword), []byte(password))
		if err == nil {
			err, user = GetUser(db, username)
		}
	}
	return err, user
}

func Signup(db *sql.DB, username, password string) error {
	fmt.Println("signup")
	var id int
	err := db.QueryRow("select id from users where username = $1", username).Scan(&id)
	// _, err := db.Query("select id from users where username = $1", username)
	fmt.Println("err", err)

	if err != sql.ErrNoRows {
		fmt.Println("err2", err)
		return err
	}

	fmt.Println("success")
	return nil
	// Update
	// _, err = db.Exec("insert into users (username, password) values (?,?)", username, password)
	// return err
}

func getUserId(db *sql.DB, username string) int {
	var id int
	err := db.QueryRow("select id from users where username = $1", username).Scan(&id)
	if err != sql.ErrNoRows {
		return id
	}
	return 0
}

// func GetUser(username, password string) (userId int) {
// 	hashPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

// 	fmt.Println(">>>", string(hashPassword), ">>>", len(string(hashPassword)))
// 	return userId
// }

// func AddUser(username, password string) {

// }
