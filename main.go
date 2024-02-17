package main

import (
	"fmt"
	"os"

	"github.com/mymmrac/telego"

	"github.com/joho/godotenv"
	tu "github.com/mymmrac/telego/telegoutil"
)

func main() {
	botToken := os.Getenv("TOKEN")

	// Create Bot with debug on
	bot, err := telego.NewBot(botToken, telego.WithDefaultDebugLogger())
	if err != nil {
		fmt.Println(err)
		return
	}

	bot.DeleteWebhook(nil)

	// Get updates channel
	updates, _ := bot.UpdatesViaLongPolling(nil)

	// Stop reviving updates from updates channel
	defer bot.StopLongPolling()

	// Loop through all updates when they came
	for update := range updates {
		// Check if update contains message
		if update.Message != nil {
			// Get chat ID from message
			chatID := tu.ID(update.Message.Chat.ID)

			// Copy sent message back to user
			_, _ = bot.CopyMessage(&telego.CopyMessageParams{
				ChatID:     chatID,
				FromChatID: chatID,
				MessageID:  update.Message.MessageID,
			})
		}
	}
}

func init() {
	// loads values from .env into the system
	if err := godotenv.Load(); err != nil {
		fmt.Println("No .env file found")
	}
}
