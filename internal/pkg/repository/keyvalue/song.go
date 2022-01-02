package keyvalue

import (
	"encoding/json"
	"go-todo-app/internal/pkg/models"
	"go-todo-app/internal/pkg/payloads"
	"go-todo-app/internal/pkg/storage"

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
