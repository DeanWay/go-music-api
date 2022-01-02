package repository

import (
	"go-music-api/internal/pkg/models"
	"go-music-api/internal/pkg/payloads"
)

type AlbumSearchParams struct {
	PriceHigh *float64
	PriceLow  *float64
	Artist    *string
	Title     *string
}

type AlbumRepository interface {
	GetAllAlbums() []models.Album
	AddAlbum(
		attrs payloads.AlbumAttributes,
		songs []payloads.SongAttributes,
	) models.Album
	FindAlbumById(id string) (models.Album, error)
	SearchAlbums(params AlbumSearchParams) []models.Album
}
