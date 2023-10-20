package pluginsdk

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
)

const (
	DiagError   = diag.Error
	DiagWarning = diag.Warning
)

type (
	Diagnostic  = diag.Diagnostic
	Diagnostics = diag.Diagnostics
)

// DiagFromErr is a wrapper around the DiagFromErr function in the Plugin SDK diag package
func DiagFromErr(err error) Diagnostics {
	if err == nil {
		return nil
	}
	return Diagnostics{
		Diagnostic{
			Severity: DiagError,
			Summary:  err.Error(),
		},
	}
}

// DiagErrorf is a wrapper around the DiagErrorf function in the Plugin SDK diag package
func DiagErrorf(format string, a ...interface{}) Diagnostics {
	return Diagnostics{
		Diagnostic{
			Severity: DiagError,
			Summary:  fmt.Sprintf(format, a...),
		},
	}
}
