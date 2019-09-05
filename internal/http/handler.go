package http

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/uudashr/marketplace/internal/category"

	"github.com/uudashr/marketplace/internal/app"

	"github.com/labstack/echo"
)

// AppService represents application service.
//go:generate mockery -name=AppService
type AppService interface {
	RegisterNewCategory(app.RegisterNewCategoryCommand) (*category.Category, error)
	RetrieveCategories() ([]*category.Category, error)
	RetrieveCategoryByID(app.RetrieveCategoryByIDCommand) (*category.Category, error)
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

	return c.JSON(http.StatusOK, categoryPayload{
		ID:   cat.ID(),
		Name: cat.Name(),
	})
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

	return e, nil
}

type registerNewCategoryPayload struct {
	Name string `json:"name"`
}

type categoryPayload struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}
