package payloads

type SongAttributes struct {
	Title           string `json:"title"`
	Artist          string `json:"artist"`
	DurationSeconds uint   `json:"durationSeconds"`
}

type SongRequestData struct {
	Attributes SongAttributes `json:"attributes"`
}

type SongRequest struct {
	Data SongRequestData `json:"data"`
}
