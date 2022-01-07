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

const songPrefix = "song"

var _ port.SongRepository = (*SongKeyValueRepository)(nil)

func (repo SongKeyValueRepository) AddSong(
	newSong entity.Song,
) error {
	songJson, err := json.Marshal(newSong)
	if err != nil {
		return err
	}
	key := withPrefix(songPrefix, newSong.Id.String())
	repo.Store.Insert(key, string(songJson))
	return nil
}

func (repo SongKeyValueRepository) GetSongById(
	id string,
) (entity.Song, error) {
	key := withPrefix(songPrefix, id)
	val, err := repo.Store.Get(key)
	if err != nil {
		return entity.Song{}, err
	}
	var song entity.Song
	json.Unmarshal([]byte(val), &song)
	return song, err
}
