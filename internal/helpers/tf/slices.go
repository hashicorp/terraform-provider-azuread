// Copyright IBM Corp. 2014, 2025
// SPDX-License-Identifier: MPL-2.0

package tf

import "strings"

// Difference returns the elements in `a` that aren't in `b`.
func Difference(a, b []string) []string {
	mb := make(map[string]struct{}, len(b))
	for _, x := range b {
		mb[x] = struct{}{}
	}
	var diff []string
	for _, x := range a {
		if _, found := mb[x]; !found {
			diff = append(diff, x)
		}
	}
	return diff
}

// FromCommaSeparated returns a []string from the supplied comma-separated-values
func FromCommaSeparated(in string) []string {
	out := make([]string, 0)
	for _, v := range strings.Split(in, ",") {
		if v = strings.TrimSpace(v); v != "" {
			out = append(out, v)
		}
	}
	return out
}

// FromSpaceSeparated returns a []string from the supplied space-separated-values
func FromSpaceSeparated(in string) []string {
	out := make([]string, 0)
	for _, v := range strings.Split(in, " ") {
		if v = strings.TrimSpace(v); v != "" {
			out = append(out, v)
		}
	}
	return out
}
