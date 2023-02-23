package users

import (
	"context"
	"errors"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/odata"
	"github.com/hashicorp/terraform-provider-azuread/internal/utils"
	"github.com/manicminer/hamilton/msgraph"
)

func assignManager(ctx context.Context, client *msgraph.UsersClient, directoryObjectsClient *msgraph.DirectoryObjectsClient, userId, managerId string) error {
	if managerId != "" {
		managerObject, _, err := directoryObjectsClient.Get(ctx, managerId, odata.Query{})
		if err != nil {
			return err
		}
		if managerObject == nil {
			return errors.New("managerObject was nil")
		}
		// TODO: remove this workaround for https://github.com/hashicorp/terraform-provider-azuread/issues/588
		//if managerObject.ODataId == nil {
		//	return tf.ErrorDiagF(errors.New("ODataId was nil"), "Could not retrieve manager principal object %q", managerId)
		//}
		managerObject.ODataId = (*odata.Id)(utils.String(fmt.Sprintf("%s/v1.0/%s/directoryObjects/%s",
			client.BaseClient.Endpoint, client.BaseClient.TenantId, managerId)))

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
