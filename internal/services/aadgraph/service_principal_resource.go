package aadgraph

import (
	"fmt"
	"log"

	"github.com/Azure/azure-sdk-for-go/services/graphrbac/1.6/graphrbac"
	"github.com/hashicorp/go-azure-helpers/response"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"github.com/terraform-providers/terraform-provider-azuread/internal/clients"
	"github.com/terraform-providers/terraform-provider-azuread/internal/services/aadgraph/graph"
	"github.com/terraform-providers/terraform-provider-azuread/internal/tf"
	"github.com/terraform-providers/terraform-provider-azuread/internal/utils"
	"github.com/terraform-providers/terraform-provider-azuread/internal/validate"
)

const servicePrincipalResourceName = "azuread_service_principal"

func ServicePrincipalResource() *schema.Resource {
	return &schema.Resource{
		Create: servicePrincipalResourceCreate,
		Read:   servicePrincipalResourceRead,
		Update: servicePrincipalResourceUpdate,
		Delete: servicePrincipalResourceDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

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

func servicePrincipalResourceCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*clients.AadClient).AadGraph.ServicePrincipalsClient
	ctx := meta.(*clients.AadClient).StopContext

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
		return fmt.Errorf("creating Service Principal for application  %q: %+v", applicationId, err)
	}
	if sp.ObjectID == nil {
		return fmt.Errorf("Service Principal	objectID is nil")
	}
	d.SetId(*sp.ObjectID)

	_, err = graph.WaitForCreationReplication(d.Timeout(schema.TimeoutCreate), func() (interface{}, error) {
		return client.Get(ctx, *sp.ObjectID)
	})
	if err != nil {
		return fmt.Errorf("waiting for Service Principal with ObjectId %q: %+v", *sp.ObjectID, err)
	}

	return servicePrincipalResourceRead(d, meta)
}

func servicePrincipalResourceUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*clients.AadClient).AadGraph.ServicePrincipalsClient
	ctx := meta.(*clients.AadClient).StopContext

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
		return fmt.Errorf("patching Service Principal with ID %q: %+v", d.Id(), err)
	}

	return servicePrincipalResourceRead(d, meta)
}

func servicePrincipalResourceRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*clients.AadClient).AadGraph.ServicePrincipalsClient
	ctx := meta.(*clients.AadClient).StopContext

	objectId := d.Id()

	app, err := client.Get(ctx, objectId)
	if err != nil {
		if utils.ResponseWasNotFound(app.Response) {
			log.Printf("[DEBUG] Service Principal with Object ID %q was not found - removing from state!", objectId)
			d.SetId("")
			return nil
		}
		return fmt.Errorf("retrieving Service Principal ID %q: %+v", objectId, err)
	}

	d.Set("application_id", app.AppID)
	d.Set("display_name", app.DisplayName)
	d.Set("object_id", app.ObjectID)
	d.Set("app_role_assignment_required", app.AppRoleAssignmentRequired)

	// tags doesn't exist as a property, so extract it
	if err := d.Set("tags", app.Tags); err != nil {
		return fmt.Errorf("setting `tags`: %+v", err)
	}

	if err := d.Set("oauth2_permissions", graph.FlattenOauth2Permissions(app.Oauth2Permissions)); err != nil {
		return fmt.Errorf("setting `oauth2_permissions`: %+v", err)
	}

	return nil
}

func servicePrincipalResourceDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*clients.AadClient).AadGraph.ServicePrincipalsClient
	ctx := meta.(*clients.AadClient).StopContext

	applicationId := d.Id()
	app, err := client.Delete(ctx, applicationId)
	if err != nil {
		if !response.WasNotFound(app.Response) {
			return fmt.Errorf("deleting Service Principal ID %q: %+v", applicationId, err)
		}
	}

	return nil
}
