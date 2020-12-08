package aadgraph

import (
	"context"
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
			return tf.ErrorDiagF(err, "Generating OAuth2 Permision for application with object ID %q", objectId)
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
			return tf.ErrorDiagPathF(nil, "application_object_id", "Application with object ID %q was not found", id.ObjectId)
		}
		return tf.ErrorDiagPathF(err, "application_object_id", "Retrieving Application with object ID %q", id.ObjectId)
	}

	var newPermissions *[]graphrbac.OAuth2Permission

	if d.IsNewResource() {
		newPermissions, err = graph.OAuth2PermissionAdd(app.Oauth2Permissions, &permission)
		if err != nil {
			if _, ok := err.(*graph.AlreadyExistsError); ok {
				return tf.ImportAsExistsDiag("azuread_application_oauth2_permission", id.String())
			}
			return tf.ErrorDiagF(err, "Failed to add OAuth2 Permission")
		}
	} else {
		if existing, _ := graph.OAuth2PermissionFindById(app, id.PermissionId); existing == nil {
			return tf.ErrorDiagPathF(nil, "role_id", "OAuth2 Permission with ID %q was not found for Application %q", id.PermissionId, id.ObjectId)
		}

		newPermissions, err = graph.OAuth2PermissionUpdate(app.Oauth2Permissions, &permission)
		if err != nil {
			return tf.ErrorDiagF(err, "Updating OAuth2 Permission with ID %q", *permission.ID)
		}
	}

	properties := graphrbac.ApplicationUpdateParameters{
		Oauth2Permissions: newPermissions,
	}
	if _, err := client.Patch(ctx, id.ObjectId, properties); err != nil {
		return tf.ErrorDiagF(err, "Updating Application with ID %q", id.ObjectId)
	}

	d.SetId(id.String())

	return applicationOAuth2PermissionResourceRead(ctx, d, meta)
}

func applicationOAuth2PermissionResourceRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*clients.AadClient).AadGraph.ApplicationsClient

	id, err := graph.ParseOAuth2PermissionId(d.Id())
	if err != nil {
		return tf.ErrorDiagPathF(err, "id", "Parsing OAuth2 Permission ID %q", d.Id())
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
		return tf.ErrorDiagPathF(err, "application_object_id", "Retrieving Application with object ID %q", id.ObjectId)
	}

	permission, err := graph.OAuth2PermissionFindById(app, id.PermissionId)
	if err != nil {
		return tf.ErrorDiagF(err, "Identifying OAuth2 Permission")
	}

	if permission == nil {
		log.Printf("[DEBUG] OAuth2 Permission %q (ID %q) was not found - removing from state!", id.PermissionId, id.ObjectId)
		d.SetId("")
		return nil
	}

	if dg := tf.Set(d, "application_object_id", id.ObjectId); dg != nil {
		return dg
	}

	if dg := tf.Set(d, "permission_id", id.PermissionId); dg != nil {
		return dg
	}

	if dg := tf.Set(d, "admin_consent_description", permission.AdminConsentDescription); dg != nil {
		return dg
	}

	if dg := tf.Set(d, "admin_consent_display_name", permission.AdminConsentDisplayName); dg != nil {
		return dg
	}

	if dg := tf.Set(d, "is_enabled", permission.IsEnabled); dg != nil {
		return dg
	}

	if dg := tf.Set(d, "type", permission.Type); dg != nil {
		return dg
	}

	if dg := tf.Set(d, "user_consent_description", permission.UserConsentDescription); dg != nil {
		return dg
	}

	if dg := tf.Set(d, "user_consent_display_name", permission.UserConsentDisplayName); dg != nil {
		return dg
	}

	if dg := tf.Set(d, "value", permission.Value); dg != nil {
		return dg
	}

	return nil
}

func applicationOAuth2PermissionResourceDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*clients.AadClient).AadGraph.ApplicationsClient

	id, err := graph.ParseOAuth2PermissionId(d.Id())
	if err != nil {
		return tf.ErrorDiagPathF(err, "id", "Parsing OAuth2 Permission ID %q", d.Id())
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
		return tf.ErrorDiagPathF(err, "application_object_id", "Retrieving Application with ID %q", id.ObjectId)
	}

	var newPermissions *[]graphrbac.OAuth2Permission

	log.Printf("[DEBUG] Disabling OAuth2 Permission %q for Application %q prior to removal", id.PermissionId, id.ObjectId)
	newPermissions, err = graph.OAuth2PermissionResultDisableById(app.Oauth2Permissions, id.PermissionId)
	if err != nil {
		return tf.ErrorDiagF(err, "Disabling OAuth2 Permission with ID %q for application %q", id.PermissionId, id.ObjectId)
	}

	properties := graphrbac.ApplicationUpdateParameters{
		Oauth2Permissions: newPermissions,
	}
	if _, err := client.Patch(ctx, id.ObjectId, properties); err != nil {
		return tf.ErrorDiagF(err, "Updating Application with ID %q", id.ObjectId)
	}

	log.Printf("[DEBUG] Removing OAuth2 Permission %q for Application %q", id.PermissionId, id.ObjectId)
	newPermissions, err = graph.OAuth2PermissionResultRemoveById(app.Oauth2Permissions, id.PermissionId)
	if err != nil {
		return tf.ErrorDiagF(err, "Removing OAuth2 Permission with ID %q for application %q", id.PermissionId, id.ObjectId)
	}

	properties = graphrbac.ApplicationUpdateParameters{
		Oauth2Permissions: newPermissions,
	}
	if _, err := client.Patch(ctx, id.ObjectId, properties); err != nil {
		return tf.ErrorDiagF(err, "Updating Application with ID %q", id.ObjectId)
	}

	return nil
}
