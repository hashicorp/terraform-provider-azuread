package serviceprincipal

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ServicePrincipalClient struct {
	Client *msgraph.Client
}

func NewServicePrincipalClientWithBaseURI(sdkApi sdkEnv.Api) (*ServicePrincipalClient, error) {
	client, err := msgraph.NewClient(sdkApi, "serviceprincipal", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ServicePrincipalClient: %+v", err)
	}

	return &ServicePrincipalClient{
		Client: client,
	}, nil
}
