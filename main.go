package main

import (
	"database/sql"
	"fmt"
	"log"
	"strconv"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	_ "github.com/lib/pq"
)

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	db_1, err := sql.Open("postgres", dbInfo)
	if err != nil {
		panic(err)
	}
	defer db_1.Close()
	sqlStatement := `
	INSERT INTO users (username, vallet, sun)
	VALUES ($1, $2, $3)
	RETURNING id`
	id := 0

	bot, err := tgbotapi.NewBotAPI("5456123319:AAHc27zB0_TrSVS6LL8h0VYV4hbRiTY1J_w")
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
			row := db_1.QueryRow("SELECT username FROM users where username= $1 and vallet = $2", update.Message.Chat.UserName, command[1])
			err = row.Scan(&update.Message.Chat.UserName)
			if err != nil {
				err = db_1.QueryRow(sqlStatement, update.Message.Chat.UserName, command[1], command[2]).Scan(&id)
			} else {
				db_1.Exec("UPDATE users SET sun = sun + $1", amount)
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
			sqlDelete := `DELETE FROM users WHERE id > 0; `
			_, err := db_1.Exec(sqlDelete)
			if err != nil {
				panic(err)

			}
		default:
			bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, "command not found"))
		}

	}
}
