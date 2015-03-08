package authen

import (
	"database/sql"
	"errors"
	"golang.org/x/crypto/bcrypt"
	"model"
)

func Authen(db *sql.DB, username, password string) (err error, user model.User) {

	err, hashPassword := getHashPassword(db, username)
	if err == nil {
		err = bcrypt.CompareHashAndPassword([]byte(hashPassword), []byte(password))
		if err == nil {
			err, user = GetUser(db, username)
		}
	}
	return err, user
}

func Signup(db *sql.DB, user model.User) error {
	if userExist(db, user.Username) {
		return errors.New("user exists")
	}

	hashPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	user.Password = string(hashPassword)
	err := addUser(db, user)
	return err
}

func userExist(db *sql.DB, username string) bool {
	var id int
	err := db.QueryRow("select id from users where username = $1", username).Scan(&id)
	return err != sql.ErrNoRows
}
