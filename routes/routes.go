package routes

import (
	product "StoreManager/controllers/products"
	sale "StoreManager/controllers/sales"
	"StoreManager/controllers/user"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func StartService() {
	router := gin.Default()
	api := router.Group("/api")
	{
		// user routes
		api.GET("/user", user.GetAllUsers)
		api.POST("/user", user.CreateUser)
		api.GET("/user/:id", user.GetUser)
		api.PUT("/user/:id", user.UpdateUser)
		api.DELETE("/user/:id", user.DeleteUser)

		// product routes
		api.GET("/product", product.GetAllProducts)
		api.POST("/product", product.CreateProduct)
		api.GET("/product/:id", product.GetProductById)
		api.PUT("/product/:id", product.UpdateProduct)
		api.DELETE("/product/:id", product.DeleteProduct)

		// sales routes
		api.GET("/sale", sale.GetAllSales)
		api.POST("/sale", sale.CreateSale)
		api.GET("/sale/:id", sale.GetSaleById)
		api.PUT("/sale/:id", sale.UpdateSale)
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
