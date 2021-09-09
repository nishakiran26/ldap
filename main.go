package main

import (
	"fmt"
	auth "ldapbackend/controller"
	"ldapbackend/database"
	models "ldapbackend/model"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {

	database.CreateConnection()
	fmt.Println("Gin router connection")
	router := gin.Default()
	router.Use(cors.Default())
	//Testing server
	router.GET("/api", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Server is Running",
		})
	})
	//Login page route
	router.POST("/login", func(c *gin.Context) {
		var loginDetails models.Login

		c.BindJSON(&loginDetails)

		status := auth.Auth(loginDetails, c)
		if status {
			c.JSON(200, gin.H{
				"status": "Successfully logged in",
			})
		}

	})
	defer database.CloseConnection()
	router.Run(":8082")
}
