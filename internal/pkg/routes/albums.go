package routes

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"go-todo-app/internal/pkg/interfaces"
	"go-todo-app/internal/pkg/models"
	"go-todo-app/internal/pkg/services"
)

type AlbumRouter struct {
	AlbumService services.AlbumService
}

func (resource AlbumRouter) GetAlbums(c *gin.Context) {
	responseList := albumsToResponse(resource.AlbumService.GetAllAlbums())
	c.IndentedJSON(http.StatusOK, responseList)
}

func (resource AlbumRouter) PostAlbums(c *gin.Context) {
	var albumRequest struct {
		Attributes interfaces.AlbumAttributes `json:"attributes"`
	}
	if err := c.BindJSON(&albumRequest); err != nil {
		badRequest(c)
		return
	}
	newAlbum := resource.AlbumService.AddAlbum(albumRequest.Attributes)
	response := albumToResponse(newAlbum)
	c.IndentedJSON(http.StatusCreated, response)
}

func (resource AlbumRouter) GetAlbumByID(c *gin.Context) {
	id := c.Param("id")
	album, err := resource.AlbumService.FindAlbumById(id)
	if err == nil {
		response := albumToResponse(album)
		c.IndentedJSON(http.StatusOK, response)
	} else {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
	}
}

func badRequest(c *gin.Context) {
	c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "bad request"})
}

func albumToResponse(album models.Album) interfaces.Document {
	return interfaces.Document{
		Id: album.Uuid.String(),
		Attributes: interfaces.AlbumAttributes{
			Title:     album.Title,
			Artist:    album.Artist,
			Price:     album.Price,
			CreatedAt: album.CreatedAt.Format(time.RFC1123),
		},
	}
}

func albumsToResponse(albums []models.Album) []interfaces.Document {
	newList := make([]interfaces.Document, len(albums), len(albums))
	for i, v := range albums {
		newList[i] = albumToResponse(v)
	}
	return newList
}
