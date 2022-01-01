package services

import (
	"fmt"
	"go-todo-app/internal/pkg/interfaces"
	"go-todo-app/internal/pkg/models"
	"go-todo-app/internal/pkg/storage"
	"time"

	"github.com/google/uuid"
)

type AlbumService struct {
	Store storage.Storage
}

func (service AlbumService) GetAllAlbums() []models.Album {
	storeList := service.Store.List()
	albums := make([]models.Album, len(storeList))
	fmt.Println(storeList)
	for i, v := range storeList {
		albums[i] = v.(models.Album)
	}
	return albums
}

func (service AlbumService) FindAlbumById(id string) (models.Album, error) {
	val, err := service.Store.GetById(id)
	album := val.(models.Album)
	return album, err
}

func (service AlbumService) AddAlbum(
	request interfaces.AlbumAttributes,
) models.Album {
	newAlbum := models.Album{
		Uuid:      uuid.New(),
		Title:     request.Title,
		Artist:    request.Artist,
		Price:     request.Price,
		CreatedAt: time.Now().UTC(),
	}
	service.Store.Insert(newAlbum.Uuid.String(), newAlbum)
	return newAlbum
}
