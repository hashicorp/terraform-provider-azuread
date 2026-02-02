package memberof

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MemberOfClient struct {
	Client *msgraph.Client
}

func NewMemberOfClientWithBaseURI(sdkApi sdkEnv.Api) (*MemberOfClient, error) {
	client, err := msgraph.NewClient(sdkApi, "memberof", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MemberOfClient: %+v", err)
	}

	return &MemberOfClient{
		Client: client,
	}, nil
}
