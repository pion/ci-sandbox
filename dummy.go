// Package sandbox is a CI sandbox.
package sandbox

import (
	"errors"

	_ "github.com/pion/transport/test" // nolint
)

var ErrDummy = errors.New("dummy")
