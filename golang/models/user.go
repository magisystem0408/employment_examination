package models

import (
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Id       string `json:"id"`
	Email    string `json:"email"`
	Password []byte `json:"password"`
}

func (user *User) SetPassword(password string) {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), 14)
	user.Password = hashedPassword
}

func (user *User) ComparePassword(password string) error {
	return bcrypt.CompareHashAndPassword(user.Password, []byte(password))

}
