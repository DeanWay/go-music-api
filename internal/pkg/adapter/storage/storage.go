package storage

import "github.com/doug-martin/goqu/v9"

type KeyValueStorage interface {
	Insert(key string, value string) error
	Get(key string) (string, error)
	List() []string
	Delete(key string) error
}

type SqlStorage interface {
	DB() *goqu.Database
}
