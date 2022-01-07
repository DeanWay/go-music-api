package keyvalue

import (
	"encoding/json"
	"go-music-api/internal/pkg/adapter/storage"
	"go-music-api/internal/pkg/domain/entity"
	"go-music-api/internal/pkg/domain/port"
)

type SongKeyValueRepository struct {
	Store storage.KeyValueStorage
}

var _ port.SongRepository = (*SongKeyValueRepository)(nil)

func (repo SongKeyValueRepository) AddSong(
	newSong entity.Song,
) error {
	songJson, err := json.Marshal(newSong)
	if err != nil {
		return err
	}
	repo.Store.Insert(newSong.Id.String(), string(songJson))
	return nil
}

func (repo SongKeyValueRepository) GetSongById(
	id string,
) (entity.Song, error) {
	val, err := repo.Store.Get(id)
	if err != nil {
		return entity.Song{}, err
	}
	var song entity.Song
	json.Unmarshal([]byte(val), &song)
	return song, err
}
