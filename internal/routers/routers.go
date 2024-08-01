package routers

import (
	"net/http"
	"strconv"

	"github.com/ikiruhawk/food-store/internal/crud"
	"github.com/ikiruhawk/food-store/internal/models"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func RunRouters() *echo.Echo {
	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:3000"},
		AllowMethods: []string{echo.GET, echo.PUT, echo.POST, echo.DELETE},
	}))
	e.GET("/", homeHandleFunc)
	e.GET("/products", productsHandleFunc)
	e.GET("/product/:id", productInfoHandleFunc)
	return e
}

func homeHandleFunc(c echo.Context) error {
	return c.JSON(http.StatusOK, nil)
}

func productsHandleFunc(c echo.Context) error {
	prs, err := crud.GetProducts()
	if err != nil {
		return err
	}
	cgs, err := crud.GetCategories()
	if err != nil {
		return err
	}
	mfs, err := crud.GetManufacturers()
	if err != nil {
		return err
	}

	data := struct {
		Products      []models.Product      `json:"products"`
		Categories    []models.Category     `json:"categories"`
		Manufacturers []models.Manufacturer `json:"manufacturers"`
	}{
		Products:      prs,
		Categories:    cgs,
		Manufacturers: mfs,
	}

	return c.JSON(http.StatusOK, data)
}

func productInfoHandleFunc(c echo.Context) error {
	pId := c.Param("id")
	id, err := strconv.Atoi(pId)
	if err != nil {
		return err
	}

	p := *crud.GetProductById(id)

	return c.JSON(http.StatusOK, p)
}
