package domain

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CreateDomainOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *stable.Domain
}

type CreateDomainOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultCreateDomainOperationOptions() CreateDomainOperationOptions {
	return CreateDomainOperationOptions{}
}

func (o CreateDomainOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o CreateDomainOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o CreateDomainOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// CreateDomain - Create domain. Adds a domain to the tenant. Important: You cannot use an associated domain with your
// Microsoft Entra tenant until ownership is verified. See List verificationDnsRecords for details. Root domains require
// verification. For example, contoso.com requires verification. If a root domain is verified, subdomains of the root
// domain are automatically verified. For example, subdomain.contoso.com is automatically be verified if contoso.com has
// been verified.
func (c DomainClient) CreateDomain(ctx context.Context, input stable.Domain, options CreateDomainOperationOptions) (result CreateDomainOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusCreated,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Path:          "/domains",
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

	var model stable.Domain
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
