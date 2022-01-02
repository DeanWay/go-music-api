package storage

type KeyValueStorage interface {
	Insert(id string, thing interface{}) error
	GetById(id string) (interface{}, error)
	List() []interface{}
	Delete(id string) error
}
