package entitlementmanagementroledefinition

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EntitlementManagementRoleDefinitionClient struct {
	Client *msgraph.Client
}

func NewEntitlementManagementRoleDefinitionClientWithBaseURI(sdkApi sdkEnv.Api) (*EntitlementManagementRoleDefinitionClient, error) {
	client, err := msgraph.NewClient(sdkApi, "entitlementmanagementroledefinition", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating EntitlementManagementRoleDefinitionClient: %+v", err)
	}

	return &EntitlementManagementRoleDefinitionClient{
		Client: client,
	}, nil
}
