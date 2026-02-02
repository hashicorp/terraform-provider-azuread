package entitlementmanagementaccesspackageaccesspackageresourcerolescope

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EntitlementManagementAccessPackageAccessPackageResourceRoleScopeClient struct {
	Client *msgraph.Client
}

func NewEntitlementManagementAccessPackageAccessPackageResourceRoleScopeClientWithBaseURI(sdkApi sdkEnv.Api) (*EntitlementManagementAccessPackageAccessPackageResourceRoleScopeClient, error) {
	client, err := msgraph.NewClient(sdkApi, "entitlementmanagementaccesspackageaccesspackageresourcerolescope", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating EntitlementManagementAccessPackageAccessPackageResourceRoleScopeClient: %+v", err)
	}

	return &EntitlementManagementAccessPackageAccessPackageResourceRoleScopeClient{
		Client: client,
	}, nil
}
