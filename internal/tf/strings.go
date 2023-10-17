package tf

import "github.com/manicminer/hamilton/msgraph"

func NullableString(input string) *msgraph.StringNullWhenEmpty {
	output := msgraph.StringNullWhenEmpty(input)
	return &output
}
