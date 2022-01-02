package memory

import "errors"

type MemoryStorage map[string]string

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

func (store MemoryStorage) List() []string {
	values := make([]string, len(store), len(store))
	i := 0
	for _, value := range store {
		values[i] = value
		i++
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
