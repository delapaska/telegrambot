package postgresql

import (
	"database/sql"
	"github.com/egorkurito/telegrambot/config"
	"log"
)

// Тут чисто коннект к бд

// connect to postgresql database
func NewPostgresDB(cfg config.CommonEnvConfigs) *sql.DB {
	db, err := sql.Open("postgres", cfg.PostgreSQL.Host)
	if err != nil {
		log.Panicf("Error open postgresql connection: %s", err)
		return nil
	}

	if err := db.Ping(); err != nil {
		log.Panicf("Error ping postgresql connection: %s", err)
		return nil
	}

	return db
}
