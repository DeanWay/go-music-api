package payloads

type Document struct {
	Id         string      `json:"id"`
	Attributes interface{} `json:"attributes"`
}

type AlbumAttributes struct {
	Title     string  `json:"title"`
	Artist    string  `json:"artist"`
	Price     float64 `json:"price"`
	CreatedAt string  `json:"createdAt"`
}
