// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package acceptance

import "github.com/hashicorp/terraform-plugin-testing/helper/resource"

// ImportStep returns a Test Step which Imports the Resource, optionally
// ignoring any fields which may not be imported (for example, as they're
// not returned from the API)
func (td TestData) ImportStep(ignore ...string) resource.TestStep {
	step := resource.TestStep{
		ResourceName:      td.ResourceName,
		ImportState:       true,
		ImportStateVerify: true,
	}

	if len(ignore) > 0 {
		step.ImportStateVerifyIgnore = ignore
	}

	return step
}

// RequiresImportErrorStep returns a Test Step which expects a Requires Import
// error to be returned when running this step
func (td TestData) RequiresImportErrorStep(config string) resource.TestStep {
	return resource.TestStep{
		Config:      config,
		ExpectError: RequiresImportError(td.ResourceType),
	}
}
