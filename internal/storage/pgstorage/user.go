package pgstorage

import (
	"fmt"

	"github.com/Masterminds/squirrel"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"github.com/tengenatari/web-referee/internal/models"
	"golang.org/x/net/context"
)

func (storage *PGstorage) getUserShard(user *User) (uint32, error) {
	if user == nil || user.Id == uuid.Nil {
		return 0, errors.New("user is nil")
	}
	return (user.Id.ID())%storage.shards + 1, nil
}

func (storage *PGstorage) CreateUser(ctx context.Context, user *models.User) error {

	userUuid, err := uuid.NewV7()
	if err != nil {
		return err
	}
	userMapped := User{
		Id:     userUuid,
		Name:   user.Name,
		Rating: user.Rating,
		TigrId: user.TigrId,
	}
	shard, err := storage.getUserShard(&userMapped)
	if err != nil {
		return err
	}

	q := squirrel.Insert(fmt.Sprintf("%s.%03d", UserTable, shard)).
		Columns(UserColumnRating, UserColumnName, UserTigrId).Values(user.Rating, user.Name, user.TigrId)
	err = storage.execQuery(ctx, q)
	if err != nil {
		return errors.Wrap(err, "Failed to create user")
	}
	return nil
}
