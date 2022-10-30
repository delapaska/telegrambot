package main

import (
	"net/http"
	"test/config"
	"test/internal/bot"
	"test/internal/repository"
	"test/pkg/db/postgresql"
)

// Основной фаил старта приложения бота, никакой логики, только инициализация нужных сервисов, подтягивания конфигов, коннект к бд и запуск сервера

func main() {
	cfg := config.GetCommonEnvConfigs()

	// Инициализация бд
	conDB := postgresql.NewPostgresDB(cfg)

	repo := repository.New(conDB)

	// Инициализация серверва
	// ...

	// Инициализия rabbitMQ
	// ...

	// Инициализация бота
	bot := bot.NewBot(*repo, http.Client{})

	bot.Run()
}
