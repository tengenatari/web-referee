package models

import "time"

type User struct {
	Id        int64
	Name      string
	Email     string
	TigrId    string
	Rating    int64
	CreatedAt time.Time
}
