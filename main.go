package main

import (
	"jwt/controllers"
	"jwt/initializers"
	"jwt/middleware"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvironmentVariables()
	initializers.ConnectToDb()
	initializers.SyncDatabase()
}

func main() {
	r := gin.Default()
	v1 := r.Group("v1")

	v1.POST("/signup", controllers.SignUp)
	v1.POST("/login", controllers.Login)
	v1.GET("/validate", middleware.RequiredAuth, controllers.Validate)

	r.Run()
}
