package applicationtemplate

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ApplicationTemplateClient struct {
	Client *msgraph.Client
}

func NewApplicationTemplateClientWithBaseURI(sdkApi sdkEnv.Api) (*ApplicationTemplateClient, error) {
	client, err := msgraph.NewClient(sdkApi, "applicationtemplate", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ApplicationTemplateClient: %+v", err)
	}

	return &ApplicationTemplateClient{
		Client: client,
	}, nil
}
