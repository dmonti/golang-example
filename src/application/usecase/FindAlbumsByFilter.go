package usecase

import (
	"example/src/application/repository"
	"example/src/domain"
)

type FindAlbumsByFilter struct {
	repository repository.AlbumRepository `required:"true"`
}

func (usecase FindAlbumsByFilter) Execute(input FindAlbumsByFilterInput) FindAlbumsByFilterOutput {
	albums := usecase.repository.FindAllByFilter(input.Filter)
	return FindAlbumsByFilterOutput{albums}
}

type FindAlbumsByFilterInput struct {
	Filter map[string][]string
}

type FindAlbumsByFilterOutput struct {
	Albums []domain.Album
}

func NewFindAlbumsByFilter(repository repository.AlbumRepository) *FindAlbumsByFilter {
	usecase := new(FindAlbumsByFilter)
	usecase.repository = repository
	return usecase
}
