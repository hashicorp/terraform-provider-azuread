package tf

import (
	"fmt"

	"github.com/hashicorp/go-cty/cty"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
)

func ImportAsExistsDiag(resourceName, id string) diag.Diagnostics {
	return diag.Diagnostics{diag.Diagnostic{
		Severity:      diag.Error,
		Summary:       fmt.Sprintf("A resource with the ID %q already exists", id),
		Detail:        fmt.Sprintf("To be managed via Terraform, this resource needs to be imported into the State. Please see the resource documentation for %q for more information.", resourceName),
		AttributePath: cty.Path{cty.GetAttrStep{Name: "id"}},
	}}
}
