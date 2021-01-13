package validate

import (
	"testing"

	"github.com/hashicorp/go-cty/cty"
)

func TestURLIsHTTPS(t *testing.T) {
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
			diags := URLIsHTTPS(tc.Url, cty.Path{})

			if len(diags) != tc.Errors {
				t.Fatalf("Expected URLIsHTTPS to have %d not %d errors for %q", tc.Errors, len(diags), tc.Url)
			}
		})
	}
}

func TestURLIsHTTPOrHTTPS(t *testing.T) {
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
			diags := URLIsHTTPOrHTTPS(tc.Url, cty.Path{})

			if len(diags) != tc.Errors {
				t.Fatalf("Expected URLIsHTTPOrHTTPS to have %d not %d errors for %q", tc.Errors, len(diags), tc.Url)
			}
		})
	}
}

func TestURLIsAppURI(t *testing.T) {
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
			Url:    "urn://www.example.com",
			Errors: 0,
		},
		{
			Url:    "ms-appx://www.example.com",
			Errors: 0,
		},
	}

	for _, tc := range cases {
		t.Run(tc.Url, func(t *testing.T) {
			diags := URLIsAppURI(tc.Url, cty.Path{})

			if len(diags) != tc.Errors {
				t.Fatalf("Expected URLIsAppURI to have %d not %d errors for %q", tc.Errors, len(diags), tc.Url)
			}
		})
	}
}
