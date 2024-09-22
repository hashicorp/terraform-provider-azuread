// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package applications

import (
	"context"
	"fmt"
	"github.com/hashicorp/go-azure-helpers/lang/response"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
	"net/http"
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/pointer"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/applications/stable/owner"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/terraform-provider-azuread/internal/helpers/applications"
	"github.com/hashicorp/terraform-provider-azuread/internal/helpers/tf"
	"github.com/hashicorp/terraform-provider-azuread/internal/helpers/tf/pluginsdk"
	"github.com/hashicorp/terraform-provider-azuread/internal/helpers/tf/validation"
	"github.com/hashicorp/terraform-provider-azuread/internal/sdk"
	"github.com/hashicorp/terraform-provider-azuread/internal/services/applications/parse"
)

type ApplicationOwnerModel struct {
	ApplicationId string `tfschema:"application_id"`
	OwnerObjectId string `tfschema:"owner_object_id"`
}

var _ sdk.Resource = ApplicationOwnerResource{}

type ApplicationOwnerResource struct{}

func (r ApplicationOwnerResource) IDValidationFunc() pluginsdk.SchemaValidateFunc {
	return parse.ValidateOwnerID
}

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
			client := metadata.Client.Applications.ApplicationOwnerClient

			var model ApplicationOwnerModel
			if err := metadata.Decode(&model); err != nil {
				return fmt.Errorf("decoding: %+v", err)
			}

			applicationId, err := stable.ParseApplicationID(model.ApplicationId)
			if err != nil {
				return err
			}

			ownerId := stable.NewApplicationIdOwnerID(applicationId.ApplicationId, model.OwnerObjectId)

			// TODO: migrate ID to use stable.ApplicationIdOwnerID
			id := parse.NewOwnerID(applicationId.ApplicationId, model.OwnerObjectId)

			tf.LockByName(applicationResourceName, applicationId.ApplicationId)
			defer tf.UnlockByName(applicationResourceName, applicationId.ApplicationId)

			o, err := applications.GetOwner(ctx, client, ownerId)
			if err != nil {
				return fmt.Errorf("checking for presence of existing %s: %+v", ownerId, err)
			}
			if o != nil {
				return metadata.ResourceRequiresImport(r.ResourceType(), id)
			}

			properties := stable.ReferenceCreate{
				ODataId: pointer.To(client.Client.BaseUri + stable.NewDirectoryObjectID(ownerId.DirectoryObjectId).ID()),
			}

			options := owner.AddOwnerRefOperationOptions{
				RetryFunc: func(resp *http.Response, _ *odata.OData) (bool, error) {
					if response.WasNotFound(resp) {
						return true, nil
					}
					return false, nil
				},
			}

			if _, err = client.AddOwnerRef(ctx, *applicationId, properties, options); err != nil {
				return fmt.Errorf("adding %s: %+v", ownerId, err)
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
			client := metadata.Client.Applications.ApplicationOwnerClient

			id, err := parse.ParseOwnerID(metadata.ResourceData.Id())
			if err != nil {
				return err
			}

			applicationId := parse.NewApplicationID(id.ApplicationId)
			ownerId := stable.NewApplicationIdOwnerID(applicationId.ApplicationId, id.OwnerId)

			tf.LockByName(applicationResourceName, id.ApplicationId)
			defer tf.UnlockByName(applicationResourceName, id.ApplicationId)

			owner, err := applications.GetOwner(ctx, client, ownerId)
			if err != nil {
				return fmt.Errorf("retrieving %s: %+v", id, err)
			}
			if owner == nil {
				return metadata.MarkAsGone(id)
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
			client := metadata.Client.Applications.ApplicationOwnerClient

			id, err := parse.ParseOwnerID(metadata.ResourceData.Id())
			if err != nil {
				return err
			}

			ownerId := stable.NewApplicationIdOwnerID(id.ApplicationId, id.OwnerId)

			tf.LockByName(applicationResourceName, id.ApplicationId)
			defer tf.UnlockByName(applicationResourceName, id.ApplicationId)

			if _, err = client.RemoveOwnerRef(ctx, ownerId, owner.DefaultRemoveOwnerRefOperationOptions()); err != nil {
				return fmt.Errorf("removing %s: %+v", id, err)
			}

			return nil
		},
	}
}
