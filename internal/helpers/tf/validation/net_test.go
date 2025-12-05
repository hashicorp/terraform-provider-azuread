// Copyright IBM Corp. 2019, 2025
// SPDX-License-Identifier: MPL-2.0

package validation

import (
	"testing"
)

func TestStringIsIpPrefix(t *testing.T) {
	cases := []struct {
		Value    string
		TestName string
		ErrCount int
	}{
		{
			Value:    "10.0.0.0/8",
			TestName: "Valid_NonRoutable1",
			ErrCount: 0,
		},
		{
			Value:    "192.168.0.0/16",
			TestName: "Valid_NonRoutable2",
			ErrCount: 0,
		},
		{
			Value:    "172.16.20.5",
			TestName: "Invalid_SingleAddress",
			ErrCount: 1,
		},
		{
			Value:    "224.0.50.8",
			TestName: "Invalid_MulticastAddress",
			ErrCount: 1,
		},
		{
			Value:    "100.64.10.0",
			TestName: "Invalid_Network",
			ErrCount: 1,
		},
	}

	for _, tc := range cases {
		t.Run(tc.TestName, func(t *testing.T) {
			warnings, errors := StringIsIpPrefix(tc.Value, "test")

			if len(warnings) > 0 {
				t.Fatalf("Expected StringIsIpPrefix to have %d not %d warnings for %q", 0, len(warnings), tc.TestName)
			}
			if len(errors) != tc.ErrCount {
				t.Fatalf("Expected StringIsIpPrefix to have %d not %d errors for %q", tc.ErrCount, len(errors), tc.TestName)
			}
		})
	}
}

func TestPrefixLengthAtLeast(t *testing.T) {
	cases := []struct {
		MinLength int
		Value     string
		TestName  string
		ErrCount  int
	}{
		{
			MinLength: 8,
			Value:     "10.0.0.0/8",
			TestName:  "Valid_Exact",
			ErrCount:  0,
		},
		{
			MinLength: 16,
			Value:     "192.168.0.0/24",
			TestName:  "Valid_Larger",
			ErrCount:  0,
		},
		{
			MinLength: 8,
			Value:     "10.0.0.0/4",
			TestName:  "Invalid_Smaller",
			ErrCount:  1,
		},
	}

	for _, tc := range cases {
		t.Run(tc.TestName, func(t *testing.T) {
			warnings, errors := PrefixLengthAtLeast(tc.MinLength)(tc.Value, "test")

			if len(warnings) > 0 {
				t.Fatalf("Expected PrefixLengthAtLeast to have %d not %d warnings for %q", 0, len(warnings), tc.TestName)
			}
			if len(errors) != tc.ErrCount {
				t.Fatalf("Expected PrefixLengthAtLeast to have %d not %d errors for %q", tc.ErrCount, len(errors), tc.TestName)
			}
		})
	}
}

func TestPrefixLengthAtMost(t *testing.T) {
	cases := []struct {
		MaxLength int
		Value     string
		TestName  string
		ErrCount  int
	}{
		{
			MaxLength: 24,
			Value:     "192.168.0.0/24",
			TestName:  "Valid_Exact",
			ErrCount:  0,
		},
		{
			MaxLength: 16,
			Value:     "10.0.0.0/8",
			TestName:  "Valid_Smaller",
			ErrCount:  0,
		},
		{
			MaxLength: 8,
			Value:     "10.0.0.0/12",
			TestName:  "Invalid_Larger",
			ErrCount:  1,
		},
	}

	for _, tc := range cases {
		t.Run(tc.TestName, func(t *testing.T) {
			warnings, errors := PrefixLengthAtMost(tc.MaxLength)(tc.Value, "test")

			if len(warnings) > 0 {
				t.Fatalf("Expected PrefixLengthAtMost to have %d not %d warnings for %q", 0, len(warnings), tc.TestName)
			}
			if len(errors) != tc.ErrCount {
				t.Fatalf("Expected PrefixLengthAtMost to have %d not %d errors for %q", tc.ErrCount, len(errors), tc.TestName)
			}
		})
	}
}

func TestPrefixLengthBetween(t *testing.T) {
	cases := []struct {
		MinLength int
		MaxLength int
		Value     string
		TestName  string
		ErrCount  int
	}{
		{
			MinLength: 16,
			MaxLength: 24,
			Value:     "192.168.0.0/24",
			TestName:  "Valid_ExactUpper",
			ErrCount:  0,
		},
		{
			MinLength: 16,
			MaxLength: 24,
			Value:     "172.16.0.0/16",
			TestName:  "Valid_ExactLower",
			ErrCount:  0,
		},
		{
			MinLength: 8,
			MaxLength: 16,
			Value:     "10.50.0.0/12",
			TestName:  "Valid_InRange",
			ErrCount:  0,
		},
		{
			MinLength: 24,
			MaxLength: 28,
			Value:     "10.0.0.0/12",
			TestName:  "Invalid_Smaller",
			ErrCount:  1,
		},
		{
			MinLength: 24,
			MaxLength: 28,
			Value:     "192.168.100.0/30",
			TestName:  "Invalid_Larger",
			ErrCount:  1,
		},
	}

	for _, tc := range cases {
		t.Run(tc.TestName, func(t *testing.T) {
			warnings, errors := PrefixLengthBetween(tc.MinLength, tc.MaxLength)(tc.Value, "test")

			if len(warnings) > 0 {
				t.Fatalf("Expected PrefixLengthBetween to have %d not %d warnings for %q", 0, len(warnings), tc.TestName)
			}
			if len(errors) != tc.ErrCount {
				t.Fatalf("Expected PrefixLengthBetween to have %d not %d errors for %q", tc.ErrCount, len(errors), tc.TestName)
			}
		})
	}
}
