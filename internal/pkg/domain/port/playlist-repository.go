package port

import "go-music-api/internal/pkg/domain/entity"

type HasAddPlaylist interface {
	AddPlaylist(entity.Playlist) error
}

type PlaylistRepository interface {
	HasAddPlaylist
}
