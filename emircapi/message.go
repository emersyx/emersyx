package emircapi

// The constant values below must be used for the Command member of the Message struct.
const(
    DISCONNECTED    = "DISCONNECTED"
    JOIN            = "JOIN"
    PART            = "PART"
    PRIVMSG         = "PRIVMSG"
    QUIT            = "QUIT"
)

// This is the basic structure for an IRC message received by the client when an event occurs. Names of the struct
// members have been taken from RFC-1459 and RFC-2812. This is the structure which implements the emcomapi.Event
// interface for IRC events.
type Message struct {
    Source      string
    Raw         string
    Command     string
    Origin      string
    Parameters  []string
}

// This method returns the source of the event.
func (m Message) GetSourceIdentifier() string {
    return m.Source;
}
