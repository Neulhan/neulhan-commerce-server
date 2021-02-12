package rest

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
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
	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowOrigins:     "http://localhost:3000, https://devgood.io",
		AllowCredentials: true,
	}))
	app.Use(recover.New())
	app.Use(logger.New())
	app.Use(middleware.NewUserMiddleware())

	//app.Use(cors.New())

	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"server": "ON AIR!"})
	})

	productAPI := app.Group("/products")
	{
		productAPI.Get("/", h.GetProducts)
		productAPI.Post("/", h.CreateProduct)
		productAPI.Get("/:id", h.GetProduct)
		productAPI.Delete("/:id", h.DeleteProduct)
		productAPI.Get("/promos", h.GetPromos)
		productAPI.Post("/update", h.UpdateProduct)
	}

	authAPI := app.Group("/auth")
	{
		authAPI.Post("/kakao", h.KakaoLogin)
		authAPI.Post("/github", h.GithubLogin)
	}

	usersAPI := app.Group("/users")
	{
		usersAPI.Get("/", h.GetUsers)
		usersAPI.Post("/charge", h.Charge)
	}
	userAPI := app.Group("/user")
	{
		userAPI.Get("/", h.GetUserInfo)
		userAPI.Delete("/quit", h.QuitUser)
		userAPI.Get("/:id/orders", h.GetOrders)
		userAPI.Post("/:id/signout", h.SignOut)
	}

	return app.Listen(address)
}
