// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package administrativeunits

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/pointer"
	"github.com/hashicorp/go-azure-helpers/lang/response"
	administrativeunitBeta "github.com/hashicorp/go-azure-sdk/microsoft-graph/administrativeunits/beta/administrativeunit"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/directory/stable/administrativeunit"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/directory/stable/administrativeunitmember"
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
	"github.com/hashicorp/go-uuid"
	"github.com/hashicorp/terraform-provider-azuread/internal/clients"
	"github.com/hashicorp/terraform-provider-azuread/internal/helpers/consistency"
	"github.com/hashicorp/terraform-provider-azuread/internal/helpers/tf"
	"github.com/hashicorp/terraform-provider-azuread/internal/helpers/tf/pluginsdk"
	"github.com/hashicorp/terraform-provider-azuread/internal/helpers/tf/validation"
)

const administrativeUnitResourceName = "azuread_administrative_unit"

func administrativeUnitResource() *pluginsdk.Resource {
	return &pluginsdk.Resource{
		CreateContext: administrativeUnitResourceCreate,
		ReadContext:   administrativeUnitResourceRead,
		UpdateContext: administrativeUnitResourceUpdate,
		DeleteContext: administrativeUnitResourceDelete,

		CustomizeDiff: administrativeUnitResourceCustomizeDiff,

		Timeouts: &pluginsdk.ResourceTimeout{
			Create: pluginsdk.DefaultTimeout(5 * time.Minute),
			Read:   pluginsdk.DefaultTimeout(5 * time.Minute),
			Update: pluginsdk.DefaultTimeout(5 * time.Minute),
			Delete: pluginsdk.DefaultTimeout(5 * time.Minute),
		},

		Importer: pluginsdk.ImporterValidatingResourceId(func(id string) error {
			if _, err := uuid.ParseUUID(id); err != nil {
				return fmt.Errorf("specified ID (%q) is not valid: %s", id, err)
			}
			return nil
		}),

		Schema: map[string]*pluginsdk.Schema{
			"display_name": {
				Description:      "The display name for the administrative unit",
				Type:             pluginsdk.TypeString,
				Required:         true,
				ValidateDiagFunc: validation.ValidateDiag(validation.StringIsNotEmpty),
			},

			"description": {
				Description: "The description for the administrative unit",
				Type:        pluginsdk.TypeString,
				Optional:    true,
			},

			"members": {
				Description: "A set of object IDs of members who should be present in this administrative unit. Supported object types are Users or Groups",
				Type:        pluginsdk.TypeSet,
				Optional:    true,
				Computed:    true,
				Set:         pluginsdk.HashString,
				Elem: &pluginsdk.Schema{
					Type:             pluginsdk.TypeString,
					ValidateDiagFunc: validation.ValidateDiag(validation.IsUUID),
				},
			},

			"prevent_duplicate_names": {
				Description: "If `true`, will return an error if an existing administrative unit is found with the same name",
				Type:        pluginsdk.TypeBool,
				Optional:    true,
				Default:     false,
			},

			"hidden_membership_enabled": {
				Description: "Whether the administrative unit and its members are hidden or publicly viewable in the directory",
				Type:        pluginsdk.TypeBool,
				Optional:    true,
			},

			"object_id": {
				Description: "The object ID of the administrative unit",
				Type:        pluginsdk.TypeString,
				Computed:    true,
			},
		},
	}
}

func administrativeUnitResourceCustomizeDiff(ctx context.Context, diff *pluginsdk.ResourceDiff, meta interface{}) error {
	client := meta.(*clients.Client).AdministrativeUnits.AdministrativeUnitClient
	ctx, cancel := context.WithTimeout(ctx, 5*time.Minute)
	defer cancel()

	// Check for duplicate names
	oldDisplayName, newDisplayName := diff.GetChange("display_name")
	if diff.Get("prevent_duplicate_names").(bool) && pluginsdk.ValueIsNotEmptyOrUnknown(newDisplayName) &&
		(oldDisplayName.(string) == "" || oldDisplayName.(string) != newDisplayName.(string)) {
		result, err := administrativeUnitFindByName(ctx, client, newDisplayName.(string))
		if err != nil {
			return fmt.Errorf("could not check for existing administrative unit(s): %+v", err)
		}
		if result != nil && len(*result) > 0 {
			for _, existingAu := range *result {
				if existingAu.Id == nil {
					return fmt.Errorf("API error: administrative unit returned with nil object ID during duplicate name check")
				}
				if diff.Id() == "" || diff.Id() == *existingAu.Id {
					return tf.ImportAsDuplicateError("azuread_administrative_unit", *existingAu.Id, newDisplayName.(string))
				}
			}
		}
	}

	return nil
}

func administrativeUnitResourceCreate(ctx context.Context, d *pluginsdk.ResourceData, meta interface{}) pluginsdk.Diagnostics {
	client := meta.(*clients.Client).AdministrativeUnits.AdministrativeUnitClient
	memberClient := meta.(*clients.Client).AdministrativeUnits.AdministrativeUnitMemberClient

	displayName := d.Get("display_name").(string)

	// Perform this check at apply time to catch any duplicate names created during the same apply
	if d.Get("prevent_duplicate_names").(bool) {
		result, err := administrativeUnitFindByName(ctx, client, displayName)
		if err != nil {
			return tf.ErrorDiagPathF(err, "name", "Could not check for existing administrative unit(s)")
		}
		if result != nil && len(*result) > 0 {
			existingAu := (*result)[0]
			if existingAu.Id == nil {
				return tf.ErrorDiagF(errors.New("API returned administrative unit with nil object ID during duplicate name check"), "Bad API response")
			}
			return tf.ImportAsDuplicateDiag("azuread_administrative_unit", *existingAu.Id, displayName)
		}
	}

	properties := stable.AdministrativeUnit{
		DisplayName: nullable.Value(displayName),
		Visibility:  nullable.Value(administrativeUnitVisibilityPublic),
	}

	if v := d.Get("description").(string); v != "" {
		properties.Description = nullable.Value(v)
	}

	if d.Get("hidden_membership_enabled").(bool) {
		properties.Visibility = nullable.Value(administrativeUnitVisibilityHiddenMembership)
	}

	resp, err := client.CreateAdministrativeUnit(ctx, properties, administrativeunit.DefaultCreateAdministrativeUnitOperationOptions())
	if err != nil {
		return tf.ErrorDiagF(err, "Creating administrative unit %q", displayName)
	}

	administrativeUnit := resp.Model
	if administrativeUnit == nil {
		return tf.ErrorDiagF(errors.New("API returned nil administrative unit"), "Bad API Response")
	}
	if administrativeUnit.Id == nil {
		return tf.ErrorDiagF(errors.New("API returned administrative unit with nil object ID"), "Bad API Response")
	}

	d.SetId(*administrativeUnit.Id)

	// Set a temporary display name as we'll attempt to patch the AU with the correct name after creating it
	uid, err := uuid.GenerateUUID()
	if err != nil {
		return tf.ErrorDiagF(err, "Failed to generate a UUID")
	}
	tempDisplayName := fmt.Sprintf("TERRAFORM_UPDATE_%s", uid)

	// Attempt to patch the newly created administrative unit with a temporary name, which will tell us whether it
	// exists yet. After, reset the name back to the correct name.
	// The SDK handles retries for us here in the event of 404, 429 or 5xx, then returns after giving up
	id := stable.NewDirectoryAdministrativeUnitID(*administrativeUnit.Id)
	updateResp, err := client.UpdateAdministrativeUnit(ctx, id, stable.AdministrativeUnit{
		DisplayName: nullable.Value(tempDisplayName),
	}, administrativeunit.UpdateAdministrativeUnitOperationOptions{
		RetryFunc: func(resp *http.Response, o *odata.OData) (bool, error) {
			return response.WasNotFound(resp), nil
		},
	})
	if err != nil {
		if response.WasNotFound(updateResp.HttpResponse) {
			return tf.ErrorDiagF(err, "Timed out whilst waiting for new %s to be replicated in Azure AD", id)
		}
		return tf.ErrorDiagF(err, "Failed to patch %s after creating", id)
	}

	// Set correct original display name
	updateResp, err = client.UpdateAdministrativeUnit(ctx, id, stable.AdministrativeUnit{
		DisplayName: nullable.Value(displayName),
	}, administrativeunit.UpdateAdministrativeUnitOperationOptions{
		RetryFunc: func(resp *http.Response, o *odata.OData) (bool, error) {
			return response.WasNotFound(resp), nil
		},
	})
	if err != nil {
		if response.WasNotFound(updateResp.HttpResponse) {
			return tf.ErrorDiagF(err, "Timed out whilst waiting for new %s to be replicated in Azure AD", id)
		}
		return tf.ErrorDiagF(err, "Failed to patch %s after creating", id)
	}

	// Add members after the administrative unit is created
	if v, ok := d.GetOk("members"); ok {
		for _, memberIdRaw := range v.(*pluginsdk.Set).List() {
			memberId := stable.NewDirectoryObjectID(memberIdRaw.(string))

			addMemberProperties := stable.ReferenceCreate{
				ODataId: pointer.To(client.Client.BaseUri + memberId.ID()),
			}

			if _, err = memberClient.AddAdministrativeUnitMemberRef(ctx, id, addMemberProperties, administrativeunitmember.DefaultAddAdministrativeUnitMemberRefOperationOptions()); err != nil {
				return tf.ErrorDiagF(err, "Could not add member %q to %s", memberId, id)
			}
		}
	}

	return administrativeUnitResourceRead(ctx, d, meta)
}

func administrativeUnitResourceUpdate(ctx context.Context, d *pluginsdk.ResourceData, meta interface{}) pluginsdk.Diagnostics {
	client := meta.(*clients.Client).AdministrativeUnits.AdministrativeUnitClient
	memberClient := meta.(*clients.Client).AdministrativeUnits.AdministrativeUnitMemberClient

	administrativeUnitId := d.Id()
	displayName := d.Get("display_name").(string)

	tf.LockByName(administrativeUnitResourceName, administrativeUnitId)
	defer tf.UnlockByName(administrativeUnitResourceName, administrativeUnitId)

	// Perform this check at apply time to catch any duplicate names created during the same apply
	if d.Get("prevent_duplicate_names").(bool) {
		result, err := administrativeUnitFindByName(ctx, client, displayName)
		if err != nil {
			return tf.ErrorDiagPathF(err, "display_name", "Could not check for existing administrative unit(s)")
		}
		if result != nil && len(*result) > 0 {
			for _, existingAU := range *result {
				if existingAU.Id == nil {
					return tf.ErrorDiagF(errors.New("API returned administrative unit with nil object ID during duplicate name check"), "Bad API response")
				}

				if *existingAU.Id != administrativeUnitId {
					return tf.ImportAsDuplicateDiag("azuread_administrative_unit", *existingAU.Id, displayName)
				}
			}
		}
	}

	id := stable.NewDirectoryAdministrativeUnitID(administrativeUnitId)
	administrativeUnit := stable.AdministrativeUnit{
		Description: nullable.Value(d.Get("description").(string)),
		DisplayName: nullable.Value(displayName),
		Visibility:  nullable.Value(administrativeUnitVisibilityPublic),
	}

	if d.Get("hidden_membership_enabled").(bool) {
		administrativeUnit.Visibility = nullable.Value(administrativeUnitVisibilityHiddenMembership)
	}

	if _, err := client.UpdateAdministrativeUnit(ctx, id, administrativeUnit, administrativeunit.DefaultUpdateAdministrativeUnitOperationOptions()); err != nil {
		return tf.ErrorDiagF(err, "Updating %s", id)
	}

	if d.HasChange("members") {
		membersResp, err := memberClient.ListAdministrativeUnitMembers(ctx, id, administrativeunitmember.DefaultListAdministrativeUnitMembersOperationOptions())
		if err != nil {
			return tf.ErrorDiagF(err, "Could not retrieve members for %s", id)
		}

		existingMembers := make([]string, 0)
		for _, member := range pointer.From(membersResp.Model) {
			existingMembers = append(existingMembers, pointer.From(member.DirectoryObject().Id))
		}
		desiredMembers := *tf.ExpandStringSlicePtr(d.Get("members").(*pluginsdk.Set).List())
		membersForRemoval := tf.Difference(existingMembers, desiredMembers)
		membersToAdd := tf.Difference(desiredMembers, existingMembers)

		for _, memberForRemoval := range membersForRemoval {
			if _, err = memberClient.RemoveAdministrativeUnitMemberRef(ctx, stable.NewDirectoryAdministrativeUnitIdMemberID(administrativeUnitId, memberForRemoval), administrativeunitmember.DefaultRemoveAdministrativeUnitMemberRefOperationOptions()); err != nil {
				return tf.ErrorDiagF(err, "Could not remove members from %s", id)
			}
		}

		for _, v := range membersToAdd {
			memberId := stable.NewDirectoryObjectID(v)

			addMemberProperties := stable.ReferenceCreate{
				ODataId: pointer.To(client.Client.BaseUri + memberId.ID()),
			}

			if _, err = memberClient.AddAdministrativeUnitMemberRef(ctx, id, addMemberProperties, administrativeunitmember.DefaultAddAdministrativeUnitMemberRefOperationOptions()); err != nil {
				return tf.ErrorDiagF(err, "Could not add member %q to %s", memberId, id)
			}
		}
	}

	return administrativeUnitResourceRead(ctx, d, meta)
}

func administrativeUnitResourceRead(ctx context.Context, d *pluginsdk.ResourceData, meta interface{}) pluginsdk.Diagnostics {
	client := meta.(*clients.Client).AdministrativeUnits.AdministrativeUnitClient
	memberClient := meta.(*clients.Client).AdministrativeUnits.AdministrativeUnitMemberClient

	id := stable.NewDirectoryAdministrativeUnitID(d.Id())
	resp, err := client.GetAdministrativeUnit(ctx, id, administrativeunit.DefaultGetAdministrativeUnitOperationOptions())
	if err != nil {
		if response.WasNotFound(resp.HttpResponse) {
			log.Printf("[DEBUG] %s was not found - removing from state", id)
			d.SetId("")
			return nil
		}
		return tf.ErrorDiagF(err, "Retrieving %s", id)
	}

	administrativeUnit := resp.Model
	tf.Set(d, "description", administrativeUnit.Description.GetOrZero())
	tf.Set(d, "display_name", administrativeUnit.DisplayName.GetOrZero())
	tf.Set(d, "object_id", id.AdministrativeUnitId)

	hiddenMembershipEnabled := strings.EqualFold(administrativeUnit.Visibility.GetOrZero(), administrativeUnitVisibilityHiddenMembership)
	tf.Set(d, "hidden_membership_enabled", hiddenMembershipEnabled)

	membersResp, err := memberClient.ListAdministrativeUnitMembers(ctx, id, administrativeunitmember.DefaultListAdministrativeUnitMembersOperationOptions())
	if err != nil {
		return tf.ErrorDiagF(err, "Could not retrieve members for %s", id)
	}

	members := make([]string, 0)
	for _, member := range pointer.From(membersResp.Model) {
		members = append(members, pointer.From(member.DirectoryObject().Id))
	}
	tf.Set(d, "members", members)

	preventDuplicates := false
	if v := d.Get("prevent_duplicate_names").(bool); v {
		preventDuplicates = v
	}
	tf.Set(d, "prevent_duplicate_names", preventDuplicates)

	return nil
}

func administrativeUnitResourceDelete(ctx context.Context, d *pluginsdk.ResourceData, meta interface{}) pluginsdk.Diagnostics {
	client := meta.(*clients.Client).AdministrativeUnits.AdministrativeUnitClient
	clientBeta := meta.(*clients.Client).AdministrativeUnits.AdministrativeUnitClientBeta
	id := beta.NewAdministrativeUnitID(d.Id())

	if _, err := clientBeta.DeleteAdministrativeUnit(ctx, id, administrativeunitBeta.DefaultDeleteAdministrativeUnitOperationOptions()); err != nil {
		return tf.ErrorDiagF(err, "Deleting %s", id)
	}

	// Wait for administrative unit object to be deleted
	if err := consistency.WaitForDeletion(ctx, func(ctx context.Context) (*bool, error) {
		if resp, err := client.GetAdministrativeUnit(ctx, stable.NewDirectoryAdministrativeUnitID(id.AdministrativeUnitId), administrativeunit.DefaultGetAdministrativeUnitOperationOptions()); err != nil {
			if response.WasNotFound(resp.HttpResponse) {
				return pointer.To(false), nil
			}
			return nil, err
		}
		return pointer.To(true), nil
	}); err != nil {
		return tf.ErrorDiagF(err, "Waiting for deletion of %s", id)
	}

	return nil
}
