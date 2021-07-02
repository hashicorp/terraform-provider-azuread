package validate

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/hashicorp/go-cty/cty"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func IsHTTPSURL(i interface{}, path cty.Path) diag.Diagnostics {
	return IsURIFunc([]string{"https"}, false)(i, path)
}

func IsHTTPOrHTTPSURL(i interface{}, path cty.Path) diag.Diagnostics {
	return IsURIFunc([]string{"http", "https"}, false)(i, path)
}

func IsAppURI(i interface{}, path cty.Path) diag.Diagnostics {
	return IsURIFunc([]string{"http", "https", "api", "ms-appx"}, true)(i, path)
}

func IsURIFunc(validURLSchemes []string, URNAllowed bool) schema.SchemaValidateDiagFunc {
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

		if URNAllowed {
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

		if u.Host == "" {
			ret = append(ret, diag.Diagnostic{
				Severity:      diag.Error,
				Summary:       "URI has no host",
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
