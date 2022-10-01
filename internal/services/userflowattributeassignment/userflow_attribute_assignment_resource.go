package userflowattributeassignment

import (
	"context"
	"errors"
	"log"
	"net/http"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/manicminer/hamilton/msgraph"
	"github.com/manicminer/hamilton/odata"

	"github.com/hashicorp/terraform-provider-azuread/internal/clients"
	"github.com/hashicorp/terraform-provider-azuread/internal/helpers"
	"github.com/hashicorp/terraform-provider-azuread/internal/tf"
	"github.com/hashicorp/terraform-provider-azuread/internal/utils"
	"github.com/hashicorp/terraform-provider-azuread/internal/validate"
)

func userflowAttributeAssignmentResource() *schema.Resource {
	return &schema.Resource{
		CreateContext: userflowAttributeAssignmentResourceCreate,
		ReadContext:   userflowAttributeAssignmentResourceRead,
		UpdateContext: userflowAttributeAssignmentResourceUpdate,
		DeleteContext: userflowAttributeAssignmentResourceDelete,

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(5 * time.Minute),
			Read:   schema.DefaultTimeout(5 * time.Minute),
			Update: schema.DefaultTimeout(5 * time.Minute),
			Delete: schema.DefaultTimeout(5 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"object_id": {
				Description: "The object ID of the assignment",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"display_name": {
				Description:      "The display name of the identityUserFlowAttribute within a user flow.",
				Type:             schema.TypeString,
				Required:         true,
				ValidateDiagFunc: validate.NoEmptyStrings,
			},
			"is_optional": {
				Description: "Determines whether the identityUserFlowAttribute is optional. true means the user doesn't have to provide a value. false means the user cannot complete sign-up without providing a value.",
				Type:        schema.TypeBool,
				Optional:    true,
			},

			"requires_verification": {
				Description: "Determines whether the identityUserFlowAttribute requires verification. This is only used for verifying the user's phone number or email address.",
				Type:        schema.TypeBool,
				Optional:    true,
			},
			"user_attribute_values": {
				Description: "The input options for the user flow attribute. Only applicable when the userInputType is radioSingleSelect, dropdownSingleSelect, or checkboxMultiSelect.",
				Type:        schema.TypeList,
				Optional:    true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"user_input_type": {
				Description: "The input type of the user flow attribute. Possible values are: textBox, dateTimeDropdown, radioSingleSelect, dropdownSingleSelect, emailBox, checkboxMultiSelect.",
				Type:        schema.TypeString,
				Optional:    true,
				ValidateFunc: validation.StringInSlice([]string{
					msgraph.UserInputTypeRadioSingleSelect,
					msgraph.UserInputTypeTextBox,
					msgraph.UserInputTypeEmailBox,
					msgraph.UserInputTypeDateTimeDropdown,
					msgraph.UserInputTypeRadioSingleSelect,
					msgraph.UserInputTypeCheckboxMultiSelect,
					msgraph.UserInputTypeDateTimeDropdown,
				}, false),
			},
			"userflow_attribute_id": {
				Description:      "The ID of the user flow attribute that should be assigned to the specified user flow",
				Type:             schema.TypeString,
				Required:         true,
				ValidateDiagFunc: validate.NoEmptyStrings,
			},
			"userflow_id": {
				Description:      "The ID of the user flow that should be assigned the user flow attribute",
				Type:             schema.TypeString,
				Required:         true,
				ValidateDiagFunc: validate.NoEmptyStrings,
			},
		},
	}
}

func userflowAttributeAssignmentResourceCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*clients.Client).B2CUserFlow.UserFlowClient

	attrAssignment := msgraph.UserFlowAttributeAssignment{
		DisplayName:          utils.String(d.Get("display_name").(string)),
		IsOptional:           utils.Bool(d.Get("is_optional").(bool)),
		RequiresVerification: utils.Bool(d.Get("requires_verification").(bool)),
		UserInputType:        utils.String(d.Get("user_input_type").(string)),
		UserAttribute: &msgraph.UserFlowAttribute{
			ID: utils.String(d.Get("userflow_attribute_id").(string)),
		},
	}
	if v, ok := d.GetOk("user_attribute_values"); ok {
		attrAssignment.UserAttributeValues = *tf.ExpandStringSlicePtr(v.([]interface{}))
	}
	userflowID := d.Get("userflow_id").(string)
	resp, _, err := client.AssignAttribute(ctx, userflowID, attrAssignment)
	if err != nil {
		return tf.ErrorDiagF(err, "Creating attribute assignment %+v", attrAssignment)
	}

	if resp.ID == nil || *resp.ID == "" {
		return tf.ErrorDiagF(errors.New("API returned nil object ID"), "Bad API Response")
	}

	d.SetId(*resp.ID)
	d.Set("userflow_id", userflowID)

	return userflowAttributeAssignmentResourceRead(ctx, d, meta)
}

func userflowAttributeAssignmentResourceUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	if d.HasChange("user_attribute_id") {
		return tf.ErrorDiagF(errors.New("Cannot update user_attribute_id"), "Cannot update user_attribute_id")
	}
	if d.HasChange("userflow_id") {
		return tf.ErrorDiagF(errors.New("Cannot update userflow_id"), "Cannot update userflow_id")
	}

	userflowId := d.Get("userflow_id").(string)
	attrAssignment := msgraph.UserFlowAttributeAssignment{
		ID:                   utils.String(d.Id()),
		DisplayName:          utils.String(d.Get("display_name").(string)),
		IsOptional:           utils.Bool(d.Get("is_optional").(bool)),
		RequiresVerification: utils.Bool(d.Get("requires_verification").(bool)),
		UserAttributeValues:  d.Get("user_attribute_values").([]string),
		UserInputType:        utils.String(d.Get("user_input_type").(string)),
	}

	client := meta.(*clients.Client).B2CUserFlow.UserFlowClient
	_, _, err := client.UpdateAttributeAssignment(ctx, userflowId, attrAssignment)
	if err != nil {
		return tf.ErrorDiagF(err, "Could not update user attribute assignment with ID: %q userflow ID: %q", d.Id(), userflowId)
	}
	return userflowAttributeAssignmentResourceRead(ctx, d, meta)
}

func userflowAttributeAssignmentResourceRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*clients.Client).B2CUserFlow.UserFlowClient

	objectId := d.Id()
	userflowID := d.Get("userflow_id").(string)

	query := odata.Query{
		Expand: odata.Expand{
			Relationship: "userAttribute",
		},
	}
	resp, status, err := client.GetAssignedAttribute(ctx, userflowID, objectId, query)
	if err != nil {
		if status == http.StatusNotFound {
			log.Printf("[DEBUG] Attribute assignment with Object ID %q was not found - removing from state!", objectId)
			d.SetId("")
			return nil
		}
		return tf.ErrorDiagF(err, "Retrieving attribute assignment with object ID: %q", objectId)
	}

	tf.Set(d, "object_id", objectId)
	tf.Set(d, "display_name", *resp.DisplayName)
	tf.Set(d, "is_optional", *resp.IsOptional)
	tf.Set(d, "requires_verification", *resp.RequiresVerification)
	tf.Set(d, "user_attribute_values", resp.UserAttributeValues)
	tf.Set(d, "user_input_type", *resp.UserInputType)
	tf.Set(d, "userflow_attribute_id", *resp.UserAttribute.ID)

	return nil
}

func userflowAttributeAssignmentResourceDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*clients.Client).B2CUserFlow.UserFlowClient

	objectId := d.Id()
	userflowId := d.Get("userflow_id").(string)

	status, err := client.RemoveAttributeAssignment(ctx, userflowId, objectId)
	if err != nil {
		if status == http.StatusNotFound {
			log.Printf("[DEBUG] Attribute assignment with Object ID %q for userflow ID %q was not found - removing from state!", objectId, userflowId)
			d.SetId("")
			return nil
		}
		return tf.ErrorDiagPathF(err, "id", "Deleting attribute assignment with object ID %q, got status %d", objectId, status)
	}

	// Wait for attribute assignment object to be deleted
	if err := helpers.WaitForDeletion(ctx, func(ctx context.Context) (*bool, error) {
		_, status, err := client.GetAssignedAttribute(ctx, userflowId, objectId, odata.Query{})
		if err != nil {
			return nil, err
		}
		return utils.Bool(status == http.StatusOK), nil
	}); err != nil {
		return tf.ErrorDiagF(err, "Waiting for deletion of attribute assignment with object ID %q for userflow ID %q", objectId, userflowId)
	}

	return nil
}
