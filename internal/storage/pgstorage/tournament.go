package pgstorage

import (
	"fmt"

	"github.com/Masterminds/squirrel"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"github.com/tengenatari/web-referee/internal/models"
	"golang.org/x/net/context"
)

func (storage *WebRefereeStorage) getTournamentShard(tournament *Tournament) (uint32, error) {
	if tournament == nil || tournament.Id == uuid.Nil {
		return 0, errors.New("tournament is nil")
	}
	fmt.Print((tournament.Id.ID())%storage.shards + 1)
	return (tournament.Id.ID())%storage.shards + 1, nil
}

func (storage *WebRefereeStorage) CreateTournament(ctx context.Context, tournament *models.Tournament) (uuid.UUID, error) {
	tournamentUuid, err := uuid.NewV7()

	if err != nil {
		return uuid.Nil, errors.Wrap(err, "failed to create tournament uuid")
	}

	tournamentMapped := Tournament{
		Id:   tournamentUuid,
		Name: tournament.Name,
		Date: tournament.Date,
	}

	shard, err := storage.getTournamentShard(&tournamentMapped)

	if err != nil {
		return uuid.Nil, errors.Wrap(err, "failed to get tournament shard")
	}

	q := squirrel.Insert(fmt.Sprintf("schema_%03d.%s", shard, TournamentTable)).
		Columns(TournamentId, TournamentName, TournamentDate).Values(tournamentMapped.Id, tournamentMapped.Name, tournamentMapped.Date).
		PlaceholderFormat(squirrel.Dollar)
	err = storage.execQuery(ctx, q)
	if err != nil {
		return uuid.Nil, errors.Wrap(err, "Failed to create tournament")
	}
	return tournamentUuid, nil
}
