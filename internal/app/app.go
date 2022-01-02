package app

import (
	"github.com/gin-gonic/gin"

	"go-todo-app/internal/pkg/examples"
	"go-todo-app/internal/pkg/repository"
	"go-todo-app/internal/pkg/repository/keyvalue"
	"go-todo-app/internal/pkg/routes"
	"go-todo-app/internal/pkg/storage/memory"
)

func App() *gin.Engine {
	appEngine := gin.Default()
	albumRouter := routes.AlbumRouter{
		AlbumRepository: initAlbumRepo(),
	}

	appEngine.GET("/album/:id", albumRouter.GetAlbumByID)
	appEngine.GET("/albums", albumRouter.ListAlbums)
	appEngine.POST("/albums", albumRouter.PostAlbums)
	appEngine.GET("/albums/search", albumRouter.SearchAlbums)
	return appEngine
}

func initAlbumRepo() repository.AlbumRepository {
	albumRepo := keyvalue.AlbumKeyValueRepo{
		Store: memory.MemoryStorage{},
	}
	examples.AddExampleAlbums(albumRepo)
	return albumRepo
}
