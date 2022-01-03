package payloads

type SongAttributes struct {
	Title           string `json:"title"`
	DurationSeconds uint   `json:"durationSeconds"`
}
