package main

import "github.com/Paulo-Lopes-Estevao/webserver-example-go/http"

func main() {
	webserver:=http.NewWebServer()
	webserver.Serve()
}
