package memory

import (
	"errors"
	"go-music-api/internal/pkg/adapter/storage"
	"strings"
)

type MemoryStorage map[string]string

var _ storage.KeyValueStorage = (*MemoryStorage)(nil)

func (store MemoryStorage) Insert(key string, value string) error {
	store[key] = value
	return nil
}

func (store MemoryStorage) Get(key string) (string, error) {
	val, ok := store[key]
	if ok {
		return val, nil
	} else {
		return "", errors.New("Not found")
	}
}

func (store MemoryStorage) List(prefix string) []string {
	values := []string{}
	for key, value := range store {
		if strings.HasPrefix(key, prefix) {
			values = append(values, value)
		}
	}
	return values
}

func (store MemoryStorage) Delete(key string) error {
	_, err := store.Get(key)
	if err != nil {
		return err
	}
	delete(store, key)
	return nil
}
