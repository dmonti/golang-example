package http

import (
	"github.com/gin-gonic/gin"
)

type HttpServer struct {
	engine *gin.Engine
}

func Run() HttpServer {
	engine := gin.New()
	engine.Use(gin.Logger())
	engine.ForwardedByClientIP = true
	engine.SetTrustedProxies([]string{"127.0.0.1"})

	for _, controller := range controllers {
		AddController(controller)
	}

	for _, route := range routes {
		engine.GET(route.relativePath, gin.HandlerFunc(route.handler))
	}

	engine.Run("localhost:8080")
	return HttpServer{engine}
}
