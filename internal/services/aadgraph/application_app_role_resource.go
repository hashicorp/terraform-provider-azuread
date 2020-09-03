package aadgraph

import (
	"fmt"
	"log"

	"github.com/Azure/azure-sdk-for-go/services/graphrbac/1.6/graphrbac"
	"github.com/hashicorp/go-uuid"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"

	"github.com/terraform-providers/terraform-provider-azuread/internal/clients"
	"github.com/terraform-providers/terraform-provider-azuread/internal/services/aadgraph/graph"
	"github.com/terraform-providers/terraform-provider-azuread/internal/tf"
	"github.com/terraform-providers/terraform-provider-azuread/internal/utils"
	"github.com/terraform-providers/terraform-provider-azuread/internal/validate"
)

func applicationAppRoleResource() *schema.Resource {
	return &schema.Resource{
		Create: applicationAppRoleResourceCreateUpdate,
		Update: applicationAppRoleResourceCreateUpdate,
		Read:   applicationAppRoleResourceRead,
		Delete: applicationAppRoleResourceDelete,

		Importer: tf.ValidateResourceIDPriorToImport(func(id string) error {
			_, err := graph.ParseAppRoleId(id)
			return err
		}),

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

func applicationAppRoleResourceCreateUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*clients.AadClient).AadGraph.ApplicationsClient
	ctx := meta.(*clients.AadClient).StopContext

	objectId := d.Get("application_object_id").(string)

	// errors should be handled by the validation
	var roleId string
	if v, ok := d.GetOk("role_id"); ok {
		roleId = v.(string)
	} else {
		rid, err := uuid.GenerateUUID()
		if err != nil {
			return fmt.Errorf("generating App Role for Object ID %q: %+v", objectId, err)
		}
		roleId = rid
	}

	allowedMemberTypesRaw := d.Get("allowed_member_types").(*schema.Set).List()
	allowedMemberTypes := make([]string, 0, len(allowedMemberTypesRaw))
	for _, a := range allowedMemberTypesRaw {
		allowedMemberTypes = append(allowedMemberTypes, a.(string))
	}

	role := graphrbac.AppRole{
		AllowedMemberTypes: &allowedMemberTypes,
		ID:                 utils.String(roleId),
		Description:        utils.String(d.Get("description").(string)),
		DisplayName:        utils.String(d.Get("display_name").(string)),
		IsEnabled:          utils.Bool(d.Get("is_enabled").(bool)),
	}

	if v, ok := d.GetOk("value"); ok {
		role.Value = utils.String(v.(string))
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

	var newRoles *[]graphrbac.AppRole

	if d.IsNewResource() {
		newRoles, err = graph.AppRoleAdd(app.AppRoles, &role)
		if err != nil {
			if _, ok := err.(*graph.AlreadyExistsError); ok {
				return tf.ImportAsExistsError("azuread_application_app_role", id.String())
			}
			return fmt.Errorf("adding App Role: %+v", err)
		}
	} else {
		if existing := graph.AppRoleFindById(app, id.RoleId); existing == nil {
			return fmt.Errorf("App Role with ID %q was not found for Application %q", id.RoleId, id.ObjectId)
		}

		newRoles, err = graph.AppRoleUpdate(app.AppRoles, &role)
		if err != nil {
			return fmt.Errorf("updating App Role: %s", err)
		}
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
	d.Set("allowed_member_types", role.AllowedMemberTypes)
	d.Set("description", role.Description)
	d.Set("display_name", role.DisplayName)
	d.Set("is_enabled", role.IsEnabled)
	d.Set("value", role.Value)

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

	log.Printf("[DEBUG] Disabling App Role %q for Application %q prior to removal", id.RoleId, id.ObjectId)
	newRoles, err := graph.AppRoleResultDisableById(app.AppRoles, id.RoleId)
	if err != nil {
		return fmt.Errorf("deleting App Role: %s", err)
	}

	properties := graphrbac.ApplicationUpdateParameters{
		AppRoles: newRoles,
	}
	if _, err := client.Patch(ctx, id.ObjectId, properties); err != nil {
		return fmt.Errorf("patching Application with ID %q: %+v", id.ObjectId, err)
	}

	log.Printf("[DEBUG] Removing App Role %q for Application %q", id.RoleId, id.ObjectId)
	properties = graphrbac.ApplicationUpdateParameters{
		AppRoles: graph.AppRoleResultRemoveById(app.AppRoles, id.RoleId),
	}
	if _, err := client.Patch(ctx, id.ObjectId, properties); err != nil {
		return fmt.Errorf("patching Application with ID %q: %+v", id.ObjectId, err)
	}

	return nil
}
