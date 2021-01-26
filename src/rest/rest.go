package rest

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/recover"
	"gorm.io/gorm"
	"log"
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
	app := iris.Default()
	app.UseRouter(recover.New())
	app.Use(middleware.CustomMiddleWare())

	app.Get("/", func(c iris.Context) {
		c.JSON(iris.Map{"server": "ON AIR!"})
	})

	app.Get("/products", h.GetProducts)
	app.Get("/promos", h.GetPromos)
	//
	usersAPI := app.Party("/users")
	{
		usersAPI.Post("/users/signin", h.SignIn)
		usersAPI.Post("/users", h.AddUser)
		usersAPI.Post("/users/charge", h.Charge)
	}
	//
	userAPI := app.Party("/user")
	{
		userAPI.Get("/user/:id/orders", h.GetOrders)
		userAPI.Post("/user/:id/signout", h.SignOut)
	}

	return app.Listen(address)
}
