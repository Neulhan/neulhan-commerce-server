package rest

import (
	"github.com/gin-gonic/gin"
)

func RunAPI(address string) error {
	h, err := NewHandler()
	if err != nil {
		return err
	}
	return RunAPIWithHandler(address, h)
}

func RunAPIWithHandler(address string, h HandlerInterface) error {
	r := gin.Default()

	h, _ = NewHandler()

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
