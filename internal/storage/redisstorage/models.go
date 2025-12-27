package redisstorage

import (
	"github.com/google/uuid"
)

type GamePair struct {
	GameID uuid.UUID `json:"game_id"`
	White  uuid.UUID `json:"white"`
	Black  uuid.UUID `json:"black"`
}

type Pairings struct {
	Pairings []GamePair `json:"pairings"`
}
