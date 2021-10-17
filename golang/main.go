package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"log"
)

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

func main() {

	app := gin.Default()
	app.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "ok",
		})
	})
	userRoutes := app.Group("/users")
	{
		userRoutes.POST("/registor", func(c *gin.Context) {
			var reqBody map[string]string
			c.Bind(&reqBody)
			user := User{
				Email: reqBody["email"],
			}
			hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(reqBody["password"]), 14)
			user.Password = string(hashedPassword)

			if nil != user.Set(user.Email, user.Password) {
				fmt.Errorf("データベースが登録できませんでした。")
			}

			c.JSON(200, gin.H{
				"message": "登録が完了しました",
			})
		})
		userRoutes.POST("/login", func(c *gin.Context) {
			var reqBody map[string]string
			c.Bind(&reqBody)
			//ハッシュ化されたパスワードが入ってる
			user := User{Email: reqBody["email"], Password: reqBody["password"]}

			password, _ := user.Get(user.Email)

			err := bcrypt.CompareHashAndPassword([]byte(password), []byte(user.Password))
			if err != nil {
				c.JSON(400, gin.H{
					"message": "正しくありません",
				})
				return
			}

			c.JSON(200, gin.H{
				"message": "ok!",
			})
		})
	}
	if err := app.Run(":5000"); err != nil {
		log.Fatalln(err.Error())
	}
}
