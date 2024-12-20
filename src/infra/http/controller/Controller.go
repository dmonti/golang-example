package controller

import "github.com/gin-gonic/gin"

type Controller struct {
	BasePath  string
	Endpoints []Endpoint
}

type Endpoint struct {
	RelativePath string
	Handler      func(*gin.Context)
}
