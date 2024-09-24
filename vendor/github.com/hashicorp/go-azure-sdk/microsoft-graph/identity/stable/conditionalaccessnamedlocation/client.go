package conditionalaccessnamedlocation

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ConditionalAccessNamedLocationClient struct {
	Client *msgraph.Client
}

func NewConditionalAccessNamedLocationClientWithBaseURI(sdkApi sdkEnv.Api) (*ConditionalAccessNamedLocationClient, error) {
	client, err := msgraph.NewClient(sdkApi, "conditionalaccessnamedlocation", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ConditionalAccessNamedLocationClient: %+v", err)
	}

	return &ConditionalAccessNamedLocationClient{
		Client: client,
	}, nil
}
