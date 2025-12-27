package models

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	Id        uuid.UUID
	Name      string
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

type Game struct {
	Id          uuid.UUID
	White       uuid.UUID
	Black       uuid.UUID
	TourNum     int64
	GameUrl     string
	ResultBlack int64
	ResultWhite int64
}
