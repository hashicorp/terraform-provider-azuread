// Copyright IBM Corp. 2019, 2025
// SPDX-License-Identifier: MPL-2.0

package validation

import (
	"testing"
)

func TestIsHTTPSURL(t *testing.T) {
	cases := []struct {
		Url    string
		Errors int
	}{
		{
			Url:    "",
			Errors: 1,
		},
		{
			Url:    "this is not a url",
			Errors: 1,
		},
		{
			Url:    "www.example.com",
			Errors: 1,
		},
		{
			Url:    "ftp://www.example.com",
			Errors: 1,
		},
		{
			Url:    "http://www.example.com",
			Errors: 1,
		},
		{
			Url:    "https://www.example.com",
			Errors: 0,
		},
	}

	for _, tc := range cases {
		t.Run(tc.Url, func(t *testing.T) {
			warnings, errors := IsHttpsUrl(tc.Url, "test")

			if len(warnings) > 0 {
				t.Fatalf("Expected URLIsHTTPS to have 0 not %d warnings for %q", len(warnings), tc.Url)
			}
			if len(errors) != tc.Errors {
				t.Fatalf("Expected URLIsHTTPS to have %d not %d errors for %q", tc.Errors, len(errors), tc.Url)
			}
		})
	}
}

func TestIsHTTPOrHTTPSURL(t *testing.T) {
	cases := []struct {
		Url    string
		Errors int
	}{
		{
			Url:    "",
			Errors: 1,
		},
		{
			Url:    "this is not a url",
			Errors: 1,
		},
		{
			Url:    "www.example.com",
			Errors: 1,
		},
		{
			Url:    "ftp://www.example.com",
			Errors: 1,
		},
		{
			Url:    "http://www.example.com",
			Errors: 0,
		},
		{
			Url:    "https://www.example.com",
			Errors: 0,
		},
	}

	for _, tc := range cases {
		t.Run(tc.Url, func(t *testing.T) {
			warnings, errors := IsHttpOrHttpsUrl(tc.Url, "test")

			if len(warnings) > 0 {
				t.Fatalf("Expected URLIsHTTPOrHTTPS to have 0 not %d warnings for %q", len(warnings), tc.Url)
			}
			if len(errors) != tc.Errors {
				t.Fatalf("Expected URLIsHTTPOrHTTPS to have %d not %d errors for %q", tc.Errors, len(errors), tc.Url)
			}
		})
	}
}

func TestIsAppURI(t *testing.T) {
	cases := []struct {
		Url    string
		Errors int
	}{
		{
			Url:    "",
			Errors: 1,
		},
		{
			Url:    "this is not a url",
			Errors: 1,
		},
		{
			Url:    "www.example.com",
			Errors: 1,
		},
		{
			Url:    "ftp://www.example.com",
			Errors: 1,
		},
		{
			Url:    "http://www.example.com",
			Errors: 0,
		},
		{
			Url:    "https://www.example.com",
			Errors: 0,
		},
		{
			Url:    "api://www.example.com",
			Errors: 0,
		},
		{
			Url:    "urn:uuid:6e8bc430-9c3a-11d9-9669-0800200c9a66",
			Errors: 0,
		},
		{
			Url:    "urn:nbn:de:bvb:19-146642",
			Errors: 0,
		},
		{
			Url:    "ms-appx://www.example.com",
			Errors: 0,
		},
	}

	for _, tc := range cases {
		t.Run(tc.Url, func(t *testing.T) {
			warnings, errors := IsAppUri(tc.Url, "test")

			if len(warnings) > 0 {
				t.Fatalf("Expected URLIsAppURI to have 0 not %d warnings for %q", len(warnings), tc.Url)
			}
			if len(errors) != tc.Errors {
				t.Fatalf("Expected URLIsAppURI to have %d not %d errors for %q", tc.Errors, len(errors), tc.Url)
			}
		})
	}
}

func TestIsUriFunc(t *testing.T) {
	cases := []struct {
		TestName           string
		Url                string
		UrnAllowed         bool
		AllowTrailingSlash bool
		ForceTrailingSlash bool
		Errors             int
		Schemes            []string
	}{
		{
			TestName:           "no path with trailing slash not allowed should error",
			Url:                "http://www.example.com/",
			UrnAllowed:         true,
			AllowTrailingSlash: false,
			ForceTrailingSlash: false,
			Errors:             1,
			Schemes:            []string{"http"},
		},
		{
			TestName:           "no path with no trailing slash valid",
			Url:                "http://www.example.com",
			UrnAllowed:         true,
			AllowTrailingSlash: false,
			ForceTrailingSlash: false,
			Errors:             0,
			Schemes:            []string{"http"},
		},
		{
			TestName:           "path with no trailing slash is valid",
			Url:                "http://www.example.com/path",
			UrnAllowed:         true,
			AllowTrailingSlash: false,
			ForceTrailingSlash: false,
			Errors:             0,
			Schemes:            []string{"http"},
		},
		{
			TestName:           "uri empty should not be valid",
			Url:                "",
			UrnAllowed:         true,
			AllowTrailingSlash: false,
			ForceTrailingSlash: false,
			Errors:             1,
			Schemes:            []string{"http"},
		},
	}

	for _, tc := range cases {
		t.Run(tc.TestName, func(t *testing.T) {
			warnings, errors := IsUriFunc(tc.Schemes, tc.UrnAllowed, tc.AllowTrailingSlash, tc.ForceTrailingSlash)(tc.Url, "test")

			if len(warnings) > 0 {
				t.Fatalf("Expected IsUriFunc() to have 0 not %d warnings for %q", len(warnings), tc.Url)
			}
			if len(errors) != tc.Errors {
				t.Fatalf("Expected IsUriFunc() to have %d not %d errors for %v", tc.Errors, len(errors), tc.Url)
			}
		})
	}
}
