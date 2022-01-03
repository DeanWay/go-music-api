package postgres

import (
	"go-music-api/internal/pkg/models"
	"go-music-api/internal/pkg/payloads"

	psqlStorage "go-music-api/internal/pkg/storage/postgres"

	"github.com/google/uuid"
)

type SongPostgresRepo struct {
	Store psqlStorage.PostgresStorage
}

func (repo SongPostgresRepo) AddSong(
	albumUuid uuid.UUID,
	attrs payloads.SongAttributes,
) models.Song {
	newSong := models.NewSong(albumUuid, attrs)
	stmt, err := repo.Store.DB.Prepare(`
		insert into song
		(uuid, albumUuid, title, durationSeconds)
		values
		($1, $2, $3, $4)
	`)
	if err != nil {
		panic(err)
	}
	_, err = stmt.Exec(
		newSong.Uuid.String(),
		newSong.AlbumUuid,
		newSong.Title,
		newSong.DurationSeconds,
	)
	if err != nil {
		panic(err)
	}
	return newSong
}

func (repo SongPostgresRepo) FindSongById(id string) (models.Song, error) {
	var song models.Song
	err := repo.Store.DB.QueryRow(
		`select
			uuid,
			albumUuid,
			title,
			durationSeconds
		from
			song
		where
			uuid = $1
		limit 1
		`,
		id,
	).Scan(
		&song.Uuid,
		&song.AlbumUuid,
		&song.Title,
		&song.DurationSeconds,
	)
	if err != nil {
		return song, err
	}
	return song, nil
}
