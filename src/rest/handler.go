package rest

import (
	"github.com/kataras/iris/v12"
	"gorm.io/gorm"
	"neulhan-commerce-server/src/dblayer"
	"neulhan-commerce-server/src/jwt"
	"neulhan-commerce-server/src/models"
	"strconv"
)

type HandlerInterface interface {
	GetProducts(c iris.Context)
	GetProduct(c iris.Context)
	CreateProduct(c iris.Context)
	UpdateProduct(c iris.Context)
	DeleteProduct(c iris.Context)
	GetPromos(c iris.Context)
	KakaoLogin(c iris.Context)
	GithubLogin(c iris.Context)
	SignOut(c iris.Context)
	QuitUser(c iris.Context)
	GetUsers(c iris.Context)
	GetUserInfo(c iris.Context)
	GetOrders(c iris.Context)
	Charge(c iris.Context)
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

func (h *Handler) GetProducts(c iris.Context) {
	if h.db == nil {
		return
	}

	products, err := h.db.GetProducts()

	if err != nil {
		c.StopWithError(iris.StatusInternalServerError, err)
	}

	c.JSON(products)
}

func (h *Handler) GetProduct(c iris.Context) {
	if h.db == nil {
		return
	}
	id, err := c.Params().GetInt("id")
	if err != nil {
		c.StopWithError(iris.StatusBadRequest, err)
		return
	}
	product, err := h.db.GetProduct(id)
	if err != nil {
		c.StopWithError(iris.StatusInternalServerError, err)
		return
	}
	c.JSON(product)
}

func (h *Handler) CreateProduct(c iris.Context) {
	if h.db == nil {
		return
	}

	var product models.Product
	err := c.ReadJSON(&product)
	if err != nil {
		c.StopWithError(iris.StatusBadRequest, err)
		return
	}

	product, err = h.db.CreateProduct(product)
	if err != nil {
		c.StopWithError(iris.StatusInternalServerError, err)
	}

	c.JSON(product)
}

func (h *Handler) UpdateProduct(c iris.Context) {
	if h.db == nil {
		return
	}

	var product models.Product
	err := c.ReadJSON(&product)
	if err != nil {
		c.StopWithError(iris.StatusBadRequest, err)
		return
	}

	product, err = h.db.UpdateProduct(product)
	if err != nil {
		c.StopWithError(iris.StatusInternalServerError, err)
	}

	c.JSON(product)
}

func (h *Handler) DeleteProduct(c iris.Context) {
	if h.db == nil {
		return
	}
	id, err := c.Params().GetInt("id")
	if err != nil {
		c.StopWithError(iris.StatusBadRequest, err)
		return
	}

	err = h.db.DeleteProduct(id)
	if err != nil {
		c.StopWithError(iris.StatusInternalServerError, err)
		return
	}
	c.JSON(iris.Map{"status": "success"})
}

func (h *Handler) GetPromos(c iris.Context) {
	if h.db == nil {
		return
	}

	promos, err := h.db.GetPromos()

	if err != nil {
		c.StopWithError(iris.StatusInternalServerError, err)
		return
	}

	c.JSON(promos)
}

func (h *Handler) KakaoLogin(c iris.Context) {
	if h.db == nil {
		return
	}

	err := auth.KakaoAuth(c)
	if err != nil {
		c.StopWithError(iris.StatusBadRequest, err)
	}
	c.JSON(iris.Map{"HELLO": "WORLD"})
}

func (h *Handler) GithubLogin(c iris.Context) {
	if h.db == nil {
		return
	}

}

func (h *Handler) AddUser(c iris.Context) {
	if h.db == nil {
		return
	}

	var user models.User

	err := c.ReadJSON(&user)
	if err != nil {
		c.StopWithError(iris.StatusBadRequest, err)
		return
	}

	user, err = h.db.AddUser(user)
	if err != nil {
		c.StopWithError(iris.StatusInternalServerError, err)
		return
	}
	c.JSON(user)
}

func (h *Handler) QuitUser(c iris.Context) {
	if h.db == nil {
		return
	}
	userID, err := c.Values().GetInt("UserID")
	if err != nil {
		c.StopWithError(iris.StatusInternalServerError, err)
	}
	err = h.db.DeleteUserByID(userID)
	if err != nil {
		c.StopWithError(iris.StatusInternalServerError, err)
	}
	c.RemoveCookie("accessToken")
	c.JSON(userID)
}

func (h *Handler) GetUserInfo(ctx iris.Context) {
	var user models.User
	userID, err := ctx.Values().GetInt("UserID")
	if err != nil {
		ctx.StopWithError(iris.StatusBadRequest, err)
	}
	user, err = h.db.GetUserByID(userID)
	ctx.JSON(user)
}

func (h *Handler) GetUsers(c iris.Context) {
	if h.db == nil {
		return
	}
	users, err := h.db.GetUsers()
	if err != nil {
		c.StopWithError(iris.StatusInternalServerError, err)
	}
	c.JSON(users)
}

func (h *Handler) GetOrders(c iris.Context) {
	if h.db == nil {
		return
	}
	p := c.Params().Get("id")

	id, err := strconv.Atoi(p)
	if err != nil {
		c.StopWithError(iris.StatusBadRequest, err)
		return
	}

	orders, err := h.db.GetUserOrdersByID(id)

	if err != nil {
		c.StopWithError(iris.StatusInternalServerError, err)
		return
	}

	c.JSON(orders)
}

func (h *Handler) Charge(c iris.Context) {
	if h.db == nil {
		return
	}
}
