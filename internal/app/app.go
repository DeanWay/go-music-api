package app

import (
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"

	"go-music-api/internal/pkg/repository"
	"go-music-api/internal/pkg/repository/postgres"
	"go-music-api/internal/pkg/routes"
	psqlStorage "go-music-api/internal/pkg/storage/postgres"
)

func App(deps *Deps) *gin.Engine {
	appEngine := gin.Default()
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

func DefaultDeps() Deps {
	postgresStorage := psqlStorage.New(
		psqlStorage.ConnectionParams{
			Username: "postgres",
			Password: "mysecretpassword",
			Host:     "localhost",
			Port:     "5432",
			Database: "music",
		},
	)

	songRepo := postgres.SongPostgresRepo{
		Store: postgresStorage,
	}
	albumRepo := postgres.AlbumPostgresRepo{
		Store:    postgresStorage,
		SongRepo: songRepo,
	}
	return Deps{
		AlbumRepo: albumRepo,
		SongRepo:  songRepo,
	}
}

func newRedisClient() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
}
