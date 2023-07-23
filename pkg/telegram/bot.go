package telegram

import (
	"github.com/Koderbek/pocket_news_bot/pkg/config"
	"github.com/Koderbek/pocket_news_bot/pkg/repository"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type Bot struct {
	bot        *tgbotapi.BotAPI
	repository *repository.Repository
	messages   config.Messages
}

func NewBot(bot *tgbotapi.BotAPI, repository *repository.Repository, messages config.Messages) *Bot {
	return &Bot{
		bot:        bot,
		repository: repository,
		messages:   messages,
	}
}

func (b *Bot) Start() error {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := b.bot.GetUpdatesChan(u)
	if err != nil {
		return err
	}

	for update := range updates {
		if update.CallbackQuery != nil {
			if err := b.handleCallbackQuery(update.CallbackQuery); err != nil {
				return err
			}

			continue
		}

		if update.Message != nil {
			if err := b.handleCommand(update.Message); err != nil {
				return err
			}

			continue
		}
	}

	return nil
}
