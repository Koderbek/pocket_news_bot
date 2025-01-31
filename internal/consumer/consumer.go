package consumer

import (
	"github.com/Koderbek/pocket_news_bot/internal/config"
	"github.com/Koderbek/pocket_news_bot/internal/model"
	"github.com/Koderbek/pocket_news_bot/internal/news"
	"github.com/Koderbek/pocket_news_bot/internal/repository"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"strings"
	"time"
)

const (
	botUrl  = "https://t.me/pocket_news_bot"
	botName = "🗞Pocket News"
)

type Consumer struct {
	bot        *tgbotapi.BotAPI
	newsClient news.Client
	repo       *repository.Repository
	cfg        config.Consumer
}

func NewConsumer(bot *tgbotapi.BotAPI, newsClient news.Client, repo *repository.Repository, cfg config.Consumer) *Consumer {
	return &Consumer{bot: bot, newsClient: newsClient, repo: repo, cfg: cfg}
}

func (c *Consumer) Start() error {
	var requestCount int8 = 0
	for {
		//TODO: Чистка таблицы sent_news - вынести в отдельную консольную команду
		if time.Now().Hour() == c.cfg.MailingTimeEnd || requestCount > c.cfg.RequestLimit {
			//Новый день. Чистим таблицу sent_news
			if err := c.repo.SentNews.Clean(); err != nil {
				return err
			}

			time.Sleep(c.cfg.DailySleep * time.Hour)
		}

		categories, err := c.repo.Category.GetAll() //Получаем все категории из category
		if err != nil {
			return err
		}

		for _, cat := range categories {
			requestCount++
			if requestCount > c.cfg.RequestLimit {
				break
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
				message = append(message, makeMessage(article))
			}

			if len(linksHash) == 0 {
				//Если отправлять нечего, то скрипаем
				continue
			}

			//Отправляем сообщение пользователям
			if err = c.sendMessage(cat, strings.Join(message, "\n\n")); err != nil {
				return err
			}

			//Сохраняем хэш отправленных сообщений
			if err = c.repo.SentNews.Save(linksHash); err != nil {
				return err
			}

			//Пауза на 30 минут
			time.Sleep(c.cfg.CategorySleep * time.Minute)
		}
	}
}

func (c *Consumer) sendMessage(cat model.Category, messageText string) error {
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
