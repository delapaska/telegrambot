package postgresql

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"test/config"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

// Тут чисто коннект к бд

// connect to postgresql database
func NewPostgresDB(cfg config.CommonEnvConfigs) *sql.DB {
	envErr := godotenv.Load("C:/telegrambot/.env")
	if envErr != nil {
		fmt.Printf("Could not load .env file")

	}
	fmt.Printf("POSTGRES HOST %s \n", os.Getenv("TOKEN"))

	db, err := sql.Open("postgres", cfg.PostgreSQL.Host)
	if err != nil {
		log.Panicf("Error open postgresql connection: %s", err)
		return nil
	}

	/*
		if err := db.Ping(); err != nil {
			log.Panicf("Error ping postgresql connection: %s", err)
			return nil
		}
	*/

	return db
}
