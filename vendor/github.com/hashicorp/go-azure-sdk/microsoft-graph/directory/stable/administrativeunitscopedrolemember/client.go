package administrativeunitscopedrolemember

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AdministrativeUnitScopedRoleMemberClient struct {
	Client *msgraph.Client
}

func NewAdministrativeUnitScopedRoleMemberClientWithBaseURI(sdkApi sdkEnv.Api) (*AdministrativeUnitScopedRoleMemberClient, error) {
	client, err := msgraph.NewClient(sdkApi, "administrativeunitscopedrolemember", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating AdministrativeUnitScopedRoleMemberClient: %+v", err)
	}

	return &AdministrativeUnitScopedRoleMemberClient{
		Client: client,
	}, nil
}
