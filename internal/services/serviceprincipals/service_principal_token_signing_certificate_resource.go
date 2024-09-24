// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package serviceprincipals

import (
	"context"
	"errors"
	"fmt"
	"log"
	"regexp"
	"strings"
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/pointer"
	"github.com/hashicorp/go-azure-helpers/lang/response"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/serviceprincipals/stable/serviceprincipal"
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
	"github.com/hashicorp/terraform-provider-azuread/internal/clients"
	"github.com/hashicorp/terraform-provider-azuread/internal/helpers/consistency"
	"github.com/hashicorp/terraform-provider-azuread/internal/helpers/credentials"
	"github.com/hashicorp/terraform-provider-azuread/internal/helpers/tf"
	"github.com/hashicorp/terraform-provider-azuread/internal/helpers/tf/pluginsdk"
	"github.com/hashicorp/terraform-provider-azuread/internal/helpers/tf/validation"
	"github.com/hashicorp/terraform-provider-azuread/internal/services/serviceprincipals/parse"
)

func servicePrincipalTokenSigningCertificateResource() *pluginsdk.Resource {
	return &pluginsdk.Resource{
		CreateContext: servicePrincipalTokenSigningCertificateResourceCreate,
		ReadContext:   servicePrincipalTokenSigningCertificateResourceRead,
		DeleteContext: servicePrincipalTokenSigningCertificateResourceDelete,

		Timeouts: &pluginsdk.ResourceTimeout{
			Create: pluginsdk.DefaultTimeout(5 * time.Minute),
			Read:   pluginsdk.DefaultTimeout(5 * time.Minute),
			Delete: pluginsdk.DefaultTimeout(5 * time.Minute),
		},

		Importer: pluginsdk.ImporterValidatingResourceId(func(id string) error {
			_, err := parse.SigningCertificateID(id)
			return err
		}),

		Schema: map[string]*pluginsdk.Schema{
			"service_principal_id": {
				Description:  "The object ID of the service principal for which this certificate should be created",
				Type:         pluginsdk.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validation.IsUUID,
			},

			"display_name": {
				Description:  "A friendly name for the certificate",
				Type:         pluginsdk.TypeString,
				Optional:     true,
				Computed:     true,
				ForceNew:     true,
				ValidateFunc: validation.StringMatch(regexp.MustCompile("^CN=.+$|^$"), ""),
			},

			"end_date": {
				Description:  "The end date until which the certificate is valid, formatted as an RFC3339 date string (e.g. `2018-01-01T01:02:03Z`). Default is 3 years from current date.",
				Type:         pluginsdk.TypeString,
				Optional:     true,
				Computed:     true,
				ForceNew:     true,
				ValidateFunc: validation.IsRFC3339Time,
			},

			"key_id": {
				Description: "A UUID used to uniquely identify the verify certificate.",
				Type:        pluginsdk.TypeString,
				Computed:    true,
			},

			"thumbprint": {
				Description: "The thumbprint of the certificate.",
				Type:        pluginsdk.TypeString,
				Computed:    true,
			},

			"start_date": {
				Description: "The start date from which the certificate is valid, formatted as an RFC3339 date string (e.g. `2018-01-01T01:02:03Z`).",
				Type:        pluginsdk.TypeString,
				Computed:    true,
			},

			"value": {
				Description: "The certificate data, which is PEM encoded but does not include the header/footer",
				Type:        pluginsdk.TypeString,
				Computed:    true,
				Sensitive:   true,
			},
		},
	}
}

func servicePrincipalTokenSigningCertificateResourceCreate(ctx context.Context, d *pluginsdk.ResourceData, meta interface{}) pluginsdk.Diagnostics {
	client := meta.(*clients.Client).ServicePrincipals.ServicePrincipalClient
	servicePrincipalId := stable.NewServicePrincipalID(d.Get("service_principal_id").(string))

	properties := serviceprincipal.AddTokenSigningCertificateRequest{}

	if v, ok := d.GetOk("display_name"); ok {
		properties.DisplayName = nullable.NoZero(v.(string))
	}

	if v, ok := d.GetOk("end_date"); ok {
		properties.EndDateTime = nullable.NoZero(v.(string))
	}

	tf.LockByName(servicePrincipalResourceName, servicePrincipalId.ServicePrincipalId)
	defer tf.UnlockByName(servicePrincipalResourceName, servicePrincipalId.ServicePrincipalId)

	resp, err := client.AddTokenSigningCertificate(ctx, servicePrincipalId, properties, serviceprincipal.DefaultAddTokenSigningCertificateOperationOptions())
	if err != nil {
		return tf.ErrorDiagF(err, "Could not add token signing certificate to %s", servicePrincipalId)
	}

	key := resp.Model
	if key == nil {
		return tf.ErrorDiagF(errors.New("model was nil"), "API error adding token signing certificate for %s", servicePrincipalId)
	}

	// Wait for the credential to appear in the service principal manifest, this can take several minutes
	timeout, _ := ctx.Deadline()
	polledForCredential, err := (&pluginsdk.StateChangeConf{ //nolint:staticcheck
		Pending:                   []string{"Waiting"},
		Target:                    []string{"Done"},
		Timeout:                   time.Until(timeout),
		MinTimeout:                1 * time.Second,
		ContinuousTargetOccurence: 5,
		Refresh: func() (interface{}, string, error) {
			resp, err := client.GetServicePrincipal(ctx, servicePrincipalId, serviceprincipal.DefaultGetServicePrincipalOperationOptions())
			if err != nil {
				return nil, "Error", err
			}

			servicePrincipal := resp.Model
			if servicePrincipal == nil {
				return nil, "Error", errors.New("model was nil")
			}

			if servicePrincipal.KeyCredentials != nil {
				for _, cred := range *servicePrincipal.KeyCredentials {
					if strings.EqualFold(cred.KeyId.GetOrZero(), key.KeyId.GetOrZero()) {
						return &cred, "Done", nil
					}
				}
			}

			return nil, "Waiting", nil
		},
	}).WaitForStateContext(ctx)

	if err != nil {
		return tf.ErrorDiagF(err, "Waiting for token signing certificate credential for %s", servicePrincipalId)
	} else if polledForCredential == nil {
		return tf.ErrorDiagF(errors.New("certificate credential not found in service principal manifest"), "Waiting for certificate credential for %s", servicePrincipalId)
	}

	// Workaround b/c the returned keyId is for the Sign key, rather than Verify key,
	// so we need to get the Verify keyId based on the customKeyIdentifier
	servicePrincipalResponse, err := client.GetServicePrincipal(ctx, servicePrincipalId, serviceprincipal.DefaultGetServicePrincipalOperationOptions())
	if err != nil {
		return tf.ErrorDiagF(err, "Retrieving %s", servicePrincipalId)
	}

	servicePrincipal := servicePrincipalResponse.Model
	if servicePrincipal == nil {
		return tf.ErrorDiagF(errors.New("model was nil"), "Retrieving %s", servicePrincipalId)
	}

	credential := credentials.GetVerifyKeyCredentialFromCustomKeyId(servicePrincipal.KeyCredentials, key.CustomKeyIdentifier.GetOrZero())
	if credential == nil {
		return tf.ErrorDiagF(errors.New("returned credential was nil"), "Could not determine key ID for newly added token signing certificate for %s", servicePrincipalId)
	}

	id := parse.NewCredentialID(servicePrincipalId.ServicePrincipalId, "tokenSigningCertificate", credential.KeyId.GetOrZero())
	d.SetId(id.String())

	return servicePrincipalTokenSigningCertificateResourceRead(ctx, d, meta)
}

func servicePrincipalTokenSigningCertificateResourceRead(ctx context.Context, d *pluginsdk.ResourceData, meta interface{}) pluginsdk.Diagnostics {
	client := meta.(*clients.Client).ServicePrincipals.ServicePrincipalClient

	id, err := parse.SigningCertificateID(d.Id())
	if err != nil {
		return tf.ErrorDiagPathF(err, "id", "Parsing certificate credential with ID %q", d.Id())
	}

	servicePrincipalId := stable.NewServicePrincipalID(id.ObjectId)

	options := serviceprincipal.GetServicePrincipalOperationOptions{
		Select: pointer.To([]string{"keyCredentials"}),
	}
	resp, err := client.GetServicePrincipal(ctx, servicePrincipalId, options)
	if err != nil {
		if response.WasNotFound(resp.HttpResponse) {
			log.Printf("[DEBUG] %s was not found - removing from state!", servicePrincipalId)
			d.SetId("")
			return nil
		}
		return tf.ErrorDiagPathF(err, "service_principal_id", "Retrieving %s", servicePrincipalId)
	}

	servicePrincipal := resp.Model
	if servicePrincipal == nil {
		return tf.ErrorDiagF(errors.New("model was nil"), "Retrieving %s", servicePrincipalId)
	}

	credential := credentials.GetKeyCredential(servicePrincipal.KeyCredentials, id.KeyId)
	if credential == nil {
		log.Printf("[DEBUG] Certificate credential %q (ID %q) was not found - removing from state!", id.KeyId, id.ObjectId)
		d.SetId("")
		return nil
	}

	tf.Set(d, "service_principal_id", id.ObjectId)
	tf.Set(d, "key_id", id.KeyId)
	tf.Set(d, "display_name", credential.DisplayName.GetOrZero())
	tf.Set(d, "value", credential.Key.GetOrZero())
	tf.Set(d, "start_date", credential.StartDateTime.GetOrZero())
	tf.Set(d, "end_date", credential.EndDateTime.GetOrZero())

	// thumbprint not available when querying service principal, so we generate it from the pem value in the Key field.
	var thumbprint string
	if credential.Key != nil {
		thumbprint, err = credentials.GetTokenSigningCertificateThumbprint(
			[]byte("-----BEGIN CERTIFICATE-----\n" + credential.Key.GetOrZero() + "\n-----END CERTIFICATE-----"))
		if err != nil {
			return tf.ErrorDiagPathF(err, "id", "parsing tokenSigningCertificate key value with ID %q", id.KeyId)
		}
	}
	tf.Set(d, "thumbprint", thumbprint)

	return nil
}

func servicePrincipalTokenSigningCertificateResourceDelete(ctx context.Context, d *pluginsdk.ResourceData, meta interface{}) pluginsdk.Diagnostics {
	client := meta.(*clients.Client).ServicePrincipals.ServicePrincipalClient

	id, err := parse.SigningCertificateID(d.Id())
	if err != nil {
		return tf.ErrorDiagPathF(err, "id", "Parsing certificate credential with ID %q", d.Id())
	}

	servicePrincipalId := stable.NewServicePrincipalID(id.ObjectId)

	tf.LockByName(servicePrincipalResourceName, servicePrincipalId.ServicePrincipalId)
	defer tf.UnlockByName(servicePrincipalResourceName, servicePrincipalId.ServicePrincipalId)

	resp, err := client.GetServicePrincipal(ctx, servicePrincipalId, serviceprincipal.DefaultGetServicePrincipalOperationOptions())
	if err != nil {
		if response.WasNotFound(resp.HttpResponse) {
			return tf.ErrorDiagPathF(fmt.Errorf("Service Principal was not found"), "service_principal_id", "Retrieving %s", servicePrincipalId)
		}
		return tf.ErrorDiagPathF(err, "service_principal_id", "Retrieving %s", servicePrincipalId)
	}

	servicePrincipal := resp.Model
	if servicePrincipal == nil {
		return tf.ErrorDiagF(errors.New("model was nil"), "Retrieving %s", servicePrincipalId)
	}

	// use CustomKeyIdentifier to determine which certs and passwords are associated
	customKeyId := ""
	newKeyCredentials := make([]stable.KeyCredential, 0)
	if servicePrincipal.KeyCredentials != nil {
		for _, cred := range *servicePrincipal.KeyCredentials {
			if !strings.EqualFold(cred.KeyId.GetOrZero(), id.KeyId) {
				customKeyId = cred.CustomKeyIdentifier.GetOrZero()
			}
		}
		for _, cred := range *servicePrincipal.KeyCredentials {
			if !strings.EqualFold(cred.CustomKeyIdentifier.GetOrZero(), customKeyId) {
				newKeyCredentials = append(newKeyCredentials, cred)
			}
		}
	}

	newPasswordCredentials := make([]stable.PasswordCredential, 0)
	if servicePrincipal.PasswordCredentials != nil {
		for _, cred := range *servicePrincipal.PasswordCredentials {
			if !strings.EqualFold(cred.CustomKeyIdentifier.GetOrZero(), customKeyId) {
				newPasswordCredentials = append(newPasswordCredentials, cred)
			}
		}
	}

	properties := stable.ServicePrincipal{
		KeyCredentials:      &newKeyCredentials,
		PasswordCredentials: &newPasswordCredentials,
	}
	if _, err := client.UpdateServicePrincipal(ctx, servicePrincipalId, properties, serviceprincipal.DefaultUpdateServicePrincipalOperationOptions()); err != nil {
		return tf.ErrorDiagF(err, "Removing token signing certificate credentials %q from %s", id.KeyId, servicePrincipalId)
	}

	// Wait for service principal token signing certificate to be deleted
	if err := consistency.WaitForDeletion(ctx, func(ctx context.Context) (*bool, error) {
		resp, err := client.GetServicePrincipal(ctx, servicePrincipalId, serviceprincipal.DefaultGetServicePrincipalOperationOptions())
		if err != nil {
			return nil, err
		}

		servicePrincipal := resp.Model
		if servicePrincipal == nil {
			return nil, errors.New("model was nil")
		}

		credential := credentials.GetKeyCredential(servicePrincipal.KeyCredentials, id.KeyId)
		return pointer.To(credential != nil), nil
	}); err != nil {
		return tf.ErrorDiagF(err, "Waiting for deletion of token signing certificate credential %q %s", id.KeyId, servicePrincipalId)
	}

	return nil
}
