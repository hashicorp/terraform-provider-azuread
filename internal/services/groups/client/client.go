// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package client

import (
	"github.com/glueckkanja/terraform-provider-azuread/internal/common"
	administrativeunitmemberBeta "github.com/hashicorp/go-azure-sdk/microsoft-graph/directory/beta/administrativeunitmember"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/directoryobjects/stable/directoryobject"
	groupBeta "github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/group"
	memberBeta "github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/member"
	memberofBeta "github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/memberof"
	ownerBeta "github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/owner"
	transitivememberBeta "github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/transitivemember"
)

// Note: Whilst it is technically possible that we could use both the Stable and Beta APIs for groups (retaining use of
// Beta APIs solely for those properties that require it), we are currently using the Beta APIs pretty much across the
// board owing to the complexity of the azuread_group resource, and known bugs when retrieving members with the Stable API.

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
	administrativeUnitMemberClientBeta, err := administrativeunitmemberBeta.NewAdministrativeUnitMemberClientWithBaseURI(o.Environment.MicrosoftGraph)
	if err != nil {
		return nil, err
	}
	o.Configure(administrativeUnitMemberClientBeta.Client)

	directoryObjectClient, err := directoryobject.NewDirectoryObjectClientWithBaseURI(o.Environment.MicrosoftGraph)
	if err != nil {
		return nil, err
	}
	o.Configure(directoryObjectClient.Client)

	// resourceBehaviorOptions & resourceProvisioningOptions fields not supported in v1.0 API
	groupClientBeta, err := groupBeta.NewGroupClientWithBaseURI(o.Environment.MicrosoftGraph)
	if err != nil {
		return nil, err
	}
	o.Configure(groupClientBeta.Client)

	// Group members not returned in full when using v1.0 API, see https://github.com/glueckkanja/terraform-provider-azuread/issues/1018
	memberClientBeta, err := memberBeta.NewMemberClientWithBaseURI(o.Environment.MicrosoftGraph)
	if err != nil {
		return nil, err
	}
	o.Configure(memberClientBeta.Client)

	memberOfClientBeta, err := memberofBeta.NewMemberOfClientWithBaseURI(o.Environment.MicrosoftGraph)
	if err != nil {
		return nil, err
	}
	o.Configure(memberOfClientBeta.Client)

	ownerClientBeta, err := ownerBeta.NewOwnerClientWithBaseURI(o.Environment.MicrosoftGraph)
	if err != nil {
		return nil, err
	}
	o.Configure(ownerClientBeta.Client)

	// Group members not returned in full when using v1.0 API, see https://github.com/glueckkanja/terraform-provider-azuread/issues/1018
	transitiveMemberClientBeta, err := transitivememberBeta.NewTransitiveMemberClientWithBaseURI(o.Environment.MicrosoftGraph)
	if err != nil {
		return nil, err
	}
	o.Configure(transitiveMemberClientBeta.Client)

	return &Client{
		AdministrativeUnitMemberClientBeta: administrativeUnitMemberClientBeta,
		DirectoryObjectClient:              directoryObjectClient,
		GroupClientBeta:                    groupClientBeta,
		GroupMemberClientBeta:              memberClientBeta,
		GroupMemberOfClientBeta:            memberOfClientBeta,
		GroupOwnerClientBeta:               ownerClientBeta,
		GroupTransitiveMemberClientBeta:    transitiveMemberClientBeta,
	}, nil
}
