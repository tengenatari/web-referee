package pgstorage

import (
	"fmt"

	"github.com/Masterminds/squirrel"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"github.com/samber/lo"
	"github.com/tengenatari/web-referee/internal/models"
	"golang.org/x/net/context"
)

func (storage *WebRefereeStorage) CreatePairing(ctx context.Context, games []*models.Game, tournamentUUID uuid.UUID) error {

	gamesMapped := lo.Map(
		games,
		func(game *models.Game, _ int) *Game {
			id, err := uuid.NewV7()
			if err != nil {
				return nil
			}
			return &Game{
				Id:      id,
				White:   game.White,
				Black:   game.Black,
				TourNum: game.TourNum,
			}
		},
	)
	shard, err := storage.getTournamentShard(&Tournament{Id: tournamentUUID})
	q := squirrel.Insert(fmt.Sprintf("schema_%03d.%s", shard, GameTable)).Columns(GameId, GameWhite, GameBlack, GameTourNum).
		PlaceholderFormat(squirrel.Dollar)
	for _, gameMapped := range gamesMapped {
		q = q.Values(gameMapped.Id, gameMapped.White, gameMapped.Black, gameMapped.TourNum)
	}
	err = storage.execQuery(ctx, q)
	if err != nil {
		return errors.Wrap(err, "create pairing games")
	}
	return nil
}

func (storage *WebRefereeStorage) GetPairing(ctx context.Context, tournamentId uuid.UUID) ([]*models.Game, error) {
	shard, err := storage.getTournamentShard(&Tournament{Id: tournamentId})
	if err != nil {
		return nil, errors.Wrap(err, "get tournament shard")
	}
	playerScheme := fmt.Sprintf("schema_%03d.%s", shard, PlayerTable)
	gameScheme := fmt.Sprintf("schema_%03d.%s", shard, GameTable)
	gameSchemeColumnId := fmt.Sprintf("%s.%s", gameScheme, GameId)
	gameSchemeColumnWhite := fmt.Sprintf("%s.%s", gameScheme, GameWhite)
	gameSchemeColumnBlack := fmt.Sprintf("%s.%s", gameScheme, GameBlack)
	q := squirrel.Select(gameSchemeColumnId, gameSchemeColumnWhite, gameSchemeColumnBlack).From(playerScheme).
		InnerJoin(gameScheme).JoinClause(fmt.Sprintf("ON %s.%s = %s.%s", playerScheme, PlayerId, gameScheme, GameWhite)).
		PlaceholderFormat(squirrel.Dollar).
		Where(squirrel.Eq{fmt.Sprintf("%s.%s", playerScheme, PlayerTournamentId): tournamentId.String()})

	rows, err := storage.query(ctx, q)
	if err != nil {
		return nil, errors.Wrap(err, "error during execution query")
	}

	defer rows.Close()

	var games []*models.Game

	for rows.Next() {
		var game models.Game
		if err := rows.Scan(&game.Id, &game.White, &game.Black); err != nil {
			return nil, errors.Wrap(err, "failed to scan row")
		}
		games = append(games, &game)
	}

	return games, nil
}
