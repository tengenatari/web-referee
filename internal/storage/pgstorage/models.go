package pgstorage

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	Id            uuid.UUID `db:"id"`
	Name          string    `db:"name"`
	Rating        int64     `db:"rating"`
	TigrId        string    `db:"tigr_id"`
	UserCreatedAt time.Time `db:"created_at"`
}

type Tournament struct {
	Id                  uuid.UUID `db:"id"`
	Name                string    `db:"name"`
	TournamentCreatedAt time.Time `db:"created_at"`
	Date                time.Time `db:"date"`
}

type Player struct {
	Id           uuid.UUID `db:"id"`
	TournamentId uuid.UUID `db:"tournament_id"`
	MacMahon     int64     `db:"name"`
	UserId       uuid.UUID `db:"user_id"`
}

const (
	UserTable        = "users"
	UserId           = "id"
	UserColumnRating = "rating"
	UserColumnName   = "name"
	UserTigrId       = "tigr_id"
	UserCreatedAt    = "created_at"

	TournamentTable     = "tournaments"
	TournamentId        = "id"
	TournamentName      = "name"
	TournamentCreatedAt = "created_at"
	TournamentDate      = "date"

	PlayerTable        = "players"
	PlayerId           = "id"
	PlayerUserId       = "user_id"
	PlayerMacMahon     = "mac_mahon"
	PlayerTournamentId = "tournament_id"
)
