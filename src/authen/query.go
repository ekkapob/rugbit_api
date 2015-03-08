package authen

import (
	"database/sql"
	"model"
)

func GetUser(db *sql.DB, username string) (err error, user model.User) {
	err = db.QueryRow("select id, firstname, lastname from users where username = $1", username).Scan(&user.Id, &user.Firstname, &user.Lastname)
	if err == nil {
		user.Username = username
	}

	return err, user
}

func getHashPassword(db *sql.DB, username string) (err error, hashPassword string) {
	err = db.QueryRow("select password from users where username = $1", username).Scan(&hashPassword)
	return err, hashPassword
}

func addUser(db *sql.DB, user model.User) error {
	stmt, err := db.Prepare("insert into users (username, password, firstname, lastname) values ($1, $2, $3, $4)")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(user.Username, user.Password, user.Firstname, user.Lastname)
	return nil
}
