// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package applications

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/hashicorp/go-azure-sdk/sdk/odata"
	"github.com/hashicorp/terraform-provider-azuread/internal/sdk"
	"github.com/hashicorp/terraform-provider-azuread/internal/services/applications/parse"
	"github.com/hashicorp/terraform-provider-azuread/internal/tf"
	"github.com/hashicorp/terraform-provider-azuread/internal/tf/pluginsdk"
	"github.com/hashicorp/terraform-provider-azuread/internal/tf/validation"
	"github.com/manicminer/hamilton/msgraph"
)

type ApplicationOwnerModel struct {
	ApplicationId string `tfschema:"application_id"`
	OwnerObjectId string `tfschema:"owner_object_id"`
}

type ApplicationOwnerResource struct{}

func (r ApplicationOwnerResource) IDValidationFunc() pluginsdk.SchemaValidateFunc {
	return parse.ValidateOwnerID
}

var _ sdk.Resource = ApplicationOwnerResource{}

func (r ApplicationOwnerResource) ResourceType() string {
	return "azuread_application_owner"
}

func (r ApplicationOwnerResource) ModelObject() interface{} {
	return &ApplicationOwnerModel{}
}

func (r ApplicationOwnerResource) Arguments() map[string]*pluginsdk.Schema {
	return map[string]*pluginsdk.Schema{
		"application_id": {
			Description:  "The resource ID of the application to which the owner should be added",
			Type:         pluginsdk.TypeString,
			Required:     true,
			ForceNew:     true,
			ValidateFunc: parse.ValidateApplicationID,
		},

		"owner_object_id": {
			Description:  "Object ID of the principal that will be granted ownership of the application",
			Type:         pluginsdk.TypeString,
			Required:     true,
			ForceNew:     true,
			ValidateFunc: validation.IsUUID,
		},
	}
}

func (r ApplicationOwnerResource) Attributes() map[string]*pluginsdk.Schema {
	return map[string]*pluginsdk.Schema{}
}

func (r ApplicationOwnerResource) Create() sdk.ResourceFunc {
	return sdk.ResourceFunc{
		Timeout: 10 * time.Minute,
		Func: func(ctx context.Context, metadata sdk.ResourceMetaData) error {
			client := metadata.Client.Applications.ApplicationsClient
			client.BaseClient.DisableRetries = true
			defer func() { client.BaseClient.DisableRetries = false }()

			var model ApplicationOwnerModel
			if err := metadata.Decode(&model); err != nil {
				return fmt.Errorf("decoding: %+v", err)
			}

			applicationId, err := parse.ParseApplicationID(model.ApplicationId)
			if err != nil {
				return err
			}

			id := parse.NewOwnerID(applicationId.ApplicationId, model.OwnerObjectId)

			tf.LockByName(applicationResourceName, id.ApplicationId)
			defer tf.UnlockByName(applicationResourceName, id.ApplicationId)

			_, status, err := client.GetOwner(ctx, id.ApplicationId, id.OwnerId)
			if err != nil && status != http.StatusNotFound {
				return fmt.Errorf("checking for presence of existing %s: %+v", id, err)
			}
			if status != http.StatusNotFound {
				return metadata.ResourceRequiresImport(r.ResourceType(), id)
			}

			// Construct an @odata.id for the $ref endpoint
			odataId := (odata.Id)(fmt.Sprintf("%s/v1.0/%s/directoryObjects/%s", client.BaseClient.Endpoint, metadata.Client.TenantID, id.OwnerId))

			properties := &msgraph.Application{
				DirectoryObject: msgraph.DirectoryObject{
					Id: &id.ApplicationId,
				},
				Owners: &msgraph.Owners{{
					Id:      &id.OwnerId,
					ODataId: &odataId,
				}},
			}

			if _, err = client.AddOwners(ctx, properties); err != nil {
				return fmt.Errorf("adding %s: %+v", id, err)
			}

			metadata.SetID(id)
			return nil
		},
	}
}

func (r ApplicationOwnerResource) Read() sdk.ResourceFunc {
	return sdk.ResourceFunc{
		Timeout: 5 * time.Minute,
		Func: func(ctx context.Context, metadata sdk.ResourceMetaData) error {
			client := metadata.Client.Applications.ApplicationsClient
			client.BaseClient.DisableRetries = true
			defer func() { client.BaseClient.DisableRetries = false }()

			id, err := parse.ParseOwnerID(metadata.ResourceData.Id())
			if err != nil {
				return err
			}

			applicationId := parse.NewApplicationID(id.ApplicationId)

			tf.LockByName(applicationResourceName, id.ApplicationId)
			defer tf.UnlockByName(applicationResourceName, id.ApplicationId)

			result, status, err := client.GetOwner(ctx, id.ApplicationId, id.OwnerId)
			if err != nil {
				if status == http.StatusNotFound {
					return metadata.MarkAsGone(id)
				}
				return fmt.Errorf("retrieving %s: %+v", id, err)
			}
			if result == nil {
				return fmt.Errorf("retrieving %s: result was nil", id)
			}

			state := ApplicationOwnerModel{
				ApplicationId: applicationId.ID(),
				OwnerObjectId: id.OwnerId,
			}

			return metadata.Encode(&state)
		},
	}
}

func (r ApplicationOwnerResource) Delete() sdk.ResourceFunc {
	return sdk.ResourceFunc{
		Timeout: 5 * time.Minute,
		Func: func(ctx context.Context, metadata sdk.ResourceMetaData) error {
			client := metadata.Client.Applications.ApplicationsClient
			client.BaseClient.DisableRetries = true
			defer func() { client.BaseClient.DisableRetries = false }()

			id, err := parse.ParseOwnerID(metadata.ResourceData.Id())
			if err != nil {
				return err
			}

			tf.LockByName(applicationResourceName, id.ApplicationId)
			defer tf.UnlockByName(applicationResourceName, id.ApplicationId)

			_, err = client.RemoveOwners(ctx, id.ApplicationId, &[]string{id.OwnerId})
			if err != nil {
				return fmt.Errorf("removing %s: %+v", id, err)
			}

			return nil
		},
	}
}
