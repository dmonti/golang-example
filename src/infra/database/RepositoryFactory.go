package database

import "example/src/domain"

func NewAlbumRepository() *AlbumInMemoryRepository {
	repository := new(AlbumInMemoryRepository)
	repository.albums = []domain.Album{
		{Id: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
		{Id: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
		{Id: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
	}
	return repository
}
