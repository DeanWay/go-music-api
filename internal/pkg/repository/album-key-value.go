package repository

import (
	"time"

	"github.com/google/uuid"

	"go-todo-app/internal/pkg/models"
	"go-todo-app/internal/pkg/payloads"
	"go-todo-app/internal/pkg/storage"
)

type AlbumKeyValueRepo struct {
	Store storage.KeyValueStorage
}

func (repo AlbumKeyValueRepo) GetAllAlbums() []models.Album {
	storeList := repo.Store.List()
	albums := make([]models.Album, len(storeList))
	for i, v := range storeList {
		albums[i] = v.(models.Album)
	}
	return albums
}

func (repo AlbumKeyValueRepo) FindAlbumById(id string) (models.Album, error) {
	val, err := repo.Store.GetById(id)
	album := val.(models.Album)
	return album, err
}

func (repo AlbumKeyValueRepo) AddAlbum(
	request payloads.AlbumAttributes,
) models.Album {
	newAlbum := models.Album{
		Uuid:      uuid.New(),
		Title:     request.Title,
		Artist:    request.Artist,
		Price:     request.Price,
		CreatedAt: time.Now().UTC(),
	}
	repo.Store.Insert(newAlbum.Uuid.String(), newAlbum)
	return newAlbum
}
