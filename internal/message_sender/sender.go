package message_sender

import (
	"github.com/Koderbek/pocket_news_bot/internal/config"
	"github.com/Koderbek/pocket_news_bot/internal/model"
	"github.com/Koderbek/pocket_news_bot/internal/news"
	"github.com/Koderbek/pocket_news_bot/internal/repository"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"strings"
	"time"
)

const (
	botUrl  = "https://t.me/pocket_news_bot"
	botName = "🗞Pocket News"
)

type Sender struct {
	bot        *tgbotapi.BotAPI
	newsClient news.Client
	repo       *repository.Repository
	cfg        config.MessageSender
}

func NewSender(bot *tgbotapi.BotAPI, newsClient news.Client, repo *repository.Repository, cfg config.MessageSender) *Sender {
	return &Sender{bot: bot, newsClient: newsClient, repo: repo, cfg: cfg}
}

func (c *Sender) Start() error {
	if time.Now().Hour() >= c.cfg.MailingTimeEnd && time.Now().Hour() < c.cfg.MailingTimeStart {
		return nil
	}

	cat, err := c.repo.Category.ForSending()
	if err != nil {
		return err
	}

	if cat == nil {
		return nil
	}

	catNews, err := c.newsClient.GetNews(cat.Code)
	if err != nil {
		return err
	}

	if err = c.repo.Category.UpdateLastSent(cat.Code); err != nil {
		return err
	}

	var linksHash []string
	message := []string{makeMessageHeader(cat)}
	i := 1
	for _, article := range catNews {
		linkHash := linkHashSum(article.Url)
		if c.repo.SentNews.IsExists(linkHash) {
			//Если новость отправляли ранее, то скипаем ее
			continue
		}

		domain, err := parseHost(article.Url)
		if err != nil {
			return err
		}

		if c.repo.DomainBlacklist.IsExists(domain) {
			//Проверяем наличие ресурса среди запрещенных
			continue
		}

		linksHash = append(linksHash, linkHash)
		message = append(message, makeMessage(article, i))
		i++
	}

	if len(linksHash) == 0 {
		return nil
	}

	//Отправляем сообщение пользователям
	if err = c.sendMessage(cat, strings.Join(message, "\n\n")); err != nil {
		return err
	}

	//Сохраняем хэш отправленных сообщений
	if err = c.repo.SentNews.Save(linksHash); err != nil {
		return err
	}

	return nil
}

func (c *Sender) sendMessage(cat *model.Category, messageText string) error {
	chatCategories, err := c.repo.ChatCategory.GetByCategoryId(cat.Id)
	if err != nil {
		return err
	}

	//Асинхронная отправка сообщений пользователям с данной категорией
	for _, chatCategory := range chatCategories {
		go func(chatId int64) {
			msg := tgbotapi.NewMessage(chatId, messageText)
			msg.ParseMode = tgbotapi.ModeHTML
			c.bot.Send(msg)
		}(chatCategory.ChatId)
	}

	return nil
}
