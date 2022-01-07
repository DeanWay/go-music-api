package routes

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"go-music-api/internal/pkg/domain/entity"
	"go-music-api/internal/pkg/domain/failure"
	"go-music-api/internal/pkg/domain/port"
	"go-music-api/internal/pkg/domain/usecase"
	"go-music-api/internal/pkg/http/payloads"
)

func ListAlbums(listAlbums usecase.ListAlbumsUseCase) gin.HandlerFunc {
	return func(c *gin.Context) {
		albums, err := listAlbums.ListAlbums()
		if err != nil {
			panic(err)
		}
		responseList := albumsToResponse(albums)
		c.IndentedJSON(http.StatusOK, responseList)
	}
}

func PostAlbums(createAlbum usecase.CreateAlbumUseCase) gin.HandlerFunc {
	return func(c *gin.Context) {
		var request payloads.AlbumRequest
		requestParseErr := c.BindJSON(&request)
		songIds, uuidParseErr := songIdsFromRequest(request)
		if requestParseErr != nil || uuidParseErr != nil {
			badRequest(c)
			return
		}
		attrs := request.Data.Attributes
		newAlbum, err := createAlbum.CreateAlbum(
			attrs.Title,
			attrs.Artist,
			attrs.Price,
			songIds,
		)
		if err != nil {
			panic(err)
		}
		response := albumToResponse(newAlbum)
		c.IndentedJSON(http.StatusCreated, response)
	}

}

func GetAlbumByID(getAlbums usecase.GetAlbumUseCase) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		album, err := getAlbums.GetAlbum(id)
		if err == nil {
			response := albumToResponse(album)
			c.IndentedJSON(http.StatusOK, response)
		} else if errors.Is(err, failure.ErrNotFound) {
			c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
		} else {
			panic(err)
		}
	}

}

func SearchAlbums(searchAlbums usecase.SearchAlbumsUseCase) gin.HandlerFunc {
	return func(c *gin.Context) {
		query := port.AlbumSearchParams{}
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

		albums, err := searchAlbums.SearchAlbums(query)
		if err != nil {
			panic(err)
		}
		c.IndentedJSON(http.StatusOK, albumsToResponse(albums))
	}
}

func albumToResourceObject(album entity.Album) payloads.ResponseResourceObject {
	songLinks := payloads.ResourceLinkage{}
	for _, songUuid := range album.SongIds {
		songLinks = append(songLinks, payloads.ResourceIdentifier{
			Id:   songUuid.String(),
			Type: "song",
		})
	}

	return payloads.ResponseResourceObject{
		Id:   album.Id.String(),
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

func albumToResponse(album entity.Album) payloads.Response {
	return payloads.MakeResponse(albumToResourceObject(album))
}

func albumsToResponse(albums []entity.Album) payloads.ListResponse {
	dataList := make([]payloads.ResponseResourceObject, len(albums), len(albums))
	for i, v := range albums {
		dataList[i] = albumToResourceObject(v)
	}
	return payloads.MakeListResponse(dataList)
}

func songIdsFromRequest(request payloads.AlbumRequest) ([]uuid.UUID, error) {
	songIds := []uuid.UUID{}
	for _, songId := range request.Data.Relationships.Songs.Data {
		songUuid, err := uuid.Parse(songId.Id)
		if err != nil {
			return songIds, err
		}
		songIds = append(songIds, songUuid)
	}
	return songIds, nil
}
