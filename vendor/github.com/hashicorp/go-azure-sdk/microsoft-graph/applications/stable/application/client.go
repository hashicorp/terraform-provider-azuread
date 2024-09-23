package application

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ApplicationClient struct {
	Client *msgraph.Client
}

func NewApplicationClientWithBaseURI(sdkApi sdkEnv.Api) (*ApplicationClient, error) {
	client, err := msgraph.NewClient(sdkApi, "application", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ApplicationClient: %+v", err)
	}

	return &ApplicationClient{
		Client: client,
	}, nil
}
