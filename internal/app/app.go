package app

import (
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"

	"go-music-api/internal/pkg/repository"
	"go-music-api/internal/pkg/repository/keyvalue"
	"go-music-api/internal/pkg/routes"
	redisStorage "go-music-api/internal/pkg/storage/redis"
)

func App() *gin.Engine {
	appEngine := gin.Default()
	deps := initDeps()
	albumRouter := routes.AlbumRouter{
		AlbumRepository: deps.AlbumRepo,
	}
	songRouter := routes.SongRouter{
		SongRepository: deps.SongRepo,
	}

	// album
	appEngine.GET("/album/:id", albumRouter.GetAlbumByID)
	appEngine.GET("/albums", albumRouter.ListAlbums)
	appEngine.POST("/albums", albumRouter.PostAlbums)
	appEngine.GET("/albums/search", albumRouter.SearchAlbums)

	// song
	appEngine.GET("/song/:id", songRouter.GetSongByID)
	return appEngine
}

type Deps struct {
	AlbumRepo repository.AlbumRepository
	SongRepo  repository.SongRepository
}

func initDeps() Deps {
	redisClient := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	songRepo := keyvalue.SongKeyValueRepo{
		Store: redisStorage.RedisStorage{
			CollectionName: "song",
			Client:         redisClient,
		},
	}
	albumRepo := keyvalue.AlbumKeyValueRepo{
		Store: redisStorage.RedisStorage{
			CollectionName: "album",
			Client:         redisClient,
		},
		SongRepo: songRepo,
	}
	return Deps{
		AlbumRepo: albumRepo,
		SongRepo:  songRepo,
	}
}
