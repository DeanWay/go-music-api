package postgres

import (
	psqlStorage "go-music-api/internal/pkg/adapter/storage/postgres"
	"go-music-api/internal/pkg/domain/entity"
	"go-music-api/internal/pkg/domain/failure"
	"go-music-api/internal/pkg/domain/port"
	"time"

	"github.com/doug-martin/goqu/v9"
	"github.com/google/uuid"
)

type AlbumPostgresRepository struct {
	Store psqlStorage.PostgresStorage
}

var _ port.AlbumRepository = (*AlbumPostgresRepository)(nil)

func (repo AlbumPostgresRepository) ListAlbums() ([]entity.Album, error) {
	query := repo.Store.DB().From("album")
	albums, err := queryToAlbums(query)
	if err != nil {
		return albums, err
	}
	return repo.populateSongIds(albums), nil
}

func (repo AlbumPostgresRepository) GetAlbumById(
	id string,
) (entity.Album, error) {
	var row albumRow
	found, err := repo.Store.DB().From(
		"album",
	).Where(
		goqu.Ex{"uuid": id},
	).ScanStruct(&row)
	album := row.toEntity()
	if !found {
		return album, failure.ErrNotFound
	}
	if err != nil {
		return album, err
	}
	return repo.populateSongIds([]entity.Album{album})[0], nil
}

func (repo AlbumPostgresRepository) AddAlbum(
	album entity.Album,
) error {
	row := albumRow{
		AlbumId:   album.Id,
		Title:     album.Title,
		Artist:    album.Artist,
		Price:     album.Price,
		CreatedAt: album.CreatedAt,
	}
	_, err := repo.Store.DB().Insert("album").Rows(
		row,
	).Executor().Exec()
	if err != nil {
		return err
	}
	albumSongRows := []albumSongRow{}
	for _, songId := range album.SongIds {
		albumSongRows = append(albumSongRows, albumSongRow{
			AlbumId: album.Id,
			SongId:  songId,
		})
	}
	_, err = repo.Store.DB().Insert("album_song").Rows(
		albumSongRows,
	).Executor().Exec()
	return err
}

func (repo AlbumPostgresRepository) SearchAlbums(
	params port.AlbumSearchParams,
) ([]entity.Album, error) {
	query := repo.Store.DB().From("album")
	if params.Title != nil {
		query = query.Where(goqu.Ex{"title": *params.Title})
	}
	if params.Artist != nil {
		query = query.Where(goqu.Ex{"artist": *params.Artist})
	}
	if params.PriceHigh != nil {
		query = query.Where(goqu.Ex{"price": goqu.Op{"lt": *params.PriceHigh}})
	}
	if params.PriceLow != nil {
		query = query.Where(goqu.Ex{"price": goqu.Op{"gt": *params.PriceLow}})
	}
	albums, err := queryToAlbums(query)
	if err != nil {
		return albums, nil
	}
	return repo.populateSongIds(albums), nil
}

func (repo AlbumPostgresRepository) populateSongIds(
	albums []entity.Album,
) []entity.Album {
	albumIds := []uuid.UUID{}
	for _, album := range albums {
		albumIds = append(albumIds, album.Id)
	}
	rows := []albumSongRow{}
	albumIdToSongIds := map[uuid.UUID][]uuid.UUID{}
	repo.Store.DB().From("album_song").Where(
		goqu.Ex{"album_uuid": albumIds},
	).ScanStructs(&rows)
	for _, row := range rows {
		albumIdToSongIds[row.AlbumId] = append(
			albumIdToSongIds[row.AlbumId],
			row.SongId,
		)
	}
	result := []entity.Album{}
	for _, album := range albums {
		album.SongIds = albumIdToSongIds[album.Id]
		result = append(result, album)
	}
	return result
}

type albumRow struct {
	AlbumId   uuid.UUID `db:"uuid"`
	Title     string    `db:"title"`
	Artist    string    `db:"artist"`
	Price     float64   `db:"price"`
	CreatedAt time.Time `db:"created_at"`
}

type albumSongRow struct {
	AlbumId uuid.UUID `db:"album_uuid"`
	SongId  uuid.UUID `db:"song_uuid"`
}

func (row albumRow) toEntity() entity.Album {
	return entity.Album{
		Id:        row.AlbumId,
		Title:     row.Title,
		Artist:    row.Artist,
		Price:     row.Price,
		CreatedAt: row.CreatedAt,
	}
}

func queryToAlbums(query *goqu.SelectDataset) ([]entity.Album, error) {
	rows := []albumRow{}
	err := query.ScanStructs(&rows)
	albums := []entity.Album{}
	if err != nil {
		return albums, err
	}
	for _, row := range rows {
		albums = append(albums, row.toEntity())
	}
	return albums, nil
}
