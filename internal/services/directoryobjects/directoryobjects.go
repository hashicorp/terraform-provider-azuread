package directoryobjects

import (
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func formatODataType(in string) string {
	return cases.Title(language.AmericanEnglish, cases.NoLower).String(strings.TrimPrefix(in, "#microsoft.graph."))
}
