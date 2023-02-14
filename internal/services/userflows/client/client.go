package client

import (
	"github.com/hashicorp/terraform-provider-azuread/internal/common"
	"github.com/manicminer/hamilton/msgraph"
)

type Client struct {
	B2CUserFlowClient        *msgraph.B2CUserFlowClient
	UserFlowAttributesClient *msgraph.UserFlowAttributesClient
}

func NewClient(o *common.ClientOptions) *Client {
	userFlowAttributeClient := msgraph.NewUserFlowAttributesClient()
	o.ConfigureClient(&userFlowAttributeClient.BaseClient)

	b2cUserFlowClient := msgraph.NewB2CUserFlowClient()
	o.ConfigureClient(&b2cUserFlowClient.BaseClient)

	// Use beta API as this resource doesn't exist in the v1.0 API
	b2cUserFlowClient.BaseClient.ApiVersion = msgraph.VersionBeta

	return &Client{
		B2CUserFlowClient:        b2cUserFlowClient,
		UserFlowAttributesClient: userFlowAttributeClient,
	}
}
