package endpoint

import (
	"errors"
	"fmt"
)

var (
	// ErrUnsupportedProtocol will return if protocol is unsupported.
	ErrUnsupportedProtocol = errors.New("unsupported protocol")
	// ErrInvalidProtocolValue means value is invalid.
	ErrInvalidValue = errors.New("invalid value")
)

// Error represents error related to endpoint.
type Error struct {
	Op       string
	Protocol string
	Values   []string

	Err error
}

func (e *Error) Error() string {
	if e.Values == nil {
		return fmt.Sprintf("%s %s: %s", e.Op, e.Protocol, e.Err.Error())
	}
	return fmt.Sprintf("%s %s %s: %s", e.Op, e.Protocol, e.Values, e.Err.Error())
}

// Unwrap implements xerrors.Wrapper
func (e *Error) Unwrap() error {
	return e.Err
}