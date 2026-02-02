package entitlementmanagementroleassignment

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EntitlementManagementRoleAssignmentClient struct {
	Client *msgraph.Client
}

func NewEntitlementManagementRoleAssignmentClientWithBaseURI(sdkApi sdkEnv.Api) (*EntitlementManagementRoleAssignmentClient, error) {
	client, err := msgraph.NewClient(sdkApi, "entitlementmanagementroleassignment", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating EntitlementManagementRoleAssignmentClient: %+v", err)
	}

	return &EntitlementManagementRoleAssignmentClient{
		Client: client,
	}, nil
}
