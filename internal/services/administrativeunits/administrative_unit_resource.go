// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package administrativeunits

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/pointer"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
	"github.com/hashicorp/go-uuid"
	"github.com/hashicorp/terraform-provider-azuread/internal/clients"
	"github.com/hashicorp/terraform-provider-azuread/internal/helpers"
	"github.com/hashicorp/terraform-provider-azuread/internal/tf"
	"github.com/hashicorp/terraform-provider-azuread/internal/tf/pluginsdk"
	"github.com/hashicorp/terraform-provider-azuread/internal/tf/validation"
	"github.com/manicminer/hamilton/msgraph"
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
	client := meta.(*clients.Client).AdministrativeUnits.AdministrativeUnitsClient

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
				if existingAu.ID == nil {
					return fmt.Errorf("API error: administrative unit returned with nil object ID during duplicate name check")
				}
				if diff.Id() == "" || diff.Id() == *existingAu.ID {
					return tf.ImportAsDuplicateError("azuread_administrative_unit", *existingAu.ID, newDisplayName.(string))
				}
			}
		}
	}

	return nil
}

func administrativeUnitResourceCreate(ctx context.Context, d *pluginsdk.ResourceData, meta interface{}) pluginsdk.Diagnostics {
	client := meta.(*clients.Client).AdministrativeUnits.AdministrativeUnitsClient
	directoryObjectsClient := meta.(*clients.Client).AdministrativeUnits.DirectoryObjectsClient
	tenantId := meta.(*clients.Client).TenantID

	displayName := d.Get("display_name").(string)

	// Perform this check at apply time to catch any duplicate names created during the same apply
	if d.Get("prevent_duplicate_names").(bool) {
		result, err := administrativeUnitFindByName(ctx, client, displayName)
		if err != nil {
			return tf.ErrorDiagPathF(err, "name", "Could not check for existing administrative unit(s)")
		}
		if result != nil && len(*result) > 0 {
			existingAu := (*result)[0]
			if existingAu.ID == nil {
				return tf.ErrorDiagF(errors.New("API returned administrative unit with nil object ID during duplicate name check"), "Bad API response")
			}
			return tf.ImportAsDuplicateDiag("azuread_administrative_unit", *existingAu.ID, displayName)
		}
	}

	// Set a temporary display name as we'll attempt to patch the AU with the correct name after creating it
	uuid, err := uuid.GenerateUUID()
	if err != nil {
		return tf.ErrorDiagF(err, "Failed to generate a UUID")
	}
	tempDisplayName := fmt.Sprintf("TERRAFORM_UPDATE_%s", uuid)

	properties := msgraph.AdministrativeUnit{
		Description: tf.NullableString(d.Get("description").(string)),
		DisplayName: pointer.To(tempDisplayName),
		Visibility:  pointer.To(msgraph.AdministrativeUnitVisibilityPublic),
	}

	if d.Get("hidden_membership_enabled").(bool) {
		properties.Visibility = pointer.To(msgraph.AdministrativeUnitVisibilityHiddenMembership)
	}

	administrativeUnit, _, err := client.Create(ctx, properties)
	if err != nil {
		return tf.ErrorDiagF(err, "Creating administrative unit %q", displayName)
	}

	if administrativeUnit.ID == nil {
		return tf.ErrorDiagF(errors.New("API returned administrative unit with nil object ID"), "Bad API Response")
	}

	d.SetId(*administrativeUnit.ID)

	// Attempt to patch the newly created administrative unit with the correct name, which will tell us whether it exists yet
	// The SDK handles retries for us here in the event of 404, 429 or 5xx, then returns after giving up
	status, err := client.Update(ctx, msgraph.AdministrativeUnit{
		ID:          administrativeUnit.ID,
		DisplayName: pointer.To(displayName),
	})
	if err != nil {
		if status == http.StatusNotFound {
			return tf.ErrorDiagF(err, "Timed out whilst waiting for new administrative unit to be replicated in Azure AD")
		}
		return tf.ErrorDiagF(err, "Failed to patch administrative unit after creating")
	}

	// Add members after the administrative unit is created
	members := make(msgraph.Members, 0)
	if v, ok := d.GetOk("members"); ok {
		for _, memberId := range v.(*pluginsdk.Set).List() {
			memberObject, _, err := directoryObjectsClient.Get(ctx, memberId.(string), odata.Query{})
			if err != nil {
				return tf.ErrorDiagF(err, "Could not retrieve member principal object %q", memberId)
			}
			if memberObject == nil {
				return tf.ErrorDiagF(errors.New("memberObject was nil"), "Could not retrieve member principal object %q", memberId)
			}
			memberObject.ODataId = (*odata.Id)(pointer.To(fmt.Sprintf("%s/v1.0/%s/directoryObjects/%s",
				client.BaseClient.Endpoint, tenantId, memberId)))

			members = append(members, *memberObject)
		}
	}
	if len(members) > 0 {
		if _, err := client.AddMembers(ctx, d.Id(), &members); err != nil {
			return tf.ErrorDiagF(err, "Could not add members to administrative unit with object ID: %q", d.Id())
		}
	}

	return administrativeUnitResourceRead(ctx, d, meta)
}

func administrativeUnitResourceUpdate(ctx context.Context, d *pluginsdk.ResourceData, meta interface{}) pluginsdk.Diagnostics {
	client := meta.(*clients.Client).AdministrativeUnits.AdministrativeUnitsClient
	directoryObjectsClient := meta.(*clients.Client).AdministrativeUnits.DirectoryObjectsClient
	tenantId := meta.(*clients.Client).TenantID

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
				if existingAU.ID == nil {
					return tf.ErrorDiagF(errors.New("API returned administrative unit with nil object ID during duplicate name check"), "Bad API response")
				}

				if *existingAU.ID != administrativeUnitId {
					return tf.ImportAsDuplicateDiag("azuread_administrative_unit", *existingAU.ID, displayName)
				}
			}
		}
	}

	administrativeUnit := msgraph.AdministrativeUnit{
		ID:          pointer.To(administrativeUnitId),
		Description: tf.NullableString(d.Get("description").(string)),
		DisplayName: pointer.To(displayName),
		Visibility:  pointer.To(msgraph.AdministrativeUnitVisibilityPublic),
	}

	if d.Get("hidden_membership_enabled").(bool) {
		administrativeUnit.Visibility = pointer.To(msgraph.AdministrativeUnitVisibilityHiddenMembership)
	}

	if _, err := client.Update(ctx, administrativeUnit); err != nil {
		return tf.ErrorDiagF(err, "Updating administrative unit with ID: %q", d.Id())
	}

	if d.HasChange("members") {
		members, _, err := client.ListMembers(ctx, *administrativeUnit.ID)
		if err != nil {
			return tf.ErrorDiagF(err, "Could not retrieve members for administrative unit with object ID: %q", d.Id())
		}

		existingMembers := *members
		desiredMembers := *tf.ExpandStringSlicePtr(d.Get("members").(*pluginsdk.Set).List())
		membersForRemoval := tf.Difference(existingMembers, desiredMembers)
		membersToAdd := tf.Difference(desiredMembers, existingMembers)

		if len(membersForRemoval) > 0 {
			if _, err = client.RemoveMembers(ctx, d.Id(), &membersForRemoval); err != nil {
				return tf.ErrorDiagF(err, "Could not remove members from administrative unit with object ID: %q", d.Id())
			}
		}

		if len(membersToAdd) > 0 {
			newMembers := make(msgraph.Members, 0)
			for _, memberId := range membersToAdd {
				memberObject, _, err := directoryObjectsClient.Get(ctx, memberId, odata.Query{})
				if err != nil {
					return tf.ErrorDiagF(err, "Could not retrieve principal object %q", memberId)
				}
				if memberObject == nil {
					return tf.ErrorDiagF(errors.New("returned memberObject was nil"), "Could not retrieve member principal object %q", memberId)
				}
				memberObject.ODataId = (*odata.Id)(pointer.To(fmt.Sprintf("%s/v1.0/%s/directoryObjects/%s",
					client.BaseClient.Endpoint, tenantId, memberId)))

				newMembers = append(newMembers, *memberObject)
			}

			if _, err := client.AddMembers(ctx, administrativeUnitId, &newMembers); err != nil {
				return tf.ErrorDiagF(err, "Could not add members to administrative unit with object ID: %q", d.Id())
			}
		}
	}

	return administrativeUnitResourceRead(ctx, d, meta)
}

func administrativeUnitResourceRead(ctx context.Context, d *pluginsdk.ResourceData, meta interface{}) pluginsdk.Diagnostics {
	client := meta.(*clients.Client).AdministrativeUnits.AdministrativeUnitsClient

	administrativeUnit, status, err := client.Get(ctx, d.Id(), odata.Query{})
	if err != nil {
		if status == http.StatusNotFound {
			log.Printf("[DEBUG] Administrative Unit with ID %q was not found - removing from state", d.Id())
			d.SetId("")
			return nil
		}
		return tf.ErrorDiagF(err, "Retrieving administrative unit with object ID: %q", d.Id())
	}

	tf.Set(d, "description", administrativeUnit.Description)
	tf.Set(d, "display_name", administrativeUnit.DisplayName)
	tf.Set(d, "object_id", administrativeUnit.ID)

	hiddenMembershipEnabled := administrativeUnit.Visibility != nil && *administrativeUnit.Visibility == msgraph.AdministrativeUnitVisibilityHiddenMembership
	tf.Set(d, "hidden_membership_enabled", hiddenMembershipEnabled)

	members, _, err := client.ListMembers(ctx, *administrativeUnit.ID)
	if err != nil {
		return tf.ErrorDiagPathF(err, "members", "Could not retrieve members for administrative unit with object ID %q", d.Id())
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
	client := meta.(*clients.Client).AdministrativeUnits.AdministrativeUnitsClient
	administrativeUnitId := d.Id()

	_, status, err := client.Get(ctx, administrativeUnitId, odata.Query{})
	if err != nil {
		if status == http.StatusNotFound {
			return tf.ErrorDiagPathF(fmt.Errorf("Administrative unit was not found"), "id", "Retrieving administrative unit with object ID %q", administrativeUnitId)
		}
		return tf.ErrorDiagPathF(err, "id", "Retrieving administrative unit with object ID: %q", administrativeUnitId)
	}

	if _, err := client.Delete(ctx, administrativeUnitId); err != nil {
		return tf.ErrorDiagF(err, "Deleting administrative unit with object ID: %q", administrativeUnitId)
	}

	// Wait for administrative unit object to be deleted
	if err := helpers.WaitForDeletion(ctx, func(ctx context.Context) (*bool, error) {
		defer func() { client.BaseClient.DisableRetries = false }()
		client.BaseClient.DisableRetries = true
		if _, status, err := client.Get(ctx, administrativeUnitId, odata.Query{}); err != nil {
			if status == http.StatusNotFound {
				return pointer.To(false), nil
			}
			return nil, err
		}
		return pointer.To(true), nil
	}); err != nil {
		return tf.ErrorDiagF(err, "Waiting for deletion of administrative unit with object ID %q", administrativeUnitId)
	}

	return nil
}
