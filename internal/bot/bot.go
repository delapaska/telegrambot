package bot

import (
	"github.com/egorkurito/telegrambot/internal/repository"
	"net/http"
)

type Bot struct {
	repo   repository.Repository
	client http.Client
	//adad
}

func NewBot(repo repository.Repository, client http.Client) *Bot {
	return &Bot{
		repo:   repo,
		client: client,
	}
}

func (b *Bot) Run() {
	case "DELETE":
		b.repo.Delete()
}
