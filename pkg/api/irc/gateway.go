package irc

import "emersyx.net/emersyx/pkg/api"

// Gateway is the interface which for an IRC peripheral and receptor. The reference implementation at
// https://github.com/emersyx/emersyx_irc follows this interface.
type Gateway interface {
	api.Peripheral
	api.Receptor
	// Quit must disconnect the Gateway from the IRC server. This must be a blocking method. When the method returns,
	// the instance must not be connected to the IRC server anymore.
	Quit() error
	// Join method must send a JOIN command to the IRC server. The argument specifies the channel to be joined. If the
	// instance is not connected to any IRC server, then an error is returned.
	Join(channel string) error
	// Privmsg must send a PRIVMSG command to the IRC server. The first argument specifies the destination (i.e.  either
	// a user or a channel) and the second argument is the actual message. If the instance is not connected to any IRC
	// server, then an error is returned.
	Privmsg(destination string, message string) error
}
