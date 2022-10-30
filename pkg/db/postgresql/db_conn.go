package postgresql

import (
	"database/sql"
	"log"
	"test/config"
)

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
