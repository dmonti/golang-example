package usecase

import (
	"example/src/application/usecase"
	"example/src/infra/database"
	"testing"
)

func TestAlbumInMemoryRepository_FindByFilter(t *testing.T) {
	repository := database.NewAlbumRepository()

	input := usecase.FindAlbumsByFilterInput{Filter: map[string][]string{"title": {"vaughan"}}}
	output := usecase.NewFindAlbumsByFilter(repository).Execute(input)
	if len(output.Albums) != 1 {
		t.Errorf("Expected 1 album, but got %v", len(output.Albums))
	} else if output.Albums[0].Title != "Sarah Vaughan and Clifford Brown" {
		t.Errorf("Expected album with title 'Sarah Vaughan and Clifford Brown', but got %v", output.Albums[0].Title)
	}
}

func TestAlbumInMemoryRepository_FindByFilter_empty(t *testing.T) {
	repository := database.NewAlbumRepository()

	input := usecase.FindAlbumsByFilterInput{Filter: map[string][]string{"title": {"invalid title"}}}
	output := usecase.NewFindAlbumsByFilter(repository).Execute(input)
	if len(output.Albums) != 0 {
		t.Errorf("Expected 0 albums, but got %v", len(output.Albums))
	}
}
