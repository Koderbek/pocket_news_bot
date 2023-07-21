package telegram

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

const (
	commandStart        = "start"
	commandEditCategory = "editcategory"
)

func (b *Bot) handleCommand(message *tgbotapi.Message) error {
	switch message.Command() {
	case commandStart:
		return b.handleStartCommand(message)
	case commandEditCategory:
		return b.handleEditCategoryCommand(message)
	default:
		return b.handleUnknownCommand(message)
	}
}

func (b *Bot) handleStartCommand(message *tgbotapi.Message) error {
	startMessage := fmt.Sprintf(b.messages.Start, "/"+commandEditCategory)
	msg := tgbotapi.NewMessage(message.Chat.ID, startMessage)
	_, err := b.bot.Send(msg)

	return err
}

func (b *Bot) handleEditCategoryCommand(message *tgbotapi.Message) error {
	msg := tgbotapi.NewMessage(message.Chat.ID, b.messages.EditCategory)
	resMsg, err := b.bot.Send(msg)
	if err != nil {
		return err
	}

	chatCategories, err := b.repository.GetAll()
	if err != nil {
		return err
	}

	var buttons []tgbotapi.InlineKeyboardButton
	for _, cat := range chatCategories {
		buttons = append(buttons, tgbotapi.NewInlineKeyboardButtonData(cat.Name, cat.Code))
	}

	msgBut := tgbotapi.NewEditMessageReplyMarkup(
		message.Chat.ID,
		resMsg.MessageID,
		tgbotapi.NewInlineKeyboardMarkup(buttons),
	)

	_, err = b.bot.Send(msgBut)

	return err
}

func (b *Bot) handleUnknownCommand(message *tgbotapi.Message) error {
	msg := tgbotapi.NewMessage(message.Chat.ID, b.messages.UnknownCommand)
	_, err := b.bot.Send(msg)
	return err
}
