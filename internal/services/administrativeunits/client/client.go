// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package client

import (
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/directory/stable/administrativeunit"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/directory/stable/administrativeunitmember"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/directory/stable/administrativeunitscopedrolemember"
	"github.com/hashicorp/terraform-provider-azuread/internal/common"
)

type Client struct {
	AdministrativeUnitClient                 *administrativeunit.AdministrativeUnitClient
	AdministrativeUnitMemberClient           *administrativeunitmember.AdministrativeUnitMemberClient
	AdministrativeUnitScopedRoleMemberClient *administrativeunitscopedrolemember.AdministrativeUnitScopedRoleMemberClient
}

func NewClient(o *common.ClientOptions) (*Client, error) {
	administrativeUnitClient, err := administrativeunit.NewAdministrativeUnitClientWithBaseURI(o.Environment.MicrosoftGraph)
	if err != nil {
		return nil, err
	}
	o.Configure(administrativeUnitClient.Client)

	memberClient, err := administrativeunitmember.NewAdministrativeUnitMemberClientWithBaseURI(o.Environment.MicrosoftGraph)
	if err != nil {
		return nil, err
	}
	o.Configure(memberClient.Client)

	scopedRoleMemberClient, err := administrativeunitscopedrolemember.NewAdministrativeUnitScopedRoleMemberClientWithBaseURI(o.Environment.MicrosoftGraph)
	if err != nil {
		return nil, err
	}
	o.Configure(scopedRoleMemberClient.Client)

	return &Client{
		AdministrativeUnitClient:                 administrativeUnitClient,
		AdministrativeUnitMemberClient:           memberClient,
		AdministrativeUnitScopedRoleMemberClient: scopedRoleMemberClient,
	}, nil
}
