package aadgraph

import (
	"fmt"
	"log"

	"github.com/Azure/azure-sdk-for-go/services/graphrbac/1.6/graphrbac"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"

	"github.com/terraform-providers/terraform-provider-azuread/internal/clients"
	"github.com/terraform-providers/terraform-provider-azuread/internal/services/aadgraph/graph"
	"github.com/terraform-providers/terraform-provider-azuread/internal/tf"
	"github.com/terraform-providers/terraform-provider-azuread/internal/utils"
	"github.com/terraform-providers/terraform-provider-azuread/internal/validate"
)

func ApplicationOAuth2PermissionResource() *schema.Resource {
	return &schema.Resource{
		Create: applicationOAuth2PermissionResourceCreate,
		Update: applicationOAuth2PermissionResourceUpdate,
		Read:   applicationOAuth2PermissionResourceRead,
		Delete: applicationOAuth2PermissionResourceDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"application_object_id": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validate.UUID,
			},

			"admin_consent_description": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validate.NoEmptyStrings,
			},

			"admin_consent_display_name": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validate.NoEmptyStrings,
			},

			"is_enabled": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},

			"permission_id": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ForceNew:     true,
				ValidateFunc: validate.UUID,
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
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validate.NoEmptyStrings,
			},

			"user_consent_display_name": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validate.NoEmptyStrings,
			},

			"value": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validate.NoEmptyStrings,
			},
		},
	}
}

func applicationOAuth2PermissionResourceCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*clients.AadClient).AadGraph.ApplicationsClient
	ctx := meta.(*clients.AadClient).StopContext

	objectId := d.Get("application_object_id").(string)

	permission, err := graph.OAuth2PermissionForResource(d)
	if err != nil {
		return fmt.Errorf("generating App Role for Object ID %q: %+v", objectId, err)
	}

	id := graph.OAuth2PermissionIdFrom(objectId, *permission.ID)

	tf.LockByName(resourceApplicationName, id.ObjectId)
	defer tf.UnlockByName(resourceApplicationName, id.ObjectId)

	// ensure the Application Object exists
	app, err := client.Get(ctx, id.ObjectId)
	if err != nil {
		if utils.ResponseWasNotFound(app.Response) {
			return fmt.Errorf("Application with ID %q was not found", id.ObjectId)
		}
		return fmt.Errorf("retrieving Application ID %q: %+v", id.ObjectId, err)
	}

	newRoles, err := graph.OAuth2PermissionAdd(app.Oauth2Permissions, permission)
	if err != nil {
		return tf.ImportAsExistsError("azuread_application_oauth2_permission", id.String())
	}

	properties := graphrbac.ApplicationUpdateParameters{
		Oauth2Permissions: newRoles,
	}
	if _, err := client.Patch(ctx, id.ObjectId, properties); err != nil {
		return fmt.Errorf("patching Application with ID %q: %+v", id.ObjectId, err)
	}

	d.SetId(id.String())

	return applicationOAuth2PermissionResourceRead(d, meta)
}

func applicationOAuth2PermissionResourceUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*clients.AadClient).AadGraph.ApplicationsClient
	ctx := meta.(*clients.AadClient).StopContext

	objectId := d.Get("application_object_id").(string)

	permission, err := graph.OAuth2PermissionForResource(d)
	if err != nil {
		return fmt.Errorf("generating App Role for Object ID %q: %+v", objectId, err)
	}

	id := graph.OAuth2PermissionIdFrom(objectId, *permission.ID)

	tf.LockByName(resourceApplicationName, id.ObjectId)
	defer tf.UnlockByName(resourceApplicationName, id.ObjectId)

	// ensure the Application Object exists
	app, err := client.Get(ctx, id.ObjectId)
	if err != nil {
		if utils.ResponseWasNotFound(app.Response) {
			return fmt.Errorf("Application with ID %q was not found", id.ObjectId)
		}
		return fmt.Errorf("retrieving Application ID %q: %+v", id.ObjectId, err)
	}

	if existing, _ := graph.OAuth2PermissionFindById(app, id.PermissionId); existing == nil {
		return fmt.Errorf("App Role with ID %q was not found for Application %q", id.PermissionId, id.ObjectId)
	}

	newPermissions, err := graph.OAuth2PermissionUpdate(app.Oauth2Permissions, permission)
	if err != nil {
		return fmt.Errorf("updating OAuth2 Permission: %s", err)
	}

	properties := graphrbac.ApplicationUpdateParameters{
		Oauth2Permissions: newPermissions,
	}
	if _, err := client.Patch(ctx, id.ObjectId, properties); err != nil {
		return fmt.Errorf("patching Application with ID %q: %+v", id.ObjectId, err)
	}

	d.SetId(id.String())

	return applicationOAuth2PermissionResourceRead(d, meta)
}

func applicationOAuth2PermissionResourceRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*clients.AadClient).AadGraph.ApplicationsClient
	ctx := meta.(*clients.AadClient).StopContext

	id, err := graph.ParseOAuth2PermissionId(d.Id())
	if err != nil {
		return fmt.Errorf("parsing App Role ID: %v", err)
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
		return fmt.Errorf("retrieving Application ID %q: %+v", id.ObjectId, err)
	}

	permission, err := graph.OAuth2PermissionFindById(app, id.PermissionId)
	if err != nil {
		return fmt.Errorf("identifying OAuth2 Permission: %s", err)
	}

	if permission == nil {
		log.Printf("[DEBUG] App Role %q (ID %q) was not found - removing from state!", id.PermissionId, id.ObjectId)
		d.SetId("")
		return nil
	}

	d.Set("application_object_id", id.ObjectId)
	d.Set("permission_id", id.PermissionId)

	if description := permission.AdminConsentDescription; description != nil {
		d.Set("admin_consent_description", description)
	}

	if displayName := permission.AdminConsentDisplayName; displayName != nil {
		d.Set("admin_consent_display_name", displayName)
	}

	if isEnabled := permission.IsEnabled; isEnabled != nil {
		d.Set("is_enabled", isEnabled)
	}

	if permissionType := permission.Type; permissionType != nil {
		d.Set("type", permissionType)
	}

	if description := permission.UserConsentDescription; description != nil {
		d.Set("user_consent_description", description)
	}

	if displayName := permission.UserConsentDisplayName; displayName != nil {
		d.Set("user_consent_display_name", displayName)
	}

	if value := permission.Value; value != nil {
		d.Set("value", value)
	}

	return nil
}

func applicationOAuth2PermissionResourceDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*clients.AadClient).AadGraph.ApplicationsClient
	ctx := meta.(*clients.AadClient).StopContext

	id, err := graph.ParseOAuth2PermissionId(d.Id())
	if err != nil {
		return fmt.Errorf("parsing App Role ID: %v", err)
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
		return fmt.Errorf("retrieving Application ID %q: %+v", id.ObjectId, err)
	}

	properties := graphrbac.ApplicationUpdateParameters{
		Oauth2Permissions: graph.OAuth2PermissionResultDisableById(app.Oauth2Permissions, id.PermissionId),
	}
	if _, err := client.Patch(ctx, id.ObjectId, properties); err != nil {
		return fmt.Errorf("patching Application with ID %q: %+v", id.ObjectId, err)
	}

	properties = graphrbac.ApplicationUpdateParameters{
		Oauth2Permissions: graph.OAuth2PermissionResultRemoveById(app.Oauth2Permissions, id.PermissionId),
	}
	if _, err := client.Patch(ctx, id.ObjectId, properties); err != nil {
		return fmt.Errorf("patching Application with ID %q: %+v", id.ObjectId, err)
	}

	return nil
}
