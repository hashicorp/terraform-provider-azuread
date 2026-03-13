package entitlementmanagementassignmentpolicy

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EntitlementManagementAssignmentPolicyClient struct {
	Client *msgraph.Client
}

func NewEntitlementManagementAssignmentPolicyClientWithBaseURI(sdkApi sdkEnv.Api) (*EntitlementManagementAssignmentPolicyClient, error) {
	client, err := msgraph.NewClient(sdkApi, "entitlementmanagementassignmentpolicy", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating EntitlementManagementAssignmentPolicyClient: %+v", err)
	}

	return &EntitlementManagementAssignmentPolicyClient{
		Client: client,
	}, nil
}
