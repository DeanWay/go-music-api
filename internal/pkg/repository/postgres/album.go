package postgres

import (
	"go-music-api/internal/pkg/models"
	"go-music-api/internal/pkg/payloads"
	"go-music-api/internal/pkg/repository"
)

type AlbumPostgresRepo struct{}

func (repo AlbumPostgresRepo) GetAllAlbums() []models.Album {
	panic("not implemented")
}

func (repo AlbumPostgresRepo) FindAlbumById(id string) (models.Album, error) {
	panic("not implemented")
}

func (repo AlbumPostgresRepo) AddAlbum(
	attrs payloads.AlbumAttributes,
) models.Album {
	panic("not implemented")
}

func (repo AlbumPostgresRepo) SearchAlbums(
	params repository.AlbumSearchParams,
) []models.Album {
	panic("not implemented")
}
