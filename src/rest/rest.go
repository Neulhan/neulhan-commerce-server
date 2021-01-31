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
	app.Use(middleware.Logger())
	app.UseRouter(middleware.Cors)

	app.Get("/", func(c iris.Context) {
		c.JSON(iris.Map{"server": "ON AIR!"})
	})

	productAPI := app.Party("/products")
	{
		productAPI.Get("/", h.GetProducts)
		productAPI.Post("/", h.CreateProduct)
		productAPI.Delete("/", h.DeleteProduct)
		productAPI.Get("/{id:int}", h.GetProduct)
		productAPI.Get("/promos", h.GetPromos)
		productAPI.Post("/update", h.UpdateProduct)
	}

	authAPI := app.Party("/auth")
	{
		authAPI.Post("/kakao", h.KakaoLogin)
		authAPI.Post("/github", h.GithubLogin)
	}

	usersAPI := app.Party("/users")
	{
		usersAPI.Post("/", h.AddUser)
		usersAPI.Post("/charge", h.Charge)
	}
	userAPI := app.Party("/user")
	{
		userAPI.Get("/:id/orders", h.GetOrders)
		userAPI.Post("/:id/signout", h.SignOut)
	}

	return app.Listen(address)
}
