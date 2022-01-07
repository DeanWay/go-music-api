package postgres

import (
	psqlStorage "go-music-api/internal/pkg/adapter/storage/postgres"
	"go-music-api/internal/pkg/domain/entity"
	"go-music-api/internal/pkg/domain/failure"
	"go-music-api/internal/pkg/domain/port"

	"github.com/doug-martin/goqu/v9"
	"github.com/google/uuid"
)

type SongPostgresRepository struct {
	Store psqlStorage.PostgresStorage
}

var _ port.SongRepository = (*SongPostgresRepository)(nil)

func (repo SongPostgresRepository) AddSong(
	newSong entity.Song,
) error {
	row := songRow{
		Id:              newSong.Id,
		Title:           newSong.Title,
		Artist:          newSong.Artist,
		DurationSeconds: newSong.DurationSeconds,
	}
	_, err := repo.Store.DB().Insert("song").Rows(row).Executor().Exec()
	return err
}

func (repo SongPostgresRepository) GetSongById(
	id string,
) (entity.Song, error) {
	var row songRow
	found, err := repo.Store.DB().From(
		"song",
	).Where(
		goqu.Ex{"uuid": id},
	).ScanStruct(&row)
	song := row.toEntity()
	if !found {
		return song, failure.ErrNotFound
	}
	if err != nil {
		return song, err
	}
	return song, nil
}

type songRow struct {
	Id              uuid.UUID `db:"uuid"`
	Title           string    `db:"title"`
	Artist          string    `db:"artist"`
	DurationSeconds uint      `db:"duration_seconds"`
}

func (row songRow) toEntity() entity.Song {
	return entity.Song{
		Id:              row.Id,
		Title:           row.Title,
		Artist:          row.Artist,
		DurationSeconds: row.DurationSeconds,
	}
}
