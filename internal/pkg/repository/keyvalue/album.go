package keyvalue

import (
	"fmt"
	"time"

	"github.com/google/uuid"

	"go-todo-app/internal/pkg/models"
	"go-todo-app/internal/pkg/payloads"
	"go-todo-app/internal/pkg/repository"
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

func (repo AlbumKeyValueRepo) FindAlbumById(
	id string,
) (models.Album, error) {
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

func (repo AlbumKeyValueRepo) SearchAlbums(
	params repository.AlbumSearchParams,
) []models.Album {
	fmt.Println(params)
	albums := repo.GetAllAlbums()
	result := []models.Album{}
	for _, album := range albums {
		filtered := filterAlbum(&album, params)
		if filtered != nil {
			result = append(result, *filtered)
		}
	}
	return result
}

func filterAlbum(
	album *models.Album,
	params repository.AlbumSearchParams,
) *models.Album {
	if params.PriceHigh != nil && (*album).Price > *params.PriceHigh {
		return nil
	}
	if params.PriceLow != nil && (*album).Price < *params.PriceLow {
		return nil
	}
	if params.Title != nil && (*album).Title != *params.Title {
		return nil
	}
	if params.Artist != nil && (*album).Artist != *params.Artist {
		return nil
	}
	return album
}
