package privilegedaccessgroupeligibilityschedule

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PrivilegedAccessGroupEligibilityScheduleClient struct {
	Client *msgraph.Client
}

func NewPrivilegedAccessGroupEligibilityScheduleClientWithBaseURI(sdkApi sdkEnv.Api) (*PrivilegedAccessGroupEligibilityScheduleClient, error) {
	client, err := msgraph.NewClient(sdkApi, "privilegedaccessgroupeligibilityschedule", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating PrivilegedAccessGroupEligibilityScheduleClient: %+v", err)
	}

	return &PrivilegedAccessGroupEligibilityScheduleClient{
		Client: client,
	}, nil
}
