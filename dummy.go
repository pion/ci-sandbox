// SPDX-FileCopyrightText: 2026 The Pion community <https://pion.ly>
// SPDX-License-Identifier: MIT

// Package sandbox is a CI sandbox.
package sandbox

import (
	"errors"
	"log"

	_ "github.com/pion/transport/v3/test" // nolint
)

// ErrDummy is a dummy.
var ErrDummy = errors.New("dummy")

// PublicAPI is a dummy public API definition.
type PublicAPI struct {
	PublicMember int
}

// PublicFunc is a dummy public method.
func (p *PublicAPI) PublicFunc() {
	log.Println("Running PublicAPI.PublicFunc")             // nolint:forbidigo
	log.Println("PublicAPI.PublicMember: ", p.PublicMember) // nolint:forbidigo
	if p.PublicMember == 1 {
		log.Println("PublicAPI.PublicMember is one") // nolint:forbidigo
	}
	if p.PublicMember != 1 {
		log.Println("PublicAPI.PublicMember is not one") // nolint:forbidigo
	}
}
