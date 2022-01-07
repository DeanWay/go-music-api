package redisstorage

import (
	"fmt"
	"go-music-api/internal/pkg/adapter/storage"

	"github.com/go-redis/redis"
)

type RedisStorage struct {
	Client *redis.Client
}

var _ storage.KeyValueStorage = (*RedisStorage)(nil)

func (store RedisStorage) Insert(key string, value string) error {
	return store.Client.Set(key, value, 0).Err()
}

func (store RedisStorage) Get(key string) (string, error) {
	cmd := store.Client.Get(key)
	return cmd.Val(), cmd.Err()
}

func (store RedisStorage) List(prefix string) []string {
	searchKey := fmt.Sprintf("%s:*", prefix)
	allKeys := store.Client.Keys(searchKey).Val()
	result := make([]string, len(allKeys), len(allKeys))
	for i, key := range allKeys {
		result[i], _ = store.Get(key)
	}
	return result
}

func (store RedisStorage) Delete(key string) error {
	return store.Client.Del(key).Err()
}
