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

func ApplicationAppRoleResource() *schema.Resource {
	return &schema.Resource{
		Create: applicationAppRoleResourceCreate,
		Update: applicationAppRoleResourceUpdate,
		Read:   applicationAppRoleResourceRead,
		Delete: applicationAppRoleResourceDelete,

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

			"allowed_member_types": {
				Type:     schema.TypeSet,
				Required: true,
				MinItems: 1,
				Elem: &schema.Schema{
					Type: schema.TypeString,
					ValidateFunc: validation.StringInSlice(
						[]string{"User", "Application"},
						false,
					),
				},
			},

			"description": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validate.NoEmptyStrings,
			},

			"display_name": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validate.NoEmptyStrings,
			},

			"is_enabled": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},

			"role_id": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ForceNew:     true,
				ValidateFunc: validate.UUID,
			},

			"value": {
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func applicationAppRoleResourceCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*clients.AadClient).AadGraph.ApplicationsClient
	ctx := meta.(*clients.AadClient).StopContext

	objectId := d.Get("application_object_id").(string)

	role, err := graph.AppRoleForResource(d)
	if err != nil {
		return fmt.Errorf("generating App Role for Object ID %q: %+v", objectId, err)
	}

	id := graph.AppRoleIdFrom(objectId, *role.ID)

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

	newRoles, err := graph.AppRoleAdd(app.AppRoles, role)
	if err != nil {
		return tf.ImportAsExistsError("azuread_application_app_role", id.String())
	}

	properties := graphrbac.ApplicationUpdateParameters{
		AppRoles: newRoles,
	}
	if _, err := client.Patch(ctx, id.ObjectId, properties); err != nil {
		return fmt.Errorf("patching Application with ID %q: %+v", id.ObjectId, err)
	}

	d.SetId(id.String())

	return applicationAppRoleResourceRead(d, meta)
}

func applicationAppRoleResourceUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*clients.AadClient).AadGraph.ApplicationsClient
	ctx := meta.(*clients.AadClient).StopContext

	objectId := d.Get("application_object_id").(string)

	role, err := graph.AppRoleForResource(d)
	if err != nil {
		return fmt.Errorf("generating App Role for Object ID %q: %+v", objectId, err)
	}

	id := graph.AppRoleIdFrom(objectId, *role.ID)

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

	if existing := graph.AppRoleFindById(app, id.RoleId); existing == nil {
		return fmt.Errorf("App Role with ID %q was not found for Application %q", id.RoleId, id.ObjectId)
	}

	properties := graphrbac.ApplicationUpdateParameters{
		AppRoles: graph.AppRoleUpdate(app.AppRoles, role),
	}
	if _, err := client.Patch(ctx, id.ObjectId, properties); err != nil {
		return fmt.Errorf("patching Application with ID %q: %+v", id.ObjectId, err)
	}

	d.SetId(id.String())

	return applicationAppRoleResourceRead(d, meta)
}

func applicationAppRoleResourceRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*clients.AadClient).AadGraph.ApplicationsClient
	ctx := meta.(*clients.AadClient).StopContext

	id, err := graph.ParseAppRoleId(d.Id())
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

	role := graph.AppRoleFindById(app, id.RoleId)
	if role == nil {
		log.Printf("[DEBUG] App Role %q (ID %q) was not found - removing from state!", id.RoleId, id.ObjectId)
		d.SetId("")
		return nil
	}

	d.Set("application_object_id", id.ObjectId)
	d.Set("role_id", id.RoleId)

	if allowedMemberTypes := role.AllowedMemberTypes; allowedMemberTypes != nil {
		d.Set("allowed_member_types", allowedMemberTypes)
	}

	if description := role.Description; description != nil {
		d.Set("description", description)
	}

	if displayName := role.DisplayName; displayName != nil {
		d.Set("display_name", displayName)
	}

	if isEnabled := role.IsEnabled; isEnabled != nil {
		d.Set("is_enabled", isEnabled)
	}

	if value := role.Value; value != nil {
		d.Set("value", value)
	}

	return nil
}

func applicationAppRoleResourceDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*clients.AadClient).AadGraph.ApplicationsClient
	ctx := meta.(*clients.AadClient).StopContext

	id, err := graph.ParseAppRoleId(d.Id())
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
		AppRoles: graph.AppRoleResultDisableById(app.AppRoles, id.RoleId),
	}
	if _, err := client.Patch(ctx, id.ObjectId, properties); err != nil {
		return fmt.Errorf("patching Application with ID %q: %+v", id.ObjectId, err)
	}

	properties = graphrbac.ApplicationUpdateParameters{
		AppRoles: graph.AppRoleResultRemoveById(app.AppRoles, id.RoleId),
	}
	if _, err := client.Patch(ctx, id.ObjectId, properties); err != nil {
		return fmt.Errorf("patching Application with ID %q: %+v", id.ObjectId, err)
	}

	return nil
}
