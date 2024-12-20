package repository

import "example/src/domain"

type AlbumRepository interface {
	FindById(id string) *domain.Album
	FindAllByFilter(filter map[string][]string) []domain.Album
}
