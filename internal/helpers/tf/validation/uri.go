// Copyright IBM Corp. 2019, 2025
// SPDX-License-Identifier: MPL-2.0

package validation

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/hashicorp/terraform-provider-azuread/internal/helpers/tf/pluginsdk"
)

func IsAppUri(i interface{}, k string) (warnings []string, errors []error) {
	return IsUriFunc([]string{"http", "https", "api", "ms-appx"}, true, false, false)(i, k)
}

func IsHttpOrHttpsUrl(i interface{}, k string) (warnings []string, errors []error) {
	return IsUriFunc([]string{"http", "https"}, false, true, false)(i, k)
}

func IsHttpsUrl(i interface{}, k string) (warnings []string, errors []error) {
	return IsUriFunc([]string{"https"}, false, true, false)(i, k)
}

func IsLogoutUrl(i interface{}, k string) (warnings []string, errors []error) {
	warnings, errors = IsUriFunc([]string{"http", "https"}, false, true, false)(i, k)
	if len(errors) > 0 {
		return
	}

	if len(i.(string)) > 255 {
		errors = append(errors, fmt.Errorf("URL must be 255 characters or less for %q", k))
	}

	return
}

func IsRedirectUriFunc(urnAllowed bool, allowAllSchemes bool) pluginsdk.SchemaValidateFunc {
	return func(i interface{}, k string) (warnings []string, errors []error) {
		// See https://docs.microsoft.com/en-us/azure/active-directory-b2c/tutorial-create-user-flows?pivots=b2c-custom-policy#register-the-proxyidentityexperienceframework-application
		var allowedSchemes []string
		if !allowAllSchemes {
			allowedSchemes = []string{"http", "https", "ms-appx-web", "brk-multihub"}
		}

		warnings, errors = IsUriFunc(allowedSchemes, urnAllowed, true, true)(i, k)
		if len(errors) > 0 {
			return
		}

		if len(i.(string)) > 256 {
			errors = append(errors, fmt.Errorf("URI must be 256 characters or less for %q", k))
		}

		return
	}
}

func IsUriFunc(validUriSchemes []string, urnAllowed bool, allowTrailingSlash bool, forceTrailingSlash bool) pluginsdk.SchemaValidateFunc {
	return func(i interface{}, k string) ([]string, []error) {
		v, ok := i.(string)
		if !ok {
			return nil, []error{fmt.Errorf("expected a string value for %q", k)}
		}

		if v == "" {
			return nil, []error{fmt.Errorf("URI must not be empty for %q", k)}
		}

		if urnAllowed {
			parts := strings.Split(v, ":")
			if len(parts) >= 3 && parts[0] == "urn" {
				return nil, nil
			}
		}

		u, err := url.Parse(v)
		if err != nil {
			return nil, []error{fmt.Errorf("URI is in an invalid format for %q", k)}
		}

		if !allowTrailingSlash && u.Path == "/" {
			return nil, []error{fmt.Errorf("URI must not have a trailing slash when there is no path segment for %q", k)}
		}

		if u.Host == "" {
			return nil, []error{fmt.Errorf("URI has no host for %q", k)}
		}

		if len(validUriSchemes) == 0 {
			return nil, nil
		}

		if forceTrailingSlash && u.Path == "" {
			return nil, []error{fmt.Errorf("URI must have a trailing slash when there is no path segment for %q", k)}
		}

		for _, s := range validUriSchemes {
			if u.Scheme == s {
				return nil, nil
			}
		}

		return nil, []error{fmt.Errorf("unexpected URI scheme for %q, expected one of: %s", k, strings.Join(validUriSchemes, ", "))}
	}
}
