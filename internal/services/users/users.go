// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package users

import (
	"context"
	"errors"
	"fmt"

	"github.com/hashicorp/go-azure-helpers/lang/pointer"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
	"github.com/manicminer/hamilton/msgraph"
)

func assignManager(ctx context.Context, client *msgraph.UsersClient, directoryObjectsClient *msgraph.DirectoryObjectsClient, tenantId, userId, managerId string) error {
	if managerId != "" {
		managerObject, _, err := directoryObjectsClient.Get(ctx, managerId, odata.Query{})
		if err != nil {
			return err
		}
		if managerObject == nil {
			return errors.New("managerObject was nil")
		}
		managerObject.ODataId = (*odata.Id)(pointer.To(fmt.Sprintf("%s/v1.0/%s/directoryObjects/%s",
			client.BaseClient.Endpoint, tenantId, managerId)))

		manager := msgraph.User{
			DirectoryObject: *managerObject,
		}

		if _, err = client.AssignManager(ctx, userId, manager); err != nil {
			return err
		}
	} else {
		if _, err := client.DeleteManager(ctx, userId); err != nil {
			return err
		}
	}

	return nil
}
