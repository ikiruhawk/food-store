package routers

import (
	"html/template"
	"io"
	"net/http"
	"strconv"

	"github.com/ikiruhawk/food-store/internal/crud"
	"github.com/ikiruhawk/food-store/internal/models"
	"github.com/labstack/echo/v4"
)

type templateRenderer struct {
	templates *template.Template
}

var tpl *templateRenderer

func (t *templateRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	if viewContext, isMap := data.(map[string]interface{}); isMap {
		viewContext["reverse"] = c.Echo().Reverse
	}

	return t.templates.ExecuteTemplate(w, name, data)
}

func RunRouters() *echo.Echo {
	e := echo.New()
	tpl = &templateRenderer{
		templates: template.Must(template.ParseGlob("views/html/*.html")),
	}
	e.Renderer = tpl
	e.GET("/", homeHandleFunc)
	e.GET("/products", productsHandleFunc)
	e.GET("/product/:id", productInfoHandleFunc)
	return e
}

func homeHandleFunc(c echo.Context) error {
	return c.Render(http.StatusOK, "index.html", nil)
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
		Products  []models.Product
		Categories []models.Category
		Manufacturers []models.Manufacturer
	}{
		Products: prs,
		Categories: cgs,
		Manufacturers: mfs,
	}

	return c.Render(http.StatusOK, "products.html", data)
}

func productInfoHandleFunc(c echo.Context) error {
	pId := c.Param("id")
	id, err := strconv.Atoi(pId)
	if err != nil {
		return err
	}

	p := *crud.GetProductById(id)

	return c.Render(http.StatusOK, "product.html", p)
}
