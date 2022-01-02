package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"go-music-api/internal/pkg/models"
	"go-music-api/internal/pkg/payloads"
	"go-music-api/internal/pkg/repository"
)

type SongRouter struct {
	SongRepository repository.SongRepository
}

func (resource SongRouter) GetSongByID(c *gin.Context) {
	id := c.Param("id")
	song, err := resource.SongRepository.FindSongById(id)
	if err == nil {
		response := songToResponse(song)
		c.IndentedJSON(http.StatusOK, response)
	} else {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "song not found"})
	}
}

func songToResourceObject(song models.Song) payloads.ResourceObject {
	return payloads.ResourceObject{
		Id:   song.Uuid.String(),
		Type: "song",
		Attributes: payloads.SongAttributes{
			Title:           song.Title,
			DurationSeconds: song.DurationSeconds,
		},
		Relationships: payloads.ResponseRelationshipMap{
			"album": payloads.ResponseRelationshipModel{
				Data: payloads.ResourceLinkage{
					{
						Id:   song.AlbumUuid.String(),
						Type: "album",
					},
				},
			},
		},
	}
}

func songToResponse(song models.Song) payloads.Response {
	return payloads.MakeResponse(songToResourceObject(song))
}
