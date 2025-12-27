package redisstorage

import (
	"encoding/json"
	"fmt"

	"github.com/google/uuid"
	"github.com/samber/lo"
	"github.com/tengenatari/web-referee/internal/models"
	"golang.org/x/net/context"
)

func (s *RedisStorage) SavePairing(ctx context.Context, tournamentID uuid.UUID, games []*models.Game) error {
	if len(games) == 0 {
		return fmt.Errorf("no pairs to save")
	}

	roundPairings := Pairings{
		Pairings: lo.Map(games, func(game *models.Game, _ int) GamePair {
			return GamePair{GameID: tournamentID, White: game.White, Black: game.Black}
		}),
	}

	data, err := json.Marshal(roundPairings)
	if err != nil {
		return fmt.Errorf("failed to marshal pairings: %w", err)
	}

	key := getKey(tournamentID)

	err = s.client.Set(ctx, key, data, s.ttl).Err()
	if err != nil {
		return fmt.Errorf("failed to save to Redis: %w", err)
	}
	fmt.Println("saved pairings to Redis to", key)
	return nil
}
func getKey(tournamentID uuid.UUID) string {
	return fmt.Sprintf("t:%s", tournamentID.String())
}
func (s *RedisStorage) GetPairing(ctx context.Context, tournamentID uuid.UUID) ([]*models.Game, error) {
	key := getKey(tournamentID)
	data, err := s.client.Get(ctx, key).Result()
	if err != nil {
		return nil, fmt.Errorf("failed to get: %w", err)
	}

	var pairings Pairings
	err = json.Unmarshal([]byte(data), &pairings)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	fmt.Println("Redis get", key)
	games := lo.Map(pairings.Pairings, func(game GamePair, _ int) *models.Game {
		return &models.Game{Id: game.GameID, White: game.White, Black: game.Black}
	})
	return games, nil
}
