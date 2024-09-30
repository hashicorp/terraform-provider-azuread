// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package credentials

const (
	KeyCredentialUsageSign   = "Sign"
	KeyCredentialUsageVerify = "Verify"
)

var PossibleValuesForKeyCredentialUsage = []string{KeyCredentialUsageSign, KeyCredentialUsageVerify}
