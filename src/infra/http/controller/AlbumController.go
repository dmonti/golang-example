package controller

import (
	"example/src/application/usecase"
	"example/src/infra/database"
	"net/http"

	"github.com/gin-gonic/gin"
)

var repository = database.NewAlbumRepository()

func NewAlbumController() Controller {
	controller := Controller{}
	controller.BasePath = "/albums"
	controller.Endpoints = []Endpoint{
		NewAlbumIdEndpoint(),
		NewAlbumFilterEndpoint(),
	}
	return controller
}

func NewAlbumIdEndpoint() Endpoint {
	endpoint := Endpoint{}
	endpoint.RelativePath = "/:id"
	endpoint.Handler = func(c *gin.Context) {
		input := usecase.FindAlbumByIdInput{Id: c.Param("id")}
		output := usecase.NewFindAlbumById(repository).Execute(input)
		c.IndentedJSON(http.StatusOK, output.Album)
	}
	return endpoint
}

func NewAlbumFilterEndpoint() Endpoint {
	endpoint := Endpoint{}
	endpoint.RelativePath = ""
	endpoint.Handler = func(c *gin.Context) {
		input := usecase.FindAlbumsByFilterInput{Filter: c.Request.URL.Query()}
		output := usecase.NewFindAlbumsByFilter(repository).Execute(input)
		c.IndentedJSON(http.StatusOK, output.Albums)
	}
	return endpoint
}
