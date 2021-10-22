package utils

import "strings"

// EscapeSingleQuote replaces all occurrences of single quote, with 2 single quotes.
// For requests that use single quotes, if any parameter values also contain single quotes,
// those must be double escaped; otherwise, the request will fail due to invalid syntax.
// https://docs.microsoft.com/en-us/graph/query-parameters#escaping-single-quotes
func EscapeSingleQuote(qparam string) string {
	return strings.ReplaceAll(qparam, `'`, `''`)
}
