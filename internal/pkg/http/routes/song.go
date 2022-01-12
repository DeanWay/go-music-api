package routes

import (
	"errors"
	"net/http"
	"net/url"

	"github.com/gin-gonic/gin"

	"go-music-api/internal/pkg/domain/entity"
	"go-music-api/internal/pkg/domain/failure"
	"go-music-api/internal/pkg/domain/usecase"
	"go-music-api/internal/pkg/http/payloads"
)

func GetSongByID(getSong usecase.GetSongUseCase) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		song, err := getSong.GetSong(id)
		if err == nil {
			response := songToResponse(song)
			c.IndentedJSON(http.StatusOK, response)
		} else if errors.Is(err, failure.ErrNotFound) {
			c.IndentedJSON(http.StatusNotFound, gin.H{"message": "song not found"})
		} else {
			panic(err)
		}
	}
}

func PostSong(createSong usecase.CreateSongUseCase) gin.HandlerFunc {
	return func(c *gin.Context) {
		var request payloads.SongRequest
		bindJsonErr := c.BindJSON(&request)
		attrs := request.Data.Attributes
		audioFileUrl, urlParseErr := url.Parse(attrs.AudioFile)
		if bindJsonErr != nil || urlParseErr != nil {
			badRequest(c)
			return
		}
		newSong, err := createSong.CreateSong(
			attrs.Title,
			attrs.Artist,
			attrs.DurationSeconds,
			*audioFileUrl,
		)
		if err != nil {
			panic(err)
		}
		response := songToResponse(newSong)
		c.IndentedJSON(http.StatusCreated, response)
	}
}

func songToResourceObject(song entity.Song) payloads.ResponseResourceObject {
	return payloads.ResponseResourceObject{
		Id:   song.Id.String(),
		Type: "song",
		Attributes: payloads.SongAttributes{
			Title:           song.Title,
			Artist:          song.Artist,
			DurationSeconds: song.DurationSeconds,
			AudioFile:       song.AudioFile.String(),
		},
	}
}

func songToResponse(song entity.Song) payloads.Response {
	return payloads.MakeResponse(songToResourceObject(song))
}
