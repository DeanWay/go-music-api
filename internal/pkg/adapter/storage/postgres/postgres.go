package postgres

import (
	"database/sql"
	"fmt"
	"go-music-api/internal/pkg/adapter/storage"

	"github.com/doug-martin/goqu/v9"
	_ "github.com/doug-martin/goqu/v9/dialect/postgres"
	_ "github.com/lib/pq"
)

type PostgresStorage struct {
	db *goqu.Database
}

var _ storage.SqlStorage = (*PostgresStorage)(nil)

func New(
	params ConnectionParams,
) PostgresStorage {

	fmt.Println(params.String())
	db, err := sql.Open(
		"postgres",
		params.String(),
	)
	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	return PostgresStorage{
		db: goqu.New("postgres", db),
	}
}

func (store PostgresStorage) DB() *goqu.Database {
	return store.db
}

type ConnectionParams struct {
	Username string
	Password string
	Host     string
	Port     string
	Database string
}

func (params ConnectionParams) String() string {
	return fmt.Sprintf(
		"postgresql://%s:%s@%s:%s/%s?sslmode=disable",
		params.Username,
		params.Password,
		params.Host,
		params.Port,
		params.Database,
	)
}
