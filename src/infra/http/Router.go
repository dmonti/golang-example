package http

import (
	"example/src/infra/http/controller"

	"github.com/gin-gonic/gin"
)

type HandlerFunc func(*gin.Context)

type Route struct {
	relativePath string
	handler      HandlerFunc
}

var routes = []Route{}

func AddRoute(path string, handler HandlerFunc) {
	routes = append(routes, Route{path, handler})
}

var controllers = []controller.Controller{
	controller.NewAlbumController(),
}

func AddController(controller controller.Controller) {
	for _, endpoint := range controller.Endpoints {
		AddRoute(controller.BasePath+endpoint.RelativePath, endpoint.Handler)
	}
}
