package claimsmappingpolicy

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ClaimsMappingPolicyClient struct {
	Client *msgraph.Client
}

func NewClaimsMappingPolicyClientWithBaseURI(sdkApi sdkEnv.Api) (*ClaimsMappingPolicyClient, error) {
	client, err := msgraph.NewClient(sdkApi, "claimsmappingpolicy", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ClaimsMappingPolicyClient: %+v", err)
	}

	return &ClaimsMappingPolicyClient{
		Client: client,
	}, nil
}
