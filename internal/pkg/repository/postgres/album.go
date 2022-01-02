package postgres

import (
	"go-todo-app/internal/pkg/models"
	"go-todo-app/internal/pkg/payloads"
	"go-todo-app/internal/pkg/repository"
)

type AlbumPostgresRepo struct{}

func (repo AlbumPostgresRepo) GetAllAlbums() []models.Album {
	panic("not implemented")
}

func (repo AlbumPostgresRepo) FindAlbumById(id string) (models.Album, error) {
	panic("not implemented")
}

func (repo AlbumPostgresRepo) AddAlbum(
	request payloads.AlbumAttributes,
) models.Album {
	panic("not implemented")
}

func (repo AlbumPostgresRepo) SearchAlbums(
	params repository.AlbumSearchParams,
) []models.Album {
	panic("not implemented")
}
