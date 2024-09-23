package user

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserClient struct {
	Client *msgraph.Client
}

func NewUserClientWithBaseURI(sdkApi sdkEnv.Api) (*UserClient, error) {
	client, err := msgraph.NewClient(sdkApi, "user", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserClient: %+v", err)
	}

	return &UserClient{
		Client: client,
	}, nil
}
