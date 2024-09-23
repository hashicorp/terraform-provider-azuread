package logo

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type LogoClient struct {
	Client *msgraph.Client
}

func NewLogoClientWithBaseURI(sdkApi sdkEnv.Api) (*LogoClient, error) {
	client, err := msgraph.NewClient(sdkApi, "logo", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating LogoClient: %+v", err)
	}

	return &LogoClient{
		Client: client,
	}, nil
}
