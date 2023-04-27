package client

import (
	"github.com/hashicorp/terraform-provider-azuread/internal/common"
	"github.com/manicminer/hamilton/msgraph"
)

type Client struct {
	UserFlowAttributesClient *msgraph.UserFlowAttributesClient
}

func NewClient(o *common.ClientOptions) *Client {
	userFlowAttributeClient := msgraph.NewUserFlowAttributesClient()
	o.ConfigureClient(&userFlowAttributeClient.BaseClient)

	return &Client{
		UserFlowAttributesClient: userFlowAttributeClient,
	}
}
