package directoryrole

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DirectoryRoleClient struct {
	Client *msgraph.Client
}

func NewDirectoryRoleClientWithBaseURI(sdkApi sdkEnv.Api) (*DirectoryRoleClient, error) {
	client, err := msgraph.NewClient(sdkApi, "directoryrole", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DirectoryRoleClient: %+v", err)
	}

	return &DirectoryRoleClient{
		Client: client,
	}, nil
}
