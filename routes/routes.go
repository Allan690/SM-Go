package routes

import (
	"StoreManager/controllers/auth"
	product "StoreManager/controllers/products"
	sale "StoreManager/controllers/sales"
	"StoreManager/controllers/user"
	middleware "StoreManager/middleware/auth"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func StartService() {
	router := gin.Default()
	api := router.Group("/api")
	{
		// user routes
		api.GET("/user", middleware.AuthorizeJWT(), user.GetAllUsers)
		api.POST("/user", user.CreateUser)
		api.GET("/user/:id", middleware.AuthorizeJWT(), user.GetUser)
		api.PUT("/user/:id", middleware.AuthorizeJWT(),  user.UpdateUser)
		api.DELETE("/user/:id", middleware.AuthorizeJWT(), user.DeleteUser)

		// product routes
		api.GET("/product",middleware.AuthorizeJWT(), product.GetAllProducts)
		api.POST("/product", middleware.AuthorizeJWT(), product.CreateProduct)
		api.GET("/product/:id", middleware.AuthorizeJWT(), product.GetProductById)
		api.PUT("/product/:id", middleware.AuthorizeJWT(), product.UpdateProduct)
		api.DELETE("/product/:id", middleware.AuthorizeJWT(), product.DeleteProduct)

		// sales routes
		api.GET("/sale", middleware.AuthorizeJWT(),sale.GetAllSales)
		api.POST("/sale", middleware.AuthorizeJWT(), sale.CreateSale)
		api.GET("/sale/:id", middleware.AuthorizeJWT(), sale.GetSaleById)
		api.PUT("/sale/:id", middleware.AuthorizeJWT(), sale.UpdateSale)

		// auth routes
		api.POST("/login", auth.LoginController)
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
