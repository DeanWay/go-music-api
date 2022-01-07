package postgres

import (
	"database/sql"
	"fmt"
	"go-music-api/internal/pkg/adapter/storage"
	"log"
	"os"

	"github.com/doug-martin/goqu/v9"
	_ "github.com/doug-martin/goqu/v9/dialect/postgres"
	_ "github.com/lib/pq"
)

type PostgresStorage struct {
	db *goqu.Database
}

var _ storage.SqlStorage = (*PostgresStorage)(nil)

func New(
	config Config,
) PostgresStorage {
	connectionParams := config.ConnectionParams
	db, err := sql.Open(
		"postgres",
		connectionParams.String(),
	)
	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	goqu.SetDefaultPrepared(true)
	goquDb := goqu.New("postgres", db)
	if config.LogSql {
		logger := log.Default()
		logger.SetOutput(os.Stdout)
		goquDb.Logger(logger)
	}
	return PostgresStorage{
		db: goquDb,
	}
}

func (store PostgresStorage) DB() *goqu.Database {
	return store.db
}

type Config struct {
	ConnectionParams ConnectionParams
	LogSql           bool
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
