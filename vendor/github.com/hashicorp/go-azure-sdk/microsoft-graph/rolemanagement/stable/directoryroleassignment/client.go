package directoryroleassignment

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DirectoryRoleAssignmentClient struct {
	Client *msgraph.Client
}

func NewDirectoryRoleAssignmentClientWithBaseURI(sdkApi sdkEnv.Api) (*DirectoryRoleAssignmentClient, error) {
	client, err := msgraph.NewClient(sdkApi, "directoryroleassignment", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DirectoryRoleAssignmentClient: %+v", err)
	}

	return &DirectoryRoleAssignmentClient{
		Client: client,
	}, nil
}
