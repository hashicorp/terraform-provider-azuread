package aadgraph

import (
	"context"
	"fmt"
	"github.com/hashicorp/go-cty/cty"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"log"

	"github.com/Azure/azure-sdk-for-go/services/graphrbac/1.6/graphrbac"
	"github.com/hashicorp/go-uuid"
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
			return diag.Diagnostics{diag.Diagnostic{
				Severity: diag.Error,
				Summary:  fmt.Sprintf("Generating App Role for application with object ID %q", objectId),
				Detail:   err.Error(),
			}}
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
			return diag.Diagnostics{diag.Diagnostic{
				Severity:      diag.Error,
				Summary:       fmt.Sprintf("Application with object ID %q was not found", id.ObjectId),
				AttributePath: cty.Path{cty.GetAttrStep{Name: "application_object_id"}},
			}}
		}
		return diag.Diagnostics{diag.Diagnostic{
			Severity:      diag.Error,
			Summary:       fmt.Sprintf("retrieving Application with object ID %q", id.ObjectId),
			Detail:        err.Error(),
			AttributePath: cty.Path{cty.GetAttrStep{Name: "application_object_id"}},
		}}
	}

	var newRoles *[]graphrbac.AppRole

	if d.IsNewResource() {
		newRoles, err = graph.AppRoleAdd(app.AppRoles, &role)
		if err != nil {
			if _, ok := err.(*graph.AlreadyExistsError); ok {
				return tf.ImportAsExistsDiag("azuread_application_app_role", id.String())
			}
			return diag.Diagnostics{diag.Diagnostic{
				Severity: diag.Error,
				Summary:  fmt.Sprintf("Failed to add App Role"),
				Detail:   err.Error(),
			}}
		}
	} else {
		if existing, _ := graph.AppRoleFindById(app, id.RoleId); existing == nil {
			return diag.Diagnostics{diag.Diagnostic{
				Severity:      diag.Error,
				Summary:       fmt.Sprintf("App Role with ID %q was not found for Application %q", id.RoleId, id.ObjectId),
				AttributePath: cty.Path{cty.GetAttrStep{Name: "role_id"}},
			}}
		}

		newRoles, err = graph.AppRoleUpdate(app.AppRoles, &role)
		if err != nil {
			return diag.Diagnostics{diag.Diagnostic{
				Severity: diag.Error,
				Summary:  fmt.Sprintf("Updating App Role with ID %q", *role.ID),
				Detail:   err.Error(),
			}}
		}
	}

	properties := graphrbac.ApplicationUpdateParameters{
		AppRoles: newRoles,
	}
	if _, err := client.Patch(ctx, id.ObjectId, properties); err != nil {
		return diag.Diagnostics{diag.Diagnostic{
			Severity: diag.Error,
			Summary:  fmt.Sprintf("Updating Application with ID %q", id.ObjectId),
			Detail:   err.Error(),
		}}
	}

	d.SetId(id.String())

	return applicationAppRoleResourceRead(ctx, d, meta)
}

func applicationAppRoleResourceRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*clients.AadClient).AadGraph.ApplicationsClient

	id, err := graph.ParseAppRoleId(d.Id())
	if err != nil {
		return diag.Diagnostics{diag.Diagnostic{
			Severity:      diag.Error,
			Summary:       fmt.Sprintf("Parsing App Role ID %q", d.Id()),
			Detail:        err.Error(),
			AttributePath: cty.Path{cty.GetAttrStep{Name: "id"}},
		}}
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
		return diag.Diagnostics{diag.Diagnostic{
			Severity:      diag.Error,
			Summary:       fmt.Sprintf("Retrieving Application with ID %q", id.ObjectId),
			Detail:        err.Error(),
			AttributePath: cty.Path{cty.GetAttrStep{Name: "application_object_id"}},
		}}
	}

	role, err := graph.AppRoleFindById(app, id.RoleId)
	if err != nil {
		return diag.Diagnostics{diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Identifying App Role",
			Detail:   err.Error(),
		}}
	}

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

func applicationAppRoleResourceDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*clients.AadClient).AadGraph.ApplicationsClient

	id, err := graph.ParseAppRoleId(d.Id())
	if err != nil {
		return diag.Diagnostics{diag.Diagnostic{
			Severity:      diag.Error,
			Summary:       fmt.Sprintf("Parsing App Role ID %q", d.Id()),
			Detail:        err.Error(),
			AttributePath: cty.Path{cty.GetAttrStep{Name: "id"}},
		}}
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
		return diag.Diagnostics{diag.Diagnostic{
			Severity:      diag.Error,
			Summary:       fmt.Sprintf("Retrieving Application with ID %q", id.ObjectId),
			Detail:        err.Error(),
			AttributePath: cty.Path{cty.GetAttrStep{Name: "application_object_id"}},
		}}
	}

	log.Printf("[DEBUG] Disabling App Role %q for Application %q prior to removal", id.RoleId, id.ObjectId)
	newRoles, err := graph.AppRoleResultDisableById(app.AppRoles, id.RoleId)
	if err != nil {
		return diag.Diagnostics{diag.Diagnostic{
			Severity: diag.Error,
			Summary:  fmt.Sprintf("Disabling App Role with ID %q for application %q", id.RoleId, id.ObjectId),
			Detail:   err.Error(),
		}}
	}

	properties := graphrbac.ApplicationUpdateParameters{
		AppRoles: newRoles,
	}
	if _, err := client.Patch(ctx, id.ObjectId, properties); err != nil {
		return diag.Diagnostics{diag.Diagnostic{
			Severity: diag.Error,
			Summary:  fmt.Sprintf("Updating Application with ID %q", id.ObjectId),
			Detail:   err.Error(),
		}}
	}

	log.Printf("[DEBUG] Removing App Role %q for Application %q", id.RoleId, id.ObjectId)
	newRoles, err = graph.AppRoleResultRemoveById(app.AppRoles, id.RoleId)
	if err != nil {
		return diag.Diagnostics{diag.Diagnostic{
			Severity: diag.Error,
			Summary:  fmt.Sprintf("Removing App Role with ID %q for application %q", id.RoleId, id.ObjectId),
			Detail:   err.Error(),
		}}
	}

	properties = graphrbac.ApplicationUpdateParameters{
		AppRoles: newRoles,
	}
	if _, err := client.Patch(ctx, id.ObjectId, properties); err != nil {
		return diag.Diagnostics{diag.Diagnostic{
			Severity: diag.Error,
			Summary:  fmt.Sprintf("Updating Application with ID %q", id.ObjectId),
			Detail:   err.Error(),
		}}
	}

	return nil
}
