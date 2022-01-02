package examples

import (
	"go-music-api/internal/pkg/payloads"
	"go-music-api/internal/pkg/repository"
)

func AddExampleAlbums(albumRepo repository.AlbumRepository) {
	albumRepo.AddAlbum(
		payloads.AlbumAttributes{
			Title:  "Blue Train",
			Artist: "John Coltrane",
			Price:  56.99,
		},
		[]payloads.SongAttributes{
			{
				Title:           "The Song",
				DurationSeconds: 316,
			},
		},
	)
	albumRepo.AddAlbum(
		payloads.AlbumAttributes{
			Title:  "Jeru",
			Artist: "Gerry Mulligan",
			Price:  17.99,
		},
		[]payloads.SongAttributes{},
	)
	albumRepo.AddAlbum(
		payloads.AlbumAttributes{
			Title:  "Sarah Vaughan and Clifford Brown",
			Artist: "Sarah Vaughan",
			Price:  39.99,
		},
		[]payloads.SongAttributes{},
	)
}
