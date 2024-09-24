package manager

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagerClient struct {
	Client *msgraph.Client
}

func NewManagerClientWithBaseURI(sdkApi sdkEnv.Api) (*ManagerClient, error) {
	client, err := msgraph.NewClient(sdkApi, "manager", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ManagerClient: %+v", err)
	}

	return &ManagerClient{
		Client: client,
	}, nil
}
