package rest

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"log"
	"net/http"
	"neulhan-commerce-server/src/config"
	"neulhan-commerce-server/src/middleware"
)

func RunAPI(address string) error {
	log.Println("RUN API...")
	h, err := NewHandler(config.DSN, &gorm.Config{})
	if err != nil {
		return err
	}
	return RunAPIWithHandler(address, h)
}

func RunAPIWithHandler(address string, h HandlerInterface) error {
	r := gin.Default()
	r.Use(middleware.CustomMiddleWare())

	r.GET("", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"server": "ON AIR!"})
	})

	r.GET("/products", h.GetProducts)
	r.GET("/promos", h.GetPromos)

	r.Group("/users")
	{
		r.POST("/users/signin", h.SignIn)
		r.POST("/users", h.AddUser)
		r.POST("/users/charge", h.Charge)
	}

	r.Group("/user")
	{
		r.GET("/user/:id/orders", h.GetOrders)
		r.POST("/user/:id/signout", h.SignOut)
	}

	return r.Run(address)
}
