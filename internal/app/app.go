package app

import (
	"github.com/gin-gonic/gin"

	"go-music-api/config"
	"go-music-api/internal/pkg/domain/usecase"
	"go-music-api/internal/pkg/http/routes"
)

func RestApiApp(deps *config.Deps) *gin.Engine {
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
