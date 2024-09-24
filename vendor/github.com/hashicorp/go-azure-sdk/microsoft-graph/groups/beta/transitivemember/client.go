package transitivemember

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TransitiveMemberClient struct {
	Client *msgraph.Client
}

func NewTransitiveMemberClientWithBaseURI(sdkApi sdkEnv.Api) (*TransitiveMemberClient, error) {
	client, err := msgraph.NewClient(sdkApi, "transitivemember", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating TransitiveMemberClient: %+v", err)
	}

	return &TransitiveMemberClient{
		Client: client,
	}, nil
}
