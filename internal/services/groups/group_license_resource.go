// Copyright IBM Corp. 2014, 2025
// SPDX-License-Identifier: MPL-2.0

package groups

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/pointer"
	"github.com/hashicorp/go-azure-helpers/lang/response"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/group"
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
	"github.com/hashicorp/terraform-provider-azuread/internal/helpers/consistency"
	"github.com/hashicorp/terraform-provider-azuread/internal/helpers/tf"
	"github.com/hashicorp/terraform-provider-azuread/internal/helpers/tf/pluginsdk"
	"github.com/hashicorp/terraform-provider-azuread/internal/helpers/tf/validation"
	"github.com/hashicorp/terraform-provider-azuread/internal/sdk"
	"github.com/hashicorp/terraform-provider-azuread/internal/services/groups/parse"
)

type GroupLicenseResourceModel struct {
	GroupId       string   `tfschema:"group_id"`
	SkuId         string   `tfschema:"sku_id"`
	DisabledPlans []string `tfschema:"disabled_plans"`
}

var _ sdk.Resource = GroupLicenseResource{}

type GroupLicenseResource struct{}

func (r GroupLicenseResource) IDValidationFunc() pluginsdk.SchemaValidateFunc {
	return parse.ValidateGroupLicenseID
}

func (r GroupLicenseResource) ResourceType() string {
	return "azuread_group_license"
}

func (r GroupLicenseResource) ModelObject() interface{} {
	return &GroupLicenseResourceModel{}
}

func (r GroupLicenseResource) Arguments() map[string]*pluginsdk.Schema {
	return map[string]*pluginsdk.Schema{
		"group_id": {
			Description:  "The object ID of the group to which the license should be assigned",
			Type:         pluginsdk.TypeString,
			Required:     true,
			ForceNew:     true,
			ValidateFunc: validation.IsUUID,
		},

		"sku_id": {
			Description:  "The unique identifier (GUID) for the SKU (license) to assign to the group",
			Type:         pluginsdk.TypeString,
			Required:     true,
			ForceNew:     true,
			ValidateFunc: validation.IsUUID,
		},

		"disabled_plans": {
			Description: "A set of unique identifiers (GUIDs) for the service plans to disable for this license",
			Type:        pluginsdk.TypeSet,
			Optional:    true,
			ForceNew:    true,
			Elem: &pluginsdk.Schema{
				Type:         pluginsdk.TypeString,
				ValidateFunc: validation.IsUUID,
			},
		},
	}
}

func (r GroupLicenseResource) Attributes() map[string]*pluginsdk.Schema {
	return map[string]*pluginsdk.Schema{}
}

func (r GroupLicenseResource) Create() sdk.ResourceFunc {
	return sdk.ResourceFunc{
		Timeout: 10 * time.Minute,
		Func: func(ctx context.Context, metadata sdk.ResourceMetaData) error {
			client := metadata.Client.Groups.GroupClientBeta

			var model GroupLicenseResourceModel
			if err := metadata.Decode(&model); err != nil {
				return fmt.Errorf("decoding: %+v", err)
			}

			groupId := beta.NewGroupID(model.GroupId)
			id := parse.NewGroupLicenseID(model.GroupId, model.SkuId)

			tf.LockByName(groupResourceName, model.GroupId)
			defer tf.UnlockByName(groupResourceName, model.GroupId)

			resp, err := client.GetGroup(ctx, groupId, group.GetGroupOperationOptions{
				Select: &[]string{"id", "assignedLicenses"},
			})
			if err != nil {
				if response.WasNotFound(resp.HttpResponse) {
					return fmt.Errorf("assigning %s: group was not found", id)
				}
				return fmt.Errorf("retrieving %s: %+v", groupId, err)
			}

			g := resp.Model
			if g == nil {
				return fmt.Errorf("retrieving %s: model was nil", groupId)
			}

			if existing := findGroupLicense(g, model.SkuId); existing != nil {
				return metadata.ResourceRequiresImport(r.ResourceType(), id)
			}

			// Microsoft Graph requires disabledPlans to be a non-null collection, so default a nil slice
			// (when disabled_plans is unset) to an empty slice rather than sending a null value.
			disabledPlans := model.DisabledPlans
			if disabledPlans == nil {
				disabledPlans = []string{}
			}

			properties := group.AssignLicenseRequest{
				AddLicenses: &[]beta.AssignedLicense{
					{
						SkuId:         nullable.Value(model.SkuId),
						DisabledPlans: &disabledPlans,
					},
				},
				RemoveLicenses: &[]string{},
			}

			options := group.AssignLicenseOperationOptions{
				RetryFunc: func(resp *http.Response, _ *odata.OData) (bool, error) {
					return response.WasNotFound(resp), nil
				},
			}

			if _, err = client.AssignLicense(ctx, groupId, properties, options); err != nil {
				return fmt.Errorf("assigning %s: %+v", id, err)
			}

			// The assignLicense action for a group is asynchronous, so wait for the license to appear on the
			// group's own assignedLicenses. This only reflects the assignment to the group object itself.
			// Propagation to group members happens separately and is not managed by this resource.
			if err = consistency.WaitForUpdate(ctx, func(ctx context.Context) (*bool, error) {
				resp, err := client.GetGroup(ctx, groupId, group.GetGroupOperationOptions{
					Select: &[]string{"id", "assignedLicenses"},
				})
				if err != nil {
					return nil, err
				}
				return pointer.To(findGroupLicense(resp.Model, model.SkuId) != nil), nil
			}); err != nil {
				return fmt.Errorf("waiting for assignment of %s: %+v", id, err)
			}

			metadata.SetID(id)
			return nil
		},
	}
}

func (r GroupLicenseResource) Read() sdk.ResourceFunc {
	return sdk.ResourceFunc{
		Timeout: 5 * time.Minute,
		Func: func(ctx context.Context, metadata sdk.ResourceMetaData) error {
			client := metadata.Client.Groups.GroupClientBeta

			id, err := parse.GroupLicenseID(metadata.ResourceData.Id())
			if err != nil {
				return err
			}

			groupId := beta.NewGroupID(id.GroupId)

			resp, err := client.GetGroup(ctx, groupId, group.GetGroupOperationOptions{
				Select: &[]string{"id", "assignedLicenses"},
			})
			if err != nil {
				if response.WasNotFound(resp.HttpResponse) {
					return metadata.MarkAsGone(id)
				}
				return fmt.Errorf("retrieving %s: %+v", groupId, err)
			}

			assignment := findGroupLicense(resp.Model, id.SkuId)
			if assignment == nil {
				return metadata.MarkAsGone(id)
			}

			state := GroupLicenseResourceModel{
				GroupId:       id.GroupId,
				SkuId:         id.SkuId,
				DisabledPlans: pointer.From(assignment.DisabledPlans),
			}

			return metadata.Encode(&state)
		},
	}
}

func (r GroupLicenseResource) Delete() sdk.ResourceFunc {
	return sdk.ResourceFunc{
		Timeout: 5 * time.Minute,
		Func: func(ctx context.Context, metadata sdk.ResourceMetaData) error {
			client := metadata.Client.Groups.GroupClientBeta

			id, err := parse.GroupLicenseID(metadata.ResourceData.Id())
			if err != nil {
				return err
			}

			groupId := beta.NewGroupID(id.GroupId)

			tf.LockByName(groupResourceName, id.GroupId)
			defer tf.UnlockByName(groupResourceName, id.GroupId)

			properties := group.AssignLicenseRequest{
				AddLicenses:    &[]beta.AssignedLicense{},
				RemoveLicenses: &[]string{id.SkuId},
			}

			if _, err = client.AssignLicense(ctx, groupId, properties, group.DefaultAssignLicenseOperationOptions()); err != nil {
				return fmt.Errorf("removing %s: %+v", id, err)
			}

			if err = consistency.WaitForDeletion(ctx, func(ctx context.Context) (*bool, error) {
				resp, err := client.GetGroup(ctx, groupId, group.GetGroupOperationOptions{
					Select: &[]string{"id", "assignedLicenses"},
				})
				if err != nil {
					if response.WasNotFound(resp.HttpResponse) {
						return pointer.To(false), nil
					}
					return nil, err
				}
				return pointer.To(findGroupLicense(resp.Model, id.SkuId) != nil), nil
			}); err != nil {
				return fmt.Errorf("waiting for removal of %s: %+v", id, err)
			}

			return nil
		},
	}
}

func findGroupLicense(g *beta.Group, skuId string) *beta.AssignedLicense {
	if g == nil || g.AssignedLicenses == nil {
		return nil
	}

	for _, license := range *g.AssignedLicenses {
		if license.SkuId.GetOrZero() == skuId {
			assignedLicense := license
			return &assignedLicense
		}
	}

	return nil
}
