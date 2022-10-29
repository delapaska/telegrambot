package main

import (
	"github.com/egorkurito/telegrambot/config"
	"github.com/egorkurito/telegrambot/internal/bot"
	"github.com/egorkurito/telegrambot/internal/repository"
	"github.com/egorkurito/telegrambot/pkg/db/postgresql"
	"net/http"
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
	bot := bot.NewBot(repo, http.Client{})

	bot.Run()
}
