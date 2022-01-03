package postgres

import (
	"go-music-api/internal/pkg/models"
	"go-music-api/internal/pkg/payloads"
	"go-music-api/internal/pkg/repository"
	psqlStorage "go-music-api/internal/pkg/storage/postgres"
	"strings"

	"github.com/google/uuid"
)

type AlbumPostgresRepo struct {
	Store    psqlStorage.PostgresStorage
	SongRepo SongPostgresRepo
}

func (repo AlbumPostgresRepo) GetAllAlbums() []models.Album {
	albums := []models.Album{}
	rows, err := repo.Store.DB.Query(`
		select
			uuid,
			title,
			artist,
			price,
			createdAt
		from
			album
	`)
	if err != nil {
		panic(err)
	}
	for rows.Next() {
		var album models.Album
		err = rows.Scan(
			&album.Uuid,
			&album.Title,
			&album.Artist,
			&album.Price,
			&album.CreatedAt,
		)
		if err != nil {
			panic(err)
		}
		album.SongUuids = repo.getSongUuids(album.Uuid)
		albums = append(albums, album)
	}
	if err = rows.Err(); err != nil {
		panic(err)
	}
	return albums
}

func (repo AlbumPostgresRepo) FindAlbumById(id string) (models.Album, error) {
	var album models.Album
	err := repo.Store.DB.QueryRow(
		`select
			uuid,
			title,
			artist,
			price,
			createdAt
		from
			album
		where
			uuid = $1
		limit 1
		`,
		id,
	).Scan(
		&album.Uuid,
		&album.Title,
		&album.Artist,
		&album.Price,
		&album.CreatedAt,
	)
	if err != nil {
		return album, err
	}
	album.SongUuids = repo.getSongUuids(album.Uuid)
	return album, nil
}

func (repo AlbumPostgresRepo) AddAlbum(
	attrs payloads.AlbumAttributes,
	songAttrs []payloads.SongAttributes,
) models.Album {
	newAlbum := models.NewAlbum(attrs)
	stmt, err := repo.Store.DB.Prepare(`
		insert into album
		(uuid, title, artist, price, createdAt)
		values
		($1, $2, $3, $4, $5)
	`)
	if err != nil {
		panic(err)
	}
	_, err = stmt.Exec(
		newAlbum.Uuid.String(),
		newAlbum.Title,
		newAlbum.Artist,
		newAlbum.Price,
		newAlbum.CreatedAt,
	)
	if err != nil {
		panic(err)
	}
	songUuids := make([]uuid.UUID, len(songAttrs), len(songAttrs))
	for i, song := range songAttrs {
		newSong := repo.SongRepo.AddSong(newAlbum.Uuid, song)
		songUuids[i] = newSong.Uuid
	}
	newAlbum.SongUuids = songUuids
	return newAlbum
}

func (repo AlbumPostgresRepo) SearchAlbums(
	params repository.AlbumSearchParams,
) []models.Album {
	query := `
	select
		uuid,
		title,
		artist,
		price,
		createdAt
	from
		album
	where
	`
	clauses := []string{}
	args := []interface{}{}
	if params.PriceHigh != nil {
		clauses = append(clauses, "album.price < $1")
		args = append(args, *&params.PriceHigh)
	} else {
		clauses = append(clauses, "$1")
		args = append(args, "true")
	}
	if params.PriceLow != nil {
		clauses = append(clauses, "album.price > $2")
		args = append(args, *&params.PriceLow)
	} else {
		clauses = append(clauses, "$2")
		args = append(args, "true")
	}
	if params.Title != nil {
		clauses = append(clauses, "album.title = $3")
		args = append(args, *&params.Title)
	} else {
		clauses = append(clauses, "$3")
		args = append(args, "true")
	}
	if params.Artist != nil {
		clauses = append(clauses, "album.artist = $4")
		args = append(args, *&params.Artist)
	} else {
		clauses = append(clauses, "$4")
		args = append(args, "true")
	}
	query += strings.Join(clauses, " and ")
	rows, err := repo.Store.DB.Query(query, args...)
	if err != nil {
		panic(err)
	}
	albums := []models.Album{}
	for rows.Next() {
		var album models.Album
		err = rows.Scan(
			&album.Uuid,
			&album.Title,
			&album.Artist,
			&album.Price,
			&album.CreatedAt,
		)
		if err != nil {
			panic(err)
		}
		album.SongUuids = repo.getSongUuids(album.Uuid)
		albums = append(albums, album)
	}
	if err = rows.Err(); err != nil {
		panic(err)
	}
	return albums
}

func (repo AlbumPostgresRepo) getSongUuids(albumUuid uuid.UUID) []uuid.UUID {
	rows, err := repo.Store.DB.Query(
		`select
			uuid
		from
			song
		where
			song.albumUuid = $1
		`,
		albumUuid,
	)
	if err != nil {
		panic(err)
	}
	result := []uuid.UUID{}
	for rows.Next() {
		var songUuid uuid.UUID
		err = rows.Scan(&songUuid)
		if err != nil {
			panic(err)
		}
		result = append(result, songUuid)
	}
	if err = rows.Err(); err != nil {
		panic(err)
	}
	return result
}
