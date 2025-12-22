package bootstrap

import (
	"fmt"
	"log"

	"github.com/tengenatari/web-referee/config"
	"github.com/tengenatari/web-referee/internal/storage/pgstorage"
)

func InitPGStorage(cfg *config.Config) *pgstorage.WebRefereeStorage {

	connectionString := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=%s",
		cfg.Database.Username, cfg.Database.Password, cfg.Database.Host, cfg.Database.Port, cfg.Database.DBName, cfg.Database.SSLMode)
	fmt.Println(cfg.Database.MigrationsPath)
	storage, err := pgstorage.NewWebRefereeStorge(connectionString, cfg.Database.MigrationsPath, cfg.Database.Shards)
	if err != nil {
		log.Panic(fmt.Sprintf("ошибка инициализации БД, %v", err))
	}
	return storage
}
