package database

import (
	"example/src/domain"
	"fmt"
	"slices"
	"strings"
)

type AlbumInMemoryRepository struct {
	albums []domain.Album
}

func (repository AlbumInMemoryRepository) FindById(id string) *domain.Album {
	albums := repository.FindAllByFilter(map[string][]string{"id": {id}})
	if len(albums) > 0 {
		return &albums[0]
	}
	return nil
}

func (repository AlbumInMemoryRepository) FindAllByFilter(filter map[string][]string) []domain.Album {
	albums := append([]domain.Album{}, repository.albums...)

	id := filter["id"]
	if len(id) > 0 {
		for _, album := range albums {
			fmt.Printf("album.Id %v contains? %v \n", album.Id, slices.Contains(id, album.Id))
			if !slices.Contains(id, album.Id) {
				albums = remove(albums, album)
			}
		}
	}

	title := filter["title"]
	if len(title) > 0 {
		for _, album := range albums {
			if !strings.Contains(strings.ToLower(album.Title), strings.ToLower(title[0])) {
				albums = remove(albums, album)
			}
		}
	}

	return albums
}

// TODO Improve array removal
func remove(albums []domain.Album, toRemove domain.Album) []domain.Album {
	result := []domain.Album{}
	for _, a := range albums {
		if !a.Equals(toRemove) {
			result = append(result, a)
		}
	}
	return result
}
