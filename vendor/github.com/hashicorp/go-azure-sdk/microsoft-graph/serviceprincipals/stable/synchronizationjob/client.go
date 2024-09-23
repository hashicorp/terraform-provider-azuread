package synchronizationjob

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SynchronizationJobClient struct {
	Client *msgraph.Client
}

func NewSynchronizationJobClientWithBaseURI(sdkApi sdkEnv.Api) (*SynchronizationJobClient, error) {
	client, err := msgraph.NewClient(sdkApi, "synchronizationjob", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating SynchronizationJobClient: %+v", err)
	}

	return &SynchronizationJobClient{
		Client: client,
	}, nil
}
