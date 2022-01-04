package api_test

import (
	"go-music-api/internal/app"
	"go-music-api/internal/pkg/models"
	"go-music-api/test/fakes/repo"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestGetAllAlbums(t *testing.T) {
	server := app.App(&app.Deps{
		AlbumRepo: repo.FakeAlbumRepo{
			Albums: []models.Album{
				{
					Uuid:   uuid.MustParse("aaaaaaaa-aaaa-aaaa-aaaa-aaaaaaaaaaaa"),
					Title:  "The album",
					Artist: "Queen",
					Price:  1.2,
					SongUuids: []uuid.UUID{
						uuid.MustParse("bbbbbbbb-bbbb-bbbb-bbbb-bbbbbbbbbbbb"),
					},
				},
			},
		},
		SongRepo: repo.FakeSongRepo{},
	})

	response := httptest.NewRecorder()
	request, _ := http.NewRequest("GET", "/albums", nil)
	server.ServeHTTP(response, request)

	expectedJSON := `
	{
		"data": [
			{
				"id": "aaaaaaaa-aaaa-aaaa-aaaa-aaaaaaaaaaaa",
				"type": "album",
				"attributes": {
					"title": "The album",
					"artist": "Queen",
					"price": 1.2
				},
				"relationships": {
					"songs": {
						"data": [
							{
								"id": "bbbbbbbb-bbbb-bbbb-bbbb-bbbbbbbbbbbb",
								"type": "song"
							}
						]
					}
				}
			}
		],
		"meta": {"count": 1}
	}
	`
	assert.Equal(t, http.StatusOK, response.Result().StatusCode)
	assert.JSONEq(t, response.Body.String(), expectedJSON)
}
