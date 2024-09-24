package directoryobject

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DirectoryObjectClient struct {
	Client *msgraph.Client
}

func NewDirectoryObjectClientWithBaseURI(sdkApi sdkEnv.Api) (*DirectoryObjectClient, error) {
	client, err := msgraph.NewClient(sdkApi, "directoryobject", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DirectoryObjectClient: %+v", err)
	}

	return &DirectoryObjectClient{
		Client: client,
	}, nil
}
