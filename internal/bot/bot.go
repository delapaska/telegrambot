package bot

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"test/config"
	"test/internal/repository"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
)

type Bot struct {
	repo   repository.Repository
	client http.Client

	//adad
}
type wallet map[string]float64

var db = map[int64]wallet{}

func NewBot(repo repository.Repository, client http.Client) *Bot {
	return &Bot{
		repo:   repo,
		client: client,
	}
}

func (b *Bot) Run() {
	envErr := godotenv.Load("C:/telegrambot/.env")
	if envErr != nil {
		fmt.Printf("Could not load .env file")

	}
	//fmt.Printf("POSTGRES HOST %s \n", os.Getenv("token"))
	fmt.Printf(config.GetCommonEnvConfigs().TgBot.TOKEN)
	bot, err := tgbotapi.NewBotAPI(os.Getenv("TOKEN"))
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil { // If we got a message
			continue
		}
		command := strings.Split(update.Message.Text, " ")
		switch command[0] {
		case "ADD":
			if len(command) != 3 {
				bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, "wrong command"))
			}
			amount, err := strconv.ParseFloat(command[2], 64)
			if err != nil {
				bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, err.Error()))
			}

			if _, ok := db[update.Message.Chat.ID]; !ok {
				db[update.Message.Chat.ID] = wallet{}
			}
			db[update.Message.Chat.ID][command[1]] += amount
			balanceText := fmt.Sprintf("%f", db[update.Message.Chat.ID][command[1]])
			bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, balanceText))
		case "SUB":
			if len(command) != 3 {
				bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, "wrong command"))
			}
			amount, err := strconv.ParseFloat(command[2], 64)

			if err != nil {
				bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, err.Error()))
			}
			if _, ok := db[update.Message.Chat.ID]; !ok {
				continue
			}
			db[update.Message.Chat.ID][command[1]] -= amount
			balanceText := fmt.Sprintf("%f", db[update.Message.Chat.ID][command[1]])
			bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, balanceText))
		case "DEL":
			if len(command) != 2 {
				bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, "wrong command"))
			}

			delete(db[update.Message.Chat.ID], command[1])
		case "SHOW":
			msg := ""
			var sum float64 = 0
			for key, value := range db[update.Message.Chat.ID] {
				price, _ := getPrice(key)
				sum += value * price
				msg += fmt.Sprintf("%s: %f [%.2f]\n", key, value, value*price)
			}
			msg += fmt.Sprintf("Total: %.2f\n", sum)
			bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, msg))
		case "DELETE":

		default:
			bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, "command not found"))
		}
	}
}
