package group

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupClient struct {
	Client *msgraph.Client
}

func NewGroupClientWithBaseURI(sdkApi sdkEnv.Api) (*GroupClient, error) {
	client, err := msgraph.NewClient(sdkApi, "group", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating GroupClient: %+v", err)
	}

	return &GroupClient{
		Client: client,
	}, nil
}
