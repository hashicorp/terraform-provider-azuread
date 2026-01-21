// Copyright IBM Corp. 2014, 2025
// SPDX-License-Identifier: MPL-2.0

package client

import (
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/directoryobjects/stable/directoryobject"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/oauth2permissiongrants/stable/oauth2permissiongrant"
	serviceprincipalBeta "github.com/hashicorp/go-azure-sdk/microsoft-graph/serviceprincipals/beta/serviceprincipal"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/serviceprincipals/stable/claimsmappingpolicy"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/serviceprincipals/stable/owner"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/serviceprincipals/stable/serviceprincipal"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/serviceprincipals/stable/synchronizationjob"
	"github.com/hashicorp/terraform-provider-azuread/internal/common"
)

type Client struct {
	ClaimsMappingPolicyClient   *claimsmappingpolicy.ClaimsMappingPolicyClient
	DirectoryObjectClient       *directoryobject.DirectoryObjectClient
	OAuth2PermissionGrantClient *oauth2permissiongrant.OAuth2PermissionGrantClient
	ServicePrincipalClient      *serviceprincipal.ServicePrincipalClient
	ServicePrincipalClientBeta  *serviceprincipalBeta.ServicePrincipalClient
	ServicePrincipalOwnerClient *owner.OwnerClient
	SynchronizationJobClient    *synchronizationjob.SynchronizationJobClient
}

func NewClient(o *common.ClientOptions) (*Client, error) {
	claimsMappingPolicyClient, err := claimsmappingpolicy.NewClaimsMappingPolicyClientWithBaseURI(o.Environment.MicrosoftGraph)
	if err != nil {
		return nil, err
	}
	o.Configure(claimsMappingPolicyClient.Client)

	directoryObjectClient, err := directoryobject.NewDirectoryObjectClientWithBaseURI(o.Environment.MicrosoftGraph)
	if err != nil {
		return nil, err
	}
	o.Configure(directoryObjectClient.Client)

	oAuth2PermissionGrantClient, err := oauth2permissiongrant.NewOAuth2PermissionGrantClientWithBaseURI(o.Environment.MicrosoftGraph)
	if err != nil {
		return nil, err
	}
	o.Configure(oAuth2PermissionGrantClient.Client)

	servicePrincipalClient, err := serviceprincipal.NewServicePrincipalClientWithBaseURI(o.Environment.MicrosoftGraph)
	if err != nil {
		return nil, err
	}
	o.Configure(servicePrincipalClient.Client)

	// Needed because v1.0 API doesn't return `samlMetadataUrl`
	servicePrincipalClientBeta, err := serviceprincipalBeta.NewServicePrincipalClientWithBaseURI(o.Environment.MicrosoftGraph)
	if err != nil {
		return nil, err
	}
	o.Configure(servicePrincipalClientBeta.Client)

	servicePrincipalOwnerClient, err := owner.NewOwnerClientWithBaseURI(o.Environment.MicrosoftGraph)
	if err != nil {
		return nil, err
	}
	o.Configure(servicePrincipalOwnerClient.Client)

	synchronizationJobClient, err := synchronizationjob.NewSynchronizationJobClientWithBaseURI(o.Environment.MicrosoftGraph)
	if err != nil {
		return nil, err
	}
	o.Configure(synchronizationJobClient.Client)

	return &Client{
		ClaimsMappingPolicyClient:   claimsMappingPolicyClient,
		DirectoryObjectClient:       directoryObjectClient,
		OAuth2PermissionGrantClient: oAuth2PermissionGrantClient,
		ServicePrincipalClient:      servicePrincipalClient,
		ServicePrincipalClientBeta:  servicePrincipalClientBeta,
		ServicePrincipalOwnerClient: servicePrincipalOwnerClient,
		SynchronizationJobClient:    synchronizationJobClient,
	}, nil
}
