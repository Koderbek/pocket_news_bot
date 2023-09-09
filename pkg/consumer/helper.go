package consumer

import (
	"crypto/md5"
	"fmt"
	"github.com/Koderbek/pocket_news_bot/pkg/model"
	"net/url"
)

func makeMessage(article model.Article) string {
	return fmt.Sprintf("<b>%s</b>\n<i>%s</i>\n<a href=\"%s\">Читать в источнике</a>", article.Title, article.Description, article.Url)
}

func makeMessageHeader(category model.Category) string {
	return fmt.Sprintf("<a href=\"%s\">%s</a> | #%s", botUrl, botName, category.Name)
}

func linkHashSum(link string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(link)))
}

func parseHost(srcUrl string) (string, error) {
	u, err := url.Parse(srcUrl)
	if err != nil {
		return "", err
	}

	return u.Hostname(), nil
}
