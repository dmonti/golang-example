package database

import (
	"example/src/infra/database"
	"testing"
)

var repository = database.NewAlbumRepository()

func TestAlbumInMemoryRepository_FindById(t *testing.T) {
	id2 := "2"
	album1 := repository.FindById(id2)
	if album1 == nil {
		t.Errorf("Expected album with #%v to be found", id2)
	} else if album1.Id != id2 {
		t.Errorf("Expected album with #%v to be found, but got #%v", id2, album1.Id)
	}

	albumNil := repository.FindById("nil")
	if albumNil != nil {
		t.Errorf("Expected nil album to be found")
	}
}

func TestAlbumInMemoryRepository_FindByFilter(t *testing.T) {
	filter := map[string][]string{"title": {"train"}}
	albums := repository.FindAllByFilter(filter)
	if len(albums) != 1 {
		t.Errorf("Expected 1 album to be found, but got %v", len(albums))
	} else if albums[0].Title != "Blue Train" {
		t.Errorf("Expected album with title 'Blue Train' to be found, but got %v", albums[0].Title)
	}
}
