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

// NewProcessorOptions calls the function with the same name exported by the specified plugin and returns the same value
// returned by the exported function.
func NewProcessorOptions(plug *plugin.Plugin) (ProcessorOptions, error) {
	if plug == nil {
		return nil, errors.New("invalid plugin handle")
	}

	f, err := plug.Lookup("NewProcessorOptions")
	if err != nil {
		return nil, errors.New("the processor plugin does not have the NewProcessorOptions symbol")
	}

	fc, ok := f.(func() (ProcessorOptions, error))
	if ok == false {
		return nil, errors.New("the NewProcessorOptions function does not have the correct signature")
	}

	return fc()
}

// NewProcessor calls the function with the same name exported by the specified plugin and returns the same value
// returned by the exported function.
func NewProcessor(plug *plugin.Plugin, options ...func(Processor) error) (Processor, error) {
	if plug == nil {
		return nil, errors.New("invalid plugin handle")
	}

	f, err := plug.Lookup("NewProcessor")
	if err != nil {
		return nil, errors.New("the processor plugin does not have the NewProcessor symbol")
	}

	fc, ok := f.(func(options ...func(Processor) error) (Processor, error))
	if ok == false {
		return nil, errors.New("the NewProcessor function does not have the correct signature")
	}

	return fc(options...)
}
