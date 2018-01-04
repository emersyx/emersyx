package emircapi

import (
	"emersyx.net/emersyx_apis/emcomapi"
)

// IRCResource is the interface which specifies the methods that must be implemented by any supported IRC resource.
type IRCResource interface {
	// Connect must start the connection process to the selected IRC server. This must be a blocking method. When the
	// method returns, the IRCResource must connected to the IRC server if the return value is nil. Otherwise, it is
	// considered that an error occured and the connection was not possible.
	Connect() error
	// IsConnected must return a boolean which is true if the bot is connected to the server, and false otherwise.
	IsConnected() bool
	// Quit must disconnect the IRCResource from the IRC server. This must be a blocking method. When the method
	// returns, the instance must not be connected to the IRC server anymore.
	Quit() error
	// Join method must send a JOIN command to the IRC server. The argument specifies the channel to be joined. If the
	// instance is not connected to any IRC server, then an error is returned.
	Join(channel string) error
	// Privmsg must send a PRIVMSG command to the IRC server. The first argument specifies the destination (i.e.  either
	// a user or a channel) and the second argument is the actual message. If the instance is not connected to any IRC
	// server, then an error is returned.
	Privmsg(destination string, message string) error
}

// IRCBot is the interface which specifies the methods that must be implemented by a complete IRC receptor and resource.
// The reference implementation at https://github.com/emersyx/emersyx_irc follows this interface.
type IRCBot interface {
	emcomapi.Receptor
	IRCResource
}
