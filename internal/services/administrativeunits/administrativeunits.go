// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package administrativeunits

import (
	"context"
	"fmt"

	"github.com/hashicorp/go-azure-helpers/lang/pointer"
	"github.com/hashicorp/go-azure-helpers/lang/response"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/directory/stable/administrativeunit"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/directory/stable/administrativeunitmember"
)

func administrativeUnitFindByName(ctx context.Context, client *administrativeunit.AdministrativeUnitClient, displayName string) (*[]stable.AdministrativeUnit, error) {
	options := administrativeunit.ListAdministrativeUnitsOperationOptions{
		Filter: pointer.To(fmt.Sprintf("displayName eq '%s'", displayName)),
	}
	resp, err := client.ListAdministrativeUnits(ctx, options)
	if err != nil {
		return nil, fmt.Errorf("unable to list Administrative Units with filter %q: %+v", *options.Filter, err)
	}

	result := make([]stable.AdministrativeUnit, 0)
	if administrativeUnits := resp.Model; administrativeUnits != nil {
		for _, au := range *administrativeUnits {
			if au.DisplayName.GetOrZero() == displayName {
				result = append(result, au)
			}
		}
	}

	return &result, nil
}

func administrativeUnitGetMember(ctx context.Context, client *administrativeunitmember.AdministrativeUnitMemberClient, id string, memberId string) (*stable.DirectoryObject, error) {
	options := administrativeunitmember.ListAdministrativeUnitMembersOperationOptions{
		Filter: pointer.To(fmt.Sprintf("id eq '%s'", memberId)),
	}

	resp, err := client.ListAdministrativeUnitMembers(ctx, stable.NewDirectoryAdministrativeUnitID(id), options)
	if err != nil {
		if response.WasNotFound(resp.HttpResponse) {
			return nil, nil
		} else {
			return nil, err
		}
	}

	if resp.Model != nil {
		for _, member := range *resp.Model {
			if member.DirectoryObject().Id != nil && *member.DirectoryObject().Id == memberId {
				return &member, nil
			}
		}
	}

	return nil, nil
}
