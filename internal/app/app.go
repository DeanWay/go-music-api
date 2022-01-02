package app

import (
	"github.com/gin-gonic/gin"

	"go-todo-app/internal/pkg/examples"
	"go-todo-app/internal/pkg/repository"
	"go-todo-app/internal/pkg/routes"
	"go-todo-app/internal/pkg/storage/memory"
)

func App() *gin.Engine {
	appEngine := gin.Default()
	albumRouter := routes.AlbumRouter{
		AlbumRepository: initAlbumRepo(),
	}

	appEngine.GET("/albums", albumRouter.GetAlbums)
	appEngine.GET("/albums/:id", albumRouter.GetAlbumByID)
	appEngine.POST("/albums", albumRouter.PostAlbums)
	return appEngine
}

func initAlbumRepo() repository.AlbumRepository {
	albumRepo := repository.AlbumKeyValueRepo{
		Store: memory.MemoryStorage{},
	}
	examples.AddExampleAlbums(albumRepo)
	return albumRepo
}
