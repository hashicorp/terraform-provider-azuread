package authenticationstrengthpolicycombinationconfiguration

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AuthenticationStrengthPolicyCombinationConfigurationClient struct {
	Client *msgraph.Client
}

func NewAuthenticationStrengthPolicyCombinationConfigurationClientWithBaseURI(sdkApi sdkEnv.Api) (*AuthenticationStrengthPolicyCombinationConfigurationClient, error) {
	client, err := msgraph.NewClient(sdkApi, "authenticationstrengthpolicycombinationconfiguration", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating AuthenticationStrengthPolicyCombinationConfigurationClient: %+v", err)
	}

	return &AuthenticationStrengthPolicyCombinationConfigurationClient{
		Client: client,
	}, nil
}
