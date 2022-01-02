package storage

type KeyValueStorage interface {
	Insert(key string, value string) error
	Get(key string) (string, error)
	List() []string
	Delete(key string) error
}
