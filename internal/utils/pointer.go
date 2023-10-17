// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package utils

import "github.com/manicminer/hamilton/msgraph"

// Deprecated: Please use pointer.To() from hashicorp/go-azure-helpers
func Bool(input bool) *bool {
	return &input
}

// Deprecated: Please use pointer.To() from hashicorp/go-azure-helpers
func Int32(input int32) *int32 {
	return &input
}

// Deprecated: Please use pointer.To() from hashicorp/go-azure-helpers
func String(input string) *string {
	return &input
}

// Deprecated: Please use tf.NullableString() instead
func NullableString(input string) *msgraph.StringNullWhenEmpty {
	output := msgraph.StringNullWhenEmpty(input)
	return &output
}
