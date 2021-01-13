package validate

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/hashicorp/go-cty/cty"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func URLIsHTTPS(i interface{}, path cty.Path) diag.Diagnostics {
	return URLWithScheme([]string{"https"})(i, path)
}

func URLIsHTTPOrHTTPS(i interface{}, path cty.Path) diag.Diagnostics {
	return URLWithScheme([]string{"http", "https"})(i, path)
}

func URLIsAppURI(i interface{}, path cty.Path) diag.Diagnostics {
	return URLWithScheme([]string{"http", "https", "api", "urn", "ms-appx"})(i, path)
}

func URLWithScheme(validSchemes []string) schema.SchemaValidateDiagFunc {
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
				Summary:       "URL must not be empty",
				AttributePath: path,
			})
			return
		}

		u, err := url.Parse(v)
		if err != nil {
			ret = append(ret, diag.Diagnostic{
				Severity:      diag.Error,
				Summary:       "URL is in an invalid format",
				Detail:        err.Error(),
				AttributePath: path,
			})
			return
		}

		if u.Host == "" {
			ret = append(ret, diag.Diagnostic{
				Severity:      diag.Error,
				Summary:       "URL has no host",
				AttributePath: path,
			})
			return
		}

		for _, s := range validSchemes {
			if u.Scheme == s {
				return
			}
		}

		ret = append(ret, diag.Diagnostic{
			Severity:      diag.Error,
			Summary:       fmt.Sprintf("Expected URL to have a schema of: %s", strings.Join(validSchemes, ", ")),
			AttributePath: path,
		})
		return
	}
}
