package authenticationstrengthpolicy

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AuthenticationStrengthPolicyClient struct {
	Client *msgraph.Client
}

func NewAuthenticationStrengthPolicyClientWithBaseURI(sdkApi sdkEnv.Api) (*AuthenticationStrengthPolicyClient, error) {
	client, err := msgraph.NewClient(sdkApi, "authenticationstrengthpolicy", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating AuthenticationStrengthPolicyClient: %+v", err)
	}

	return &AuthenticationStrengthPolicyClient{
		Client: client,
	}, nil
}
