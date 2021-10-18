package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"trunk_exem/employment_examination/golang/database"
)

func CreateUser(c *gin.Context) {
	var reqBody map[string]string
	c.Bind(&reqBody)
	user := database.User{Email: reqBody["email"]}
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(reqBody["password"]), 14)
	user.Password = string(hashedPassword)
	if nil != user.Set(user.Email, user.Password) {
		fmt.Errorf("データベースに登録できませんでした。")
	}

	c.JSON(200, gin.H{
		"message": "登録が完了しました。",
	})

}

func LoginUser(c *gin.Context) {
	var reqBody map[string]string
	c.Bind(&reqBody)
	user := database.User{Email: reqBody["email"], Password: reqBody["password"]}
	password, _ := user.Get(user.Email)
	err := bcrypt.CompareHashAndPassword([]byte(password), []byte(user.Password))
	if err != nil {
		c.JSON(400, gin.H{
			"message": "正しくありません。",
		})
		return
	}
	c.JSON(200, gin.H{
		"message": "ok!",
	})

}
