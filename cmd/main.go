package main

import (
	"github.com/Paulo-Lopes-Estevao/webserver-example-go/http"
	"github.com/Paulo-Lopes-Estevao/webserver-example-go/model"
)

func main() {
	webserver:=http.NewWebServer()
	webserver.Products = &model.Products{}
	webserver.Serve()
}
