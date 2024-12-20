package usecase

import (
	"example/src/application/repository"
	"example/src/domain"
)

type FindAlbumById struct {
	repository repository.AlbumRepository `required:"true"`
}

func (usecase FindAlbumById) Execute(input FindAlbumByIdInput) FindAlbumByIdOutput {
	album := usecase.repository.FindById(input.Id)
	return FindAlbumByIdOutput{album}
}

type FindAlbumByIdInput struct {
	Id string
}

type FindAlbumByIdOutput struct {
	Album *domain.Album
}

func NewFindAlbumById(repository repository.AlbumRepository) *FindAlbumById {
	usecase := new(FindAlbumById)
	usecase.repository = repository
	return usecase
}
