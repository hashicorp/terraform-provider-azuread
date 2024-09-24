package administrativeunit

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AdministrativeUnitClient struct {
	Client *msgraph.Client
}

func NewAdministrativeUnitClientWithBaseURI(sdkApi sdkEnv.Api) (*AdministrativeUnitClient, error) {
	client, err := msgraph.NewClient(sdkApi, "administrativeunit", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating AdministrativeUnitClient: %+v", err)
	}

	return &AdministrativeUnitClient{
		Client: client,
	}, nil
}
