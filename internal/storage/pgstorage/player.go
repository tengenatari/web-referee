package pgstorage

import (
	"fmt"

	"github.com/Masterminds/squirrel"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"github.com/tengenatari/web-referee/internal/models"
	"golang.org/x/net/context"
)

func (storage *WebRefereeStorage) CreatePlayer(ctx context.Context, player *models.Player) (uuid.UUID, error) {
	playerUUID, err := uuid.NewV7()
	if err != nil {
		return uuid.Nil, errors.Wrap(err, "CreatePlayer")
	}

	playerMapped := Player{
		Id:           playerUUID,
		TournamentId: player.TournamentId,
		UserId:       player.UserId,
		MacMahon:     player.MacMahon,
	}

	shard, err := storage.getTournamentShard(&Tournament{Id: player.TournamentId})

	if err != nil {
		return uuid.Nil, errors.Wrap(err, "failed to get tournament shard")
	}

	q := squirrel.Insert(fmt.Sprintf("schema_%03d.%s", shard, PlayerTable)).
		Columns(PlayerId, PlayerTournamentId, PlayerUserId, PlayerMacMahon).
		Values(playerMapped.Id, playerMapped.TournamentId, playerMapped.UserId, playerMapped.MacMahon).
		PlaceholderFormat(squirrel.Dollar)
	err = storage.execQuery(ctx, q)
	if err != nil {
		return uuid.Nil, errors.Wrap(err, "Failed to create tournament")
	}
	return playerUUID, nil
}

func (storage *WebRefereeStorage) GetPlayersByTournamentId(ctx context.Context, tournamentId uuid.UUID) ([]*models.Player, error) {

	shard, err := storage.getTournamentShard(&Tournament{Id: tournamentId})
	if err != nil {
		return nil, errors.Wrap(err, "Failed to get tournament shard")
	}
	q := squirrel.Select(PlayerId, PlayerTournamentId, PlayerUserId).From(fmt.Sprintf("schema_%03d.%s", shard, PlayerTable)).
		PlaceholderFormat(squirrel.Dollar).
		Where(squirrel.Eq{"tournament_id": tournamentId.String()})

	rows, err := storage.query(ctx, q)

	if err != nil {
		return nil, errors.Wrap(err, "Failed to get players")
	}

	defer rows.Close()

	var players []*models.Player
	for rows.Next() {
		var player models.Player
		if err := rows.Scan(&player.Id, &player.TournamentId, &player.UserId); err != nil {
			return nil, errors.Wrap(err, "failed to scan row")
		}
		players = append(players, &player)
	}
	return players, nil
}
