// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package credentials

type CredentialError struct {
	str  string
	attr string
}

func (e CredentialError) Attr() string {
	return e.attr
}

func (e CredentialError) Error() string {
	return e.str
}
