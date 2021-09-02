package utils

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

// EnsureStringInSlice ensures the given string is contained in a slice
func EnsureStringInSlice(sl []string, in string) []string {
	for _, s := range sl {
		if strings.EqualFold(s, in) {
			return sl
		}
	}
	return append(sl, in)
}
