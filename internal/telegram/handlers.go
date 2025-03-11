package telegram

import (
	"fmt"
	"github.com/Koderbek/pocket_news_bot/internal/model"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"strings"
)

const (
	commandStart         = "start"
	commandHelp          = "help"
	commandEditCategory  = "editcategory"
	messageModifiedError = "message is not modified"
)

func (b *Bot) handleCommand(message *tgbotapi.Message) error {
	switch message.Command() {
	case commandStart:
		return b.handleStartCommand(message)
	case commandHelp:
		return b.handleStartCommand(message)
	case commandEditCategory:
		return b.handleEditCategoryCommand(message)
	default:
		return b.handleUnknownCommand(message)
	}
}

func (b *Bot) handleStartCommand(message *tgbotapi.Message) error {
	startMessage := fmt.Sprintf(b.messages.StartCommand, "/"+commandEditCategory)
	msg := tgbotapi.NewMessage(message.Chat.ID, startMessage)
	_, err := b.bot.Send(msg)

	return err
}

func (b *Bot) handleEditCategoryCommand(message *tgbotapi.Message) error {
	msg := tgbotapi.NewMessage(message.Chat.ID, b.messages.EditCategoryCommand)
	resMsg, err := b.bot.Send(msg)
	if err != nil {
		return err
	}

	buttons, err := b.chatCategoryButtons(message.Chat.ID)
	if err != nil {
		return err
	}

	msgBut := tgbotapi.NewEditMessageReplyMarkup(message.Chat.ID, resMsg.MessageID, *buttons)
	_, err = b.bot.Send(msgBut)
	return err
}

func (b *Bot) handleUnknownCommand(message *tgbotapi.Message) error {
	msg := tgbotapi.NewMessage(message.Chat.ID, b.messages.UnknownCommand)
	_, err := b.bot.Send(msg)
	return err
}

func (b *Bot) handleCallbackQuery(callbackQuery *tgbotapi.CallbackQuery) error {
	err := b.editChatCategory(callbackQuery)
	if err != nil {
		return err
	}

	message := callbackQuery.Message
	buttons, err := b.chatCategoryButtons(message.Chat.ID)
	if err != nil {
		return err
	}

	msgBut := tgbotapi.NewEditMessageReplyMarkup(message.Chat.ID, message.MessageID, *buttons)
	_, err = b.bot.Send(msgBut)

	if err != nil && strings.Contains(err.Error(), messageModifiedError) {
		return nil
	}

	return err
}

func (b *Bot) chatCategoryButtons(chatId int64) (*tgbotapi.InlineKeyboardMarkup, error) {
	chatCategories, err := b.repository.ChatCategory.GetByChatId(chatId)
	if err != nil {
		return nil, err
	}

	categories, err := b.repository.Category.GetAll()
	if err != nil {
		return nil, err
	}

	var buttons [][]tgbotapi.InlineKeyboardButton
	for _, cat := range categories {
		btnName := "⬜️ " + cat.Name
		for _, chatCat := range chatCategories {
			if chatCat.CategoryId == cat.Id {
				btnName = "✅ " + cat.Name
				break
			}
		}

		buttons = append(
			buttons,
			tgbotapi.NewInlineKeyboardRow(tgbotapi.NewInlineKeyboardButtonData(btnName, cat.Code)),
		)
	}

	allCat := model.AllCategory()
	buttons = append(
		buttons,
		tgbotapi.NewInlineKeyboardRow(tgbotapi.NewInlineKeyboardButtonData(allCat.Name, allCat.Code)),
	)

	inlineKeyboardMarkup := tgbotapi.NewInlineKeyboardMarkup(buttons...)
	return &inlineKeyboardMarkup, nil
}

func (b *Bot) editChatCategory(callbackQuery *tgbotapi.CallbackQuery) error {
	catCode := callbackQuery.Data
	chatId := callbackQuery.Message.Chat.ID
	allCat := model.AllCategory()
	if catCode == allCat.Code {
		allCats, err := b.repository.Category.GetAll()
		if err != nil {
			return err
		}

		for _, cat := range allCats {
			if b.repository.ChatCategory.HasChatCategory(chatId, cat.Id) {
				continue
			}

			err = b.repository.ChatCategory.Add(chatId, cat.Id, callbackQuery.Message.Chat.UserName)
			if err != nil {
				return err
			}
		}

		return nil
	}

	cat, err := b.repository.Category.GetByCode(catCode)
	if err != nil {
		return err
	}

	if b.repository.ChatCategory.HasChatCategory(chatId, cat.Id) {
		err = b.repository.ChatCategory.Deactivate(chatId, cat.Id)
	} else {
		err = b.repository.ChatCategory.Add(chatId, cat.Id, callbackQuery.Message.Chat.UserName)
	}

	return err
}
