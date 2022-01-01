package memory

import "errors"

type MemoryStorage map[string]interface{}

func (store MemoryStorage) Insert(id string, thing interface{}) error {
	store[id] = thing
	return nil
}

func (store MemoryStorage) GetById(id string) (interface{}, error) {
	val, ok := store[id]
	if ok {
		return val, nil
	} else {
		return nil, errors.New("Not found")
	}
}

func (store MemoryStorage) List() []interface{} {
	values := make([]interface{}, len(store), len(store))
	i := 0
	for _, value := range store {
		values[i] = value
		i++
	}
	return values
}

func (store MemoryStorage) Delete(id string) error {
	_, err := store.GetById(id)
	if err != nil {
		return err
	}
	delete(store, id)
	return nil
}
