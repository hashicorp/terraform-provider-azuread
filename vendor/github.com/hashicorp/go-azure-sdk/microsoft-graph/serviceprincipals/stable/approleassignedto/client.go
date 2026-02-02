package approleassignedto

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AppRoleAssignedToClient struct {
	Client *msgraph.Client
}

func NewAppRoleAssignedToClientWithBaseURI(sdkApi sdkEnv.Api) (*AppRoleAssignedToClient, error) {
	client, err := msgraph.NewClient(sdkApi, "approleassignedto", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating AppRoleAssignedToClient: %+v", err)
	}

	return &AppRoleAssignedToClient{
		Client: client,
	}, nil
}
