package repository

import (
	"go-todo-app/internal/pkg/models"
	"go-todo-app/internal/pkg/payloads"
)

type AlbumRepository interface {
	GetAllAlbums() []models.Album
	AddAlbum(attrs payloads.AlbumAttributes) models.Album
	FindAlbumById(id string) (models.Album, error)
}
