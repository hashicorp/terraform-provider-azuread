// Copyright IBM Corp. 2014, 2025
// SPDX-License-Identifier: MPL-2.0

package tf

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func Set(d *schema.ResourceData, attr string, value interface{}) diag.Diagnostics {
	//lintignore:R001
	if err := d.Set(attr, value); err != nil {
		return ErrorDiagPathF(err, attr, "Could not set attribute")
	}
	return nil
}
