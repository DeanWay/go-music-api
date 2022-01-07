package examples

import (
	"fmt"
	"go-music-api/internal/pkg/domain/entity"
	"go-music-api/internal/pkg/domain/port"
	"time"

	"github.com/google/uuid"
)

func AddExampleAlbums(albumRepo port.AlbumRepository) {
	err := albumRepo.AddAlbum(
		entity.Album{
			Id:        uuid.New(),
			Title:     "Blue Train",
			Artist:    "John Coltrane",
			Price:     56.99,
			CreatedAt: time.Now().UTC(),
		},
	)
	err = albumRepo.AddAlbum(
		entity.Album{
			Id:        uuid.New(),
			Title:     "Jeru",
			Artist:    "Gerry Mulligan",
			Price:     17.99,
			CreatedAt: time.Now().UTC(),
		},
	)
	err = albumRepo.AddAlbum(
		entity.Album{
			Id:        uuid.New(),
			Title:     "Sarah Vaughan and Clifford Brown",
			Artist:    "Sarah Vaughan",
			Price:     39.99,
			CreatedAt: time.Now().UTC(),
		},
	)
	if err != nil {
		fmt.Println(err)
	}
}
