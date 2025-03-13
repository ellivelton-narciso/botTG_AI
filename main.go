package main

import (
	"botTG_AI/codeai"
	"log"
	"os"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Aviso: Arquivo .env não encontrado ou não carregado corretamente")
	}

	openaiKey := os.Getenv("OPENAI_API_KEY")
	if openaiKey == "" {
		log.Panic("Erro: OPENAI_API_KEY não foi definido")
	}

	botToken := os.Getenv("TELEGRAM_BOT_TOKEN")
	if botToken == "" {
		log.Panic("Erro: TELEGRAM_BOT_TOKEN não foi definido")
	}

	bot, err := tgbotapi.NewBotAPI(botToken)
	if err != nil {
		log.Panic(err)
	}

	debugStr := os.Getenv("TELEGRAM_BOT_DEBUG")
	debugMode, err := strconv.ParseBool(debugStr)
	if err != nil {
		log.Printf("Aviso: TELEGRAM_BOT_DEBUG inválido ou não definido, assumindo 'false'")
		debugMode = false
	}
	bot.Debug = debugMode

	log.Printf("Bot iniciado como %s | Debug: %t", bot.Self.UserName, bot.Debug)

	updateConfig := tgbotapi.NewUpdate(0)
	updateConfig.Timeout = 60

	updates := bot.GetUpdatesChan(updateConfig)

	for update := range updates {
		if update.Message == nil {
			continue
		}

		log.Printf("Mensagem recebida de %s: %s", update.Message.From.UserName, update.Message.Text)

		if update.Message.IsCommand() {
			switch update.Message.Command() {
			case "start":
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Hello World")
				bot.Send(msg)
			case "generate":
				prompt := update.Message.CommandArguments()
				if prompt == "" {
					msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Por favor, forneça um prompt após /generate")
					bot.Send(msg)
					continue
				}
				code := codeai.GenerateCode(prompt, openaiKey, debugMode)

				msg := tgbotapi.NewMessage(update.Message.Chat.ID, code)
				bot.Send(msg)
			}
		}
	}
}
