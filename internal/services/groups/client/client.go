// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package client

import (
	administrativeunitmemberBeta "github.com/hashicorp/go-azure-sdk/microsoft-graph/directory/beta/administrativeunitmember"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/directoryobjects/stable/directoryobject"
	groupBeta "github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/group"
	memberBeta "github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/member"
	memberofBeta "github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/memberof"
	ownerBeta "github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/owner"
	transitivememberBeta "github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/transitivemember"
	"github.com/hashicorp/terraform-provider-azuread/internal/common"
)

type Client struct {
	AdministrativeUnitMemberClientBeta *administrativeunitmemberBeta.AdministrativeUnitMemberClient
	DirectoryObjectClient              *directoryobject.DirectoryObjectClient
	GroupClientBeta                    *groupBeta.GroupClient
	GroupMemberClientBeta              *memberBeta.MemberClient
	GroupMemberOfClientBeta            *memberofBeta.MemberOfClient
	GroupOwnerClientBeta               *ownerBeta.OwnerClient
	GroupTransitiveMemberClientBeta    *transitivememberBeta.TransitiveMemberClient
}

func NewClient(o *common.ClientOptions) (*Client, error) {
	administrativeUnitMemberClient, err := administrativeunitmemberBeta.NewAdministrativeUnitMemberClientWithBaseURI(o.Environment.MicrosoftGraph)
	if err != nil {
		return nil, err
	}
	o.Configure(administrativeUnitMemberClient.Client)

	directoryObjectClient, err := directoryobject.NewDirectoryObjectClientWithBaseURI(o.Environment.MicrosoftGraph)
	if err != nil {
		return nil, err
	}
	o.Configure(directoryObjectClient.Client)

	// resourceBehaviorOptions & resourceProvisioningOptions fields not supported in v1.0 API
	groupClient, err := groupBeta.NewGroupClientWithBaseURI(o.Environment.MicrosoftGraph)
	if err != nil {
		return nil, err
	}
	o.Configure(groupClient.Client)

	// Group members not returned in full when using v1.0 API, see https://github.com/hashicorp/terraform-provider-azuread/issues/1018
	memberClient, err := memberBeta.NewMemberClientWithBaseURI(o.Environment.MicrosoftGraph)
	if err != nil {
		return nil, err
	}
	o.Configure(memberClient.Client)

	memberOfClient, err := memberofBeta.NewMemberOfClientWithBaseURI(o.Environment.MicrosoftGraph)
	if err != nil {
		return nil, err
	}
	o.Configure(memberOfClient.Client)

	ownerClient, err := ownerBeta.NewOwnerClientWithBaseURI(o.Environment.MicrosoftGraph)
	if err != nil {
		return nil, err
	}
	o.Configure(ownerClient.Client)

	// Group members not returned in full when using v1.0 API, see https://github.com/hashicorp/terraform-provider-azuread/issues/1018
	transitiveMemberClient, err := transitivememberBeta.NewTransitiveMemberClientWithBaseURI(o.Environment.MicrosoftGraph)
	if err != nil {
		return nil, err
	}
	o.Configure(transitiveMemberClient.Client)

	return &Client{
		AdministrativeUnitMemberClientBeta: administrativeUnitMemberClient,
		DirectoryObjectClient:              directoryObjectClient,
		GroupClientBeta:                    groupClient,
		GroupMemberClientBeta:              memberClient,
		GroupMemberOfClientBeta:            memberOfClient,
		GroupOwnerClientBeta:               ownerClient,
		GroupTransitiveMemberClientBeta:    transitiveMemberClient,
	}, nil
}
