package app

import (
	"github.com/gin-gonic/gin"

	"go-todo-app/internal/pkg/examples"
	"go-todo-app/internal/pkg/routes"
	"go-todo-app/internal/pkg/services"
	"go-todo-app/internal/pkg/storage/memory"
)

func App() *gin.Engine {
	appEngine := gin.Default()
	albumRouter := routes.AlbumRouter{
		AlbumService: initAlbumSerice(),
	}

	appEngine.GET("/albums", albumRouter.GetAlbums)
	appEngine.GET("/albums/:id", albumRouter.GetAlbumByID)
	appEngine.POST("/albums", albumRouter.PostAlbums)
	return appEngine
}

func initAlbumSerice() services.AlbumService {
	albumService := services.AlbumService{
		Store: memory.MemoryStorage{},
	}
	examples.AddExampleAlbums(albumService)
	return albumService
}
