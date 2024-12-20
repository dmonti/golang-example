package domain

import (
	"example/src/domain"
	"testing"
)

func TestAlbum_equals(t *testing.T) {
	album := domain.Album{Id: "1"}
	other := domain.Album{Id: "2"}

	if album.Equals(other) {
		t.Errorf("Expected false, but got true")
	}
}
