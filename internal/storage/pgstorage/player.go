package pgstorage

import (
	"fmt"

	"github.com/Masterminds/squirrel"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"github.com/tengenatari/web-referee/internal/models"
	"golang.org/x/net/context"
)

func (storage *WebRefereeStorage) CreatePlayer(ctx context.Context, player *models.Player) error {
	playerUUID, err := uuid.NewV7()
	if err != nil {
		return errors.Wrap(err, "CreatePlayer")
	}

	playerMapped := Player{
		Id:           playerUUID,
		TournamentId: player.TournamentId,
		UserId:       player.UserId,
		MacMahon:     player.MacMahon,
	}

	shard, err := storage.getTournamentShard(&Tournament{Id: player.TournamentId})

	if err != nil {
		return errors.Wrap(err, "failed to get tournament shard")
	}

	q := squirrel.Insert(fmt.Sprintf("schema_%03d.%s", shard, PlayerTable)).
		Columns(PlayerId, PlayerTournamentId, PlayerUserId, PlayerMacMahon).
		Values(playerMapped.Id, playerMapped.TournamentId, playerMapped.UserId, playerMapped.MacMahon).
		PlaceholderFormat(squirrel.Dollar)
	err = storage.execQuery(ctx, q)
	if err != nil {
		return errors.Wrap(err, "Failed to create tournament")
	}
	return nil
}
