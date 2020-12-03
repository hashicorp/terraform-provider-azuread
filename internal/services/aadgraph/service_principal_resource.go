package aadgraph

import (
	"context"
	"fmt"
	"log"

	"github.com/Azure/azure-sdk-for-go/services/graphrbac/1.6/graphrbac"
	"github.com/hashicorp/go-azure-helpers/response"
	"github.com/hashicorp/go-uuid"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
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
				Type:             schema.TypeString,
				Required:         true,
				ForceNew:         true,
				ValidateDiagFunc: validate.UUID,
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

			"app_roles": graph.SchemaAppRolesComputed(),

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
		return tf.ErrorDiag("Could not create service principal", err.Error(), "")
	}
	if sp.ObjectID == nil || *sp.ObjectID == "" {
		return tf.ErrorDiag("Bad API response", "ObjectID returned for service principal is nil", "")
	}
	d.SetId(*sp.ObjectID)

	_, err = graph.WaitForCreationReplication(ctx, d.Timeout(schema.TimeoutCreate), func() (interface{}, error) {
		return client.Get(ctx, *sp.ObjectID)
	})
	if err != nil {
		return tf.ErrorDiag(fmt.Sprintf("Waiting for service principal with object ID: %q", *sp.ObjectID), err.Error(), "")
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
		return tf.ErrorDiag(fmt.Sprintf("Updating service principal with object ID: %q", d.Id()), err.Error(), "")
	}

	return servicePrincipalResourceRead(ctx, d, meta)
}

func servicePrincipalResourceRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*clients.AadClient).AadGraph.ServicePrincipalsClient

	objectId := d.Id()

	sp, err := client.Get(ctx, objectId)
	if err != nil {
		if utils.ResponseWasNotFound(sp.Response) {
			log.Printf("[DEBUG] Service Principal with Object ID %q was not found - removing from state!", objectId)
			d.SetId("")
			return nil
		}

		return tf.ErrorDiag(fmt.Sprintf("retrieving service principal with object ID: %q", d.Id()), err.Error(), "")
	}

	if err := d.Set("object_id", sp.ObjectID); err != nil {
		return tf.ErrorDiag("Could not set attribute", err.Error(), "object_id")
	}

	if err := d.Set("application_id", sp.AppID); err != nil {
		return tf.ErrorDiag("Could not set attribute", err.Error(), "application_id")
	}

	if err := d.Set("display_name", sp.DisplayName); err != nil {
		return tf.ErrorDiag("Could not set attribute", err.Error(), "display_name")
	}

	if err := d.Set("app_role_assignment_required", sp.AppRoleAssignmentRequired); err != nil {
		return tf.ErrorDiag("Could not set attribute", err.Error(), "app_role_assignment_required")
	}

	if err := d.Set("tags", sp.Tags); err != nil {
		return tf.ErrorDiag("Could not set attribute", err.Error(), "tags")
	}

	if err := d.Set("app_roles", graph.FlattenAppRoles(sp.AppRoles)); err != nil {
		return tf.ErrorDiag("Could not set attribute", err.Error(), "app_roles")
	}

	if err := d.Set("oauth2_permissions", graph.FlattenOauth2Permissions(sp.Oauth2Permissions)); err != nil {
		return tf.ErrorDiag("Could not set attribute", err.Error(), "oauth2_permissions")
	}

	return nil
}

func servicePrincipalResourceDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*clients.AadClient).AadGraph.ServicePrincipalsClient

	applicationId := d.Id()
	app, err := client.Delete(ctx, applicationId)
	if err != nil {
		if !response.WasNotFound(app.Response) {
			return tf.ErrorDiag(fmt.Sprintf("Deleting service principal with object ID: %q", d.Id()), err.Error(), "")
		}
	}

	return nil
}
