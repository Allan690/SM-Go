package routes

import (
	"StoreManager/controllers/user"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func StartService() {
	router := gin.Default()
	api := router.Group("/api")
	{
		api.GET("/user", user.GetAllUsers)
		api.POST("/user", user.CreateUser)
		api.GET("/user/:id", user.GetUser)
		api.PUT("/user/:id", user.UpdateUser)
		api.DELETE("/user/:id", user.DeleteUser)
	}
	router.NoRoute(func(c *gin.Context) {
		c.AbortWithStatus(http.StatusNotFound)
	})
	err := router.Run(":8000")
	if err != nil {
		fmt.Print(err)
		panic("An error occurred when running this application")
	}
}