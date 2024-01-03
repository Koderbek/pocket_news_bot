package consumer

import (
	"fmt"
	"github.com/Koderbek/pocket_news_bot/pkg/model"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestMakeMessage(t *testing.T) {
	testCases := []struct {
		name   string
		param  model.Article
		result string
	}{
		{
			name: "case-1: valid result",
			param: model.Article{
				Title:       "Тест",
				Description: "Текст",
				Url:         "https://www.youtube.com/",
			},
			result: "<b>Тест</b>\n<i>Текст</i>\n<a href=\"https://www.youtube.com/\">Читать в источнике</a>",
		},
		{
			name:   "case-2: empty param",
			param:  model.Article{},
			result: "<b></b>\n<i></i>\n<a href=\"\">Читать в источнике</a>",
		},
	}

	for _, tc := range testCases {
		res := makeMessage(tc.param)
		require.Equal(t, tc.result, res)
	}
}

func TestMakeMessageHeader(t *testing.T) {
	testCases := []struct {
		name   string
		param  model.Category
		result string
	}{
		{
			name: "case-1: valid result",
			param: model.Category{
				Id:   0,
				Code: "test",
				Name: "Тест",
			},
			result: fmt.Sprintf("<a href=\"%s\">%s</a> | #Тест", botUrl, botName),
		},
		{
			name:   "case-2: empty param",
			param:  model.Category{},
			result: fmt.Sprintf("<a href=\"%s\">%s</a> | #", botUrl, botName),
		},
	}

	for _, tc := range testCases {
		res := makeMessageHeader(tc.param)
		require.Equal(t, tc.result, res)
	}
}

func TestLinkHashSum(t *testing.T) {
	testCases := []struct {
		name   string
		param  string
		result string
	}{
		{
			name:   "case-1: valid result",
			param:  "https://www.youtube.com/watch",
			result: "7e427d2ee465eef6c3f7eae4e4c17de4",
		},
		{
			name:   "case-2: empty param",
			param:  "",
			result: "d41d8cd98f00b204e9800998ecf8427e",
		},
	}

	for _, tc := range testCases {
		res := linkHashSum(tc.param)
		require.Equal(t, tc.result, res)
	}
}

func TestParseHost(t *testing.T) {
	testCases := []struct {
		name       string
		param      string
		result     string
		shouldFail bool
	}{
		{
			name:       "case-1: valid result",
			param:      "https://www.youtube.com/watch",
			result:     "www.youtube.com",
			shouldFail: false,
		},
		{
			name:       "case-2: empty param",
			param:      "",
			result:     "",
			shouldFail: false,
		},
		{
			name:       "case-3: parse error",
			param:      "123456gblsqfd3921--__-323---__:://",
			result:     "",
			shouldFail: true,
		},
	}

	for _, tc := range testCases {
		res, err := parseHost(tc.param)

		require.Equal(t, tc.result, res)
		if tc.shouldFail {
			require.Error(t, err)
		} else {
			require.NoError(t, err)
		}
	}
}
