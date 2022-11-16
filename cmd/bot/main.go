package main

import (
	"net/http"
	"test/internal/repository/repository_impl"

	"test/config"
	"test/internal/bot"
	"test/pkg/db/postgresql"
	//	"github.com/joho/godotenv"
)

// Основной фаил старта приложения бота, никакой логики, только инициализация нужных сервисов, подтягивания конфигов, коннект к бд и запуск сервера

func main() {
	cfg := config.GetCommonEnvConfigs()

	//fmt.Printf("POSTGRES HOST %s \n", os.Getenv("token"))

	// Инициализация бд
	conDB := postgresql.NewPostgresDB(cfg)

	repo := repository_impl.New(conDB)

	// Инициализация серверва
	// ...

	// Инициализия rabbitMQ
	// ...

	// Инициализация бота
	bot := bot.NewBot(repo, http.Client{})

	bot.Run()
}
