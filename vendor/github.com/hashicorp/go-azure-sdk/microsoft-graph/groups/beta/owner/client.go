package owner

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type OwnerClient struct {
	Client *msgraph.Client
}

func NewOwnerClientWithBaseURI(sdkApi sdkEnv.Api) (*OwnerClient, error) {
	client, err := msgraph.NewClient(sdkApi, "owner", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating OwnerClient: %+v", err)
	}

	return &OwnerClient{
		Client: client,
	}, nil
}
