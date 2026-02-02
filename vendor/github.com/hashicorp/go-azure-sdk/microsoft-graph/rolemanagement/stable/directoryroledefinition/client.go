package directoryroledefinition

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DirectoryRoleDefinitionClient struct {
	Client *msgraph.Client
}

func NewDirectoryRoleDefinitionClientWithBaseURI(sdkApi sdkEnv.Api) (*DirectoryRoleDefinitionClient, error) {
	client, err := msgraph.NewClient(sdkApi, "directoryroledefinition", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DirectoryRoleDefinitionClient: %+v", err)
	}

	return &DirectoryRoleDefinitionClient{
		Client: client,
	}, nil
}
