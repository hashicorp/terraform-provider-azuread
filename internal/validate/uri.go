package validate

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/hashicorp/go-cty/cty"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func IsAppUri(i interface{}, path cty.Path) diag.Diagnostics {
	return IsUriFunc([]string{"http", "https", "api", "ms-appx"}, true, false, false)(i, path)
}

func IsHttpOrHttpsUrl(i interface{}, path cty.Path) diag.Diagnostics {
	return IsUriFunc([]string{"http", "https"}, false, true, false)(i, path)
}

func IsHttpsUrl(i interface{}, path cty.Path) diag.Diagnostics {
	return IsUriFunc([]string{"https"}, false, true, false)(i, path)
}

func IsLogoutUrl(i interface{}, path cty.Path) (ret diag.Diagnostics) {
	ret = IsUriFunc([]string{"http", "https"}, false, true, false)(i, path)
	if len(ret) > 0 {
		return
	}

	if len(i.(string)) > 255 {
		ret = append(ret, diag.Diagnostic{
			Severity:      diag.Error,
			Summary:       "URL must be 255 characters or less",
			AttributePath: path,
		})
	}

	return
}

func IsRedirectUriFunc(urnAllowed bool, publicClient bool) schema.SchemaValidateDiagFunc {
	return func(i interface{}, path cty.Path) (ret diag.Diagnostics) {
		// See https://docs.microsoft.com/en-us/azure/active-directory-b2c/tutorial-create-user-flows?pivots=b2c-custom-policy#register-the-proxyidentityexperienceframework-application
		var allowedSchemes []string
		if !publicClient {
			allowedSchemes = []string{"http", "https", "ms-appx-web"}
		}

		ret = IsUriFunc(allowedSchemes, urnAllowed, true, true)(i, path)
		if len(ret) > 0 {
			return
		}

		if len(i.(string)) > 256 {
			ret = append(ret, diag.Diagnostic{
				Severity:      diag.Error,
				Summary:       "URI must be 256 characters or less",
				AttributePath: path,
			})
		}

		return
	}
}

func IsUriFunc(validURLSchemes []string, urnAllowed bool, allowTrailingSlash bool, forceTrailingSlash bool) schema.SchemaValidateDiagFunc {
	return func(i interface{}, path cty.Path) (ret diag.Diagnostics) {
		v, ok := i.(string)
		if !ok {
			ret = append(ret, diag.Diagnostic{
				Severity:      diag.Error,
				Summary:       "Expected a string value",
				AttributePath: path,
			})
			return
		}

		if v == "" {
			ret = append(ret, diag.Diagnostic{
				Severity:      diag.Error,
				Summary:       "URI must not be empty",
				AttributePath: path,
			})
			return
		}

		if urnAllowed {
			parts := strings.Split(v, ":")
			if len(parts) >= 3 && parts[0] == "urn" {
				return
			}
		}

		u, err := url.Parse(v)
		if err != nil {
			ret = append(ret, diag.Diagnostic{
				Severity:      diag.Error,
				Summary:       "URI is in an invalid format",
				Detail:        err.Error(),
				AttributePath: path,
			})
			return
		}

		if !allowTrailingSlash && u.Path == "/" {
			ret = append(ret, diag.Diagnostic{
				Severity:      diag.Error,
				Summary:       "URI must not have a trailing slash when there is no path segment",
				AttributePath: path,
			})
			return
		}

		if u.Host == "" {
			ret = append(ret, diag.Diagnostic{
				Severity:      diag.Error,
				Summary:       "URI has no host",
				AttributePath: path,
			})
			return
		}

		if validURLSchemes == nil {
			return
		}

		if forceTrailingSlash && u.Path == "" {
			ret = append(ret, diag.Diagnostic{
				Severity:      diag.Error,
				Summary:       "URI must have a trailing slash when there is no path segment",
				AttributePath: path,
			})
			return
		}

		for _, s := range validURLSchemes {
			if u.Scheme == s {
				return
			}
		}

		ret = append(ret, diag.Diagnostic{
			Severity:      diag.Error,
			Summary:       fmt.Sprintf("Expected URI to have a scheme of: %s", strings.Join(validURLSchemes, ", ")),
			AttributePath: path,
		})

		return
	}
}
