// Package sandbox is a CI sandbox.
package sandbox

import (
	"errors"
	"strconv"

	_ "github.com/pion/transport/test" // nolint
)

// ErrDummy is a dummy.
var ErrDummy = errors.New("dummy")

func codeQL() {
	res, _ := strconv.Atoi("10")
	num := uint8(res)
	_ = num
}
