package pgstorage

import (
	"fmt"

	"github.com/Masterminds/squirrel"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"github.com/tengenatari/web-referee/internal/models"
	"golang.org/x/net/context"
)

func (storage *WebRefereeStorage) getUserShard(user *User) (uint32, error) {
	if user == nil || user.Id == uuid.Nil {
		return 0, errors.New("user is nil")
	}
	return (user.Id.ID())%storage.shards + 1, nil
}

func (storage *WebRefereeStorage) CreateUser(ctx context.Context, user *models.User) error {

	userUuid, err := uuid.NewV7()
	if err != nil {
		return errors.Wrap(err, "error creating user uuid")
	}

	userMapped := User{
		Id:     userUuid,
		Name:   user.Name,
		Rating: user.Rating,
		TigrId: user.TigrId,
	}
	shard, err := storage.getUserShard(&userMapped)
	if err != nil {
		return errors.Wrap(err, "error getting user shard")
	}
	fmt.Println(userUuid, shard)

	q := squirrel.Insert(fmt.Sprintf("schema_%03d.%s", shard, UserTable)).
		Columns(UserId, UserColumnRating, UserColumnName, UserTigrId).Values(userMapped.Id, userMapped.Rating, userMapped.Name, userMapped.TigrId).
		PlaceholderFormat(squirrel.Dollar)
	err = storage.execQuery(ctx, q)
	if err != nil {
		return errors.Wrap(err, "Failed to create user")
	}
	return nil
}
