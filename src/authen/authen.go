package authen

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

func GetUser(username, password string) (userId int) {
	hashPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	fmt.Println(">>>", string(hashPassword), ">>>", len(string(hashPassword)))
	return userId
}

func AddUser(username, password string) {

}
