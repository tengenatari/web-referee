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

const (
	UserTable        = "users"
	UserColumnRating = "rating"
	UserColumnName   = "name"
	UserTigrId       = "tigr_id"
	UserCreatedAt    = "created_at"
)
