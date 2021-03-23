package main

import (
	"github.com/gin-gonic/gin"
	"ofcode.dev/labala-backend/config"
	"ofcode.dev/labala-backend/engine"
)

func main() {
	db := config.DBInit()
	InDB := &engine.InDB{DB: db}
	router := gin.Default()

	auth := router.Group("/auth")
	{
		auth.POST("/login", InDB.Login)
		auth.POST("/register", InDB.CreateUser)
	}

	product := router.Group("/product")
	{
		product.GET("/", InDB.AvailableSoon)
		product.GET("/:q1", InDB.GetProductUsingID)
		product.POST("/", InDB.CreateProduct)
		product.PUT("/:q1", InDB.UpdateProduct)
		product.DELETE("/:q1", InDB.DeleteProduct)
	}

	feed := router.Group("/feed")
	{
		feed.GET("/dashboard", InDB.AvailableSoon)
		feed.GET("/search/:q1", InDB.AvailableSoon)
	}

	user := router.Group("/user")
	{
		user.GET("/profile", InDB.AvailableSoon)
	}

	err := router.Run(":9099")
	if err != nil {
		panic(err)
	}
}
