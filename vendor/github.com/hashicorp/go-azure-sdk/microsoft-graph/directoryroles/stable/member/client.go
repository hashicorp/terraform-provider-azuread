package member

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MemberClient struct {
	Client *msgraph.Client
}

func NewMemberClientWithBaseURI(sdkApi sdkEnv.Api) (*MemberClient, error) {
	client, err := msgraph.NewClient(sdkApi, "member", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MemberClient: %+v", err)
	}

	return &MemberClient{
		Client: client,
	}, nil
}
