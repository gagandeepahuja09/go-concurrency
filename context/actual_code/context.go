package context

import "time"

type Context interface {
	// Deadline returns the time when work done on behalf of this context
	// should be cancelled. Deadline returns ok = false if no deadline is set.
	Deadline() (deadline time.Time, ok bool)

	// Done returns a channel that's closed when work done on behalf of this channel
	// should be cancelled. That will close all its children goroutines.
	// Done may return nil if this context can never be canceled.
	// Successive calls to done return the same value.
	Done() <-chan struct{}

	Err() error

	Value()
}
