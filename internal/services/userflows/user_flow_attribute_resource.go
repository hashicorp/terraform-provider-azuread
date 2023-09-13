// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package userflows

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/hashicorp/go-azure-sdk/sdk/odata"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-provider-azuread/internal/clients"
	"github.com/hashicorp/terraform-provider-azuread/internal/helpers"
	"github.com/hashicorp/terraform-provider-azuread/internal/tf"
	"github.com/hashicorp/terraform-provider-azuread/internal/tf/pluginsdk"
	"github.com/hashicorp/terraform-provider-azuread/internal/tf/validation"
	"github.com/hashicorp/terraform-provider-azuread/internal/utils"
	"github.com/manicminer/hamilton/msgraph"
)

func userFlowAttributeResource() *pluginsdk.Resource {
	return &pluginsdk.Resource{
		CreateContext: userFlowAttributeResourceCreate,
		ReadContext:   userFlowAttributeResourceRead,
		UpdateContext: userFlowAttributeResourceUpdate,
		DeleteContext: userFlowAttributeResourceDelete,

		Timeouts: &pluginsdk.ResourceTimeout{
			Create: pluginsdk.DefaultTimeout(5 * time.Minute),
			Read:   pluginsdk.DefaultTimeout(5 * time.Minute),
			Update: pluginsdk.DefaultTimeout(5 * time.Minute),
			Delete: pluginsdk.DefaultTimeout(5 * time.Minute),
		},

		Schema: map[string]*pluginsdk.Schema{
			"display_name": {
				Description: "The display name of the user flow attribute.",
				Type:        pluginsdk.TypeString,
				Required:    true,
				ForceNew:    true,
			},

			"data_type": {
				Description: "The data type of the user flow attribute",
				Type:        pluginsdk.TypeString,
				Required:    true,
				ForceNew:    true,
				ValidateFunc: validation.StringInSlice([]string{
					msgraph.UserflowAttributeDataTypeBoolean,
					msgraph.UserflowAttributeDataTypeDateTime,
					msgraph.UserflowAttributeDataTypeInt64,
					msgraph.UserflowAttributeDataTypeString,
					msgraph.UserflowAttributeDataTypeStringCollection,
				}, false),
			},

			"description": {
				Description: "The description of the user flow attribute that is shown to the user at the time of sign-up",
				Type:        pluginsdk.TypeString,
				Required:    true,
			},

			"attribute_type": {
				Description: "The type of the user flow attribute",
				Type:        pluginsdk.TypeString,
				Computed:    true,
			},
		},
	}
}

func userFlowAttributeResourceCreate(ctx context.Context, d *pluginsdk.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*clients.Client).UserFlows.UserFlowAttributesClient

	displayName := d.Get("display_name").(string)

	query := odata.Query{Filter: fmt.Sprintf("displayName eq '%s'", displayName)}
	if result, _, err := client.List(ctx, query); err == nil {
		for _, r := range *result {
			if r.ID != nil && r.DisplayName != nil && strings.EqualFold(*r.DisplayName, displayName) {
				return tf.ImportAsExistsDiag("azuread_user_flow_attribute", *r.ID)
			}
		}
	} else {
		return tf.ErrorDiagF(err, "Checking for existing user flow attribute: %q", displayName)
	}

	attr := msgraph.UserFlowAttribute{
		DataType:    utils.String(d.Get("data_type").(string)),
		Description: utils.String(d.Get("description").(string)),
		DisplayName: utils.String(displayName),
	}

	userFlowAttr, _, err := client.Create(ctx, attr)
	if err != nil {
		return tf.ErrorDiagF(err, "Creating user flow attribute %q", displayName)
	}

	if userFlowAttr.ID == nil || *userFlowAttr.ID == "" {
		return tf.ErrorDiagF(errors.New("API returned user flow attribute with nil ID"), "Bad API Response")
	}

	d.SetId(*userFlowAttr.ID)

	return userFlowAttributeResourceRead(ctx, d, meta)
}

func userFlowAttributeResourceUpdate(ctx context.Context, d *pluginsdk.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*clients.Client).UserFlows.UserFlowAttributesClient
	id := d.Id()

	attr := msgraph.UserFlowAttribute{
		ID:          utils.String(id),
		Description: utils.String(d.Get("description").(string)),
	}

	if _, err := client.Update(ctx, attr); err != nil {
		return tf.ErrorDiagF(err, "Could not update user flow attribute with ID: %q", id)
	}

	return userFlowAttributeResourceRead(ctx, d, meta)
}

func userFlowAttributeResourceRead(ctx context.Context, d *pluginsdk.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*clients.Client).UserFlows.UserFlowAttributesClient
	id := d.Id()

	userFlowAttr, status, err := client.Get(ctx, id, odata.Query{})
	if err != nil {
		if status == http.StatusNotFound {
			log.Printf("[DEBUG] User flow attribute with ID %q was not found - removing from state!", id)
			d.SetId("")
			return nil
		}
		return tf.ErrorDiagF(err, "Retrieving user flow attribute with ID: %q", id)
	}

	tf.Set(d, "attribute_type", userFlowAttr.UserFlowAttributeType)
	tf.Set(d, "data_type", userFlowAttr.DataType)
	tf.Set(d, "description", userFlowAttr.Description)
	tf.Set(d, "display_name", userFlowAttr.DisplayName)

	return nil
}

func userFlowAttributeResourceDelete(ctx context.Context, d *pluginsdk.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*clients.Client).UserFlows.UserFlowAttributesClient
	id := d.Id()

	_, status, err := client.Get(ctx, id, odata.Query{})
	if err != nil {
		if status == http.StatusNotFound {
			return tf.ErrorDiagPathF(fmt.Errorf("user flow attribute was not found"), "id", "Retrieving user with ID %q", id)
		}

		return tf.ErrorDiagPathF(err, "id", "Retrieving user flow attribute with ID %q", id)
	}

	status, err = client.Delete(ctx, id)
	if err != nil {
		return tf.ErrorDiagPathF(err, "id", "Deleting user flow attribute with ID %q, got status %d", id, status)
	}

	if err := helpers.WaitForDeletion(ctx, func(ctx context.Context) (*bool, error) {
		defer func() { client.BaseClient.DisableRetries = false }()
		client.BaseClient.DisableRetries = true
		if _, status, err := client.Get(ctx, id, odata.Query{}); err != nil {
			if status == http.StatusNotFound {
				return utils.Bool(false), nil
			}
			return nil, err
		}
		return utils.Bool(true), nil
	}); err != nil {
		return tf.ErrorDiagF(err, "Waiting for deletion of user flow attribute with ID %q", id)
	}

	return nil
}
