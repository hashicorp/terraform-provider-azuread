package b2cuserflow

import (
	"github.com/manicminer/hamilton/msgraph"

	"github.com/hashicorp/terraform-provider-azuread/internal/common"
)

type Client struct {
	UserFlowClient *msgraph.B2CUserFlowClient
}

func NewClient(o *common.ClientOptions) *Client {
	userFlowClient := msgraph.NewB2CUserFlowClient(o.TenantID)
	o.ConfigureClient(&userFlowClient.BaseClient)

	// Use beta API as this resource doesn't exist in the v1.0 API
	userFlowClient.BaseClient.ApiVersion = msgraph.VersionBeta

	return &Client{
		UserFlowClient: userFlowClient,
	}
}
