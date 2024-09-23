package synchronizationsecret

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SynchronizationSecretClient struct {
	Client *msgraph.Client
}

func NewSynchronizationSecretClientWithBaseURI(sdkApi sdkEnv.Api) (*SynchronizationSecretClient, error) {
	client, err := msgraph.NewClient(sdkApi, "synchronizationsecret", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating SynchronizationSecretClient: %+v", err)
	}

	return &SynchronizationSecretClient{
		Client: client,
	}, nil
}
