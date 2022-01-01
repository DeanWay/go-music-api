package app

import (
	"github.com/gin-gonic/gin"

	"go-todo-app/internal/pkg/interfaces"
	"go-todo-app/internal/pkg/routes"
	"go-todo-app/internal/pkg/services"
	"go-todo-app/internal/pkg/storage/memory"
)

func App() {
	router := gin.Default()
	albumRouter := routes.AlbumRouter{
		AlbumService: initAlbumSerice(),
	}

	router.GET("/albums", albumRouter.GetAlbums)
	router.GET("/albums/:id", albumRouter.GetAlbumByID)
	router.POST("/albums", albumRouter.PostAlbums)
	router.Run("localhost:8000")
}

func initAlbumSerice() services.AlbumService {
	albumService := services.AlbumService{
		Store: memory.MemoryStorage{},
	}
	albumService.AddAlbum(interfaces.AlbumAttributes{
		Title:  "Blue Train",
		Artist: "John Coltrane",
		Price:  56.99,
	})
	albumService.AddAlbum(interfaces.AlbumAttributes{
		Title:  "Jeru",
		Artist: "Gerry Mulligan",
		Price:  17.99,
	})
	albumService.AddAlbum(interfaces.AlbumAttributes{
		Title:  "Sarah Vaughan and Clifford Brown",
		Artist: "Sarah Vaughan",
		Price:  39.99,
	})
	return albumService
}
