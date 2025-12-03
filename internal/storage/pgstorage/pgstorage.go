package pgstorage

import (
	"fmt"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/pkg/errors"
	"golang.org/x/net/context"
)

type PGstorage struct {
	db *pgxpool.Pool
}

func NewPGStorge(connString string, migrationsPath string) (*PGstorage, error) {
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
		db: db,
	}

	return storage, nil
}

func applyMigrations(connString, migrationsPath string) error {

	m, err := migrate.New(
		"file://"+migrationsPath,
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
