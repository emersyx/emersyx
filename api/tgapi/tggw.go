package tgapi

import (
	"emersyx.net/emersyx/api"
)

// TelegramParameters is the interface type which should be implemented to enable setting of parameter values when
// making calls to the Telegram Bot API.
type TelegramParameters interface {
	// ChatID must set the chat_id parameter. In the official documentation this parameter is of type Integer or String.
	// If you wish to use an Integer as parameter, then simply pass a go string object with the contents of the Integer.
	ChatID(value string) error
	// Text must set the text parameter.
	Text(value string) error
	// ParseMode must set the parse_mode parameter.
	ParseMode(value string) error
	// DisableWebPagePreview must set the disable_web_page_preview parameter.
	DisableWebPagePreview(value bool) error
	// DisableNotification must set the disable_notification parameter.
	DisableNotification(value bool) error
	// ReplyToMessageID must set the reply_to_message_id parameter.
	ReplyToMessageID(value int64) error
	// ReplyMarkup must set the reply_markup parameter. In the official documentation, this parameter is of type
	// InlineKeyboardMarkup, ReplyKeyboardMarkup, ReplyKeyboardRemove or ForceReply. Based on the documentation, the
	// parameter is expected to be JSON-serialized. This method takes a string as argument which should be a JSON string
	// of the appropriate object.
	ReplyMarkup(value string) error
}

// The TelegramGateway interface specifies the methods that must be implemented by a complete Telegram peripheral and
// receptor. The reference implementation at https://github.com/emersyx/emersyx_telegram follows this interface.
type TelegramGateway interface {
	api.Peripheral
	api.Receptor
	// This method must call the getMe method from the Telegram Bot API. It must return the response formatted into a
	// User instance.
	GetMe() (User, error)
	// This method must call the sendMessage method from the Telegram Bot API. The method must return the response
	// formatted into a Message object.
	SendMessage(params TelegramParameters) (Message, error)
	// This method must return a new TelegramParameters instance which can be used as argument to the other methods of
	// the TelegramGateway type.
	NewTelegramParameters() TelegramParameters
}
