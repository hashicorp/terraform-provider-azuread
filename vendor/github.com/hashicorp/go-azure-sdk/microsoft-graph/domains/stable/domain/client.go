package domain

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DomainClient struct {
	Client *msgraph.Client
}

func NewDomainClientWithBaseURI(sdkApi sdkEnv.Api) (*DomainClient, error) {
	client, err := msgraph.NewClient(sdkApi, "domain", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DomainClient: %+v", err)
	}

	return &DomainClient{
		Client: client,
	}, nil
}
