package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"go-music-api/internal/pkg/models"
	"go-music-api/internal/pkg/payloads"
	"go-music-api/internal/pkg/repository"
)

type AlbumRouter struct {
	AlbumRepository repository.AlbumRepository
}

func (resource AlbumRouter) ListAlbums(c *gin.Context) {
	responseList := albumsToResponse(resource.AlbumRepository.GetAllAlbums())
	c.IndentedJSON(http.StatusOK, responseList)
}

func (resource AlbumRouter) PostAlbums(c *gin.Context) {
	var albumRequest struct {
		Attributes payloads.AlbumAttributes `json:"attributes"`
	}
	if err := c.BindJSON(&albumRequest); err != nil {
		badRequest(c)
		return
	}
	newAlbum := resource.AlbumRepository.AddAlbum(
		albumRequest.Attributes,
		[]payloads.SongAttributes{},
	)
	response := albumToResponse(newAlbum)
	c.IndentedJSON(http.StatusCreated, response)
}

func (resource AlbumRouter) GetAlbumByID(c *gin.Context) {
	id := c.Param("id")
	album, err := resource.AlbumRepository.FindAlbumById(id)
	if err == nil {
		response := albumToResponse(album)
		c.IndentedJSON(http.StatusOK, response)
	} else {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
	}
}

func (resource AlbumRouter) SearchAlbums(c *gin.Context) {
	query := repository.AlbumSearchParams{}
	priceHigh, err := parseFloatQuery(c, "price-high")
	if err != nil {
		badRequest(c)
		return
	}
	query.PriceHigh = priceHigh

	priceLow, err := parseFloatQuery(c, "price-low")
	if err != nil {
		badRequest(c)
		return
	}
	query.PriceLow = priceLow

	title, present := c.GetQuery("title")
	if present {
		query.Title = &title
	}

	artist, present := c.GetQuery("artist")
	if present {
		query.Artist = &artist
	}

	albums := resource.AlbumRepository.SearchAlbums(query)
	c.IndentedJSON(http.StatusOK, albumsToResponse(albums))
}

func albumToResourceObject(album models.Album) payloads.ResourceObject {
	songLinks := payloads.ResourceLinkage{}
	for _, songUuid := range album.SongUuids {
		songLinks = append(songLinks, payloads.ResourceIdentifier{
			Id:   songUuid.String(),
			Type: "song",
		})
	}

	return payloads.ResourceObject{
		Id:   album.Uuid.String(),
		Type: "album",
		Attributes: payloads.AlbumAttributes{
			Title:  album.Title,
			Artist: album.Artist,
			Price:  album.Price,
		},
		Relationships: payloads.ResponseRelationshipMap{
			"songs": payloads.ResponseRelationshipModel{
				Data: songLinks,
			},
		},
	}
}

func albumToResponse(album models.Album) payloads.Response {
	return payloads.MakeResponse(albumToResourceObject(album))
}

func albumsToResponse(albums []models.Album) payloads.ListResponse {
	dataList := make([]payloads.ResourceObject, len(albums), len(albums))
	for i, v := range albums {
		dataList[i] = albumToResourceObject(v)
	}
	return payloads.MakeListResponse(dataList)
}
