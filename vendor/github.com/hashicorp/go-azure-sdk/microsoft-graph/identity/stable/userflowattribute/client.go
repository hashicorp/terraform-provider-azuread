package userflowattribute

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserFlowAttributeClient struct {
	Client *msgraph.Client
}

func NewUserFlowAttributeClientWithBaseURI(sdkApi sdkEnv.Api) (*UserFlowAttributeClient, error) {
	client, err := msgraph.NewClient(sdkApi, "userflowattribute", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserFlowAttributeClient: %+v", err)
	}

	return &UserFlowAttributeClient{
		Client: client,
	}, nil
}
