package validate

import (
	"fmt"
	"time"

	"github.com/hashicorp/go-cty/cty"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-provider-azuread/internal/tf/validation"
)

func ScheduleStartDate(i interface{}, path cty.Path) (ret diag.Diagnostics) {
	v, ok := i.(string)
	if !ok {
		ret = append(ret, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Expected a string value",
		})
		return
	}

	warnings, errs := validation.IsRFC3339Time(v, "start_date")
	if len(warnings) > 0 || len(errs) > 0 {
		for _, warning := range warnings {
			ret = append(ret, diag.Diagnostic{
				Severity: diag.Warning,
				Summary:  warning,
			})
		}
		for _, err := range errs {
			ret = append(ret, diag.Diagnostic{
				Severity: diag.Error,
				Summary:  err.Error(),
			})
		}
		return
	}

	t, err := time.Parse(time.RFC3339, v)
	if err != nil {
		ret = append(ret, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  fmt.Sprintf("Failed to parse start date: %v+", err),
		})
	}

	if t.Before(time.Now()) {
		ret = append(ret, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Start date must be in the future",
		})
	}

	return
}

func ScheduleExpiryDate(i interface{}, path cty.Path) (ret diag.Diagnostics) {
	v, ok := i.(string)
	if !ok {
		ret = append(ret, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Expected a string value",
		})
		return
	}

	warnings, errs := validation.IsRFC3339Time(v, "expiration_date")
	if len(warnings) > 0 || len(errs) > 0 {
		for _, warning := range warnings {
			ret = append(ret, diag.Diagnostic{
				Severity: diag.Warning,
				Summary:  warning,
			})
		}
		for _, err := range errs {
			ret = append(ret, diag.Diagnostic{
				Severity: diag.Error,
				Summary:  err.Error(),
			})
		}
		return
	}

	t, err := time.Parse(time.RFC3339, v)
	if err != nil {
		ret = append(ret, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  fmt.Sprintf("Failed to parse start date: %v+", err),
		})
	}

	if t.Before(time.Now().Add(time.Minute * 5)) {
		ret = append(ret, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Expiry date must be at least 5 minutes in the future",
		})
	}

	return
}
