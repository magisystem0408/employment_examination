package routes

import (
	"github.com/gin-gonic/gin"
	"trunk_exem/employment_examination/golang/controllers"
)

func Setup(app *gin.Engine)  {
	app.GET("/healthx",controllers.GetHealthz)

	userRoutes :=app.Group("/users")
	{
		userRoutes.GET("/",controllers.GetUser)
		userRoutes.POST("/register",controllers.CreateUser)
		userRoutes.POST("/login",controllers.LoginUser)
	}
}

