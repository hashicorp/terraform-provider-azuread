package invitation

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type InvitationClient struct {
	Client *msgraph.Client
}

func NewInvitationClientWithBaseURI(sdkApi sdkEnv.Api) (*InvitationClient, error) {
	client, err := msgraph.NewClient(sdkApi, "invitation", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating InvitationClient: %+v", err)
	}

	return &InvitationClient{
		Client: client,
	}, nil
}
