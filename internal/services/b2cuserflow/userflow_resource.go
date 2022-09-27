package b2cuserflow

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/hashicorp/go-uuid"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/manicminer/hamilton/msgraph"
	"github.com/manicminer/hamilton/odata"

	"github.com/hashicorp/terraform-provider-azuread/internal/clients"
	"github.com/hashicorp/terraform-provider-azuread/internal/tf"
	"github.com/hashicorp/terraform-provider-azuread/internal/validate"
)

func b2cUserflowResource() *schema.Resource {
	return &schema.Resource{
		CreateContext: b2cuserflowResourceCreate,
		ReadContext:   b2cuserflowResourceRead,
		UpdateContext: b2cuserflowResourceUpdate,
		DeleteContext: b2cuserflowResourceDelete,

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(5 * time.Minute),
			Read:   schema.DefaultTimeout(5 * time.Minute),
			Update: schema.DefaultTimeout(5 * time.Minute),
			Delete: schema.DefaultTimeout(5 * time.Minute),
		},

		Importer: tf.ValidateResourceIDPriorToImport(func(id string) error {
			if _, err := uuid.ParseUUID(id); err != nil {
				return fmt.Errorf("specified ID (%q) is not valid: %s", id, err)
			}
			return nil
		}),

		Schema: map[string]*schema.Schema{
			"name": {
				Description:      "The name of the user flow. This is a required value and is immutable after it's created. The name will be prefixed with the value of B2C_1_ after creation.",
				Type:             schema.TypeString,
				Required:         true,
				ValidateDiagFunc: validate.NoEmptyStrings,
			},
			"user_flow_type": {
				Description: "The type of user flow. The supported values for userFlowType are: signUp, signIn, signUpOrSignIn, passwordReset, profileUpdate, resourceOwner.",
				Type:        schema.TypeString,
				Required:    true,
				ValidateFunc: validation.StringInSlice([]string{
					string("signUp"),
					string("signIn"),
					string("signUpOrSignIn"),
					string("passwordReset"),
					string("profileUpdate"),
					string("resourceOwner"),
				}, false),
			},

			"user_flow_type_version": {
				Description: "The version of the user flow",
				Type:        schema.TypeFloat,
				Required:    true,
			},
			"default_language_tag": {
				Description: "Indicates the default language of the b2cIdentityUserFlow that is used when no ui_locale tag is specified in the request. This field is RFC 5646 compliant.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"is_language_customization_enabled": {
				Description: " The property that determines whether language customization is enabled within the B2C user flow. Language customization is not enabled by default for B2C user flows.",
				Type:        schema.TypeBool,
				Optional:    true,
			},
		},
	}
}

func b2cuserflowResourceCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*clients.Client).B2CUserFlow.UserFlowClient
	id := d.Get("name").(string)
	userflowType := d.Get("user_flow_type").(string)
	userflowTypeVersion := float32(d.Get("user_flow_type_version").(float64))
	defaultTag := d.Get("default_language_tag").(string)
	isLanguageCustomizationEnabled := d.Get("is_language_customization_enabled").(bool)

	userflow := msgraph.B2CUserFlow{
		ID:                             &id,
		UserFlowType:                   &userflowType,
		UserFlowTypeVersion:            &userflowTypeVersion,
		DefaultLanguageTag:             &defaultTag,
		IsLanguageCustomizationEnabled: &isLanguageCustomizationEnabled,
	}
	userflowResp, _, err := client.Create(ctx, userflow)
	if err != nil {
		return tf.ErrorDiagF(err, "Creating userflow %q", userflow)
	}

	if userflowResp.ID == nil || *userflowResp.ID == "" {
		return tf.ErrorDiagF(errors.New("API returned nil object ID"), "Bad API Response")
	}

	d.SetId(fmt.Sprintf("B2C_1_%s", *userflow.ID))
	return b2cuserflowResourceRead(ctx, d, meta)
}

func b2cuserflowResourceUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return nil
}

func b2cuserflowResourceRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*clients.Client).B2CUserFlow.UserFlowClient

	objectId := d.Id()

	userflow, status, err := client.Get(ctx, objectId, odata.Query{})
	if err != nil {
		if status == http.StatusNotFound {
			log.Printf("[DEBUG] Userflow with Object ID %q was not found - removing from state!", objectId)
			d.SetId("")
			return nil
		}
		return tf.ErrorDiagF(err, "Retrieving userflow with object ID: %q", objectId)
	}

	tf.Set(d, "id", *userflow.ID)
	tf.Set(d, "user_flow_type", *userflow.UserFlowType)
	tf.Set(d, "user_flow_type_version", *userflow.UserFlowTypeVersion)
	tf.Set(d, "default_language_tag", *userflow.DefaultLanguageTag)
	tf.Set(d, "is_language_customization_enabled", *userflow.IsLanguageCustomizationEnabled)
	return nil
}

func b2cuserflowResourceDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return nil
}
