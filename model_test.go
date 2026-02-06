// SPDX-FileCopyrightText: 2025 NOI Techpark <digital@noi.bz.it>
//
// SPDX-License-Identifier: MPL-2.0

package netex

import (
	"encoding/xml"
	"os"
	"strings"
	"testing"
)

func TestFleetMembers_Empty(t *testing.T) {
	fleet := Fleet{
		Id:      NewId("test"),
		Version: "1.0",
	}
	fleet.Members = Just([]Ref{NewRef("test", "test", "version")})

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

func Test_unmarshal(t *testing.T) {
	// Contains some ISO 8601 (and notably not RFC 3339) dates
	data, err := os.ReadFile("./testdata/sta.xml")
	if err != nil {
		t.Fatalf("failed opening trains file: %s", err)
	}

	var delivery PublicationDelivery
	err = xml.Unmarshal(data, &delivery)
	if err != nil {
		t.Fatalf("failed unmarshal: %s", err)
	}
}
