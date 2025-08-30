// SPDX-FileCopyrightText: 2023 The Pion community <https://pion.ly>
// SPDX-License-Identifier: MIT

package sandbox

import (
	"fmt"
	"testing"

	_ "golang.org/x/net/nettest"
)

func TestDummy(*testing.T) {
	p := &PublicAPI{}
	p.PublicFunc()
	fmt.Println("test 2")
}
