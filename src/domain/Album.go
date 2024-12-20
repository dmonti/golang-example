package domain

type Album struct {
	Id     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

func (album Album) Equals(other Album) bool {
	return album.Id == other.Id
}
