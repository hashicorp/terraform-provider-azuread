// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package validation

import (
	"testing"
)

func TestISO639Language(t *testing.T) {
	cases := []struct {
		Value    string
		TestName string
		ErrCount int
	}{
		{
			Value:    "fr",
			TestName: "Language",
			ErrCount: 0,
		},
		{
			Value:    "es-MX",
			TestName: "LanguageLocale",
			ErrCount: 0,
		},
		{
			Value:    "cr-belter",
			TestName: "UnrecognisedButValid",
			ErrCount: 0,
		},
		{
			Value:    "en-en-GB",
			TestName: "Malformed",
			ErrCount: 1,
		},
		{
			Value:    "192",
			TestName: "NumericInvalid",
			ErrCount: 1,
		},
		{
			Value:    "@#$%^&*",
			TestName: "Gibberish",
			ErrCount: 1,
		},
	}

	for _, tc := range cases {
		t.Run(tc.TestName, func(t *testing.T) {
			_, errs := ISO639Language(tc.Value, "test")

			if len(errs) != tc.ErrCount {
				t.Fatalf("Expected ISO639Language to have %d not %d errors for %q", tc.ErrCount, len(errs), tc.TestName)
			}
		})
	}
}
