// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package helpers

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/hashicorp/terraform-provider-azuread/internal/acceptance/types"
	"github.com/hashicorp/terraform-provider-azuread/internal/clients"
)

func ExistsInAzure(client *clients.Client, testResource types.TestResource, resourceName string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		ctx := client.StopContext

		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("%q was not found in the state", resourceName)
		}

		result, err := testResource.Exists(ctx, client, rs.Primary)
		if err != nil {
			return fmt.Errorf("running exists func for %q: %+v", resourceName, err)
		}
		if result == nil {
			return fmt.Errorf("received nil for exists for %q", resourceName)
		}

		if !*result {
			return fmt.Errorf("%q did not exist", resourceName)
		}

		return nil
	}
}
