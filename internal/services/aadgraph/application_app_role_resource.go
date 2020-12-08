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

func applicationAppRoleResource() *schema.Resource {
	return &schema.Resource{
		CreateContext: applicationAppRoleResourceCreateUpdate,
		UpdateContext: applicationAppRoleResourceCreateUpdate,
		ReadContext:   applicationAppRoleResourceRead,
		DeleteContext: applicationAppRoleResourceDelete,

		Importer: tf.ValidateResourceIDPriorToImport(func(id string) error {
			_, err := graph.ParseAppRoleId(id)
			return err
		}),

		Schema: map[string]*schema.Schema{
			"application_object_id": {
				Type:             schema.TypeString,
				Required:         true,
				ForceNew:         true,
				ValidateDiagFunc: validate.UUID,
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
				Type:             schema.TypeString,
				Required:         true,
				ValidateDiagFunc: validate.NoEmptyStrings,
			},

			"display_name": {
				Type:             schema.TypeString,
				Required:         true,
				ValidateDiagFunc: validate.NoEmptyStrings,
			},

			"is_enabled": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},

			"role_id": {
				Type:             schema.TypeString,
				Optional:         true,
				Computed:         true,
				ForceNew:         true,
				ValidateDiagFunc: validate.UUID,
			},

			"value": {
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func applicationAppRoleResourceCreateUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*clients.AadClient).AadGraph.ApplicationsClient

	objectId := d.Get("application_object_id").(string)

	// errors should be handled by the validation
	var roleId string
	if v, ok := d.GetOk("role_id"); ok {
		roleId = v.(string)
	} else {
		rid, err := uuid.GenerateUUID()
		if err != nil {
			return tf.ErrorDiagF(err, "Generating App Role for application with object ID %q", objectId)
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
			return tf.ErrorDiagPathF(nil, "application_object_id", "Application with object ID %q was not found", id.ObjectId)
		}
		return tf.ErrorDiagPathF(err, "application_object_id", "Retrieving Application with object ID %q", id.ObjectId)
	}

	var newRoles *[]graphrbac.AppRole

	if d.IsNewResource() {
		newRoles, err = graph.AppRoleAdd(app.AppRoles, &role)
		if err != nil {
			if _, ok := err.(*graph.AlreadyExistsError); ok {
				return tf.ImportAsExistsDiag("azuread_application_app_role", id.String())
			}
			return tf.ErrorDiagF(err, "Failed to add App Role")
		}
	} else {
		if existing, _ := graph.AppRoleFindById(app, id.RoleId); existing == nil {
			return tf.ErrorDiagPathF(nil, "role_id", "App Role with ID %q was not found for Application %q", id.RoleId, id.ObjectId)
		}

		newRoles, err = graph.AppRoleUpdate(app.AppRoles, &role)
		if err != nil {
			return tf.ErrorDiagF(err, "Updating App Role with ID %q", *role.ID)
		}
	}

	properties := graphrbac.ApplicationUpdateParameters{
		AppRoles: newRoles,
	}
	if _, err := client.Patch(ctx, id.ObjectId, properties); err != nil {
		return tf.ErrorDiagF(err, "Updating Application with ID %q", id.ObjectId)
	}

	d.SetId(id.String())

	return applicationAppRoleResourceRead(ctx, d, meta)
}

func applicationAppRoleResourceRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*clients.AadClient).AadGraph.ApplicationsClient

	id, err := graph.ParseAppRoleId(d.Id())
	if err != nil {
		return tf.ErrorDiagPathF(err, "id", "Parsing App Role ID %q", d.Id())
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

	role, err := graph.AppRoleFindById(app, id.RoleId)
	if err != nil {
		return tf.ErrorDiagF(err, "Identifying App Role")
	}

	if role == nil {
		log.Printf("[DEBUG] App Role %q (ID %q) was not found - removing from state!", id.RoleId, id.ObjectId)
		d.SetId("")
		return nil
	}

	if dg := tf.Set(d, "application_object_id", id.ObjectId); dg != nil {
		return dg
	}

	if dg := tf.Set(d, "role_id", id.RoleId); dg != nil {
		return dg
	}

	if dg := tf.Set(d, "allowed_member_types", role.AllowedMemberTypes); dg != nil {
		return dg
	}

	if dg := tf.Set(d, "description", role.Description); dg != nil {
		return dg
	}

	if dg := tf.Set(d, "display_name", role.DisplayName); dg != nil {
		return dg
	}

	if dg := tf.Set(d, "is_enabled", role.IsEnabled); dg != nil {
		return dg
	}

	if dg := tf.Set(d, "value", role.Value); dg != nil {
		return dg
	}

	return nil
}

func applicationAppRoleResourceDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*clients.AadClient).AadGraph.ApplicationsClient

	id, err := graph.ParseAppRoleId(d.Id())
	if err != nil {
		return tf.ErrorDiagPathF(err, "id", "Parsing App Role ID %q", d.Id())
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

	log.Printf("[DEBUG] Disabling App Role %q for Application %q prior to removal", id.RoleId, id.ObjectId)
	newRoles, err := graph.AppRoleResultDisableById(app.AppRoles, id.RoleId)
	if err != nil {
		return tf.ErrorDiagF(err, "Disabling App Role with ID %q for application %q", id.RoleId, id.ObjectId)
	}

	properties := graphrbac.ApplicationUpdateParameters{
		AppRoles: newRoles,
	}
	if _, err := client.Patch(ctx, id.ObjectId, properties); err != nil {
		return tf.ErrorDiagF(err, "Updating Application with ID %q", id.ObjectId)
	}

	log.Printf("[DEBUG] Removing App Role %q for Application %q", id.RoleId, id.ObjectId)
	newRoles, err = graph.AppRoleResultRemoveById(app.AppRoles, id.RoleId)
	if err != nil {
		return tf.ErrorDiagF(err, "Removing App Role with ID %q for application %q", id.RoleId, id.ObjectId)
	}

	properties = graphrbac.ApplicationUpdateParameters{
		AppRoles: newRoles,
	}
	if _, err := client.Patch(ctx, id.ObjectId, properties); err != nil {
		return tf.ErrorDiagF(err, "Updating Application with ID %q", id.ObjectId)
	}

	return nil
}
