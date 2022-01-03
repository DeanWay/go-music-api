package postgres

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type PostgresStorage struct {
	DB *sql.DB
}

type ConnectionParams struct {
	Username string
	Password string
	Host     string
	Port     string
	Database string
}

func New(
	params ConnectionParams,
) PostgresStorage {
	db, err := sql.Open(
		"postgres",
		fmt.Sprintf(
			"postgresql://%s:%s@%s:%s/%s?sslmode=disable",
			params.Username,
			params.Password,
			params.Host,
			params.Port,
			params.Database,
		),
	)
	if err != nil {
		panic(err)
	}
	return PostgresStorage{
		DB: db,
	}
}
