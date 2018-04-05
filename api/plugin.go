package emcomapi

import (
	"errors"
	"plugin"
)

// NewRouterOptions calls the function with the same name exported by the specified plugin and returns the same value
// returned by the exported function.
func NewRouterOptions(plug *plugin.Plugin) (RouterOptions, error) {
	if plug == nil {
		return nil, errors.New("invalid plugin handle")
	}

	f, err := plug.Lookup("NewRouterOptions")
	if err != nil {
		return nil, errors.New("the router plugin does not have the NewRouterOptions symbol")
	}

	fc, ok := f.(func() RouterOptions)
	if ok == false {
		return nil, errors.New("the NewRouterOptions function does not have the correct signature")
	}

	return fc(), nil
}

// NewRouter calls the function with the same name exported by the specified plugin and returns the same value returned
// by the exported function.
func NewRouter(plug *plugin.Plugin) (Router, error) {
	if plug == nil {
		return nil, errors.New("invalid plugin handle")
	}

	f, err := plug.Lookup("NewRouter")
	if err != nil {
		return nil, errors.New("the router plugin does not have the NewRouter symbol")
	}

	fc, ok := f.(func() (Router, error))
	if ok == false {
		return nil, errors.New("the NewRouter function does not have the correct signature")
	}

	return fc()
}

// NewPeripheralOptions calls the function with the same name exported by the specified plugin and returns the same
// value returned by the exported function.
func NewPeripheralOptions(plug *plugin.Plugin) (PeripheralOptions, error) {
	if plug == nil {
		return nil, errors.New("invalid plugin handle")
	}

	f, err := plug.Lookup("NewPeripheralOptions")
	if err != nil {
		return nil, errors.New("the peripheral plugin does not have the NewPeripheralOptions symbol")
	}

	fc, ok := f.(func() (PeripheralOptions, error))
	if ok == false {
		return nil, errors.New("the NewPeripheralOptions function does not have the correct signature")
	}

	return fc()
}

// NewPeripheral calls the function with the same name exported by the specified plugin and returns the same value
// returned by the exported function.
func NewPeripheral(plug *plugin.Plugin, options ...func(Peripheral) error) (Peripheral, error) {
	if plug == nil {
		return nil, errors.New("invalid plugin handle")
	}

	f, err := plug.Lookup("NewPeripheral")
	if err != nil {
		return nil, errors.New("the peripheral plugin does not have the NewPeripheral symbol")
	}

	fc, ok := f.(func(options ...func(Peripheral) error) (Peripheral, error))
	if ok == false {
		return nil, errors.New("the NewPeripheral function does not have the correct signature")
	}

	return fc(options...)
}
