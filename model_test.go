// SPDX-FileCopyrightText: 2025 NOI Techpark <digital@noi.bz.it>
//
// SPDX-License-Identifier: MPL-2.0

package netex_test

import (
	"encoding/xml"
	"strings"
	"testing"

	"github.com/noi-techpark/go-netex"
)

func TestFleetMembers_Empty(t *testing.T) {
	fleet := netex.Fleet{
		Id:      netex.NewId("test"),
		Version: "1.0",
	}
	fleet.Members = netex.Just([]netex.Ref{netex.NewRef("test", "test", "version")})

	b, _ := xml.Marshal(fleet)
	if !strings.Contains(string(b), "members") {
		t.Error("xml doesn't contain members even though it should", string(b))
	}

	fleet.Members = nil
	b, _ = xml.Marshal(fleet)
	if strings.Contains(string(b), "members") {
		t.Error("xml contains members even though it shouldn't", string(b))
	}
}
