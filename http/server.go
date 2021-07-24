package http

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"github.com/Paulo-Lopes-Estevao/webserver-example-go/webserver/model"
)

type WebServer struct {
	Products *model.Product
}

func NewWebServer() *WebServer  {
	return &WebServer{}
}

func (w WebServer)Serve()  {
	e:= echo.New()
	e.GET("/products",w.getAll)
	e.POST("/products",w.createProduct)
	e.Start(":8080")
}

func (w WebServer) getAll(e echo.Context) error  {
	return e.JSON(http.StatusOK,w.Products)
}

func (w WebServer)createProduct(c echo.Context) error  {
	product := model.NewProduct()
	if err := c.Bind(product); err != nil {
		return err
	}
	w.Products.Add(product)
	return e.JSON(http.StatusCreated,product)
}