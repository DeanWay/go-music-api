package redisstorage

import (
	"fmt"

	"github.com/go-redis/redis"
)

type RedisStorage struct {
	CollectionName string
	Client         *redis.Client
}

func (store RedisStorage) Insert(key string, value string) error {
	return store.Client.Set(store.withPrefix(key), value, 0).Err()
}

func (store RedisStorage) Get(key string) (string, error) {
	return store.getWithoutPrefix(store.withPrefix(key))
}

func (store RedisStorage) getWithoutPrefix(key string) (string, error) {
	cmd := store.Client.Get(key)
	return cmd.Val(), cmd.Err()
}

func (store RedisStorage) List() []string {
	allKeys := store.Client.Keys(store.withPrefix("*")).Val()
	result := make([]string, len(allKeys), len(allKeys))
	for i, key := range allKeys {
		result[i], _ = store.getWithoutPrefix(key)
	}
	return result
}

func (store RedisStorage) Delete(key string) error {
	return store.Client.Del(store.withPrefix(key)).Err()
}

func (store RedisStorage) withPrefix(key string) string {
	return fmt.Sprintf("%s:%s", store.CollectionName, key)
}
