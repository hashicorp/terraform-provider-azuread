package aadgraph

import (
	"context"
	"fmt"
	"github.com/hashicorp/go-cty/cty"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"log"

	"github.com/Azure/azure-sdk-for-go/services/graphrbac/1.6/graphrbac"
	"github.com/hashicorp/go-azure-helpers/response"
	"github.com/hashicorp/go-uuid"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/terraform-providers/terraform-provider-azuread/internal/clients"
	"github.com/terraform-providers/terraform-provider-azuread/internal/services/aadgraph/graph"
	"github.com/terraform-providers/terraform-provider-azuread/internal/tf"
	"github.com/terraform-providers/terraform-provider-azuread/internal/utils"
	"github.com/terraform-providers/terraform-provider-azuread/internal/validate"
)

const servicePrincipalResourceName = "azuread_service_principal"

func servicePrincipalResource() *schema.Resource {
	return &schema.Resource{
		CreateContext: servicePrincipalResourceCreate,
		ReadContext:   servicePrincipalResourceRead,
		UpdateContext: servicePrincipalResourceUpdate,
		DeleteContext: servicePrincipalResourceDelete,

		Importer: tf.ValidateResourceIDPriorToImport(func(id string) error {
			if _, err := uuid.ParseUUID(id); err != nil {
				return fmt.Errorf("specified ID (%q) is not valid: %s", id, err)
			}
			return nil
		}),

		Schema: map[string]*schema.Schema{
			"application_id": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validate.UUID,
			},

			"app_role_assignment_required": {
				Type:     schema.TypeBool,
				Optional: true,
			},

			"display_name": {
				Type:     schema.TypeString,
				Computed: true,
			},

			"object_id": {
				Type:     schema.TypeString,
				Computed: true,
			},

			"oauth2_permissions": graph.SchemaOauth2PermissionsComputed(),

			"tags": {
				Type:     schema.TypeSet,
				Optional: true,
				Set:      schema.HashString,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
		},
	}
}

func servicePrincipalResourceCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*clients.AadClient).AadGraph.ServicePrincipalsClient

	applicationId := d.Get("application_id").(string)

	properties := graphrbac.ServicePrincipalCreateParameters{
		AppID: utils.String(applicationId),
		// there's no way of retrieving this, and there's no way of changing it
		// given there's no way to change it - we'll just default this to true
		AccountEnabled: utils.Bool(true),
	}

	if v, ok := d.GetOk("app_role_assignment_required"); ok {
		properties.AppRoleAssignmentRequired = utils.Bool(v.(bool))
	}

	if v, ok := d.GetOk("tags"); ok {
		properties.Tags = tf.ExpandStringSlicePtr(v.(*schema.Set).List())
	}

	sp, err := client.Create(ctx, properties)
	if err != nil {
		return diag.Diagnostics{diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Could not create service principal",
			Detail:   err.Error(),
		}}
	}
	if sp.ObjectID == nil || *sp.ObjectID == "" {
		return diag.Diagnostics{diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Bad API response",
			Detail:   "ObjectID returned for service principal is nil",
		}}
	}
	d.SetId(*sp.ObjectID)

	_, err = graph.WaitForCreationReplication(d.Timeout(schema.TimeoutCreate), func() (interface{}, error) {
		return client.Get(ctx, *sp.ObjectID)
	})
	if err != nil {
		return diag.Diagnostics{diag.Diagnostic{
			Severity: diag.Error,
			Summary:  fmt.Sprintf("Waiting for service principal with object ID: %q", *sp.ObjectID),
			Detail:   err.Error(),
		}}
	}

	return servicePrincipalResourceRead(ctx, d, meta)
}

func servicePrincipalResourceUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*clients.AadClient).AadGraph.ServicePrincipalsClient

	var properties graphrbac.ServicePrincipalUpdateParameters

	if d.HasChange("app_role_assignment_required") {
		properties.AppRoleAssignmentRequired = utils.Bool(d.Get("app_role_assignment_required").(bool))
	}

	if d.HasChange("tags") {
		if v, ok := d.GetOk("tags"); ok {
			properties.Tags = tf.ExpandStringSlicePtr(v.(*schema.Set).List())
		} else {
			empty := []string{} // clear tags with empty array
			properties.Tags = &empty
		}
	}

	if _, err := client.Update(ctx, d.Id(), properties); err != nil {
		return diag.Diagnostics{diag.Diagnostic{
			Severity: diag.Error,
			Summary:  fmt.Sprintf("Updating service principal with object ID: %q", d.Id()),
			Detail:   err.Error(),
		}}
	}

	return servicePrincipalResourceRead(ctx, d, meta)
}

func servicePrincipalResourceRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*clients.AadClient).AadGraph.ServicePrincipalsClient

	objectId := d.Id()

	app, err := client.Get(ctx, objectId)
	if err != nil {
		if utils.ResponseWasNotFound(app.Response) {
			log.Printf("[DEBUG] Service Principal with Object ID %q was not found - removing from state!", objectId)
			d.SetId("")
			return nil
		}

		return diag.Diagnostics{diag.Diagnostic{
			Severity: diag.Error,
			Summary:  fmt.Sprintf("retrieving service principal with object ID: %q", d.Id()),
			Detail:   err.Error(),
		}}
	}

	d.Set("application_id", app.AppID)
	d.Set("display_name", app.DisplayName)
	d.Set("object_id", app.ObjectID)
	d.Set("app_role_assignment_required", app.AppRoleAssignmentRequired)

	// tags doesn't exist as a property, so extract it
	if err := d.Set("tags", app.Tags); err != nil {
		return diag.Diagnostics{diag.Diagnostic{
			Severity:      diag.Error,
			Summary:       "Could not set attribute",
			Detail:        err.Error(),
			AttributePath: cty.Path{cty.GetAttrStep{Name: "tags"}},
		}}
	}

	if err := d.Set("oauth2_permissions", graph.FlattenOauth2Permissions(app.Oauth2Permissions)); err != nil {
		return diag.Diagnostics{diag.Diagnostic{
			Severity:      diag.Error,
			Summary:       "Could not set attribute",
			Detail:        err.Error(),
			AttributePath: cty.Path{cty.GetAttrStep{Name: "oauth2_permissions"}},
		}}
	}

	return nil
}

func servicePrincipalResourceDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*clients.AadClient).AadGraph.ServicePrincipalsClient

	applicationId := d.Id()
	app, err := client.Delete(ctx, applicationId)
	if err != nil {
		if !response.WasNotFound(app.Response) {
			return diag.Diagnostics{diag.Diagnostic{
				Severity: diag.Error,
				Summary:  fmt.Sprintf("Deleting service principal with object ID: %q", d.Id()),
				Detail:   err.Error(),
			}}
		}
	}

	return nil
}
