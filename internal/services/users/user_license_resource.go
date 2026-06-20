// Copyright IBM Corp. 2014, 2025
// SPDX-License-Identifier: MPL-2.0

package users

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/pointer"
	"github.com/hashicorp/go-azure-helpers/lang/response"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/user"
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
	"github.com/hashicorp/terraform-provider-azuread/internal/helpers/consistency"
	"github.com/hashicorp/terraform-provider-azuread/internal/helpers/tf"
	"github.com/hashicorp/terraform-provider-azuread/internal/helpers/tf/pluginsdk"
	"github.com/hashicorp/terraform-provider-azuread/internal/helpers/tf/validation"
	"github.com/hashicorp/terraform-provider-azuread/internal/sdk"
	"github.com/hashicorp/terraform-provider-azuread/internal/services/users/parse"
)

type UserLicenseResourceModel struct {
	UserId        string   `tfschema:"user_id"`
	SkuId         string   `tfschema:"sku_id"`
	DisabledPlans []string `tfschema:"disabled_plans"`
}

var _ sdk.Resource = UserLicenseResource{}

type UserLicenseResource struct{}

func (r UserLicenseResource) IDValidationFunc() pluginsdk.SchemaValidateFunc {
	return parse.ValidateUserLicenseID
}

func (r UserLicenseResource) ResourceType() string {
	return "azuread_user_license"
}

func (r UserLicenseResource) ModelObject() interface{} {
	return &UserLicenseResourceModel{}
}

func (r UserLicenseResource) Arguments() map[string]*pluginsdk.Schema {
	return map[string]*pluginsdk.Schema{
		"user_id": {
			Description:  "The object ID of the user to which the license should be assigned",
			Type:         pluginsdk.TypeString,
			Required:     true,
			ForceNew:     true,
			ValidateFunc: validation.IsUUID,
		},

		"sku_id": {
			Description:  "The unique identifier (GUID) for the SKU (license) to assign to the user",
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

func (r UserLicenseResource) Attributes() map[string]*pluginsdk.Schema {
	return map[string]*pluginsdk.Schema{}
}

func (r UserLicenseResource) Create() sdk.ResourceFunc {
	return sdk.ResourceFunc{
		Timeout: 10 * time.Minute,
		Func: func(ctx context.Context, metadata sdk.ResourceMetaData) error {
			client := metadata.Client.Users.UserClient

			var model UserLicenseResourceModel
			if err := metadata.Decode(&model); err != nil {
				return fmt.Errorf("decoding: %+v", err)
			}

			userId := stable.NewUserID(model.UserId)
			id := parse.NewUserLicenseID(model.UserId, model.SkuId)

			tf.LockByName(userResourceName, model.UserId)
			defer tf.UnlockByName(userResourceName, model.UserId)

			resp, err := client.GetUser(ctx, userId, user.GetUserOperationOptions{
				Select: &[]string{"id", "usageLocation", "assignedLicenses", "licenseAssignmentStates"},
			})
			if err != nil {
				if response.WasNotFound(resp.HttpResponse) {
					return fmt.Errorf("assigning %s: user was not found", id)
				}
				return fmt.Errorf("retrieving %s: %+v", userId, err)
			}

			u := resp.Model
			if u == nil {
				return fmt.Errorf("retrieving %s: model was nil", userId)
			}

			// Microsoft Graph rejects assignLicense for users without a usage location set, due to legal
			// requirements to check the availability of services in a given country. Surface a clear error
			// rather than relying on the opaque error returned by the API.
			if u.UsageLocation.GetOrZero() == "" {
				return fmt.Errorf("assigning %s: the user has no `usage_location` set, which is required before licenses can be assigned. Set `usage_location` on the `azuread_user` resource or user object first", id)
			}

			if existing := findDirectLicenseAssignment(u, model.SkuId); existing != nil {
				return metadata.ResourceRequiresImport(r.ResourceType(), id)
			}

			// Microsoft Graph requires disabledPlans to be a non-null collection, so default a nil slice
			// (when disabled_plans is unset) to an empty slice rather than sending a null value.
			disabledPlans := model.DisabledPlans
			if disabledPlans == nil {
				disabledPlans = []string{}
			}

			properties := user.AssignLicenseRequest{
				AddLicenses: &[]stable.AssignedLicense{
					{
						SkuId:         nullable.Value(model.SkuId),
						DisabledPlans: &disabledPlans,
					},
				},
				RemoveLicenses: &[]string{},
			}

			options := user.AssignLicenseOperationOptions{
				RetryFunc: func(resp *http.Response, _ *odata.OData) (bool, error) {
					return response.WasNotFound(resp), nil
				},
			}

			if _, err = client.AssignLicense(ctx, userId, properties, options); err != nil {
				return fmt.Errorf("assigning %s: %+v", id, err)
			}

			if err = consistency.WaitForUpdate(ctx, func(ctx context.Context) (*bool, error) {
				resp, err := client.GetUser(ctx, userId, user.GetUserOperationOptions{
					Select: &[]string{"id", "assignedLicenses", "licenseAssignmentStates"},
				})
				if err != nil {
					return nil, err
				}
				return pointer.To(findDirectLicenseAssignment(resp.Model, model.SkuId) != nil), nil
			}); err != nil {
				return fmt.Errorf("waiting for assignment of %s: %+v", id, err)
			}

			metadata.SetID(id)
			return nil
		},
	}
}

func (r UserLicenseResource) Read() sdk.ResourceFunc {
	return sdk.ResourceFunc{
		Timeout: 5 * time.Minute,
		Func: func(ctx context.Context, metadata sdk.ResourceMetaData) error {
			client := metadata.Client.Users.UserClient

			id, err := parse.UserLicenseID(metadata.ResourceData.Id())
			if err != nil {
				return err
			}

			userId := stable.NewUserID(id.UserId)

			resp, err := client.GetUser(ctx, userId, user.GetUserOperationOptions{
				Select: &[]string{"id", "assignedLicenses", "licenseAssignmentStates"},
			})
			if err != nil {
				if response.WasNotFound(resp.HttpResponse) {
					return metadata.MarkAsGone(id)
				}
				return fmt.Errorf("retrieving %s: %+v", userId, err)
			}

			// Only consider direct user assignments as managed by this resource - licenses inherited via
			// group-based licensing carry a non-null `assignedByGroup` and must not be adopted here.
			assignment := findDirectLicenseAssignment(resp.Model, id.SkuId)
			if assignment == nil {
				return metadata.MarkAsGone(id)
			}

			state := UserLicenseResourceModel{
				UserId:        id.UserId,
				SkuId:         id.SkuId,
				DisabledPlans: pointer.From(assignment.DisabledPlans),
			}

			return metadata.Encode(&state)
		},
	}
}

func (r UserLicenseResource) Delete() sdk.ResourceFunc {
	return sdk.ResourceFunc{
		Timeout: 5 * time.Minute,
		Func: func(ctx context.Context, metadata sdk.ResourceMetaData) error {
			client := metadata.Client.Users.UserClient

			id, err := parse.UserLicenseID(metadata.ResourceData.Id())
			if err != nil {
				return err
			}

			userId := stable.NewUserID(id.UserId)

			tf.LockByName(userResourceName, id.UserId)
			defer tf.UnlockByName(userResourceName, id.UserId)

			properties := user.AssignLicenseRequest{
				AddLicenses:    &[]stable.AssignedLicense{},
				RemoveLicenses: &[]string{id.SkuId},
			}

			if _, err = client.AssignLicense(ctx, userId, properties, user.DefaultAssignLicenseOperationOptions()); err != nil {
				return fmt.Errorf("removing %s: %+v", id, err)
			}

			// Wait for the license to be removed from the user, as the Graph API is eventually consistent
			if err = consistency.WaitForDeletion(ctx, func(ctx context.Context) (*bool, error) {
				resp, err := client.GetUser(ctx, userId, user.GetUserOperationOptions{
					Select: &[]string{"id", "assignedLicenses", "licenseAssignmentStates"},
				})
				if err != nil {
					if response.WasNotFound(resp.HttpResponse) {
						return pointer.To(false), nil
					}
					return nil, err
				}
				return pointer.To(findDirectLicenseAssignment(resp.Model, id.SkuId) != nil), nil
			}); err != nil {
				return fmt.Errorf("waiting for removal of %s: %+v", id, err)
			}

			return nil
		},
	}
}

// findDirectLicenseAssignment returns the directly-assigned license matching the given SKU ID, or nil if
// the user has no such direct assignment. Licenses inherited via group-based licensing (which carry a
// non-null `assignedByGroup` in their assignment state) are not considered, so this resource never adopts
// a license it did not assign directly.
func findDirectLicenseAssignment(u *stable.User, skuId string) *stable.AssignedLicense {
	if u == nil {
		return nil
	}

	// Determine whether a direct (non group-based) assignment exists for this SKU
	directlyAssigned := false
	if u.LicenseAssignmentStates != nil {
		for _, state := range *u.LicenseAssignmentStates {
			if state.SkuId.GetOrZero() == skuId && state.AssignedByGroup.GetOrZero() == "" {
				directlyAssigned = true
				break
			}
		}
	}

	if !directlyAssigned {
		return nil
	}

	if u.AssignedLicenses != nil {
		for _, license := range *u.AssignedLicenses {
			if license.SkuId.GetOrZero() == skuId {
				assignedLicense := license
				return &assignedLicense
			}
		}
	}

	return nil
}
