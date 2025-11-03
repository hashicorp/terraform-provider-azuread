package federatedidentitycredential

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type FederatedIdentityCredentialClient struct {
	Client *msgraph.Client
}

func NewFederatedIdentityCredentialClientWithBaseURI(sdkApi sdkEnv.Api) (*FederatedIdentityCredentialClient, error) {
	client, err := msgraph.NewClient(sdkApi, "federatedidentitycredential", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating FederatedIdentityCredentialClient: %+v", err)
	}

	return &FederatedIdentityCredentialClient{
		Client: client,
	}, nil
}
