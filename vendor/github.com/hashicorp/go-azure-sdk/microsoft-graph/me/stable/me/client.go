package me

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MeClient struct {
	Client *msgraph.Client
}

func NewMeClientWithBaseURI(sdkApi sdkEnv.Api) (*MeClient, error) {
	client, err := msgraph.NewClient(sdkApi, "me", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MeClient: %+v", err)
	}

	return &MeClient{
		Client: client,
	}, nil
}
