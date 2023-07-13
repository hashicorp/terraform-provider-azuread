// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package administrativeunits

import (
	"context"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/odata"
	"github.com/manicminer/hamilton/msgraph"
)

func administrativeUnitFindByName(ctx context.Context, client *msgraph.AdministrativeUnitsClient, displayName string) (*[]msgraph.AdministrativeUnit, error) {
	query := odata.Query{
		Filter: fmt.Sprintf("displayName eq '%s'", displayName),
	}
	administrativeUnits, _, err := client.List(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("unable to list Administrative Units with filter %q: %+v", query.Filter, err)
	}

	result := make([]msgraph.AdministrativeUnit, 0)
	if administrativeUnits != nil {
		for _, au := range *administrativeUnits {
			if au.DisplayName != nil && *au.DisplayName == displayName {
				result = append(result, au)
			}
		}
	}

	return &result, nil
}
