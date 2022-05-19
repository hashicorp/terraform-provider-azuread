package applications

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
	"github.com/hashicorp/terraform-provider-azuread/internal/utils"
)

const applicationExtensionPropertyResourceName = "azuread_application_extension_property"

func applicationExtensionPropertyResource() *schema.Resource {
	return &schema.Resource{
		CreateContext: applicationExtensionPropertyResourceCreate,
		ReadContext:   applicationExtensionPropertyResourceRead,
		UpdateContext: applicationExtensionPropertyResourceUpdate,
		DeleteContext: applicationExtensionPropertyResourceDelete,

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(10 * time.Minute),
			Read:   schema.DefaultTimeout(5 * time.Minute),
			Update: schema.DefaultTimeout(10 * time.Minute),
			Delete: schema.DefaultTimeout(5 * time.Minute),
		},

		Importer: tf.ValidateResourceIDPriorToImport(func(id string) error {
			if _, err := uuid.ParseUUID(id); err != nil {
				return fmt.Errorf("specified ID (%q) is not valid: %s", id, err)
			}
			return nil
		}),

		SchemaVersion: 1,

		Schema: map[string]*schema.Schema{
			"id": {
				Description: "The ID",
				Type:        schema.TypeString,
				Computed:    true,
			},

			"application_id": {
				Description: "The application ID to be used to link the extension property",
				Type:        schema.TypeString,
				Required:    true,
			},

			"name": {
				Description: "The extension property name",
				Type:        schema.TypeString,
				Required:    true,
			},

			"app_display_name": {
				Description: "The display name for the application",
				Type:        schema.TypeString,
				Computed:    true,
			},

			"data_type": {
				Description: "The extension property data type",
				Type:        schema.TypeString,
				Required:    true,
				ValidateFunc: validation.StringInSlice([]string{
					msgraph.ApplicationExtensionDataTypeBinary,
					msgraph.ApplicationExtensionDataTypeBoolean,
					msgraph.ApplicationExtensionDataTypeDateTime,
					msgraph.ApplicationExtensionDataTypeInteger,
					msgraph.ApplicationExtensionDataTypeLargeInteger,
					msgraph.ApplicationExtensionDataTypeString,
				}, false),
			},

			"target_objects": {
				Type:     schema.TypeList,
				Required: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
					ValidateFunc: validation.StringInSlice([]string{
						msgraph.ExtensionSchemaTargetTypeAdministrativeUnit,
						msgraph.ExtensionSchemaTargetTypeContact,
						msgraph.ExtensionSchemaTargetTypeDevice,
						msgraph.ExtensionSchemaTargetTypeEvent,
						msgraph.ExtensionSchemaTargetTypeGroup,
						msgraph.ExtensionSchemaTargetTypeMessage,
						msgraph.ExtensionSchemaTargetTypeOrganization,
						msgraph.ExtensionSchemaTargetTypePost,
						msgraph.ExtensionSchemaTargetTypeUser,
					}, false),
				},
			},

			"is_synced_from_on_premises": {
				Type:     schema.TypeBool,
				Computed: true,
			},
		},
	}
}

func applicationExtensionPropertyResourceCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*clients.Client).Applications.ApplicationsClient
	applicationId := d.Get("application_id").(string)
	targetObjects := d.Get("target_objects").([]interface{})

	// Create a new application
	properties := msgraph.ApplicationExtension{
		Name:          utils.String(d.Get("name").(string)),
		DataType:      d.Get("data_type").(msgraph.ApplicationExtensionDataType),
		TargetObjects: tf.ExpandStringSlicePtr(targetObjects),
	}

	appExt, _, err := client.CreateExtension(ctx, properties, applicationId)
	if err != nil {
		return tf.ErrorDiagF(err, "Could not create extension property")
	}

	if appExt.Id == nil || *appExt.Id == "" {
		return tf.ErrorDiagF(errors.New("Bad API response"), "ID returned for extension property is nil/empty")
	}

	d.SetId(*appExt.Id)

	return applicationExtensionPropertyResourceRead(ctx, d, meta)
}

func applicationExtensionPropertyResourceUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	// Update isn't supported
	err := errors.New("Extension property can't be updated once created")
	return tf.ErrorDiagF(err, "Could not update application with object ID: %q", d.Id())
}

func applicationExtensionPropertyResourceRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*clients.Client).Applications.ApplicationsClient
	applicationId := d.Get("application_id").(string)

	log.Printf("[DEBUG] Before listing extensions")
	appExts, status, err := client.ListExtensions(ctx, applicationId, odata.Query{
		Filter: "id eq '" + d.Id() + "'",
	})
	if err != nil {
		if status == http.StatusNotFound {
			log.Printf("[DEBUG] Extension property with ID %q was not found in Application %q - removing from state", d.Id(), applicationId)
			d.SetId("")
			return nil
		}

		return tf.ErrorDiagPathF(err, "id", "Retrieving extension property with ID %q in application id %q", d.Id(), applicationId)
	}

	appExt := (*appExts)[0]

	tf.Set(d, "name", appExt.Name)
	tf.Set(d, "app_display_name", appExt.AppDisplayName)
	tf.Set(d, "data_type", appExt.DataType)
	tf.Set(d, "is_synced_from_on_premises", appExt.IsSyncedFromOnPremises)
	tf.Set(d, "target_objects", appExt.TargetObjects)

	return nil
}

func applicationExtensionPropertyResourceDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*clients.Client).Applications.ApplicationsClient
	extensionId := d.Id()
	appId := d.Get("application_id").(string)

	status, err := client.DeleteExtension(ctx, appId, extensionId)
	if err != nil {
		if status == http.StatusNotFound {
			return tf.ErrorDiagPathF(fmt.Errorf("Extension property was not found"), "id", "Retrieving Extension property with ID %q", extensionId)
		}

		return tf.ErrorDiagPathF(err, "id", "Retrieving extension property with ID %q on application ID %q", extensionId, appId)
	}

	return nil
}
