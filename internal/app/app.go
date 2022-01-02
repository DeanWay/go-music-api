package app

import (
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"

	"go-music-api/internal/pkg/repository"
	"go-music-api/internal/pkg/repository/keyvalue"
	"go-music-api/internal/pkg/routes"
	"go-music-api/internal/pkg/storage"
	"go-music-api/internal/pkg/storage/memory"
	redisStorage "go-music-api/internal/pkg/storage/redis"
)

func App() *gin.Engine {
	appEngine := gin.Default()
	albumRouter := routes.AlbumRouter{
		AlbumRepository: initAlbumRepo(),
	}

	appEngine.GET("/album/:id", albumRouter.GetAlbumByID)
	appEngine.GET("/albums", albumRouter.ListAlbums)
	appEngine.POST("/albums", albumRouter.PostAlbums)
	appEngine.GET("/albums/search", albumRouter.SearchAlbums)
	return appEngine
}

func initAlbumRepo() repository.AlbumRepository {
	albumRepo := keyvalue.AlbumKeyValueRepo{
		Store: redisStore("album"),
		SongRepo: keyvalue.SongKeyValueRepo{
			Store: redisStore("song"),
		},
	}
	return albumRepo
}

func memoryStore() storage.KeyValueStorage {
	return memory.MemoryStorage{}
}

func redisStore(collectionName string) storage.KeyValueStorage {
	return redisStorage.RedisStorage{
		CollectionName: collectionName,
		Client: redis.NewClient(&redis.Options{
			Addr:     "localhost:6379",
			Password: "", // no password set
			DB:       0,  // use default DB
		}),
	}
}
