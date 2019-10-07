package main

import "github.com/gin-gonic/gin"

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("views/*.html")
	router.Static("/assets", "./assets")

	user := router.Group("/user")
	{
		user.POST("/signup", router.UserSignUp)
		user.POST("/login", router.User.login)
	}

	router.GET("/")
}
