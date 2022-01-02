package keyvalue

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"

	"go-todo-app/internal/pkg/models"
	"go-todo-app/internal/pkg/payloads"
	"go-todo-app/internal/pkg/repository"
	"go-todo-app/internal/pkg/storage"
)

type AlbumKeyValueRepo struct {
	Store    storage.KeyValueStorage
	SongRepo SongKeyValueRepo
}

func (repo AlbumKeyValueRepo) GetAllAlbums() []models.Album {
	storeList := repo.Store.List()
	albums := make([]models.Album, len(storeList))
	for i, v := range storeList {
		var album models.Album
		json.Unmarshal([]byte(v), &album)
		albums[i] = album
	}
	return albums
}

func (repo AlbumKeyValueRepo) FindAlbumById(
	id string,
) (models.Album, error) {
	val, err := repo.Store.Get(id)
	var album models.Album
	json.Unmarshal([]byte(val), &album)
	return album, err
}

func (repo AlbumKeyValueRepo) AddAlbum(
	attrs payloads.AlbumAttributes,
	songAttrs []payloads.SongAttributes,
) models.Album {
	newAlbum := models.Album{
		Uuid:      uuid.New(),
		Title:     attrs.Title,
		Artist:    attrs.Artist,
		Price:     attrs.Price,
		CreatedAt: time.Now().UTC(),
	}
	songUuids := make([]uuid.UUID, len(songAttrs), len(songAttrs))
	for i, song := range songAttrs {
		newSong := repo.SongRepo.AddSong(newAlbum.Uuid, song)
		songUuids[i] = newSong.Uuid
	}
	newAlbum.SongUuids = songUuids
	albumJson, _ := json.Marshal(newAlbum)
	repo.Store.Insert(newAlbum.Uuid.String(), string(albumJson))
	return newAlbum
}

func (repo AlbumKeyValueRepo) SearchAlbums(
	params repository.AlbumSearchParams,
) []models.Album {
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
