package userflowattributes

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/davecgh/go-spew/spew"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/manicminer/hamilton/msgraph"
	"github.com/manicminer/hamilton/odata"

	"github.com/hashicorp/terraform-provider-azuread/internal/clients"
	"github.com/hashicorp/terraform-provider-azuread/internal/helpers"
	"github.com/hashicorp/terraform-provider-azuread/internal/tf"
	"github.com/hashicorp/terraform-provider-azuread/internal/utils"
)

func userflowAttributeResource() *schema.Resource {
	return &schema.Resource{
		CreateContext: userflowAttributeResourceCreate,
		ReadContext:   userflowAttributeResourceRead,
		UpdateContext: userflowAttributeResourceUpdate,
		DeleteContext: userflowAttributeResourceDelete,

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(5 * time.Minute),
			Read:   schema.DefaultTimeout(5 * time.Minute),
			Update: schema.DefaultTimeout(5 * time.Minute),
			Delete: schema.DefaultTimeout(5 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"object_id": {
				Description: "The object ID of the userflow attribute",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"display_name": {
				Description: "The display name of the user flow attribute.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"description": {
				Description: "The description of the user flow attribute that's shown to the user at the time of sign-up.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"userflow_attribute_type": {
				Description: "The type of the user flow attribute. This is a read-only attribute that is automatically set. Depending on the type of attribute, the values for this property will be builtIn, custom, or required.",
				Type:        schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{
					string("builtIn"),
					string("custom"),
					string("required"),
				}, false),
				Required: true,
			},
			"data_type": {
				Description: "The data type of the user flow attribute. This cannot be modified after the custom user flow attribute is created. The supported values for dataType are: string, boolean, int64, stringCollection, dateTime.",
				Type:        schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{
					string(msgraph.UserflowAttributeDataTypeString),
					string(msgraph.UserflowAttributeDataTypeInt64),
					string(msgraph.UserflowAttributeDataTypeBoolean),
				}, false),
				Required: true,
			},
		},
	}

}

func userflowAttributeResourceCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*clients.Client).UserflowAttributes.Client

	attr := msgraph.UserFlowAttribute{
		DisplayName:           utils.String(d.Get("display_name").(string)),
		DataType:              utils.String(msgraph.UserflowAttributeDataType(d.Get("data_type").(string))),
		UserFlowAttributeType: utils.String(d.Get("userflow_attribute_type").(string)),
		Description:           utils.String(d.Get("description").(string)),
	}
	userflowAttr, _, err := client.Create(ctx, attr)
	if err != nil {
		return tf.ErrorDiagF(err, "Creating userflow attribute %q", spew.Sdump(attr))
	}

	if userflowAttr.ID == nil || *userflowAttr.ID == "" {
		return tf.ErrorDiagF(errors.New("API returned userflow attribute with nil object ID"), "Bad API Response")
	}

	d.SetId(*userflowAttr.ID)

	return userflowAttributeResourceRead(ctx, d, meta)
}

func userflowAttributeResourceUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*clients.Client).UserflowAttributes.Client

	if d.HasChange("display_name") {
		return tf.ErrorDiagF(errors.New("cannot update display_name"), "Could not update userflow attribute with ID: %q", d.Id())
	}
	if d.HasChange("userflow_attribute_type") {
		return tf.ErrorDiagF(errors.New("cannot update userflow_attribute_type"), "Could not update userflow attribute with ID: %q", d.Id())
	}
	if d.HasChange("data_type") {
		return tf.ErrorDiagF(errors.New("cannot update data_type"), "Could not update userflow attribute with ID: %q", d.Id())
	}

	if !d.HasChange("description") {
		// Nothing to do here.
		return userflowAttributeResourceRead(ctx, d, meta)
	}

	attr := msgraph.UserFlowAttribute{
		ID:          utils.String(d.Id()),
		Description: utils.String(d.Get("description").(string)),
	}

	if _, err := client.Update(ctx, attr); err != nil {
		return tf.ErrorDiagF(err, "Could not update userflow attribute with ID: %q", d.Id())
	}

	return userflowAttributeResourceRead(ctx, d, meta)
}

func userflowAttributeResourceRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*clients.Client).UserflowAttributes.Client

	objectId := d.Id()

	userflowAttr, status, err := client.Get(ctx, objectId, odata.Query{})
	if err != nil {
		if status == http.StatusNotFound {
			log.Printf("[DEBUG] Userflow attribute with Object ID %q was not found - removing from state!", objectId)
			d.SetId("")
			return nil
		}
		return tf.ErrorDiagF(err, "Retrieving userflow attribute with object ID: %q", objectId)
	}

	tf.Set(d, "object_id", *userflowAttr.ID)
	tf.Set(d, "display_name", *userflowAttr.DisplayName)
	tf.Set(d, "description", *userflowAttr.Description)
	tf.Set(d, "userflow_attribute_type", *userflowAttr.UserFlowAttributeType)
	tf.Set(d, "data_type", *userflowAttr.DataType)

	return nil
}

func userflowAttributeResourceDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*clients.Client).UserflowAttributes.Client
	objectId := d.Id()

	_, status, err := client.Get(ctx, objectId, odata.Query{})
	if err != nil {
		if status == http.StatusNotFound {
			return tf.ErrorDiagPathF(fmt.Errorf("Userflow attribute was not found"), "id", "Retrieving user with object ID %q", objectId)
		}

		return tf.ErrorDiagPathF(err, "id", "Retrieving userflow attribute with object ID %q", objectId)
	}

	status, err = client.Delete(ctx, objectId)
	if err != nil {
		return tf.ErrorDiagPathF(err, "id", "Deleting userflow attribute with object ID %q, got status %d", objectId, status)
	}

	// Wait for user object to be deleted
	if err := helpers.WaitForDeletion(ctx, func(ctx context.Context) (*bool, error) {
		client.BaseClient.DisableRetries = true
		if _, status, err := client.Get(ctx, objectId, odata.Query{}); err != nil {
			if status == http.StatusNotFound {
				return utils.Bool(false), nil
			}
			return nil, err
		}
		return utils.Bool(true), nil
	}); err != nil {
		return tf.ErrorDiagF(err, "Waiting for deletion of userflow attribute with object ID %q", objectId)
	}

	return nil
}
