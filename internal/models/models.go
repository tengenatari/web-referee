package models

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	Id        uuid.UUID
	Name      string
	Email     string
	TigrId    string
	Rating    int64
	CreatedAt time.Time
}

type Tournament struct {
	Id        uuid.UUID
	Name      string
	Date      time.Time
	CreatedAt time.Time
}

type Player struct {
	Id           uuid.UUID
	TournamentId uuid.UUID
	UserId       uuid.UUID
	MacMahon     int64
}
