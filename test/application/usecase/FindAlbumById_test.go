package usecase

import (
	"example/src/application/usecase"
	"example/src/infra/database"
	"testing"
)

func TestFindAlbumById(t *testing.T) {
	repository := database.NewAlbumRepository()

	input := usecase.FindAlbumByIdInput{Id: "3"}
	output := usecase.NewFindAlbumById(repository).Execute(input)
	if output.Album == nil {
		t.Errorf("Expected album with #%v, but got %v", input.Id, output.Album)
	} else if output.Album.Id != "3" {
		t.Errorf("Expected album with #%v, but got %v", input.Id, output.Album)
	}
}
