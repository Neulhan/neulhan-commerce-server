package rest

import (
	"github.com/kataras/iris/v12"
	"gorm.io/gorm"
	"neulhan-commerce-server/src/dblayer"
	"neulhan-commerce-server/src/models"
	"strconv"
)

type HandlerInterface interface {
	GetProducts(c iris.Context)
	GetPromos(c iris.Context)
	AddUser(c iris.Context)
	SignIn(c iris.Context)
	SignOut(c iris.Context)
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

func (h *Handler) SignIn(c iris.Context) {
	if h.db == nil {
		return
	}

	var customer models.Customer

	err := c.ReadJSON(&customer)
	if err != nil {
		c.StopWithError(iris.StatusBadRequest, err)
		return
	}

	customer, err = h.db.SignInUser(customer.Email, customer.Pass)
	if err != nil {
		c.StopWithError(iris.StatusInternalServerError, err)
		return
	}
	c.JSON(customer)
}

func (h *Handler) AddUser(c iris.Context) {
	if h.db == nil {
		return
	}

	var customer models.Customer

	err := c.ReadJSON(&customer)
	if err != nil {
		c.StopWithError(iris.StatusBadRequest, err)
		return
	}

	customer, err = h.db.AddUser(customer)
	if err != nil {
		c.StopWithError(iris.StatusInternalServerError, err)
		return
	}
	c.JSON(customer)
}

func (h *Handler) SignOut(c iris.Context) {
	if h.db == nil {
		return
	}
	p := c.Params().Get("id")

	id, err := strconv.Atoi(p)
	if err != nil {
		c.StopWithError(iris.StatusBadRequest, err)
		return
	}
	err = h.db.SignOutUserByID(id)
	if err != nil {
		c.StopWithError(iris.StatusInternalServerError, err)
		return
	}
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

	orders, err := h.db.GetCustomerOrdersByID(id)

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
