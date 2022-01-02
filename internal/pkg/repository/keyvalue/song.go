package keyvalue

import (
	"encoding/json"
	"go-music-api/internal/pkg/models"
	"go-music-api/internal/pkg/payloads"
	"go-music-api/internal/pkg/storage"

	"github.com/google/uuid"
)

type SongKeyValueRepo struct {
	Store storage.KeyValueStorage
}

func (repo SongKeyValueRepo) AddSong(
	albumUuid uuid.UUID,
	attrs payloads.SongAttributes,
) models.Song {
	newSong := models.Song{
		Uuid:            uuid.New(),
		AlbumUuid:       albumUuid,
		Title:           attrs.Title,
		DurationSeconds: attrs.DurationSeconds,
	}
	songJson, _ := json.Marshal(newSong)
	repo.Store.Insert(newSong.Uuid.String(), string(songJson))
	return newSong
}

func (repo SongKeyValueRepo) FindSongById(
	id string,
) (models.Song, error) {
	val, err := repo.Store.Get(id)
	var song models.Song
	json.Unmarshal([]byte(val), &song)
	return song, err
}
