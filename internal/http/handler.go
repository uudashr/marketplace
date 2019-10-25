package http

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/shopspring/decimal"

	"github.com/uudashr/marketplace/internal/store"

	"github.com/uudashr/marketplace/internal/product"

	"github.com/uudashr/marketplace/internal/app"

	"github.com/labstack/echo"
)

// AppService represents application service.
//go:generate mockery -name=AppService
type AppService interface {
	RegisterNewCategory(app.RegisterNewCategoryCommand) (*product.Category, error)
	RetrieveCategories() ([]*product.Category, error)
	RetrieveCategoryByID(app.RetrieveCategoryByIDCommand) (*product.Category, error)
	RegisterNewStore(app.RegisterNewStoreCommand) (*store.Store, error)
	OfferNewProduct(app.OfferNewProductCommand) (*product.Product, error)
}

type delegate struct {
	appService AppService
}

func (d *delegate) registerNewCategory(c echo.Context) error {
	var p registerNewCategoryPayload
	if err := c.Bind(&p); err != nil {
		return err
	}

	cat, err := d.appService.RegisterNewCategory(app.RegisterNewCategoryCommand{
		Name: p.Name,
	})
	if err != nil {
		return err
	}

	c.Response().Header().Add("Location", fmt.Sprintf("/categories/%s", cat.ID()))
	return c.NoContent(http.StatusCreated)
}

func (d *delegate) retrieveCategories(c echo.Context) error {
	cats, err := d.appService.RetrieveCategories()
	if err != nil {
		return err
	}

	out := make([]categoryPayload, len(cats))
	for i, c := range cats {
		out[i] = categoryPayload{
			ID:   c.ID(),
			Name: c.Name(),
		}
	}
	return c.JSON(http.StatusOK, out)
}

func (d *delegate) retrieveCategoryByID(c echo.Context) error {
	paramID := c.Param("id")
	cat, err := d.appService.RetrieveCategoryByID(app.RetrieveCategoryByIDCommand{
		ID: paramID,
	})
	if err != nil {
		return err
	}

	if cat == nil {
		return echo.ErrNotFound
	}

	return c.JSON(http.StatusOK, categoryPayload{
		ID:   cat.ID(),
		Name: cat.Name(),
	})
}

func (d *delegate) registerNewStore(c echo.Context) error {
	var p registerNewStorePayload
	if err := c.Bind(&p); err != nil {
		return err
	}

	str, err := d.appService.RegisterNewStore(app.RegisterNewStoreCommand{
		Name: p.Name,
	})
	if err != nil {
		return err
	}

	c.Response().Header().Add("Location", fmt.Sprintf("/stores/%s", str.ID()))
	return c.NoContent(http.StatusCreated)
}

func (d *delegate) offerNewProduct(c echo.Context) error {
	paramID := c.Param("id")

	var p offerNewProductPayload
	if err := c.Bind(&p); err != nil {
		return err
	}

	prd, err := d.appService.OfferNewProduct(app.OfferNewProductCommand{
		StoreID:     paramID,
		CategoryID:  p.CategoryID,
		Name:        p.Name,
		Price:       p.Price,
		Description: p.Description,
		Quantity:    p.Quantity,
	})

	if err != nil {
		return err
	}

	c.Response().Header().Add("Location", fmt.Sprintf("/products/%s", prd.ID()))
	return c.NoContent(http.StatusCreated)
}

func (d *delegate) checkHealthz(c echo.Context) error {
	return c.NoContent(http.StatusOK)
}

// NewHandler constructs new http handler.
func NewHandler(appService AppService) (http.Handler, error) {
	if appService == nil {
		return nil, errors.New("nil appService")
	}

	e := echo.New()
	d := &delegate{
		appService: appService,
	}

	e.POST("/categories", d.registerNewCategory)
	e.GET("/categories", d.retrieveCategories)
	e.GET("/categories/:id", d.retrieveCategoryByID)

	e.POST("/stores", d.registerNewStore)

	e.POST("/stores/:id/products", d.offerNewProduct)

	e.GET("/healthz", d.checkHealthz)

	return e, nil
}

type registerNewCategoryPayload struct {
	Name string `json:"name"`
}

type categoryPayload struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type registerNewStorePayload struct {
	Name string `json:"name"`
}

type offerNewProductPayload struct {
	CategoryID  string          `json:"categoryId"`
	Name        string          `json:"name"`
	Price       decimal.Decimal `json:"price"`
	Description string          `json:"description"`
	Quantity    int             `json:"quantity"`
}
