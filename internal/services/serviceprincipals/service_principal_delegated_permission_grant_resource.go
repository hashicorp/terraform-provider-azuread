package serviceprincipals

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/manicminer/hamilton/msgraph"
	"github.com/manicminer/hamilton/odata"

	"github.com/hashicorp/terraform-provider-azuread/internal/clients"
	"github.com/hashicorp/terraform-provider-azuread/internal/tf"
	"github.com/hashicorp/terraform-provider-azuread/internal/utils"
	"github.com/hashicorp/terraform-provider-azuread/internal/validate"
)

func servicePrincipalDelegatedPermissionGrantResource() *schema.Resource {
	return &schema.Resource{
		CreateContext: servicePrincipalDelegatedPermissionGrantResourceCreate,
		UpdateContext: servicePrincipalDelegatedPermissionGrantResourceUpdate,
		ReadContext:   servicePrincipalDelegatedPermissionGrantResourceRead,
		DeleteContext: servicePrincipalDelegatedPermissionGrantResourceDelete,

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(5 * time.Minute),
			Read:   schema.DefaultTimeout(5 * time.Minute),
			Update: schema.DefaultTimeout(5 * time.Minute),
			Delete: schema.DefaultTimeout(5 * time.Minute),
		},

		Importer: tf.ValidateResourceIDPriorToImport(func(id string) error {
			if len(id) == 0 {
				return fmt.Errorf("specified ID is not valid: %q", id)
			}
			return nil
		}),

		Schema: map[string]*schema.Schema{
			"claim_values": {
				Description: "A set of claim values for delegated permission scopes which should be included in access tokens for the resource",
				Type:        schema.TypeSet,
				Required:    true,
				MinItems:    1,
				Elem: &schema.Schema{
					Type:             schema.TypeString,
					ValidateDiagFunc: validate.NoEmptyStrings,
				},
			},

			"resource_service_principal_object_id": {
				Description:      "The object ID of the service principal representing the resource to be accessed",
				Type:             schema.TypeString,
				Required:         true,
				ForceNew:         true,
				ValidateDiagFunc: validate.UUID,
			},

			"service_principal_object_id": {
				Description:      "The object ID of the service principal for which this delegated permission grant should be created",
				Type:             schema.TypeString,
				Required:         true,
				ForceNew:         true,
				ValidateDiagFunc: validate.UUID,
			},

			"user_object_id": {
				Description:      "The object ID of the user on behalf of whom the service principal is authorized to access the resource",
				Type:             schema.TypeString,
				Optional:         true,
				ForceNew:         true,
				ValidateDiagFunc: validate.UUID,
			},
		},
	}
}

func servicePrincipalDelegatedPermissionGrantResourceCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*clients.Client).ServicePrincipals.DelegatedPermissionGrantsClient
	servicePrincipalsClient := meta.(*clients.Client).ServicePrincipals.ServicePrincipalsClient

	servicePrincipalId := d.Get("service_principal_object_id").(string)
	resourceId := d.Get("resource_service_principal_object_id").(string)

	if _, status, err := servicePrincipalsClient.Get(ctx, servicePrincipalId, odata.Query{}); err != nil {
		if status == http.StatusNotFound {
			return tf.ErrorDiagPathF(err, "principal_object_id", "Service principal with object ID %q was not found)", servicePrincipalId)
		}
		return tf.ErrorDiagF(err, "Could not retrieve service principal with object ID %q", servicePrincipalId)
	}

	if _, status, err := servicePrincipalsClient.Get(ctx, resourceId, odata.Query{}); err != nil {
		if status == http.StatusNotFound {
			return tf.ErrorDiagPathF(err, "principal_object_id", "Service principal not found for resource (Object ID: %q)", resourceId)
		}
		return tf.ErrorDiagF(err, "Could not retrieve service principal for resource (Object ID: %q)", resourceId)
	}

	properties := msgraph.DelegatedPermissionGrant{
		ClientId:   utils.String(servicePrincipalId),
		ResourceId: utils.String(resourceId),
		Scopes:     tf.ExpandStringSlicePtr(d.Get("claim_values").(*schema.Set).List()),
	}

	if v, ok := d.GetOk("user_object_id"); ok && v.(string) != "" {
		properties.PrincipalId = utils.String(v.(string))
		properties.ConsentType = utils.String(msgraph.DelegatedPermissionGrantConsentTypePrincipal)
	} else {
		properties.ConsentType = utils.String(msgraph.DelegatedPermissionGrantConsentTypeAllPrincipals)
	}

	delegatedPermissionGrant, _, err := client.Create(ctx, properties)
	if err != nil {
		return tf.ErrorDiagF(err, "Could not create delegated permission grant")
	}

	if delegatedPermissionGrant.Id == nil || *delegatedPermissionGrant.Id == "" {
		return tf.ErrorDiagF(errors.New("ID returned for delegated permission grant is nil"), "Bad API response")
	}

	d.SetId(*delegatedPermissionGrant.Id)

	return servicePrincipalDelegatedPermissionGrantResourceRead(ctx, d, meta)
}

func servicePrincipalDelegatedPermissionGrantResourceUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*clients.Client).ServicePrincipals.DelegatedPermissionGrantsClient

	properties := msgraph.DelegatedPermissionGrant{
		Id:     utils.String(d.Id()),
		Scopes: tf.ExpandStringSlicePtr(d.Get("claim_values").(*schema.Set).List()),
	}

	if _, err := client.Update(ctx, properties); err != nil {
		return tf.ErrorDiagF(err, "Could not update delegated permission grant")
	}

	return servicePrincipalDelegatedPermissionGrantResourceRead(ctx, d, meta)
}

func servicePrincipalDelegatedPermissionGrantResourceRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*clients.Client).ServicePrincipals.DelegatedPermissionGrantsClient

	delegatedPermissionGrant, status, err := client.Get(ctx, d.Id(), odata.Query{})
	if err != nil {
		if status == http.StatusNoContent {
			log.Printf("[DEBUG] Delegated Permission Grant with ID %q was not found - removing from state", d.Id())
			d.SetId("")
			return nil
		}
		return tf.ErrorDiagPathF(err, "id", "Retrieving Delegated Permission Grant with ID %q", d.Id())
	}

	tf.Set(d, "claim_values", delegatedPermissionGrant.Scopes)
	tf.Set(d, "resource_service_principal_object_id", delegatedPermissionGrant.ResourceId)
	tf.Set(d, "service_principal_object_id", delegatedPermissionGrant.ClientId)
	tf.Set(d, "user_object_id", delegatedPermissionGrant.PrincipalId)

	return nil
}

func servicePrincipalDelegatedPermissionGrantResourceDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*clients.Client).ServicePrincipals.DelegatedPermissionGrantsClient

	id := d.Id()

	if status, err := client.Delete(ctx, id); err != nil {
		return tf.ErrorDiagPathF(err, "id", "Deleting delegated permission grant with ID %q, got status %d", id, status)
	}

	return nil
}
