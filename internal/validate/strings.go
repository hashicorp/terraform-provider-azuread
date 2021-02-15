package validate

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/hashicorp/go-cty/cty"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
)

// NoEmptyStrings validates that the string is not just whitespace characters (equal to [\r\n\t\f\v ])
func NoEmptyStrings(i interface{}, path cty.Path) (ret diag.Diagnostics) {
	v, ok := i.(string)
	if !ok {
		ret = append(ret, diag.Diagnostic{
			Severity:      diag.Error,
			Summary:       "Expected a string value",
			AttributePath: path,
		})
		return
	}

	if strings.TrimSpace(v) == "" {
		ret = append(ret, diag.Diagnostic{
			Severity:      diag.Error,
			Summary:       "Value must not be empty",
			AttributePath: path,
		})
	}

	return
}

// StringIsEmailAddress validates that the given string is a valid email address (foo@bar.com)
func StringIsEmailAddress(i interface{}, path cty.Path) (ret diag.Diagnostics) {
	v, ok := i.(string)
	if !ok {
		ret = append(ret, diag.Diagnostic{
			Severity:      diag.Error,
			Summary:       "Expected a string value",
			AttributePath: path,
		})
		return
	}

	if strings.TrimSpace(v) == "" {
		ret = append(ret, diag.Diagnostic{
			Severity:      diag.Error,
			Summary:       "Value must not be empty",
			AttributePath: path,
		})
	}

	regExIsEmailAddress := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	if !regExIsEmailAddress.MatchString(v) {
		ret = append(ret, diag.Diagnostic{
			Severity:      diag.Error,
			Summary:       "Value must be a valid email address",
			AttributePath: path,
		})
	}

	return
}

//wrapper to use validator in ValidateDiagFunc

func ValidatorWrapper(validateFunc func(interface{}, string) ([]string, []error)) schema.SchemaValidateDiagFunc {
	return func(i interface{}, path cty.Path) diag.Diagnostics {
		warnings, errs := validateFunc(i, fmt.Sprintf("%+v", path))
		var diags diag.Diagnostics
		for _, warning := range warnings {
			diags = append(diags, diag.Diagnostic{
				Severity: diag.Warning,
				Summary:  warning,
			})
		}
		for _, err := range errs {
			diags = append(diags, diag.Diagnostic{
				Severity: diag.Error,
				Summary:  err.Error(),
			})
		}
		return diags
	}
}
