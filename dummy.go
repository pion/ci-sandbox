// SPDX-FileCopyrightText: 2023 The Pion community <https://pion.ly>
// SPDX-License-Identifier: MIT

// Package sandbox is a CI sandbox.
package sandbox

import (
	"errors"

	_ "github.com/pion/transport/v3/test" // nolint
)

// ErrDummy is a dummy.
var ErrDummy = errors.New("dummy")
