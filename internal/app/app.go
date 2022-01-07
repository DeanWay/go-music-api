package app

import (
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"

	"go-music-api/internal/pkg/adapter/repository/postgres"
	psqlStorage "go-music-api/internal/pkg/adapter/storage/postgres"
	"go-music-api/internal/pkg/domain/port"
	"go-music-api/internal/pkg/domain/usecase"
	"go-music-api/internal/pkg/http/routes"
)

func App(deps *Deps) *gin.Engine {
	appEngine := gin.Default()

	// album
	appEngine.GET(
		"/album/:id",
		routes.GetAlbumByID(usecase.GetAlbumUseCase{
			AlbumRepository: deps.AlbumRepository,
		}),
	)
	appEngine.GET(
		"/albums",
		routes.ListAlbums(usecase.ListAlbumsUseCase{
			AlbumRepository: deps.AlbumRepository,
		}),
	)
	appEngine.POST(
		"/albums",
		routes.PostAlbums(usecase.CreateAlbumUseCase{
			AlbumRepository: deps.AlbumRepository,
		}),
	)
	appEngine.GET(
		"/albums/search",
		routes.SearchAlbums(usecase.SearchAlbumsUseCase{
			AlbumRepository: deps.AlbumRepository,
		}),
	)

	// song
	appEngine.GET(
		"/song/:id",
		routes.GetSongByID(usecase.GetSongUseCase{
			SongRepository: deps.SongRepository,
		}),
	)
	appEngine.POST(
		"/songs",
		routes.PostSong(usecase.CreateSongUseCase{
			SongRepository: deps.SongRepository,
		}),
	)
	return appEngine
}

type Deps struct {
	AlbumRepository port.AlbumRepository
	SongRepository  port.SongRepository
}

func DefaultDeps() Deps {
	postgresStorage := psqlStorage.New(
		psqlStorage.Config{
			ConnectionParams: psqlStorage.ConnectionParams{
				Username: "postgres",
				Password: "mysecretpassword",
				Host:     "localhost",
				Port:     "5432",
				Database: "music",
			},
			LogSql: true,
		},
	)

	songRepo := postgres.SongPostgresRepository{
		Store: postgresStorage,
	}
	albumRepo := postgres.AlbumPostgresRepository{
		Store: postgresStorage,
	}
	return Deps{
		AlbumRepository: albumRepo,
		SongRepository:  songRepo,
	}
}

func newRedisClient() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
}
