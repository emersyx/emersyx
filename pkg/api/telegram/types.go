package telegram

// The structs below have been automatically generated using the tgdocs2go tool. The tool is part of the emersyx project
// and can be found at https://github.com/emersyx/tgdocs2go.

// Update type with fields compatible to the Telegram Bot API.
// https://core.telegram.org/bots/api#Update
type Update struct {
	UpdateID           int64               `json:"update_id"`
	Message            *Message            `json:"message"`
	EditedMessage      *Message            `json:"edited_message"`
	ChannelPost        *Message            `json:"channel_post"`
	EditedChannelPost  *Message            `json:"edited_channel_post"`
	InlineQuery        *InlineQuery        `json:"inline_query"`
	ChosenInlineResult *ChosenInlineResult `json:"chosen_inline_result"`
	CallbackQuery      *CallbackQuery      `json:"callback_query"`
	ShippingQuery      *ShippingQuery      `json:"shipping_query"`
	PreCheckoutQuery   *PreCheckoutQuery   `json:"pre_checkout_query"`
}

// User type with fields compatible to the Telegram Bot API.
// https://core.telegram.org/bots/api#User
type User struct {
	ID           int64  `json:"id"`
	IsBot        bool   `json:"is_bot"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	Username     string `json:"username"`
	LanguageCode string `json:"language_code"`
}

// Chat type with fields compatible to the Telegram Bot API.
// https://core.telegram.org/bots/api#Chat
type Chat struct {
	ID                          int64      `json:"id"`
	Type                        string     `json:"type"`
	Title                       string     `json:"title"`
	Username                    string     `json:"username"`
	FirstName                   string     `json:"first_name"`
	LastName                    string     `json:"last_name"`
	AllMembersAreAdministrators bool       `json:"all_members_are_administrators"`
	Photo                       *ChatPhoto `json:"photo"`
	Description                 string     `json:"description"`
	InviteLink                  string     `json:"invite_link"`
}

// Message type with fields compatible to the Telegram Bot API.
// https://core.telegram.org/bots/api#message
type Message struct {
	MessageID             int64              `json:"message_id"`
	From                  *User              `json:"from"`
	Date                  int64              `json:"date"`
	Chat                  *Chat              `json:"chat"`
	ForwardFrom           *User              `json:"forward_from"`
	ForwardFromChat       *Chat              `json:"forward_from_chat"`
	ForwardFromMessageID  int64              `json:"forward_from_message_id"`
	ForwardDate           int64              `json:"forward_date"`
	ReplyToMessage        *Message           `json:"reply_to_message"`
	EditDate              int64              `json:"edit_date"`
	Text                  string             `json:"text"`
	Entities              *[]MessageEntity   `json:"entities"`
	Audio                 *Audio             `json:"audio"`
	Document              *Document          `json:"document"`
	Game                  *Game              `json:"game"`
	Photo                 *[]PhotoSize       `json:"photo"`
	Sticker               *Sticker           `json:"sticker"`
	Video                 *Video             `json:"video"`
	Voice                 *Voice             `json:"voice"`
	VideoNote             *VideoNote         `json:"video_note"`
	NewChatMembers        *[]User            `json:"new_chat_members"`
	Caption               string             `json:"caption"`
	Contact               *Contact           `json:"contact"`
	Location              *Location          `json:"location"`
	Venue                 *Venue             `json:"venue"`
	NewChatMember         *User              `json:"new_chat_member"`
	LeftChatMember        *User              `json:"left_chat_member"`
	NewChatTitle          string             `json:"new_chat_title"`
	NewChatPhoto          *[]PhotoSize       `json:"new_chat_photo"`
	DeleteChatPhoto       bool               `json:"delete_chat_photo"`
	GroupChatCreated      bool               `json:"group_chat_created"`
	SupergroupChatCreated bool               `json:"supergroup_chat_created"`
	ChannelChatCreated    bool               `json:"channel_chat_created"`
	MigrateToChatID       int64              `json:"migrate_to_chat_id"`
	MigrateFromChatID     int64              `json:"migrate_from_chat_id"`
	PinnedMessage         *Message           `json:"pinned_message"`
	Invoice               *Invoice           `json:"invoice"`
	SuccessfulPayment     *SuccessfulPayment `json:"successful_payment"`
}

// MessageEntity type with fields compatible to the Telegram Bot API.
// https://core.telegram.org/bots/api#MessageEntity
type MessageEntity struct {
	Type   string `json:"type"`
	Offset int64  `json:"offset"`
	Length int64  `json:"length"`
	URL    string `json:"url"`
	User   *User  `json:"user"`
}

// PhotoSize type with fields compatible to the Telegram Bot API.
// https://core.telegram.org/bots/api#PhotoSize
type PhotoSize struct {
	FileID   string `json:"file_id"`
	Width    int64  `json:"width"`
	Height   int64  `json:"height"`
	FileSize int64  `json:"file_size"`
}

// Audio type with fields compatible to the Telegram Bot API.
// https://core.telegram.org/bots/api#Audio
type Audio struct {
	FileID    string `json:"file_id"`
	Duration  int64  `json:"duration"`
	Performer string `json:"performer"`
	Title     string `json:"title"`
	MimeType  string `json:"mime_type"`
	FileSize  int64  `json:"file_size"`
}

// Document type with fields compatible to the Telegram Bot API.
// https://core.telegram.org/bots/api#Document
type Document struct {
	FileID   string     `json:"file_id"`
	Thumb    *PhotoSize `json:"thumb"`
	FileName string     `json:"file_name"`
	MimeType string     `json:"mime_type"`
	FileSize int64      `json:"file_size"`
}

// Video type with fields compatible to the Telegram Bot API.
// https://core.telegram.org/bots/api#Video
type Video struct {
	FileID   string     `json:"file_id"`
	Width    int64      `json:"width"`
	Height   int64      `json:"height"`
	Duration int64      `json:"duration"`
	Thumb    *PhotoSize `json:"thumb"`
	MimeType string     `json:"mime_type"`
	FileSize int64      `json:"file_size"`
}

// Voice type with fields compatible to the Telegram Bot API.
// https://core.telegram.org/bots/api#Voice
type Voice struct {
	FileID   string `json:"file_id"`
	Duration int64  `json:"duration"`
	MimeType string `json:"mime_type"`
	FileSize int64  `json:"file_size"`
}

// VideoNote type with fields compatible to the Telegram Bot API.
// https://core.telegram.org/bots/api#VideoNote
type VideoNote struct {
	FileID   string     `json:"file_id"`
	Length   int64      `json:"length"`
	Duration int64      `json:"duration"`
	Thumb    *PhotoSize `json:"thumb"`
	FileSize int64      `json:"file_size"`
}

// Contact type with fields compatible to the Telegram Bot API.
// https://core.telegram.org/bots/api#Contact
type Contact struct {
	PhoneNumber string `json:"phone_number"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	UserID      int64  `json:"user_id"`
}

// Location type with fields compatible to the Telegram Bot API.
// https://core.telegram.org/bots/api#Location
type Location struct {
	Longitude *float64 `json:"longitude"`
	Latitude  *float64 `json:"latitude"`
}

// Venue type with fields compatible to the Telegram Bot API.
// https://core.telegram.org/bots/api#Venue
type Venue struct {
	Location     *Location `json:"location"`
	Title        string    `json:"title"`
	Address      string    `json:"address"`
	FoursquareID string    `json:"foursquare_id"`
}

// UserProfilePhotos type with fields compatible to the Telegram Bot API.
// https://core.telegram.org/bots/api#UserProfilePhotos
type UserProfilePhotos struct {
	TotalCount int64          `json:"total_count"`
	Photos     *[][]PhotoSize `json:"photos"`
}

// File type with fields compatible to the Telegram Bot API.
// https://core.telegram.org/bots/api#File
type File struct {
	FileID   string `json:"file_id"`
	FileSize int64  `json:"file_size"`
	FilePath string `json:"file_path"`
}

// ReplyKeyboardMarkup type with fields compatible to the Telegram Bot API.
// https://core.telegram.org/bots/api#ReplyKeyboardMarkup
type ReplyKeyboardMarkup struct {
	Keyboard        *[][]KeyboardButton `json:"keyboard"`
	ResizeKeyboard  bool                `json:"resize_keyboard"`
	OneTimeKeyboard bool                `json:"one_time_keyboard"`
	Selective       bool                `json:"selective"`
}

// KeyboardButton type with fields compatible to the Telegram Bot API.
// https://core.telegram.org/bots/api#KeyboardButton
type KeyboardButton struct {
	Text            string `json:"text"`
	RequestContact  bool   `json:"request_contact"`
	RequestLocation bool   `json:"request_location"`
}

// ReplyKeyboardRemove type with fields compatible to the Telegram Bot API.
// https://core.telegram.org/bots/api#ReplyKeyboardRemove
type ReplyKeyboardRemove struct {
	RemoveKeyboard bool `json:"remove_keyboard"`
	Selective      bool `json:"selective"`
}

// InlineKeyboardMarkup type with fields compatible to the Telegram Bot API.
// https://core.telegram.org/bots/api#InlineKeyboardMarkup
type InlineKeyboardMarkup struct {
	InlineKeyboard *[][]InlineKeyboardButton `json:"inline_keyboard"`
}

// CallbackQuery type with fields compatible to the Telegram Bot API.
// https://core.telegram.org/bots/api#CallbackQuery
type CallbackQuery struct {
	ID              string   `json:"id"`
	From            *User    `json:"from"`
	Message         *Message `json:"message"`
	InlineMessageID string   `json:"inline_message_id"`
	ChatInstance    string   `json:"chat_instance"`
	Data            string   `json:"data"`
	GameShortName   string   `json:"game_short_name"`
}

// ForceReply type with fields compatible to the Telegram Bot API.
// https://core.telegram.org/bots/api#ForceReply
type ForceReply struct {
	ForceReply bool `json:"force_reply"`
	Selective  bool `json:"selective"`
}

// ChatPhoto type with fields compatible to the Telegram Bot API.
// https://core.telegram.org/bots/api#ChatPhoto
type ChatPhoto struct {
	SmallFileID string `json:"small_file_id"`
	BigFileID   string `json:"big_file_id"`
}

// ChatMember type with fields compatible to the Telegram Bot API.
// https://core.telegram.org/bots/api#ChatMember
type ChatMember struct {
	User                  *User  `json:"user"`
	Status                string `json:"status"`
	UntilDate             int64  `json:"until_date"`
	CanBeEdited           bool   `json:"can_be_edited"`
	CanChangeInfo         bool   `json:"can_change_info"`
	CanPostMessages       bool   `json:"can_post_messages"`
	CanEditMessages       bool   `json:"can_edit_messages"`
	CanDeleteMessages     bool   `json:"can_delete_messages"`
	CanInviteUsers        bool   `json:"can_invite_users"`
	CanRestrictMembers    bool   `json:"can_restrict_members"`
	CanPinMessages        bool   `json:"can_pin_messages"`
	CanPromoteMembers     bool   `json:"can_promote_members"`
	CanSendMessages       bool   `json:"can_send_messages"`
	CanSendMediaMessages  bool   `json:"can_send_media_messages"`
	CanSendOtherMessages  bool   `json:"can_send_other_messages"`
	CanAddWebPagePreviews bool   `json:"can_add_web_page_previews"`
}

// ResponseParameters type with fields compatible to the Telegram Bot API.
// https://core.telegram.org/bots/api#ResponseParameters
type ResponseParameters struct {
	MigrateToChatID int64 `json:"migrate_to_chat_id"`
	RetryAfter      int64 `json:"retry_after"`
}

// Game type with fields compatible to the Telegram Bot API.
// https://core.telegram.org/bots/api#Game
type Game struct {
	Title        string           `json:"title"`
	Description  string           `json:"description"`
	Photo        *[]PhotoSize     `json:"photo"`
	Text         string           `json:"text"`
	TextEntities *[]MessageEntity `json:"text_entities"`
	Animation    *Animation       `json:"animation"`
}

// Animation type with fields compatible to the Telegram Bot API.
// https://core.telegram.org/bots/api#Animation
type Animation struct {
	FileID   string     `json:"file_id"`
	Thumb    *PhotoSize `json:"thumb"`
	FileName string     `json:"file_name"`
	MimeType string     `json:"mime_type"`
	FileSize int64      `json:"file_size"`
}

// Sticker type with fields compatible to the Telegram Bot API.
// https://core.telegram.org/bots/api#Sticker
type Sticker struct {
	FileID       string        `json:"file_id"`
	Width        int64         `json:"width"`
	Height       int64         `json:"height"`
	Thumb        *PhotoSize    `json:"thumb"`
	Emoji        string        `json:"emoji"`
	SetName      string        `json:"set_name"`
	MaskPosition *MaskPosition `json:"mask_position"`
	FileSize     int64         `json:"file_size"`
}

// MaskPosition type with fields compatible to the Telegram Bot API.
// https://core.telegram.org/bots/api#MaskPosition
type MaskPosition struct {
	Point  string  `json:"point"`
	XShift float64 `json:"x_shift"`
	YShift float64 `json:"y_shift"`
	Scale  float64 `json:"scale"`
}

// Invoice type with fields compatible to the Telegram Bot API.
// https://core.telegram.org/bots/api#Invoice
type Invoice struct {
	Title          string `json:"title"`
	Description    string `json:"description"`
	StartParameter string `json:"start_parameter"`
	Currency       string `json:"currency"`
	TotalAmount    int64  `json:"total_amount"`
}

// SuccessfulPayment type with fields compatible to the Telegram Bot API.
// https://core.telegram.org/bots/api#SuccessfulPayment
type SuccessfulPayment struct {
	Currency                string     `json:"currency"`
	TotalAmount             int64      `json:"total_amount"`
	InvoicePayload          string     `json:"invoice_payload"`
	ShippingOptionID        string     `json:"shipping_option_id"`
	OrderInfo               *OrderInfo `json:"order_info"`
	TelegramPaymentChargeID string     `json:"telegram_payment_charge_id"`
	ProviderPaymentChargeID string     `json:"provider_payment_charge_id"`
}

// OrderInfo type with fields compatible to the Telegram Bot API.
// https://core.telegram.org/bots/api#OrderInfo
type OrderInfo struct {
	Name            string           `json:"name"`
	PhoneNumber     string           `json:"phone_number"`
	Email           string           `json:"email"`
	ShippingAddress *ShippingAddress `json:"shipping_address"`
}

// ShippingAddress type with fields compatible to the Telegram Bot API.
// https://core.telegram.org/bots/api#ShippingAddress
type ShippingAddress struct {
	CountryCode string `json:"country_code"`
	State       string `json:"state"`
	City        string `json:"city"`
	StreetLine1 string `json:"street_line1"`
	StreetLine2 string `json:"street_line2"`
	PostCode    string `json:"post_code"`
}

// InlineKeyboardButton type with fields compatible to the Telegram Bot API.
// https://core.telegram.org/bots/api#InlineKeyboardButton
type InlineKeyboardButton struct {
	Text                         string        `json:"text"`
	URL                          string        `json:"url"`
	CallbackData                 string        `json:"callback_data"`
	SwitchInlineQuery            string        `json:"switch_inline_query"`
	SwitchInlineQueryCurrentChat string        `json:"switch_inline_query_current_chat"`
	CallbackGame                 *CallbackGame `json:"callback_game"`
	Pay                          bool          `json:"pay"`
}

// CallbackGame type with fields compatible to the Telegram Bot API.
// https://core.telegram.org/bots/api#CallbackGame
type CallbackGame struct {
	UserID             int64  `json:"user_id"`
	Score              int64  `json:"score"`
	Force              bool   `json:"force"`
	DisableEditMessage bool   `json:"disable_edit_message"`
	ChatID             int64  `json:"chat_id"`
	MessageID          int64  `json:"message_id"`
	InlineMessageID    string `json:"inline_message_id"`
}

// InlineQuery type with fields compatible to the Telegram Bot API.
// https://core.telegram.org/bots/api#InlineQuery
type InlineQuery struct {
	ID       string    `json:"id"`
	From     *User     `json:"from"`
	Location *Location `json:"location"`
	Query    string    `json:"query"`
	Offset   string    `json:"offset"`
}

// ChosenInlineResult type with fields compatible to the Telegram Bot API.
// https://core.telegram.org/bots/api#ChosenInlineResult
type ChosenInlineResult struct {
	ResultID        string    `json:"result_id"`
	From            *User     `json:"from"`
	Location        *Location `json:"location"`
	InlineMessageID string    `json:"inline_message_id"`
	Query           string    `json:"query"`
}

// ShippingQuery type with fields compatible to the Telegram Bot API.
// https://core.telegram.org/bots/api#ShippingQuery
type ShippingQuery struct {
	ID              string           `json:"id"`
	From            *User            `json:"from"`
	InvoicePayload  string           `json:"invoice_payload"`
	ShippingAddress *ShippingAddress `json:"shipping_address"`
}

// PreCheckoutQuery type with fields compatible to the Telegram Bot API.
// https://core.telegram.org/bots/api#PreCheckoutQuery
type PreCheckoutQuery struct {
	ID               string     `json:"id"`
	From             *User      `json:"from"`
	Currency         string     `json:"currency"`
	TotalAmount      int64      `json:"total_amount"`
	InvoicePayload   string     `json:"invoice_payload"`
	ShippingOptionID string     `json:"shipping_option_id"`
	OrderInfo        *OrderInfo `json:"order_info"`
}
