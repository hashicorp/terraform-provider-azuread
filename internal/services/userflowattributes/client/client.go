package client

import (
	"github.com/manicminer/hamilton/msgraph"

	"github.com/hashicorp/terraform-provider-azuread/internal/common"
)

type Client struct {
	Client *msgraph.UserFlowAttributesClient
}

func NewClient(o *common.ClientOptions) *Client {
	userflowAttributeClient := msgraph.NewUserFlowAttributesClient(o.TenantID)
	o.ConfigureClient(&userflowAttributeClient.BaseClient)

	return &Client{userflowAttributeClient}
}
