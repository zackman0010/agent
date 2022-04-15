package component

import (
	"context"
	"fmt"
	"net/http"

	"github.com/go-kit/log"
)

// Options are provided to a Component when it is being constructed.
type Options struct {
	// ID of the component. Guaranteed to be globally unique across all
	// components.
	ComponentID string
	Logger      log.Logger
}

// Component is a flow component. Flow components run in the background and
// optionally emit state.
type Component[Config any] interface {
	// Run starts the component, blocking until the provided context is canceled
	// or an error occurs.
	//
	// Components which have an output state may invoke onStateChange any time to
	// queue re-processing the state of the component.
	Run(ctx context.Context, onStateChange func()) error

	// CurrentState returns the current state of the component. Components may
	// return nil if there is no output state.
	//
	// CurrentState may be called at any time and must be safe to call
	// concurrently while the component updates its internal state.
	CurrentState() interface{}

	// Config returns the loaded Config of the component.
	Config() Config
}

// UpdatableComponent is an optional extension interface that Components may
// implement. Components that do not implement UpdatableComponent are updated
// by being shut down and replaced with a new instance constructed with the
// newest config.
type UpdatableComponent[Config any] interface {
	Component[Config]

	// Update provides a new Config to the component. An error may be returned if
	// the provided config object is invalid.
	Update(c Config) error
}

// HTTPComponent is an optional extension interface that Components which wish
// to register HTTP endpoints may implement.
type HTTPComponent[Config any] interface {
	Component[Config]

	// ComponentHandler returns an http.Handler for the current component.
	// ComponentHandler may return nil to avoid registering any handlers.
	// ComponentHandler will only be invoked once per component.
	//
	// Each Component has a unique HTTP path prefix where its handler can be
	// reached. This prefix is trimmed when invoking the http.Handler. Use
	// HTTPPrefix to determine what that prefix is.
	ComponentHandler() (http.Handler, error)
}

// HTTPPrefix returns the URL path prefix assigned to a specific componentID.
// The path returned by HTTPPrefix ends in a trailing slash.
func HTTPPrefix(componentID string) string {
	return fmt.Sprintf("/component/%s/", componentID)
}