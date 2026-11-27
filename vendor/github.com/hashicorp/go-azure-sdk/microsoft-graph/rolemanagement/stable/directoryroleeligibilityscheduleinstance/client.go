package directoryroleeligibilityscheduleinstance

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DirectoryRoleEligibilityScheduleInstanceClient struct {
	Client *msgraph.Client
}

func NewDirectoryRoleEligibilityScheduleInstanceClientWithBaseURI(sdkApi sdkEnv.Api) (*DirectoryRoleEligibilityScheduleInstanceClient, error) {
	client, err := msgraph.NewClient(sdkApi, "directoryroleeligibilityscheduleinstance", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DirectoryRoleEligibilityScheduleInstanceClient: %+v", err)
	}

	return &DirectoryRoleEligibilityScheduleInstanceClient{
		Client: client,
	}, nil
}
