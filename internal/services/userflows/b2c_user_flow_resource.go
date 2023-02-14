package userflows

import (
	"context"
	"errors"
	"fmt"
	"log"
	"math"
	"net/http"
	"regexp"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/hashicorp/terraform-provider-azuread/internal/clients"
	"github.com/hashicorp/terraform-provider-azuread/internal/helpers"
	"github.com/hashicorp/terraform-provider-azuread/internal/tf"
	"github.com/hashicorp/terraform-provider-azuread/internal/utils"
	"github.com/hashicorp/terraform-provider-azuread/internal/validate"
	"github.com/manicminer/hamilton/msgraph"
	"github.com/manicminer/hamilton/odata"
)

func b2cUserflowResource() *schema.Resource {
	return &schema.Resource{
		CreateContext: b2cUserFlowResourceCreate,
		ReadContext:   b2cUserFlowResourceRead,
		UpdateContext: b2cUserFlowResourceUpdate,
		DeleteContext: b2cUserFlowResourceDelete,

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(5 * time.Minute),
			Read:   schema.DefaultTimeout(5 * time.Minute),
			Update: schema.DefaultTimeout(5 * time.Minute),
			Delete: schema.DefaultTimeout(5 * time.Minute),
		},

		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: map[string]*schema.Schema{
			"name": {
				Description:      "The name of the user flow. The name must be prefixed with `B2C_1_`",
				Type:             schema.TypeString,
				Required:         true,
				ForceNew:         true,
				ValidateDiagFunc: validate.StringMatches(regexp.MustCompile("^B2C_1_.+"), "must have the prefix `B2C_1_`"),
			},

			"type": {
				Description: "The type of user flow",
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				ValidateFunc: validation.StringInSlice([]string{
					string("signUp"),
					string("signIn"),
					string("signUpOrSignIn"),
					string("passwordReset"),
					string("profileUpdate"),
					string("resourceOwner"),
				}, false),
			},

			"version": {
				Description:      "The version of the user flow",
				Type:             schema.TypeFloat,
				Required:         true,
				ForceNew:         true,
				ValidateDiagFunc: validate.ValidateDiag(validation.FloatAtMost(math.MaxFloat32)),
			},

			"default_language_tag": {
				Description: "Indicates the default language of the b2cIdentityUserFlow that is used when no ui_locale tag is specified in the request. This field is RFC 5646 compliant. Defaults to `en`",
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "en",
			},

			"language_customization_enabled": {
				Description: " The property that determines whether language customization is enabled within the B2C user flow. Language customization is not enabled by default for B2C user flows.",
				Type:        schema.TypeBool,
				Optional:    true,
			},
		},
	}
}

func b2cUserFlowResourceCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*clients.Client).UserFlows.UserFlowClient

	name := d.Get("name").(string)
	version := float32(d.Get("version").(float64))

	userFlow := msgraph.B2CUserFlow{
		ID:                             &name,
		UserFlowTypeVersion:            &version,
		UserFlowType:                   utils.String(d.Get("type").(string)),
		DefaultLanguageTag:             utils.String(d.Get("default_language_tag").(string)),
		IsLanguageCustomizationEnabled: utils.Bool(d.Get("language_customization_enabled").(bool)),
	}

	resp, _, err := client.Create(ctx, userFlow)
	if err != nil {
		return tf.ErrorDiagF(err, "Creating B2C user flow %q", name)
	}

	if resp.ID == nil || *resp.ID == "" {
		return tf.ErrorDiagF(errors.New("API returned nil ID"), "Bad API Response")
	}

	d.SetId(*resp.ID)

	return b2cUserFlowResourceRead(ctx, d, meta)
}

func b2cUserFlowResourceUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*clients.Client).UserFlows.UserFlowClient
	id := d.Id()

	userFlow := msgraph.B2CUserFlow{
		ID:                             &id,
		DefaultLanguageTag:             utils.String(d.Get("default_language_tag").(string)),
		IsLanguageCustomizationEnabled: utils.Bool(d.Get("language_customization_enabled").(bool)),
	}

	_, err := client.Update(ctx, userFlow)
	if err != nil {
		return tf.ErrorDiagF(err, "Could not update B2C User Flow with ID: %q", id)
	}

	return b2cUserFlowResourceRead(ctx, d, meta)
}

func b2cUserFlowResourceRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*clients.Client).UserFlows.UserFlowClient
	id := d.Id()

	userFlow, status, err := client.Get(ctx, id, odata.Query{})
	if err != nil {
		if status == http.StatusNotFound {
			log.Printf("[DEBUG] B2C User Flow with ID %q was not found - removing from state!", id)
			d.SetId("")

			return nil
		}

		return tf.ErrorDiagF(err, "Retrieving B2C User Flow with ID: %q", id)
	}

	tf.Set(d, "name", *userFlow.ID)
	tf.Set(d, "type", *userFlow.UserFlowType)
	tf.Set(d, "version", *userFlow.UserFlowTypeVersion)
	tf.Set(d, "default_language_tag", *userFlow.DefaultLanguageTag)
	tf.Set(d, "language_customization_enabled", *userFlow.IsLanguageCustomizationEnabled)

	return nil
}

func b2cUserFlowResourceDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*clients.Client).UserFlows.UserFlowClient
	id := d.Id()

	status, err := client.Delete(ctx, id)
	if err != nil {
		if status == http.StatusNotFound {
			log.Printf("[DEBUG] Userflow with Object ID %q was not found - removing from state!", id)
			return tf.ErrorDiagPathF(fmt.Errorf("B2C User Flow was not found"), "id", "Retrieving user with object ID %q", id)
		}

		return tf.ErrorDiagPathF(err, "id", "Deleting B2C User Flow with ID %q, got status %d", id, status)
	}

	if err := helpers.WaitForDeletion(ctx, func(ctx context.Context) (*bool, error) {
		client.BaseClient.DisableRetries = true
		if _, status, err := client.Get(ctx, id, odata.Query{}); err != nil {
			if status == http.StatusNotFound {
				return utils.Bool(false), nil
			}
			return nil, err
		}
		return utils.Bool(true), nil
	}); err != nil {
		return tf.ErrorDiagF(err, "Waiting for deletion of B2C user flow with ID %q", id)
	}

	return nil
}
