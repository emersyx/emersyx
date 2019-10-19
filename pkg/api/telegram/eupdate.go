package telegram

// EUpdate is the implementation of an Event type for the Telegram gateway in emersyx. Although the Telegram Bot API
// offers the Update type which has similar purposes, it was decided to make use of an EUpdate type in order to have the
// standard types follow the official documentation. As such, the mandatory source field and GetSourceIdentifier method
// are declared for the EUpdate type.
type EUpdate struct {
	Update
	Source string
}

// GetSourceIdentifier returns the identifier of the Gateway instance which generated the emersyx event.
func (u EUpdate) GetSourceIdentifier() string {
	return u.Source
}
