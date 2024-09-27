package serviceprincipal

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AddTokenSigningCertificateOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *stable.SelfSignedCertificate
}

type AddTokenSigningCertificateOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultAddTokenSigningCertificateOperationOptions() AddTokenSigningCertificateOperationOptions {
	return AddTokenSigningCertificateOperationOptions{}
}

func (o AddTokenSigningCertificateOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o AddTokenSigningCertificateOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o AddTokenSigningCertificateOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// AddTokenSigningCertificate - Invoke action addTokenSigningCertificate. Create a self-signed signing certificate and
// return a selfSignedCertificate object, which is the public part of the generated certificate. The self-signed signing
// certificate is composed of the following objects, which are added to the servicePrincipal: + The keyCredentials
// object with the following objects: + A private key object with usage set to Sign. + A public key object with usage
// set to Verify. + The passwordCredentials object. All the objects have the same value of customKeyIdentifier. The
// passwordCredential is used to open the PFX file (private key). It and the associated private key object have the same
// value of keyId. When set during creation through the displayName property, the subject of the certificate cannot be
// updated. The startDateTime is set to the same time the certificate is created using the action. The endDateTime can
// be up to three years after the certificate is created.
func (c ServicePrincipalClient) AddTokenSigningCertificate(ctx context.Context, id stable.ServicePrincipalId, input AddTokenSigningCertificateRequest, options AddTokenSigningCertificateOperationOptions) (result AddTokenSigningCertificateOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusAccepted,
			http.StatusCreated,
			http.StatusNoContent,
			http.StatusOK,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/addTokenSigningCertificate", id.ID()),
		RetryFunc:     options.RetryFunc,
	}

	req, err := c.Client.NewRequest(ctx, opts)
	if err != nil {
		return
	}

	if err = req.Marshal(input); err != nil {
		return
	}

	var resp *client.Response
	resp, err = req.Execute(ctx)
	if resp != nil {
		result.OData = resp.OData
		result.HttpResponse = resp.Response
	}
	if err != nil {
		return
	}

	var model stable.SelfSignedCertificate
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
