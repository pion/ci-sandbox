// SPDX-FileCopyrightText: 2026 The Pion community <https://pion.ly>
// SPDX-License-Identifier: MIT

package sandbox

import (
	"testing"

	_ "golang.org/x/net/nettest"
)

func TestDummy(*testing.T) {
	p := &PublicAPI{}
	p.PublicFunc()
}
