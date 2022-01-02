package examples

import (
	"go-todo-app/internal/pkg/payloads"
	"go-todo-app/internal/pkg/repository"
)

func AddExampleAlbums(albumRepo repository.AlbumRepository) {
	albumRepo.AddAlbum(payloads.AlbumAttributes{
		Title:  "Blue Train",
		Artist: "John Coltrane",
		Price:  56.99,
	})
	albumRepo.AddAlbum(payloads.AlbumAttributes{
		Title:  "Jeru",
		Artist: "Gerry Mulligan",
		Price:  17.99,
	})
	albumRepo.AddAlbum(payloads.AlbumAttributes{
		Title:  "Sarah Vaughan and Clifford Brown",
		Artist: "Sarah Vaughan",
		Price:  39.99,
	})
}
