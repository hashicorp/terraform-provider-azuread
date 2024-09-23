// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package client

import (
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/directoryobjects/stable/directoryobject"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/directoryroles/stable/directoryrole"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/directoryroles/stable/member"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/directoryroletemplates/stable/directoryroletemplate"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/stable/directoryroleassignment"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/stable/directoryroledefinition"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/stable/directoryroleeligibilityschedulerequest"
	"github.com/hashicorp/terraform-provider-azuread/internal/common"
)

type Client struct {
	DirectoryObjectClient                         *directoryobject.DirectoryObjectClient
	DirectoryRoleAssignmentClient                 *directoryroleassignment.DirectoryRoleAssignmentClient
	DirectoryRoleClient                           *directoryrole.DirectoryRoleClient
	DirectoryRoleDefinitionClient                 *directoryroledefinition.DirectoryRoleDefinitionClient
	DirectoryRoleEligibilityScheduleRequestClient *directoryroleeligibilityschedulerequest.DirectoryRoleEligibilityScheduleRequestClient
	DirectoryRoleMemberClient                     *member.MemberClient
	DirectoryRoleTemplateClient                   *directoryroletemplate.DirectoryRoleTemplateClient
}

func NewClient(o *common.ClientOptions) (*Client, error) {
	directoryObjectClient, err := directoryobject.NewDirectoryObjectClientWithBaseURI(o.Environment.MicrosoftGraph)
	if err != nil {
		return nil, err
	}
	o.Configure(directoryObjectClient.Client)

	directoryRoleClient, err := directoryrole.NewDirectoryRoleClientWithBaseURI(o.Environment.MicrosoftGraph)
	if err != nil {
		return nil, err
	}
	o.Configure(directoryRoleClient.Client)

	directoryRoleAssignmentClient, err := directoryroleassignment.NewDirectoryRoleAssignmentClientWithBaseURI(o.Environment.MicrosoftGraph)
	if err != nil {
		return nil, err
	}
	o.Configure(directoryRoleAssignmentClient.Client)

	directoryRoleDefinitionClient, err := directoryroledefinition.NewDirectoryRoleDefinitionClientWithBaseURI(o.Environment.MicrosoftGraph)
	if err != nil {
		return nil, err
	}
	o.Configure(directoryRoleDefinitionClient.Client)

	directoryRoleMemberClient, err := member.NewMemberClientWithBaseURI(o.Environment.MicrosoftGraph)
	if err != nil {
		return nil, err
	}
	o.Configure(directoryRoleMemberClient.Client)

	directoryRoleEligibilityScheduleRequestClient, err := directoryroleeligibilityschedulerequest.NewDirectoryRoleEligibilityScheduleRequestClientWithBaseURI(o.Environment.MicrosoftGraph)
	if err != nil {
		return nil, err
	}
	o.Configure(directoryRoleEligibilityScheduleRequestClient.Client)

	directoryRoleTemplateClient, err := directoryroletemplate.NewDirectoryRoleTemplateClientWithBaseURI(o.Environment.MicrosoftGraph)
	if err != nil {
		return nil, err
	}
	o.Configure(directoryRoleTemplateClient.Client)

	return &Client{
		DirectoryObjectClient:                         directoryObjectClient,
		DirectoryRoleAssignmentClient:                 directoryRoleAssignmentClient,
		DirectoryRoleClient:                           directoryRoleClient,
		DirectoryRoleDefinitionClient:                 directoryRoleDefinitionClient,
		DirectoryRoleEligibilityScheduleRequestClient: directoryRoleEligibilityScheduleRequestClient,
		DirectoryRoleMemberClient:                     directoryRoleMemberClient,
		DirectoryRoleTemplateClient:                   directoryRoleTemplateClient,
	}, nil
}
