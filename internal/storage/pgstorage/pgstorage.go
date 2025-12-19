package pgstorage

import (
	"fmt"

	"github.com/Masterminds/squirrel"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/pkg/errors"
	"golang.org/x/net/context"
)

type PGstorage struct {
	db     *pgxpool.Pool
	shards uint32
}

func NewPGStorge(connString string, migrationsPath string, shards uint32) (*PGstorage, error) {
	err := applyMigrations(connString, migrationsPath)

	if err != nil {
		return nil, errors.Wrap(err, "Failed to apply migrations")
	}

	config, err := pgxpool.ParseConfig(connString)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to parse config")
	}

	db, err := pgxpool.NewWithConfig(context.Background(), config)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to connect to database")
	}
	storage := &PGstorage{
		db:     db,
		shards: shards,
	}

	return storage, nil
}

func applyMigrations(connString, migrationsPath string) error {

	m, err := migrate.New(
		fmt.Sprintf("file://%s", migrationsPath),
		connString,
	)
	if err != nil {
		return errors.Wrap(err, "ошибка создания мигратора")
	}

	defer func(m *migrate.Migrate) {
		err, _ := m.Close()
		if err != nil {

		}
	}(m)

	if err := m.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		return errors.Wrap(err, "ошибка применения миграций")
	}

	fmt.Println("Миграции успешно применены")
	return nil
}

func (storage *PGstorage) execQuery(ctx context.Context, query squirrel.Sqlizer) error {
	queryText, args, err := query.ToSql()
	if err != nil {
		return errors.Wrap(err, "Generate query error")
	}
	_, err = storage.db.Exec(ctx, queryText, args...)
	if err != nil {
		err = errors.Wrap(err, "Exec query error")
	}
	return err
}

func (storage *PGstorage) query(ctx context.Context, query squirrel.Sqlizer) (pgx.Rows, error) {
	queryText, args, err := query.ToSql()
	if err != nil {
		return nil, errors.Wrap(err, "Generate query error")
	}
	rows, err := storage.db.Query(ctx, queryText, args...)
	if err != nil {
		err = errors.Wrap(err, "Rows Query error")
		return nil, err
	}
	return rows, nil
}
