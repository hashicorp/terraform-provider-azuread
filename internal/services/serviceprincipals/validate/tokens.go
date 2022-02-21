package validate

import (
	"regexp"

	"github.com/hashicorp/go-cty/cty"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
)

// RoleScopeClaimValue checks whether a value is valid for use in a `role` or `scp` claim, as used in App Roles and OAuth 2.0 Permission Scopes in Applications.
// See https://docs.microsoft.com/en-us/graph/api/resources/approle?view=graph-rest-beta and https://docs.microsoft.com/en-us/graph/api/resources/permissionscope?view=graph-rest-beta
func RoleScopeClaimValue(i interface{}, path cty.Path) (ret diag.Diagnostics) {
	v, ok := i.(string)
	if !ok {
		ret = append(ret, diag.Diagnostic{
			Severity:      diag.Error,
			Summary:       "Expected a string value",
			AttributePath: path,
		})
		return
	}

	if len(v) > 120 {
		ret = append(ret, diag.Diagnostic{
			Severity:      diag.Error,
			Summary:       "Value must be between 0-120 characters in length",
			AttributePath: path,
		})
	}

	if len(v) > 0 && string(v[0]) == "." {
		ret = append(ret, diag.Diagnostic{
			Severity:      diag.Error,
			Summary:       "Value must not begin with a period",
			AttributePath: path,
		})
	}

	if regexp.MustCompile(`[^A-Za-z0-9:!#$%&'()*+,./:;<=>?@\[\]^+_{|}~` + "`" + `-]`).MatchString(v) {
		ret = append(ret, diag.Diagnostic{
			Severity:      diag.Error,
			Summary:       "Value must be alphanumeric with these allowed characters: !#$%&'()*+,-./:;<=>?@[]^+_`{|}~",
			AttributePath: path,
		})
	}

	return // nolint:nakedret
}
