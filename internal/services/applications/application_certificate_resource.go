// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package applications

import (
	"context"
	"errors"
	"log"
	"strings"
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/pointer"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/applications/stable/application"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/terraform-provider-azuread/internal/clients"
	"github.com/hashicorp/terraform-provider-azuread/internal/helpers/consistency"
	"github.com/hashicorp/terraform-provider-azuread/internal/helpers/credentials"
	"github.com/hashicorp/terraform-provider-azuread/internal/helpers/tf"
	"github.com/hashicorp/terraform-provider-azuread/internal/helpers/tf/pluginsdk"
	"github.com/hashicorp/terraform-provider-azuread/internal/helpers/tf/validation"
	"github.com/hashicorp/terraform-provider-azuread/internal/services/applications/parse"
)

func applicationCertificateResource() *pluginsdk.Resource {
	return &pluginsdk.Resource{
		CreateContext: applicationCertificateResourceCreate,
		ReadContext:   applicationCertificateResourceRead,
		DeleteContext: applicationCertificateResourceDelete,

		Timeouts: &pluginsdk.ResourceTimeout{
			Create: pluginsdk.DefaultTimeout(10 * time.Minute),
			Read:   pluginsdk.DefaultTimeout(5 * time.Minute),
			Update: pluginsdk.DefaultTimeout(10 * time.Minute),
			Delete: pluginsdk.DefaultTimeout(5 * time.Minute),
		},

		Importer: pluginsdk.ImporterValidatingResourceId(func(id string) error {
			_, err := parse.CertificateID(id)
			return err
		}),

		Schema: map[string]*pluginsdk.Schema{
			"application_id": {
				Description:  "The resource ID of the application for which this certificate should be created",
				Type:         pluginsdk.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: stable.ValidateApplicationID,
			},

			"encoding": {
				Description: "Specifies the encoding used for the supplied certificate data",
				Type:        pluginsdk.TypeString,
				Optional:    true,
				ForceNew:    true,
				Default:     "pem",
				ValidateFunc: validation.StringInSlice([]string{
					"base64",
					"hex",
					"pem",
				}, false),
			},

			"key_id": {
				Description:  "A UUID used to uniquely identify this certificate. If omitted, a random UUID will be automatically generated",
				Type:         pluginsdk.TypeString,
				Optional:     true,
				Computed:     true,
				ForceNew:     true,
				ValidateFunc: validation.IsUUID,
			},

			"start_date": {
				Description:  "The start date from which the certificate is valid, formatted as an RFC3339 date string (e.g. `2018-01-01T01:02:03Z`). If this isn't specified, the current date and time are use",
				Type:         pluginsdk.TypeString,
				Optional:     true,
				Computed:     true,
				ForceNew:     true,
				ValidateFunc: validation.IsRFC3339Time,
			},

			"end_date": {
				Description:   "The end date until which the certificate is valid, formatted as an RFC3339 date string (e.g. `2018-01-01T01:02:03Z`). If omitted, the API will decide a suitable expiry date, which is typically around 2 years from the start date",
				Type:          pluginsdk.TypeString,
				Optional:      true,
				Computed:      true,
				ForceNew:      true,
				ConflictsWith: []string{"end_date_relative"},
				ValidateFunc:  validation.IsRFC3339Time,
			},

			"end_date_relative": {
				Description:   "A relative duration for which the certificate is valid until, for example `240h` (10 days) or `2400h30m`",
				Type:          pluginsdk.TypeString,
				Optional:      true,
				ForceNew:      true,
				ConflictsWith: []string{"end_date"},
				ValidateFunc:  validation.StringIsNotEmpty,
			},

			"type": {
				Description: "The type of key/certificate",
				Type:        pluginsdk.TypeString,
				Optional:    true,
				ForceNew:    true,
				ValidateFunc: validation.StringInSlice([]string{
					"AsymmetricX509Cert",
					"Symmetric",
				}, false),
			},

			"value": {
				Description: "The certificate data, which can be PEM encoded, base64 encoded DER or hexadecimal encoded DER. See also the `encoding` argument",
				Type:        pluginsdk.TypeString,
				Required:    true,
				ForceNew:    true,
				Sensitive:   true,
			},
		},
	}
}

func applicationCertificateResourceCreate(ctx context.Context, d *pluginsdk.ResourceData, meta interface{}) pluginsdk.Diagnostics {
	client := meta.(*clients.Client).Applications.ApplicationClient

	applicationId, err := stable.ParseApplicationID(d.Get("application_id").(string))
	if err != nil {
		return tf.ErrorDiagPathF(err, "application_id", "Parsing `application_id`")
	}

	credential, err := credentials.KeyCredentialForResource(d)
	if err != nil {
		attr := ""
		if kerr, ok := err.(credentials.CredentialError); ok {
			attr = kerr.Attr()
		}
		return tf.ErrorDiagPathF(err, attr, "Generating certificate credentials for %s", applicationId)
	}

	if credential.KeyId == nil {
		return tf.ErrorDiagF(errors.New("keyId for certificate credential is nil"), "Creating certificate credential")
	}
	id := parse.NewCredentialID(applicationId.ApplicationId, "certificate", credential.KeyId.GetOrZero())

	tf.LockByName(applicationResourceName, id.ObjectId)
	defer tf.UnlockByName(applicationResourceName, id.ObjectId)

	resp, err := client.GetApplication(ctx, *applicationId, application.DefaultGetApplicationOperationOptions())
	if err != nil {
		return tf.ErrorDiagPathF(err, "application_id", "Retrieving %s", applicationId)
	}

	app := resp.Model
	if app == nil {
		return tf.ErrorDiagF(errors.New("model was nil"), "Retrieving %s", applicationId)
	}

	newCredentials := make([]stable.KeyCredential, 0)
	if app.KeyCredentials != nil {
		for _, cred := range *app.KeyCredentials {
			if strings.EqualFold(cred.KeyId.GetOrZero(), credential.KeyId.GetOrZero()) {
				return tf.ImportAsExistsDiag("azuread_application_certificate", id.String())
			}
			newCredentials = append(newCredentials, cred)
		}
	}

	newCredentials = append(newCredentials, *credential)

	properties := stable.Application{
		Id:             &id.ObjectId,
		KeyCredentials: &newCredentials,
	}
	if _, err = client.UpdateApplication(ctx, *applicationId, properties, application.DefaultUpdateApplicationOperationOptions()); err != nil {
		return tf.ErrorDiagF(err, "Adding certificate for %s", applicationId)
	}

	// Wait for the credential to appear in the application manifest, this can take several minutes
	timeout, _ := ctx.Deadline()
	polledForCredential, err := (&pluginsdk.StateChangeConf{ //nolint:staticcheck
		Pending:                   []string{"Waiting"},
		Target:                    []string{"Done"},
		Timeout:                   time.Until(timeout),
		MinTimeout:                1 * time.Second,
		ContinuousTargetOccurence: 5,
		Refresh: func() (interface{}, string, error) {
			resp, err := client.GetApplication(ctx, *applicationId, application.DefaultGetApplicationOperationOptions())
			if err != nil {
				return nil, "Error", err
			}
			app := resp.Model
			if app == nil {
				return nil, "Error", errors.New("model was nil")
			}

			if app.KeyCredentials != nil {
				for _, cred := range *app.KeyCredentials {
					if strings.EqualFold(cred.KeyId.GetOrZero(), id.KeyId) {
						return &cred, "Done", nil
					}
				}
			}

			return nil, "Waiting", nil
		},
	}).WaitForStateContext(ctx)

	if err != nil {
		return tf.ErrorDiagF(err, "Waiting for certificate credential for %s", applicationId)
	} else if polledForCredential == nil {
		return tf.ErrorDiagF(errors.New("certificate credential not found in application manifest"), "Waiting for certificate credential for %s", applicationId)
	}

	d.SetId(id.String())

	return applicationCertificateResourceRead(ctx, d, meta)
}

func applicationCertificateResourceRead(ctx context.Context, d *pluginsdk.ResourceData, meta interface{}) pluginsdk.Diagnostics {
	client := meta.(*clients.Client).Applications.ApplicationClient

	id, err := parse.CertificateID(d.Id())
	if err != nil {
		return tf.ErrorDiagPathF(err, "id", "Parsing certificate credential with ID %q", d.Id())
	}

	applicationId := stable.NewApplicationID(id.ObjectId)

	resp, err := client.GetApplication(ctx, applicationId, application.DefaultGetApplicationOperationOptions())
	if err != nil {
		return tf.ErrorDiagPathF(err, "application_id", "Retrieving %s", applicationId)
	}

	app := resp.Model
	if app == nil {
		return tf.ErrorDiagF(errors.New("model was nil"), "Retrieving %s", applicationId)
	}

	credential := credentials.GetKeyCredential(app.KeyCredentials, id.KeyId)
	if credential == nil {
		log.Printf("[DEBUG] Certificate credential %q (ID %q) was not found - removing from state!", id.KeyId, id.ObjectId)
		d.SetId("")
		return nil
	}

	tf.Set(d, "application_id", applicationId.ID())
	tf.Set(d, "key_id", id.KeyId)
	tf.Set(d, "type", credential.Type.GetOrZero())
	tf.Set(d, "start_date", credential.StartDateTime.GetOrZero())
	tf.Set(d, "end_date", credential.EndDateTime.GetOrZero())

	return nil
}

func applicationCertificateResourceDelete(ctx context.Context, d *pluginsdk.ResourceData, meta interface{}) pluginsdk.Diagnostics {
	client := meta.(*clients.Client).Applications.ApplicationClient

	id, err := parse.CertificateID(d.Id())
	if err != nil {
		return tf.ErrorDiagPathF(err, "id", "Parsing certificate credential with ID %q", d.Id())
	}

	tf.LockByName(applicationResourceName, id.ObjectId)
	defer tf.UnlockByName(applicationResourceName, id.ObjectId)

	applicationId := stable.NewApplicationID(id.ObjectId)

	resp, err := client.GetApplication(ctx, applicationId, application.DefaultGetApplicationOperationOptions())
	if err != nil {
		return tf.ErrorDiagPathF(err, "application_id", "Retrieving %s", applicationId)
	}

	app := resp.Model
	if app == nil {
		return tf.ErrorDiagF(errors.New("model was nil"), "Retrieving %s", applicationId)
	}

	newCredentials := make([]stable.KeyCredential, 0)
	if app.KeyCredentials != nil {
		for _, cred := range *app.KeyCredentials {
			if !strings.EqualFold(cred.KeyId.GetOrZero(), id.KeyId) {
				newCredentials = append(newCredentials, cred)
			}
		}
	}

	properties := stable.Application{
		Id:             &id.ObjectId,
		KeyCredentials: &newCredentials,
	}
	if _, err = client.UpdateApplication(ctx, applicationId, properties, application.DefaultUpdateApplicationOperationOptions()); err != nil {
		return tf.ErrorDiagF(err, "Removing certificate credential %q from application with object ID %q", id.KeyId, id.ObjectId)
	}

	// Wait for application certificate to be deleted
	if err := consistency.WaitForDeletion(ctx, func(ctx context.Context) (*bool, error) {
		resp, err := client.GetApplication(ctx, applicationId, application.DefaultGetApplicationOperationOptions())
		if err != nil {
			return nil, err
		}
		app := resp.Model
		if app == nil {
			return nil, errors.New("model was nil")
		}

		credential := credentials.GetKeyCredential(app.KeyCredentials, id.KeyId)
		if credential == nil {
			return pointer.To(false), nil
		}

		return pointer.To(true), nil
	}); err != nil {
		return tf.ErrorDiagF(err, "Waiting for deletion of certificate credential %q from application with object ID %q", id.KeyId, id.ObjectId)
	}

	return nil
}
