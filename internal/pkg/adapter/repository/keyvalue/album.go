package keyvalue

import (
	"encoding/json"

	"go-music-api/internal/pkg/adapter/storage"
	"go-music-api/internal/pkg/domain/entity"
	"go-music-api/internal/pkg/domain/port"
)

type AlbumKeyValueRepository struct {
	Store storage.KeyValueStorage
}

var _ port.AlbumRepository = (*AlbumKeyValueRepository)(nil)

func (repo AlbumKeyValueRepository) ListAlbums() ([]entity.Album, error) {
	storeList := repo.Store.List()
	albums := make([]entity.Album, len(storeList))
	for i, v := range storeList {
		var album entity.Album
		err := json.Unmarshal([]byte(v), &album)
		if err != nil {
			return albums, err
		}
		albums[i] = album
	}
	return albums, nil
}

func (repo AlbumKeyValueRepository) GetAlbumById(
	id string,
) (entity.Album, error) {
	val, err := repo.Store.Get(id)
	if err != nil {
		return entity.Album{}, err
	}
	var album entity.Album
	err = json.Unmarshal([]byte(val), &album)
	return album, err
}

func (repo AlbumKeyValueRepository) AddAlbum(
	newAlbum entity.Album,
) error {
	albumJson, _ := json.Marshal(newAlbum)
	repo.Store.Insert(newAlbum.Id.String(), string(albumJson))
	return nil
}

func (repo AlbumKeyValueRepository) SearchAlbums(
	params port.AlbumSearchParams,
) ([]entity.Album, error) {
	albums, err := repo.ListAlbums()
	if err != nil {
		return albums, err
	}
	result := []entity.Album{}
	for _, album := range albums {
		filtered := filterAlbum(&album, params)
		if filtered != nil {
			result = append(result, *filtered)
		}
	}
	return result, nil
}

func filterAlbum(
	album *entity.Album,
	params port.AlbumSearchParams,
) *entity.Album {
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
