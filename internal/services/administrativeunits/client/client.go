// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package client

import (
	"github.com/glueckkanja/terraform-provider-azuread/internal/common"
	administrativeunitBeta "github.com/hashicorp/go-azure-sdk/microsoft-graph/administrativeunits/beta/administrativeunit"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/directory/stable/administrativeunit"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/directory/stable/administrativeunitmember"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/directory/stable/administrativeunitscopedrolemember"
)

type Client struct {
	AdministrativeUnitClient                 *administrativeunit.AdministrativeUnitClient
	AdministrativeUnitClientBeta             *administrativeunitBeta.AdministrativeUnitClient
	AdministrativeUnitMemberClient           *administrativeunitmember.AdministrativeUnitMemberClient
	AdministrativeUnitScopedRoleMemberClient *administrativeunitscopedrolemember.AdministrativeUnitScopedRoleMemberClient
}

func NewClient(o *common.ClientOptions) (*Client, error) {
	administrativeUnitClient, err := administrativeunit.NewAdministrativeUnitClientWithBaseURI(o.Environment.MicrosoftGraph)
	if err != nil {
		return nil, err
	}
	o.Configure(administrativeUnitClient.Client)

	// Beta API needed to delete administrative units - the stable API is broken and returns 404 with `{"Message":"The OData path is invalid."}`
	administrativeUnitClientBeta, err := administrativeunitBeta.NewAdministrativeUnitClientWithBaseURI(o.Environment.MicrosoftGraph)
	if err != nil {
		return nil, err
	}
	o.Configure(administrativeUnitClientBeta.Client)

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
		AdministrativeUnitClientBeta:             administrativeUnitClientBeta,
		AdministrativeUnitMemberClient:           memberClient,
		AdministrativeUnitScopedRoleMemberClient: scopedRoleMemberClient,
	}, nil
}
