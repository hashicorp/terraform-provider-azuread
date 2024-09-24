// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package validation

import (
	"testing"
)

func TestStringIsEmailAddress(t *testing.T) {
	cases := []struct {
		Value    string
		TestName string
		ErrCount int
	}{
		{
			Value:    "j.doe@hashicorp.com",
			TestName: "Valid_EmailAddress",
			ErrCount: 0,
		},
		{
			Value:    "j.doehashicorp.com",
			TestName: "Invalid_EmailAddress_NoAtChar",
			ErrCount: 1,
		},
		{
			Value:    "j/doe@ha$hicorp.com",
			TestName: "Invalid_EmailAddress_InvalidChars",
			ErrCount: 1,
		},
	}

	for _, tc := range cases {
		t.Run(tc.TestName, func(t *testing.T) {
			_, errs := StringIsEmailAddress(tc.Value, "test")

			if len(errs) != tc.ErrCount {
				t.Fatalf("Expected StringIsEmailAddress to have %d not %d errors for %q", tc.ErrCount, len(errs), tc.TestName)
			}
		})
	}
}
