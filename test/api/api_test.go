package api_test

import (
	"encoding/json"
	"go-music-api/config"
	"go-music-api/internal/app"
	"go-music-api/internal/pkg/domain/entity"
	"go-music-api/internal/pkg/http/payloads"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func PostAlbumsTestCase(deps *config.Deps) func(*testing.T) {
	return func(t *testing.T) {
		server := app.RestApiApp(deps)
		response := httptest.NewRecorder()
		request, _ := http.NewRequest(
			"POST",
			"/albums",
			strings.NewReader(`
			{
				"data": {
					"attributes": {
						"title": "The album",
						"artist": "Queen",
						"price": 1.2
					}
				}
			}
			`),
		)
		server.ServeHTTP(response, request)
		assert.Equal(t, http.StatusCreated, response.Result().StatusCode)
		var responseBody payloads.Response
		err := json.Unmarshal(response.Body.Bytes(), &responseBody)
		assert.Nil(t, err)
		createdAlbum, err := deps.AlbumRepository.GetAlbumById(responseBody.Data.Id)
		assert.Nil(t, err)
		assert.Equal(t, "The album", createdAlbum.Title)
		assert.Equal(t, "Queen", createdAlbum.Artist)
		assert.Equal(t, 1.2, createdAlbum.Price)
	}
}

func GetAlbumsTestCase(deps *config.Deps) func(*testing.T) {
	return func(t *testing.T) {
		server := app.RestApiApp(deps)

		deps.AlbumRepository.AddAlbum(
			entity.Album{
				Id:     uuid.MustParse("aaaaaaaa-aaaa-aaaa-aaaa-aaaaaaaaaaaa"),
				Title:  "The album",
				Artist: "Queen",
				Price:  1.2,
				SongIds: []uuid.UUID{
					uuid.MustParse("bbbbbbbb-bbbb-bbbb-bbbb-bbbbbbbbbbbb"),
				},
			},
		)

		response := httptest.NewRecorder()
		request, _ := http.NewRequest("GET", "/albums", nil)
		server.ServeHTTP(response, request)

		assert.Equal(t, http.StatusOK, response.Result().StatusCode)
		var responseBody payloads.ListResponse
		err := json.Unmarshal(response.Body.Bytes(), &responseBody)
		assert.Nil(t, err)
		var responseAttributes payloads.AlbumAttributes
		attrsJson, err := json.Marshal(responseBody.Data[0].Attributes)
		assert.Nil(t, err)
		err = json.Unmarshal(attrsJson, &responseAttributes)
		assert.Nil(t, err)
	}
}

func TestAlbumsApi(t *testing.T) {
	for _, deps := range createDeps() {
		t.Run("create album", PostAlbumsTestCase(deps))
		t.Run("get albums", GetAlbumsTestCase(deps))
	}
}

func createDeps() []*config.Deps {
	inMemoryDeps := config.InMemoryDeps()
	postgresDeps := config.PostgresDeps()
	return []*config.Deps{
		&inMemoryDeps,
		&postgresDeps,
	}
}
