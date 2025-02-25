// SPDX-FileCopyrightText: NOI Techpark <digital@noi.bz.it>
// SPDX-License-Identifier: MPL-2.0

package netex

import (
	"regexp"
	"strings"
)

// As per NeTEx spec, IDs must only contain non-accented charaters, numbers, hyphens and underscores
var idInvalid = regexp.MustCompile(`[^a-zA-Z0-9_-]`)

// Creates a NeTEx ID, converting all invalid characters to underscore "_".
func NewId(segments ...string) string {
	validSegments := []string{}
	for _, s := range segments {
		validSegments = append(validSegments, idInvalid.ReplaceAllString(s, "_"))
	}
	return strings.Join(validSegments, ":")
}

// Create a NeTEx frame ID, "edp:..."
func NewFrameId(segments ...string) string {
	return NewId(append([]string{"edp"}, segments...)...)
}

// Create a Ref of type <name>Ref
func NewRef(name string, id string, version string) Ref {
	r := Ref{}
	r.Ref = id
	r.Version = version
	r.XMLName.Local = name + "Ref"
	return r
}

// Create a TypeOfFrameRef with "epip:<name>"
func NewTypeOfFrameRef(name string, version string) Ref {
	r := Ref{}
	r.Ref = "epip:" + name
	r.Version = version
	r.XMLName.Local = "TypeOfFrameRef"
	return r
}
