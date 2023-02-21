package serviceprincipals

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"regexp"
	"strings"
	"time"

	"github.com/hashicorp/go-azure-sdk/sdk/odata"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/hashicorp/terraform-provider-azuread/internal/clients"
	"github.com/hashicorp/terraform-provider-azuread/internal/helpers"
	"github.com/hashicorp/terraform-provider-azuread/internal/services/serviceprincipals/parse"
	"github.com/hashicorp/terraform-provider-azuread/internal/tf"
	"github.com/hashicorp/terraform-provider-azuread/internal/utils"
	"github.com/hashicorp/terraform-provider-azuread/internal/validate"
	"github.com/manicminer/hamilton/msgraph"
)

func servicePrincipalTokenSigningCertificateResource() *schema.Resource {
	return &schema.Resource{
		CreateContext: servicePrincipalTokenSigningCertificateResourceCreate,
		ReadContext:   servicePrincipalTokenSigningCertificateResourceRead,
		DeleteContext: servicePrincipalTokenSigningCertificateResourceDelete,

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(5 * time.Minute),
			Read:   schema.DefaultTimeout(5 * time.Minute),
			Update: schema.DefaultTimeout(5 * time.Minute),
			Delete: schema.DefaultTimeout(5 * time.Minute),
		},

		Importer: tf.ValidateResourceIDPriorToImport(func(id string) error {
			_, err := parse.SigningCertificateID(id)
			return err
		}),

		Schema: map[string]*schema.Schema{
			"service_principal_id": {
				Description:      "The object ID of the service principal for which this certificate should be created",
				Type:             schema.TypeString,
				Required:         true,
				ForceNew:         true,
				ValidateDiagFunc: validate.UUID,
			},

			"display_name": {
				Description:      "A friendly name for the certificate",
				Type:             schema.TypeString,
				Optional:         true,
				Computed:         true,
				ForceNew:         true,
				ValidateDiagFunc: validate.ValidateDiag(validation.StringMatch(regexp.MustCompile("^CN=.+$|^$"), "")),
			},

			"end_date": {
				Description:  "The end date until which the certificate is valid, formatted as an RFC3339 date string (e.g. `2018-01-01T01:02:03Z`). Default is 3 years from current date.",
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ForceNew:     true,
				ValidateFunc: validation.IsRFC3339Time,
			},

			"key_id": {
				Description: "A UUID used to uniquely identify the verify certificate.",
				Type:        schema.TypeString,
				Computed:    true,
			},

			"thumbprint": {
				Description: "The thumbprint of the certificate.",
				Type:        schema.TypeString,
				Computed:    true,
			},

			"start_date": {
				Description: "The start date from which the certificate is valid, formatted as an RFC3339 date string (e.g. `2018-01-01T01:02:03Z`).",
				Type:        schema.TypeString,
				Computed:    true,
			},

			"value": {
				Description: "The certificate data, which is PEM encoded but does not include the header/footer",
				Type:        schema.TypeString,
				Computed:    true,
				Sensitive:   true,
			},
		},
	}
}

func servicePrincipalTokenSigningCertificateResourceCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*clients.Client).ServicePrincipals.ServicePrincipalsClient
	objectId := d.Get("service_principal_id").(string)

	keyCreds := msgraph.KeyCredential{}
	if v, ok := d.GetOk("display_name"); ok {
		keyCreds.DisplayName = utils.String(v.(string))
	}

	if v, ok := d.GetOk("end_date"); ok {
		endDate, err := time.Parse(time.RFC3339, v.(string))
		if err != nil {
			tf.ErrorDiagF(err, "Unable to parse the provided end_date %q: %+v", v, err)
		}
		keyCreds.EndDateTime = &endDate
	}

	tf.LockByName(servicePrincipalResourceName, objectId)
	defer tf.UnlockByName(servicePrincipalResourceName, objectId)

	key, _, err := client.AddTokenSigningCertificate(ctx, objectId, keyCreds)
	if err != nil {
		return tf.ErrorDiagF(err, "Could not add token signing certificate to service principal with object ID: %q", objectId)
	}

	// Wait for the credential to appear in the service principal manifest, this can take several minutes
	timeout, _ := ctx.Deadline()
	polledForCredential, err := (&resource.StateChangeConf{
		Pending:                   []string{"Waiting"},
		Target:                    []string{"Done"},
		Timeout:                   time.Until(timeout),
		MinTimeout:                1 * time.Second,
		ContinuousTargetOccurence: 5,
		Refresh: func() (interface{}, string, error) {
			servicePrincipal, _, err := client.Get(ctx, objectId, odata.Query{})
			if err != nil {
				return nil, "Error", err
			}

			if servicePrincipal.KeyCredentials != nil {
				for _, cred := range *servicePrincipal.KeyCredentials {
					if cred.KeyId != nil && strings.EqualFold(*cred.KeyId, *key.KeyId) {
						return &cred, "Done", nil
					}
				}
			}

			return nil, "Waiting", nil
		},
	}).WaitForStateContext(ctx)

	if err != nil {
		return tf.ErrorDiagF(err, "Waiting for token_signing_certificate credential for service principal with object ID %q", objectId)
	} else if polledForCredential == nil {
		return tf.ErrorDiagF(errors.New("certificate credential not found in service principal manifest"), "Waiting for certificate credential for service principal with object ID %q", objectId)
	}

	// Workaround b/c the returned keyId is for the Sign key, rather than Verify key,
	// so we need to get the Verify keyId based on the customKeyIdentifier
	servicePrincipal, _, err := client.Get(ctx, objectId, odata.Query{})
	if err != nil {
		return tf.ErrorDiagF(err, "Could not get service principal with object ID: %q", objectId)
	}
	credential := helpers.GetVerifyKeyCredentialFromCustomKeyId(servicePrincipal.KeyCredentials, *key.CustomKeyIdentifier)

	if credential == nil {
		return tf.ErrorDiagF(errors.New("returned credential was nil"), "Could not determine key ID for newly added token signing certificate on service principal %q", objectId)
	}
	id := parse.NewCredentialID(objectId, "tokenSigningCertificate", *credential.KeyId)

	d.SetId(id.String())

	return servicePrincipalTokenSigningCertificateResourceRead(ctx, d, meta)
}

func servicePrincipalTokenSigningCertificateResourceRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*clients.Client).ServicePrincipals.ServicePrincipalsClient

	id, err := parse.SigningCertificateID(d.Id())
	if err != nil {
		return tf.ErrorDiagPathF(err, "id", "Parsing certificate credential with ID %q", d.Id())
	}

	servicePrincipal, status, err := client.Get(ctx, id.ObjectId, odata.Query{
		Select: []string{"keyCredentials"},
	})
	if err != nil {
		if status == http.StatusNotFound {
			log.Printf("[DEBUG] Service Principal with ID %q for %s credential %q was not found - removing from state!", id.ObjectId, id.KeyType, id.KeyId)
			d.SetId("")
			return nil
		}
		return tf.ErrorDiagPathF(err, "service_principal_id", "Retrieving service principal with object ID %q", id.ObjectId)
	}

	credential := helpers.GetKeyCredential(servicePrincipal.KeyCredentials, id.KeyId)
	if credential == nil {
		log.Printf("[DEBUG] Certificate credential %q (ID %q) was not found - removing from state!", id.KeyId, id.ObjectId)
		d.SetId("")
		return nil
	}

	tf.Set(d, "service_principal_id", id.ObjectId)
	tf.Set(d, "key_id", id.KeyId)
	tf.Set(d, "display_name", credential.DisplayName)
	tf.Set(d, "value", credential.Key)

	startDate := ""
	if v := credential.StartDateTime; v != nil {
		startDate = v.Format(time.RFC3339)
	}
	tf.Set(d, "start_date", startDate)

	endDate := ""
	if v := credential.EndDateTime; v != nil {
		endDate = v.Format(time.RFC3339)
	}
	tf.Set(d, "end_date", endDate)

	// thumbprint not available when querying service principal, so we generate it from the pem value in the Key field.
	thumbprint, err := helpers.GetTokenSigningCertificateThumbprint(
		[]byte("-----BEGIN CERTIFICATE-----\n" + *credential.Key + "\n-----END CERTIFICATE-----"))
	if err != nil {
		return tf.ErrorDiagPathF(err, "id", "parsing tokenSigningCertificate key value with ID %q", id.KeyId)
	}

	tf.Set(d, "thumbprint", thumbprint)

	return nil
}

func servicePrincipalTokenSigningCertificateResourceDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*clients.Client).ServicePrincipals.ServicePrincipalsClient

	id, err := parse.SigningCertificateID(d.Id())
	if err != nil {
		return tf.ErrorDiagPathF(err, "id", "Parsing certificate credential with ID %q", d.Id())
	}

	tf.LockByName(servicePrincipalResourceName, id.ObjectId)
	defer tf.UnlockByName(servicePrincipalResourceName, id.ObjectId)

	app, status, err := client.Get(ctx, id.ObjectId, odata.Query{})
	if err != nil {
		if status == http.StatusNotFound {
			return tf.ErrorDiagPathF(fmt.Errorf("Service Principal was not found"), "service_principal_id", "Retrieving service principal with object ID %q", id.ObjectId)
		}
		return tf.ErrorDiagPathF(err, "service_principal_id", "Retrieving service principal with object ID %q", id.ObjectId)
	}

	// use CustomKeyIdentifier to determine which certs and passwords
	// are associated together.
	customKeyId := ""
	newKeyCredentials := make([]msgraph.KeyCredential, 0)
	if app.KeyCredentials != nil {
		for _, cred := range *app.KeyCredentials {
			if cred.KeyId != nil && !strings.EqualFold(*cred.KeyId, id.KeyId) {
				customKeyId = *cred.CustomKeyIdentifier
			}
		}
		for _, cred := range *app.KeyCredentials {
			if cred.KeyId != nil && !strings.EqualFold(*cred.CustomKeyIdentifier, customKeyId) {
				newKeyCredentials = append(newKeyCredentials, cred)
			}
		}
	}

	newPasswordCredentials := make([]msgraph.PasswordCredential, 0)
	if app.PasswordCredentials != nil {
		for _, cred := range *app.PasswordCredentials {
			if cred.KeyId != nil && !strings.EqualFold(*cred.CustomKeyIdentifier, customKeyId) {
				newPasswordCredentials = append(newPasswordCredentials, cred)
			}
		}
	}

	properties := msgraph.ServicePrincipal{
		DirectoryObject: msgraph.DirectoryObject{
			Id: &id.ObjectId,
		},
		KeyCredentials:      &newKeyCredentials,
		PasswordCredentials: &newPasswordCredentials,
	}
	if _, err := client.Update(ctx, properties); err != nil {
		return tf.ErrorDiagF(err, "Removing token signing certificate credentials %q from service principal with object ID %q", id.KeyId, id.ObjectId)
	}

	// Wait for service principal token signing certificate to be deleted
	if err := helpers.WaitForDeletion(ctx, func(ctx context.Context) (*bool, error) {
		client.BaseClient.DisableRetries = true

		servicePrincipal, _, err := client.Get(ctx, id.ObjectId, odata.Query{})
		if err != nil {
			return nil, err
		}

		credential := helpers.GetKeyCredential(servicePrincipal.KeyCredentials, id.KeyId)
		if credential == nil {
			return utils.Bool(false), nil
		}

		return utils.Bool(true), nil
	}); err != nil {
		return tf.ErrorDiagF(err, "Waiting for deletion of token signing certificate credential %q from service principal with object ID %q", id.KeyId, id.ObjectId)
	}

	return nil
}
