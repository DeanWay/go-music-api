package repo

import (
	"errors"
	"go-music-api/internal/pkg/models"
	"go-music-api/internal/pkg/payloads"
	"go-music-api/internal/pkg/repository"
)

type FakeAlbumRepo struct {
	Albums []models.Album
}

func (repo FakeAlbumRepo) GetAllAlbums() []models.Album {
	return repo.Albums
}

func (repo FakeAlbumRepo) AddAlbum(
	attrs payloads.AlbumAttributes,
	songs []payloads.SongAttributes,
) models.Album {
	album := models.NewAlbum(attrs)
	repo.Albums = append(repo.Albums, album)
	return album
}

func (repo FakeAlbumRepo) FindAlbumById(id string) (models.Album, error) {
	var found *models.Album
	for _, album := range repo.Albums {
		if album.Uuid.String() == id {
			found = &album
		}
	}
	if found == nil {
		return models.Album{}, errors.New("not found")
	} else {
		return *found, nil
	}
}

func (repo FakeAlbumRepo) SearchAlbums(
	params repository.AlbumSearchParams,
) []models.Album {
	return repo.Albums
}
