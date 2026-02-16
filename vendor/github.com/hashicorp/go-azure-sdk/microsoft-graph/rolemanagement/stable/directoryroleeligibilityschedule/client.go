package directoryroleeligibilityschedule

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DirectoryRoleEligibilityScheduleClient struct {
	Client *msgraph.Client
}

func NewDirectoryRoleEligibilityScheduleClientWithBaseURI(sdkApi sdkEnv.Api) (*DirectoryRoleEligibilityScheduleClient, error) {
	client, err := msgraph.NewClient(sdkApi, "directoryroleeligibilityschedule", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DirectoryRoleEligibilityScheduleClient: %+v", err)
	}

	return &DirectoryRoleEligibilityScheduleClient{
		Client: client,
	}, nil
}
