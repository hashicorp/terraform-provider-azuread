package directoryroletemplate

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DirectoryRoleTemplateClient struct {
	Client *msgraph.Client
}

func NewDirectoryRoleTemplateClientWithBaseURI(sdkApi sdkEnv.Api) (*DirectoryRoleTemplateClient, error) {
	client, err := msgraph.NewClient(sdkApi, "directoryroletemplate", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DirectoryRoleTemplateClient: %+v", err)
	}

	return &DirectoryRoleTemplateClient{
		Client: client,
	}, nil
}
