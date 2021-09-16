package validate

import (
	"testing"

	"github.com/hashicorp/go-cty/cty"
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
			diags := IsHttpsUrl(tc.Url, cty.Path{})

			if len(diags) != tc.Errors {
				t.Fatalf("Expected URLIsHTTPS to have %d not %d errors for %q", tc.Errors, len(diags), tc.Url)
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
			diags := IsHttpOrHttpsUrl(tc.Url, cty.Path{})

			if len(diags) != tc.Errors {
				t.Fatalf("Expected URLIsHTTPOrHTTPS to have %d not %d errors for %q", tc.Errors, len(diags), tc.Url)
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
			diags := IsAppUri(tc.Url, cty.Path{})

			if len(diags) != tc.Errors {
				t.Fatalf("Expected URLIsAppURI to have %d not %d errors for %q", tc.Errors, len(diags), tc.Url)
			}
		})
	}
}
