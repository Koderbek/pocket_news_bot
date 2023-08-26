package consumer

import (
	"fmt"
	"github.com/Koderbek/pocket_news_bot/pkg/config"
	"github.com/Koderbek/pocket_news_bot/pkg/model"
	"github.com/Koderbek/pocket_news_bot/pkg/news"
	"github.com/Koderbek/pocket_news_bot/pkg/repository"
	"github.com/Koderbek/pocket_news_bot/pkg/rkn"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"strings"
	"time"
)

type Consumer struct {
	bot        *tgbotapi.BotAPI
	newsClient news.Client
	rknClient  rkn.Client
	repo       *repository.Repository
	cfg        config.Consumer
}

func NewConsumer(bot *tgbotapi.BotAPI, newsClient news.Client, rknClient rkn.Client, repo *repository.Repository, cfg config.Consumer) *Consumer {
	return &Consumer{bot: bot, newsClient: newsClient, rknClient: rknClient, repo: repo, cfg: cfg}
}

func (c *Consumer) Start() error {
	var requestCount int8 = 0
	for {
		if time.Now().Hour() == c.cfg.MailingTimeEnd || requestCount > c.cfg.RequestLimit {
			time.Sleep(c.cfg.DailySleep * time.Hour)
		}

		categories, err := c.repo.Category.GetAll()
		if err != nil {
			return err
		}

		for _, cat := range categories {
			requestCount++
			if requestCount > c.cfg.RequestLimit {
				break
			}

			news, err := c.newsClient.GetNews(cat.Code)
			if err != nil {
				return err
			}

			var message []string
			message = append(message, "#"+cat.Name)
			for _, article := range news {
				//isForbidden, err := c.rknClient.IsForbidden(article.Url)
				//if err != nil {
				//	return err
				//}
				//
				//if isForbidden {
				//	continue
				//}

				message = append(message, makeMessage(article))
			}

			chatCategories, err := c.repo.ChatCategory.GetByCategoryId(cat.Id)
			if err != nil {
				return err
			}

			//Отправка сообщений пользователям с данной категорией
			messageText := strings.Join(message, "\n\n")
			for _, chatCategory := range chatCategories {
				go func(chatId int64) {
					msg := tgbotapi.NewMessage(chatId, messageText)
					msg.ParseMode = tgbotapi.ModeHTML
					c.bot.Send(msg)
				}(chatCategory.ChatId)
			}

			time.Sleep(c.cfg.CategorySleep * time.Minute)
		}
	}
}

func makeMessage(article model.Article) string {
	return fmt.Sprintf("<b>%s</b>\n<i>%s</i>\n<a href=\"%s\">Читать в источнике</a>", article.Title, article.Description, article.Url)
}
