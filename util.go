// SPDX-FileCopyrightText: 2025 NOI Techpark <digital@noi.bz.it>
//
// SPDX-License-Identifier: MPL-2.0

package netex

import (
	"regexp"
)

// As per NeTEx spec, IDs must only contain non-accented charaters, numbers, hyphens and underscores
var idInvalid = regexp.MustCompile(`[^a-zA-Z0-9_-]`)

func CreateID(segments ...string) string {
	id := "IT:ITH10"
	for _, s := range segments {
		id += (":" + idInvalid.ReplaceAllString(s, "_"))
	}
	return id
}

func CreateFrameId(segments ...string) string {
	return "edp:" + CreateID(segments...)
}

func MkRef(tp string, id string) Ref {
	r := Ref{}
	r.Ref = id
	r.Version = "1"
	r.XMLName.Local = tp + "Ref"
	return r
}

func MkTypeOfFrameRef(tp string) Ref {
	r := Ref{}
	r.Ref = "epip:" + tp
	r.Version = "1"
	r.XMLName.Local = "TypeOfFrameRef"
	return r
}
