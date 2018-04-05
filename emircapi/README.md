# emersyx IRC APIs

## Go plugin API

The interfaces, structs and functions defined in these sources must be followed by IRC gateway implementations. A
complete go plugin needs to follow the rules below:

* implement the `IRCGateway` interface
* use of `Message` as the default emersyx event type
* export the `NewPeripheral` and `NewPeripheralOptions` functions, as required by the wrappers with the same names
  defined in the `emcomapi/plugin.go` file

## Using implementations

Implementations of this API are to be distributed as either one go plugin file (e.g. one `.so` file for linux platforms)
or as source code which can be built into one go plugin.

The wrapper functions `NewPeripheral` and `NewPeripheralOptions` can be used to load the plugin files and call the
exported functions.

## Example implementation

An example implementation of an IRC gateway for emersyx can be found in the [emersyx_irc][1] repository.

[1]: https://github.com/emersyx/emersyx_irc
