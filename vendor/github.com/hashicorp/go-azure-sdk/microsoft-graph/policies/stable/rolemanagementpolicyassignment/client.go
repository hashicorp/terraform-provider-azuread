package rolemanagementpolicyassignment

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RoleManagementPolicyAssignmentClient struct {
	Client *msgraph.Client
}

func NewRoleManagementPolicyAssignmentClientWithBaseURI(sdkApi sdkEnv.Api) (*RoleManagementPolicyAssignmentClient, error) {
	client, err := msgraph.NewClient(sdkApi, "rolemanagementpolicyassignment", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating RoleManagementPolicyAssignmentClient: %+v", err)
	}

	return &RoleManagementPolicyAssignmentClient{
		Client: client,
	}, nil
}
