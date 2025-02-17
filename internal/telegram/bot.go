package telegram

import (
	"github.com/Koderbek/pocket_news_bot/internal/config"
	"github.com/Koderbek/pocket_news_bot/internal/repository"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"time"
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

	// Создаем rate limiter: максимум 5 запросов в 10 секунд
	rateLimiter := NewUserRateLimiter(b.messages.RateLimit, b.messages.RateLimitInterval*time.Second)
	for update := range updates {
		var msg *tgbotapi.Message
		if update.CallbackQuery != nil {
			msg = update.CallbackQuery.Message
		} else {
			msg = update.Message
		}

		// Проверяем, не превышен ли лимит запросов
		if msg != nil && !rateLimiter.Allow(msg.From.ID) {
			msgCfg := tgbotapi.NewMessage(msg.Chat.ID, b.messages.ManyRequestsCommand)
			b.bot.Send(msgCfg)

			continue
		}

		if update.CallbackQuery != nil {
			if err := b.handleCallbackQuery(update.CallbackQuery); err != nil {
				return err
			}
		} else if update.Message != nil {
			if err := b.handleCommand(update.Message); err != nil {
				return err
			}
		}
	}

	return nil
}
