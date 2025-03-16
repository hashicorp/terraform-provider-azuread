// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package groups

import (
	"context"
	"strings"
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/response"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/group"
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
	"github.com/hashicorp/terraform-provider-azuread/internal/clients"
	"github.com/hashicorp/terraform-provider-azuread/internal/helpers/tf"
	"github.com/hashicorp/terraform-provider-azuread/internal/helpers/tf/pluginsdk"
	"github.com/hashicorp/terraform-provider-azuread/internal/helpers/tf/validation"
	"github.com/hashicorp/terraform-provider-azuread/internal/services/groups/parse"
)

func groupLicenseAssignmentResource() *pluginsdk.Resource {
	return &pluginsdk.Resource{
		CreateContext: groupLicenseAssignmentResourceCreate,
		ReadContext:   groupLicenseAssignmentResourceRead,
		DeleteContext: groupLicenseAssignmentResourceDelete,

		Timeouts: &pluginsdk.ResourceTimeout{
			Create: pluginsdk.DefaultTimeout(5 * time.Minute),
			Read:   pluginsdk.DefaultTimeout(5 * time.Minute),
			Delete: pluginsdk.DefaultTimeout(5 * time.Minute),
		},

		Importer: pluginsdk.ImporterValidatingResourceId(func(id string) error {
			_, err := parse.GroupLicenseAssignmentID(id)
			return err
		}),

		Schema: map[string]*pluginsdk.Schema{
			"group_object_id": {
				Description:  "The object ID of the group you want to add the member to",
				Type:         pluginsdk.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validation.IsUUID,
			},

			"sku_id": {
				Description:  "The unique identifier for the SKU. Corresponds to the skuId from subscribedSkus or companySubscription.",
				Type:         pluginsdk.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validation.IsUUID,
			},

			"disabled_plans": {
				Description: "A collection of the unique identifiers for plans that have been disabled. IDs are available in servicePlans > servicePlanId in the tenant's subscribedSkus or serviceStatus > servicePlanId in the tenant's companySubscription.",
				Type:        pluginsdk.TypeSet,
				Optional:    true,
				ForceNew:    true,
				Elem: &pluginsdk.Schema{
					Type:         pluginsdk.TypeString,
					ValidateFunc: validation.IsUUID,
				},
			},
		},
	}
}

func groupLicenseAssignmentResourceCreate(ctx context.Context, d *pluginsdk.ResourceData, meta interface{}) pluginsdk.Diagnostics {
	client := meta.(*clients.Client).Groups.GroupClientBeta

	groupId := beta.NewGroupID(d.Get("group_object_id").(string))
	resourceId := parse.NewGroupLicenseAssignmentID(groupId.GroupId, d.Get("sku_id").(string))

	resp, err := client.GetGroup(ctx, groupId, group.GetGroupOperationOptions{
		Select: &[]string{
			"assignedLicenses",
		},
	})
	if err != nil {
		if response.WasNotFound(resp.HttpResponse) {
			return tf.ErrorDiagPathF(nil, "group_object_id", "%s was not found", groupId)
		}
		return tf.ErrorDiagPathF(err, "group_object_id", "Retrieving %s", groupId)
	}

	license := getGroupLicense(resp.Model.AssignedLicenses, resourceId.SKUId)
	if license != nil {
		return tf.ImportAsExistsDiag("azuread_group_license_assignment", resourceId.String())
	}

	if _, err := client.AssignLicense(ctx, groupId, group.AssignLicenseRequest{
		AddLicenses: &[]beta.AssignedLicense{
			{
				SkuId:         nullable.Value(resourceId.SKUId),
				DisabledPlans: tf.ExpandStringSlicePtr(d.Get("disabled_plans").(*pluginsdk.Set).List()),
			},
		},
	}, group.DefaultAssignLicenseOperationOptions()); err != nil {
		return tf.ErrorDiagF(err, "Assigning license to %s", groupId)
	}

	d.SetId(resourceId.String())

	return groupLicenseAssignmentResourceRead(ctx, d, meta)
}

func groupLicenseAssignmentResourceRead(ctx context.Context, d *pluginsdk.ResourceData, meta interface{}) pluginsdk.Diagnostics {
	client := meta.(*clients.Client).Groups.GroupClientBeta

	resourceId, err := parse.GroupLicenseAssignmentID(d.Id())
	if err != nil {
		return tf.ErrorDiagPathF(err, "id", "Parsing Group License Assignment ID %q", d.Id())
	}

	resp, err := client.GetGroup(ctx, beta.NewGroupID(resourceId.GroupId), group.GetGroupOperationOptions{
		Select: &[]string{
			"assignedLicenses",
		},
	})
	if err != nil {
		if response.WasNotFound(resp.HttpResponse) {
			return tf.ErrorDiagPathF(nil, "group_object_id", "%s was not found", resourceId.GroupId)
		}
		return tf.ErrorDiagPathF(err, "group_object_id", "Retrieving %s", resourceId.GroupId)
	}

	license := getGroupLicense(resp.Model.AssignedLicenses, resourceId.SKUId)

	if license == nil {
		return tf.ErrorDiagF(err, "Retrieving license %s for group with object ID: %s", resourceId.SKUId, resourceId.GroupId)
	}

	tf.Set(d, "group_object_id", resourceId.GroupId)
	tf.Set(d, "sku_id", resourceId.SKUId)
	tf.Set(d, "disabled_plans", tf.FlattenStringSlicePtr(license.DisabledPlans))

	return nil
}

func groupLicenseAssignmentResourceDelete(ctx context.Context, d *pluginsdk.ResourceData, meta interface{}) pluginsdk.Diagnostics {
	client := meta.(*clients.Client).Groups.GroupClientBeta

	resourceId, err := parse.GroupLicenseAssignmentID(d.Id())
	if err != nil {
		return tf.ErrorDiagPathF(err, "id", "Parsing Group License Assignment ID %q", d.Id())
	}

	resp, err := client.GetGroup(ctx, beta.NewGroupID(resourceId.GroupId), group.GetGroupOperationOptions{
		Select: &[]string{
			"assignedLicenses",
		},
	})
	if err != nil {
		if response.WasNotFound(resp.HttpResponse) {
			// Group is already deleted
			return nil
		}
		return tf.ErrorDiagPathF(err, "group_object_id", "Retrieving %s", resourceId.GroupId)
	}
	license := getGroupLicense(resp.Model.AssignedLicenses, resourceId.SKUId)

	if license == nil {
		// License is already removed
		return nil
	}

	if _, err := client.AssignLicense(ctx, beta.NewGroupID(resourceId.GroupId), group.AssignLicenseRequest{
		RemoveLicenses: &[]string{resourceId.SKUId},
	}, group.DefaultAssignLicenseOperationOptions()); err != nil {
		return tf.ErrorDiagF(err, "Removing license %s to %s", resourceId.SKUId, resourceId.GroupId)
	}

	return nil
}

func getGroupLicense(licenses *[]beta.AssignedLicense, skuId string) *beta.AssignedLicense {
	if licenses != nil {
		for _, v := range *licenses {
			if strings.EqualFold(v.SkuId.GetOrZero(), skuId) {
				return &v
			}
		}
	}

	return nil
}
