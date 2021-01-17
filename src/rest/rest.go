package rest

import "github.com/gin-gonic/gin"

func RunAPI(address string) {
	r := gin.Default()

	r.GET("/products", func(c *gin.Context) {

	})

	r.GET("/promos", func(c *gin.Context) {

	})

	r.POST("/users/signin", func(context *gin.Context) {

	})

	r.POST("/users", func(context *gin.Context) {

	})

	r.POST("/user/:id/signout", func(context *gin.Context) {

	})

	r.GET("/user/:id/orders", func(context *gin.Context) {

	})

	r.POST("/users/charge", func(context *gin.Context) {

	})
}
