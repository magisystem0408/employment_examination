package database

import "fmt"

type User struct {
	Email    string
	Password string
}

var store = make(map[string]string)

type DB interface {
	Get(key string) (value string, err error)
	Set(key, value string) error
}

func (u User) Set(key, value string) error {
	store[key] = value
	fmt.Println(value)
	return nil
}

func (u User) Get(key string) (value string, err error) {
	value = store[key]
	fmt.Println(value)
	return value, nil
}

type AuthController struct {
	db DB
}
