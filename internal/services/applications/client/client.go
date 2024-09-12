// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package client

import (
	applicationBeta "github.com/hashicorp/go-azure-sdk/microsoft-graph/applications/beta/application"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/applications/stable/application"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/applications/stable/federatedidentitycredential"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/applications/stable/logo"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/applications/stable/owner"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/applicationtemplates/stable/applicationtemplate"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/directoryobjects/stable/directoryobject"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/serviceprincipals/stable/serviceprincipal"
	"github.com/hashicorp/terraform-provider-azuread/internal/common"
)

type Client struct {
	ApplicationClient                      *application.ApplicationClient
	ApplicationClientBeta                  *applicationBeta.ApplicationClient
	ApplicationLogoClient                  *logo.LogoClient
	ApplicationOwnerClient                 *owner.OwnerClient
	ApplicationFederatedIdentityCredential *federatedidentitycredential.FederatedIdentityCredentialClient
	ApplicationTemplateClient              *applicationtemplate.ApplicationTemplateClient
	ServicePrincipalClient                 *serviceprincipal.ServicePrincipalClient
}

func NewClient(o *common.ClientOptions) (*Client, error) {
	applicationClient, err := application.NewApplicationClientWithBaseURI(o.Environment.MicrosoftGraph)
	if err != nil {
		return nil, err
	}
	o.Configure(applicationClient.Client)

	// See https://github.com/microsoftgraph/msgraph-metadata/issues/273
	applicationClientBeta, err := applicationBeta.NewApplicationClientWithBaseURI(o.Environment.MicrosoftGraph)
	if err != nil {
		return nil, err
	}
	o.Configure(applicationClientBeta.Client)

	applicationLogoClient, err := logo.NewLogoClientWithBaseURI(o.Environment.MicrosoftGraph)
	if err != nil {
		return nil, err
	}
	o.Configure(applicationLogoClient.Client)

	applicationOwnerClient, err := owner.NewOwnerClientWithBaseURI(o.Environment.MicrosoftGraph)
	if err != nil {
		return nil, err
	}
	o.Configure(applicationOwnerClient.Client)

	applicationFederatedIdentityCredentialClient, err := federatedidentitycredential.NewFederatedIdentityCredentialClientWithBaseURI(o.Environment.MicrosoftGraph)
	if err != nil {
		return nil, err
	}
	o.Configure(applicationFederatedIdentityCredentialClient.Client)

	applicationTemplateClient, err := applicationtemplate.NewApplicationTemplateClientWithBaseURI(o.Environment.MicrosoftGraph)
	if err != nil {
		return nil, err
	}
	o.Configure(applicationTemplateClient.Client)

	directoryObjectClient, err := directoryobject.NewDirectoryObjectClientWithBaseURI(o.Environment.MicrosoftGraph)
	if err != nil {
		return nil, err
	}
	o.Configure(directoryObjectClient.Client)

	servicePrincipalClient, err := serviceprincipal.NewServicePrincipalClientWithBaseURI(o.Environment.MicrosoftGraph)
	if err != nil {
		return nil, err
	}
	o.Configure(servicePrincipalClient.Client)

	return &Client{
		ApplicationClient:                      applicationClient,
		ApplicationClientBeta:                  applicationClientBeta,
		ApplicationLogoClient:                  applicationLogoClient,
		ApplicationOwnerClient:                 applicationOwnerClient,
		ApplicationFederatedIdentityCredential: applicationFederatedIdentityCredentialClient,
		ApplicationTemplateClient:              applicationTemplateClient,
		ServicePrincipalClient:                 servicePrincipalClient,
	}, nil
}
