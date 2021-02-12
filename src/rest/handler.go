package rest

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"neulhan-commerce-server/src/dblayer"
	"neulhan-commerce-server/src/jwt"
	"neulhan-commerce-server/src/models"
	"strconv"
	"time"
)

type HandlerInterface interface {
	GetProducts(c *fiber.Ctx) error
	GetProduct(c *fiber.Ctx) error
	CreateProduct(c *fiber.Ctx) error
	UpdateProduct(c *fiber.Ctx) error
	DeleteProduct(c *fiber.Ctx) error
	GetPromos(c *fiber.Ctx) error
	KakaoLogin(c *fiber.Ctx) error
	GithubLogin(c *fiber.Ctx) error
	SignOut(c *fiber.Ctx) error
	QuitUser(c *fiber.Ctx) error
	GetUsers(c *fiber.Ctx) error
	GetUserInfo(c *fiber.Ctx) error
	GetOrders(c *fiber.Ctx) error
	Charge(c *fiber.Ctx) error
}

type Handler struct {
	db dblayer.DBLayer
}

func NewHandler(dbName string, conf *gorm.Config) (*Handler, error) {
	db, err := dblayer.NewORM(dbName, conf)
	if err != nil {
		return nil, err
	}
	return &Handler{
		db: db,
	}, nil
}

func (h *Handler) GetProducts(c *fiber.Ctx) error {
	if h.db == nil {
		return errors.New("DatabaseNotConnected")
	}

	products, err := h.db.GetProducts()

	if err != nil {
		return c.Status(400).JSON(err)
	}

	return c.JSON(products)
}

func (h *Handler) GetProduct(c *fiber.Ctx) error {
	if h.db == nil {
		return errors.New("DatabaseNotConnected")
	}
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return err
	}
	product, err := h.db.GetProduct(id)
	if err != nil {
		return err
	}
	return c.JSON(product)
}

func (h *Handler) CreateProduct(c *fiber.Ctx) error {
	if h.db == nil {
		return errors.New("DatabaseNotConnected")
	}

	var product models.Product

	if err := c.BodyParser(&product); err != nil {
		return err
	}

	if _, err := h.db.CreateProduct(product); err != nil {
		return err
	}

	return c.JSON(product)
}

func (h *Handler) UpdateProduct(c *fiber.Ctx) error {
	if h.db == nil {
		return errors.New("DatabaseNotConnected")
	}

	var product models.Product

	if err := c.BodyParser(&product); err != nil {
		return err
	}

	if _, err := h.db.UpdateProduct(product); err != nil {
		return err
	}

	return c.JSON(product)
}

func (h *Handler) DeleteProduct(c *fiber.Ctx) error {
	if h.db == nil {
		return errors.New("DatabaseNotConnected")
	}
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return err
	}

	err = h.db.DeleteProduct(id)
	if err != nil {
		return err
	}
	return c.JSON(fiber.Map{"status": "success"})
}

func (h *Handler) GetPromos(c *fiber.Ctx) error {
	if h.db == nil {
		return errors.New("DatabaseNotConnected")
	}

	promos, err := h.db.GetPromos()

	if err != nil {
		return err
	}

	return c.JSON(promos)
}

func (h *Handler) KakaoLogin(c *fiber.Ctx) error {
	if h.db == nil {
		return errors.New("DatabaseNotConnected")
	}
	var user models.User

	if err := c.BodyParser(&user); err != nil {
		return err
	}

	foundUser, err := h.db.GetUserBySocialID(user.SocialID)
	if err != nil {
		return err
	}

	var accessUser models.User
	if foundUser.ID == 0 {
		accessUser, err = h.db.AddUser(user)
		if err != nil {
			return err
		}
	} else {
		accessUser = foundUser
	}
	token, err := jwt.CreateToken(int(accessUser.ID))
	if err != nil {
		return err
	}

	cookie := new(fiber.Cookie)
	cookie.Name = "accessToken"
	cookie.Value = token
	cookie.Expires = time.Now().Add(24 * time.Hour)
	cookie.HTTPOnly = false
	c.Cookie(cookie)

	return c.JSON(fiber.Map{
		"result": "success",
	})
}

type GithubLoginData struct {
	Code string `json:"code"`
}

func (h *Handler) GithubLogin(c *fiber.Ctx) error {
	if h.db == nil {
		return errors.New("DatabaseNotConnected")
	}
	var data GithubLoginData
	//err := c.ReadJSON(&data)
	//if err != nil {
	//	return err
	//}
	//resp, err := http.Get("https://api.github.com/user")
	//if err != nil {
	//}
	return c.JSON(fiber.Map{"HELLO": data.Code})
}

func (h *Handler) SignOut(c *fiber.Ctx) error {
	if h.db == nil {
		return errors.New("DatabaseNotConnected")
	}
	p := c.Params("id")

	id, err := strconv.Atoi(p)
	if err != nil {
		return err
	}
	err = h.db.SignOutUserByID(id)
	if err != nil {
		return err
	}
	return nil
}

func (h *Handler) QuitUser(c *fiber.Ctx) error {
	if h.db == nil {
		return errors.New("DatabaseNotConnected")
	}
	userID := c.Locals("UserID").(int)

	if err := h.db.DeleteUserByID(userID); err != nil {
		return err
	}
	c.ClearCookie("accessToken")
	return c.JSON(userID)
}

func (h *Handler) GetUserInfo(ctx *fiber.Ctx) error {
	var user models.User
	userID := ctx.Locals("UserID").(int)

	if _, err := h.db.GetUserByID(userID); err != nil {
		return err
	}

	return ctx.JSON(user)
}

func (h *Handler) GetUsers(c *fiber.Ctx) error {
	if h.db == nil {
		return errors.New("DatabaseNotConnected")
	}
	users, err := h.db.GetUsers()
	if err != nil {
		return err
	}
	return c.JSON(users)
}

func (h *Handler) GetOrders(c *fiber.Ctx) error {
	if h.db == nil {
		return errors.New("DatabaseNotConnected")
	}
	p := c.Params("id")

	id, err := strconv.Atoi(p)
	if err != nil {
		return err
	}

	orders, err := h.db.GetUserOrdersByID(id)

	if err != nil {
		return err
	}

	return c.JSON(orders)
}

func (h *Handler) Charge(c *fiber.Ctx) error {
	if h.db == nil {
		return errors.New("DatabaseNotConnected")
	}
	return nil
}
