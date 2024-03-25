// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package tf

import "github.com/manicminer/hamilton/msgraph"

func NullableString(input string) *msgraph.StringNullWhenEmpty {
	output := msgraph.StringNullWhenEmpty(input)
	return &output
}
