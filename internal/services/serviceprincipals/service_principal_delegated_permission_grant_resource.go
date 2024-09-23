// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package serviceprincipals

import (
	"context"
	"errors"
	"fmt"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/pointer"
	"github.com/hashicorp/go-azure-helpers/lang/response"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/oauth2permissiongrants/stable/oauth2permissiongrant"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/serviceprincipals/stable/serviceprincipal"
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
	"github.com/hashicorp/terraform-provider-azuread/internal/clients"
	"github.com/hashicorp/terraform-provider-azuread/internal/helpers/tf"
	"github.com/hashicorp/terraform-provider-azuread/internal/helpers/tf/pluginsdk"
	"github.com/hashicorp/terraform-provider-azuread/internal/helpers/tf/validation"
)

func servicePrincipalDelegatedPermissionGrantResource() *pluginsdk.Resource {
	return &pluginsdk.Resource{
		CreateContext: servicePrincipalDelegatedPermissionGrantResourceCreate,
		UpdateContext: servicePrincipalDelegatedPermissionGrantResourceUpdate,
		ReadContext:   servicePrincipalDelegatedPermissionGrantResourceRead,
		DeleteContext: servicePrincipalDelegatedPermissionGrantResourceDelete,

		Timeouts: &pluginsdk.ResourceTimeout{
			Create: pluginsdk.DefaultTimeout(5 * time.Minute),
			Read:   pluginsdk.DefaultTimeout(5 * time.Minute),
			Update: pluginsdk.DefaultTimeout(5 * time.Minute),
			Delete: pluginsdk.DefaultTimeout(5 * time.Minute),
		},

		Importer: pluginsdk.ImporterValidatingResourceId(func(id string) error {
			if len(id) == 0 {
				return fmt.Errorf("specified ID is not valid: %q", id)
			}
			return nil
		}),

		Schema: map[string]*pluginsdk.Schema{
			"claim_values": {
				Description: "A set of claim values for delegated permission scopes which should be included in access tokens for the resource",
				Type:        pluginsdk.TypeSet,
				Required:    true,
				MinItems:    1,
				Elem: &pluginsdk.Schema{
					Type:         pluginsdk.TypeString,
					ValidateFunc: validation.StringIsNotEmpty,
				},
			},

			"resource_service_principal_object_id": {
				Description:  "The object ID of the service principal representing the resource to be accessed",
				Type:         pluginsdk.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validation.IsUUID,
			},

			"service_principal_object_id": {
				Description:  "The object ID of the service principal for which this delegated permission grant should be created",
				Type:         pluginsdk.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validation.IsUUID,
			},

			"user_object_id": {
				Description:  "The object ID of the user on behalf of whom the service principal is authorized to access the resource",
				Type:         pluginsdk.TypeString,
				Optional:     true,
				ForceNew:     true,
				ValidateFunc: validation.IsUUID,
			},
		},
	}
}

func servicePrincipalDelegatedPermissionGrantResourceCreate(ctx context.Context, d *pluginsdk.ResourceData, meta interface{}) pluginsdk.Diagnostics {
	client := meta.(*clients.Client).ServicePrincipals.OAuth2PermissionGrantClient
	servicePrincipalClient := meta.(*clients.Client).ServicePrincipals.ServicePrincipalClient

	servicePrincipalId := stable.NewServicePrincipalID(d.Get("service_principal_object_id").(string))
	resourcePrincipalId := stable.NewServicePrincipalID(d.Get("resource_service_principal_object_id").(string))

	if resp, err := servicePrincipalClient.GetServicePrincipal(ctx, servicePrincipalId, serviceprincipal.DefaultGetServicePrincipalOperationOptions()); err != nil {
		if response.WasNotFound(resp.HttpResponse) {
			return tf.ErrorDiagPathF(err, "principal_object_id", "%s was not found)", servicePrincipalId)
		}
		return tf.ErrorDiagF(err, "Could not retrieve %s", servicePrincipalId)
	}

	if resp, err := servicePrincipalClient.GetServicePrincipal(ctx, resourcePrincipalId, serviceprincipal.DefaultGetServicePrincipalOperationOptions()); err != nil {
		if response.WasNotFound(resp.HttpResponse) {
			return tf.ErrorDiagPathF(err, "principal_object_id", "%s not found for resource", resourcePrincipalId)
		}
		return tf.ErrorDiagF(err, "Could not retrieve %s for resource", resourcePrincipalId)
	}

	properties := stable.OAuth2PermissionGrant{
		ClientId:   servicePrincipalId.ServicePrincipalId,
		ResourceId: pointer.To(resourcePrincipalId.ServicePrincipalId),
		Scope:      nullable.NoZero(strings.Join(tf.ExpandStringSlice(d.Get("claim_values").(*pluginsdk.Set).List()), " ")),
	}

	if v, ok := d.GetOk("user_object_id"); ok && v.(string) != "" {
		properties.PrincipalId = nullable.NoZero(v.(string))
		properties.ConsentType = nullable.Value(DelegatedPermissionGrantConsentTypePrincipal)
	} else {
		properties.ConsentType = nullable.Value(DelegatedPermissionGrantConsentTypeAllPrincipals)
	}

	options := oauth2permissiongrant.CreateOAuth2PermissionGrantOperationOptions{
		RetryFunc: func(resp *http.Response, o *odata.OData) (bool, error) {
			if response.WasNotFound(resp) {
				return true, nil
			} else if response.WasBadRequest(resp) && o != nil && o.Error != nil {
				return o.Error.Match("does not exist or one of its queried reference-property objects are not present"), nil
			}
			return false, nil
		},
	}

	resp, err := client.CreateOAuth2PermissionGrant(ctx, properties, options)
	if err != nil {
		return tf.ErrorDiagF(err, "Could not create delegated permission grant")
	}

	delegatedPermissionGrant := resp.Model
	if delegatedPermissionGrant == nil {
		return tf.ErrorDiagF(errors.New("model was nil"), "Could not create delegated permission grant")
	}
	if delegatedPermissionGrant.Id == nil || *delegatedPermissionGrant.Id == "" {
		return tf.ErrorDiagF(errors.New("ID returned for delegated permission grant is nil"), "Bad API response")
	}

	id := stable.NewOAuth2PermissionGrantID(*delegatedPermissionGrant.Id)
	d.SetId(id.OAuth2PermissionGrantId)

	return servicePrincipalDelegatedPermissionGrantResourceRead(ctx, d, meta)
}

func servicePrincipalDelegatedPermissionGrantResourceUpdate(ctx context.Context, d *pluginsdk.ResourceData, meta interface{}) pluginsdk.Diagnostics {
	client := meta.(*clients.Client).ServicePrincipals.OAuth2PermissionGrantClient
	id := stable.NewOAuth2PermissionGrantID(d.Id())

	properties := stable.OAuth2PermissionGrant{
		Scope: nullable.NoZero(strings.Join(tf.ExpandStringSlice(d.Get("claim_values").(*pluginsdk.Set).List()), " ")),
	}

	if _, err := client.UpdateOAuth2PermissionGrant(ctx, id, properties, oauth2permissiongrant.DefaultUpdateOAuth2PermissionGrantOperationOptions()); err != nil {
		return tf.ErrorDiagF(err, "Updating %s", id)
	}

	return servicePrincipalDelegatedPermissionGrantResourceRead(ctx, d, meta)
}

func servicePrincipalDelegatedPermissionGrantResourceRead(ctx context.Context, d *pluginsdk.ResourceData, meta interface{}) pluginsdk.Diagnostics {
	client := meta.(*clients.Client).ServicePrincipals.OAuth2PermissionGrantClient
	id := stable.NewOAuth2PermissionGrantID(d.Id())

	resp, err := client.GetOAuth2PermissionGrant(ctx, id, oauth2permissiongrant.DefaultGetOAuth2PermissionGrantOperationOptions())
	if err != nil {
		if response.WasNotFound(resp.HttpResponse) {
			log.Printf("[DEBUG] %s was not found - removing from state", id)
			d.SetId("")
			return nil
		}
		return tf.ErrorDiagPathF(err, "id", "Retrieving Delegated Permission Grant with ID %q", d.Id())
	}

	delegatedPermissionGrant := resp.Model
	if delegatedPermissionGrant == nil {
		return tf.ErrorDiagF(errors.New("model was nil"), "Retrieving %s", id)
	}

	tf.Set(d, "claim_values", tf.FromSpaceSeparated(delegatedPermissionGrant.Scope.GetOrZero()))
	tf.Set(d, "resource_service_principal_object_id", pointer.From(delegatedPermissionGrant.ResourceId))
	tf.Set(d, "service_principal_object_id", delegatedPermissionGrant.ClientId)
	tf.Set(d, "user_object_id", delegatedPermissionGrant.PrincipalId.GetOrZero())

	return nil
}

func servicePrincipalDelegatedPermissionGrantResourceDelete(ctx context.Context, d *pluginsdk.ResourceData, meta interface{}) pluginsdk.Diagnostics {
	client := meta.(*clients.Client).ServicePrincipals.OAuth2PermissionGrantClient
	id := stable.NewOAuth2PermissionGrantID(d.Id())

	if _, err := client.DeleteOAuth2PermissionGrant(ctx, id, oauth2permissiongrant.DefaultDeleteOAuth2PermissionGrantOperationOptions()); err != nil {
		return tf.ErrorDiagPathF(err, "id", "Deleting %s", id)
	}

	return nil
}
