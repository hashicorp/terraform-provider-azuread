package rolemanagementpolicy

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RoleManagementPolicyClient struct {
	Client *msgraph.Client
}

func NewRoleManagementPolicyClientWithBaseURI(sdkApi sdkEnv.Api) (*RoleManagementPolicyClient, error) {
	client, err := msgraph.NewClient(sdkApi, "rolemanagementpolicy", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating RoleManagementPolicyClient: %+v", err)
	}

	return &RoleManagementPolicyClient{
		Client: client,
	}, nil
}
