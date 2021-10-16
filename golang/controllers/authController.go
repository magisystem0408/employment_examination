package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"trunk_exem/employment_examination/golang/database"
	"trunk_exem/employment_examination/golang/models"
)

func CreateUser(c *gin.Context) {
	var reqBody map[string]string
	c.Bind(&reqBody)

	fmt.Println(reqBody)

	user := models.User{
		Id:    uuid.New().String(),
		Email: reqBody["email"],
	}
	user.SetPassword(reqBody["password"])
	database.DB = append(database.DB, user)

	c.JSON(200, gin.H{
		"message": database.DB,
	})
}

func LoginUser(c *gin.Context) {
	var reqBody map[string]string
	c.Bind(&reqBody)

	for i := range database.DB {
		if database.DB[i].Email == reqBody["email"] {

			if err := database.DB[i].ComparePassword(reqBody["password"]); err != nil {
				c.JSON(400, gin.H{
					"message": "miss",
				})
				return
			}

			c.JSON(200, gin.H{
				"message": "ok!",
			})
			return
		}
	}
	c.JSON(404, gin.H{
		"error":   true,
		"message": "invalid",
	})
}

func GetUser(c *gin.Context) {
	c.JSON(200, gin.H{
		"massage": database.DB,
	})

}
