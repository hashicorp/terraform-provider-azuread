package aadgraph

import (
	"context"
	"fmt"
	"log"

	"github.com/Azure/azure-sdk-for-go/services/graphrbac/1.6/graphrbac"
	"github.com/hashicorp/go-uuid"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"

	"github.com/terraform-providers/terraform-provider-azuread/internal/clients"
	"github.com/terraform-providers/terraform-provider-azuread/internal/services/aadgraph/graph"
	"github.com/terraform-providers/terraform-provider-azuread/internal/tf"
	"github.com/terraform-providers/terraform-provider-azuread/internal/utils"
	"github.com/terraform-providers/terraform-provider-azuread/internal/validate"
)

func applicationOAuth2PermissionResource() *schema.Resource {
	return &schema.Resource{
		CreateContext: applicationOAuth2PermissionResourceCreateUpdate,
		UpdateContext: applicationOAuth2PermissionResourceCreateUpdate,
		ReadContext:   applicationOAuth2PermissionResourceRead,
		DeleteContext: applicationOAuth2PermissionResourceDelete,

		Importer: tf.ValidateResourceIDPriorToImport(func(id string) error {
			_, err := graph.ParseOAuth2PermissionId(id)
			return err
		}),

		Schema: map[string]*schema.Schema{
			"application_object_id": {
				Type:             schema.TypeString,
				Required:         true,
				ForceNew:         true,
				ValidateDiagFunc: validate.UUID,
			},

			"admin_consent_description": {
				Type:             schema.TypeString,
				Required:         true,
				ValidateDiagFunc: validate.NoEmptyStrings,
			},

			"admin_consent_display_name": {
				Type:             schema.TypeString,
				Required:         true,
				ValidateDiagFunc: validate.NoEmptyStrings,
			},

			"is_enabled": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},

			"permission_id": {
				Type:             schema.TypeString,
				Optional:         true,
				Computed:         true,
				ForceNew:         true,
				ValidateDiagFunc: validate.UUID,
			},

			"type": {
				Type:     schema.TypeString,
				Required: true,
				ValidateFunc: validation.StringInSlice(
					[]string{"Admin", "User"},
					false,
				),
			},

			"user_consent_description": {
				Type:             schema.TypeString,
				Required:         true,
				ValidateDiagFunc: validate.NoEmptyStrings,
			},

			"user_consent_display_name": {
				Type:             schema.TypeString,
				Required:         true,
				ValidateDiagFunc: validate.NoEmptyStrings,
			},

			"value": {
				Type:             schema.TypeString,
				Required:         true,
				ValidateDiagFunc: validate.NoEmptyStrings,
			},
		},
	}
}

func applicationOAuth2PermissionResourceCreateUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*clients.AadClient).AadGraph.ApplicationsClient

	objectId := d.Get("application_object_id").(string)

	// errors should be handled by the validation
	var permissionId string

	if v, ok := d.GetOk("permission_id"); ok {
		permissionId = v.(string)
	} else {
		pid, err := uuid.GenerateUUID()
		if err != nil {
			return tf.ErrorDiag(fmt.Sprintf("Generating App Role for application with object ID %q", objectId), err.Error(), "")
		}
		permissionId = pid
	}

	permission := graphrbac.OAuth2Permission{
		AdminConsentDescription: utils.String(d.Get("admin_consent_description").(string)),
		AdminConsentDisplayName: utils.String(d.Get("admin_consent_display_name").(string)),
		ID:                      utils.String(permissionId),
		IsEnabled:               utils.Bool(d.Get("is_enabled").(bool)),
		Type:                    utils.String(d.Get("type").(string)),
		UserConsentDescription:  utils.String(d.Get("user_consent_description").(string)),
		UserConsentDisplayName:  utils.String(d.Get("user_consent_display_name").(string)),
		Value:                   utils.String(d.Get("value").(string)),
	}

	id := graph.OAuth2PermissionIdFrom(objectId, *permission.ID)

	tf.LockByName(resourceApplicationName, id.ObjectId)
	defer tf.UnlockByName(resourceApplicationName, id.ObjectId)

	// ensure the Application Object exists
	app, err := client.Get(ctx, id.ObjectId)
	if err != nil {
		if utils.ResponseWasNotFound(app.Response) {
			return tf.ErrorDiag(fmt.Sprintf("Application with object ID %q was not found", objectId), "", "application_object_id")
		}
		return tf.ErrorDiag(fmt.Sprintf("Retrieving application with object ID %q", objectId), err.Error(), "application_object_id")
	}

	var newPermissions *[]graphrbac.OAuth2Permission

	if d.IsNewResource() {
		newPermissions, err = graph.OAuth2PermissionAdd(app.Oauth2Permissions, &permission)
		if err != nil {
			if _, ok := err.(*graph.AlreadyExistsError); ok {
				return tf.ImportAsExistsDiag("azuread_application_oauth2_permission", id.String())
			}
			return tf.ErrorDiag("Failed to add OAuth2 Permission", err.Error(), "")
		}
	} else {
		if existing, _ := graph.OAuth2PermissionFindById(app, id.PermissionId); existing == nil {
			return tf.ErrorDiag(fmt.Sprintf("OAuth2 Permission with ID %q was not found for Application %q", id.PermissionId, id.ObjectId), "", "role_id")
		}

		newPermissions, err = graph.OAuth2PermissionUpdate(app.Oauth2Permissions, &permission)
		if err != nil {
			return tf.ErrorDiag(fmt.Sprintf("Updating App Role with ID %q", *permission.ID), err.Error(), "")
		}
	}

	properties := graphrbac.ApplicationUpdateParameters{
		Oauth2Permissions: newPermissions,
	}
	if _, err := client.Patch(ctx, id.ObjectId, properties); err != nil {
		return tf.ErrorDiag(fmt.Sprintf("Updating Application with ID %q", id.ObjectId), err.Error(), "name")
	}

	d.SetId(id.String())

	return applicationOAuth2PermissionResourceRead(ctx, d, meta)
}

func applicationOAuth2PermissionResourceRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*clients.AadClient).AadGraph.ApplicationsClient

	id, err := graph.ParseOAuth2PermissionId(d.Id())
	if err != nil {
		return tf.ErrorDiag(fmt.Sprintf("Parsing OAuth2 Permission ID %q", d.Id()), err.Error(), "id")
	}

	// ensure the Application Object exists
	app, err := client.Get(ctx, id.ObjectId)
	if err != nil {
		// the parent Application has been removed - skip it
		if utils.ResponseWasNotFound(app.Response) {
			log.Printf("[DEBUG] Application with Object ID %q was not found - removing from state!", id.ObjectId)
			d.SetId("")
			return nil
		}
		return tf.ErrorDiag(fmt.Sprintf("Retrieving Application with ID %q", id.ObjectId), err.Error(), "application_object_id")
	}

	permission, err := graph.OAuth2PermissionFindById(app, id.PermissionId)
	if err != nil {
		return tf.ErrorDiag("Identifying OAuth2 Permission", err.Error(), "permission_id")
	}

	if permission == nil {
		log.Printf("[DEBUG] OAuth2 Permission %q (ID %q) was not found - removing from state!", id.PermissionId, id.ObjectId)
		d.SetId("")
		return nil
	}

	if err := d.Set("application_object_id", id.ObjectId); err != nil {
		return tf.ErrorDiag("Could not set attribute", err.Error(), "application_object_id")
	}

	if err := d.Set("permission_id", id.PermissionId); err != nil {
		return tf.ErrorDiag("Could not set attribute", err.Error(), "permission_id")
	}

	if err := d.Set("admin_consent_description", permission.AdminConsentDescription); err != nil {
		return tf.ErrorDiag("Could not set attribute", err.Error(), "admin_consent_description")
	}

	if err := d.Set("admin_consent_display_name", permission.AdminConsentDisplayName); err != nil {
		return tf.ErrorDiag("Could not set attribute", err.Error(), "admin_consent_display_name")
	}

	if err := d.Set("is_enabled", permission.IsEnabled); err != nil {
		return tf.ErrorDiag("Could not set attribute", err.Error(), "is_enabled")
	}

	if err := d.Set("type", permission.Type); err != nil {
		return tf.ErrorDiag("Could not set attribute", err.Error(), "type")
	}

	if err := d.Set("user_consent_description", permission.UserConsentDescription); err != nil {
		return tf.ErrorDiag("Could not set attribute", err.Error(), "user_consent_description")
	}

	if err := d.Set("user_consent_display_name", permission.UserConsentDisplayName); err != nil {
		return tf.ErrorDiag("Could not set attribute", err.Error(), "user_consent_display_name")
	}

	if err := d.Set("value", permission.Value); err != nil {
		return tf.ErrorDiag("Could not set attribute", err.Error(), "value")
	}

	return nil
}

func applicationOAuth2PermissionResourceDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*clients.AadClient).AadGraph.ApplicationsClient

	id, err := graph.ParseOAuth2PermissionId(d.Id())
	if err != nil {
		return tf.ErrorDiag(fmt.Sprintf("Parsing OAuth2 Permission ID %q", d.Id()), err.Error(), "id")
	}

	tf.LockByName(resourceApplicationName, id.ObjectId)
	defer tf.UnlockByName(resourceApplicationName, id.ObjectId)

	// ensure the parent Application exists
	app, err := client.Get(ctx, id.ObjectId)
	if err != nil {
		// the parent Application has been removed - skip it
		if utils.ResponseWasNotFound(app.Response) {
			log.Printf("[DEBUG] Application with Object ID %q was not found - removing from state!", id.ObjectId)
			return nil
		}
		return tf.ErrorDiag(fmt.Sprintf("Retrieving Application with ID %q", id.ObjectId), err.Error(), "application_object_id")
	}

	var newPermissions *[]graphrbac.OAuth2Permission

	log.Printf("[DEBUG] Disabling OAuth2 Permission %q for Application %q prior to removal", id.PermissionId, id.ObjectId)
	newPermissions, err = graph.OAuth2PermissionResultDisableById(app.Oauth2Permissions, id.PermissionId)
	if err != nil {
		return tf.ErrorDiag(fmt.Sprintf("Disabling OAuth2 Permission with ID %q for application %q", id.PermissionId, id.ObjectId), err.Error(), "")
	}

	properties := graphrbac.ApplicationUpdateParameters{
		Oauth2Permissions: newPermissions,
	}
	if _, err := client.Patch(ctx, id.ObjectId, properties); err != nil {
		return tf.ErrorDiag(fmt.Sprintf("Updating Application with ID %q", id.ObjectId), err.Error(), "")
	}

	log.Printf("[DEBUG] Removing OAuth2 Permission %q for Application %q", id.PermissionId, id.ObjectId)
	newPermissions, err = graph.OAuth2PermissionResultRemoveById(app.Oauth2Permissions, id.PermissionId)
	if err != nil {
		return tf.ErrorDiag(fmt.Sprintf("Removing OAuth2 Permission with ID %q for application %q", id.PermissionId, id.ObjectId), err.Error(), "")
	}

	properties = graphrbac.ApplicationUpdateParameters{
		Oauth2Permissions: newPermissions,
	}
	if _, err := client.Patch(ctx, id.ObjectId, properties); err != nil {
		return tf.ErrorDiag(fmt.Sprintf("Updating Application with ID %q", id.ObjectId), err.Error(), "")
	}

	return nil
}
