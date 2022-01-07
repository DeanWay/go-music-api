package payloads

type AlbumAttributes struct {
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

type AlbumRequestRelationships struct {
	Songs RequestRelationshipModel `json:"songs"`
}

type AlbumRequestData struct {
	Attributes    AlbumAttributes           `json:"attributes"`
	Relationships AlbumRequestRelationships `json:"relationships"`
}

type AlbumRequest struct {
	Data AlbumRequestData `json:"data"`
}
