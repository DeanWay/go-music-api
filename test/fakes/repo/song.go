package repo

import (
	"errors"
	"go-music-api/internal/pkg/models"
	"go-music-api/internal/pkg/payloads"

	"github.com/google/uuid"
)

type FakeSongRepo struct {
	Songs []models.Song
}

func (repo FakeSongRepo) AddSong(
	albumUuid uuid.UUID,
	attrs payloads.SongAttributes,
) models.Song {
	newSong := models.NewSong(albumUuid, attrs)
	repo.Songs = append(repo.Songs, newSong)
	return newSong
}

func (repo FakeSongRepo) FindSongById(id string) (models.Song, error) {
	var found *models.Song
	for _, song := range repo.Songs {
		if song.Uuid.String() == id {
			found = &song
		}
	}
	if found == nil {
		return models.Song{}, errors.New("not found")
	} else {
		return *found, nil
	}
}
