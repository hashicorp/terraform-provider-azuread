package client

import (
	"github.com/manicminer/hamilton/msgraph"

	"github.com/hashicorp/terraform-provider-azuread/internal/common"
)

type Client struct {
	NamedLocationsClient *msgraph.NamedLocationsClient
	PoliciesClient       *msgraph.ConditionalAccessPolicyClient
}

func NewClient(o *common.ClientOptions) *Client {
	namedLocationsClient := msgraph.NewNamedLocationsClient(o.TenantID)
	o.ConfigureClient(&namedLocationsClient.BaseClient)

	return &Client{
		NamedLocationsClient: namedLocationsClient,
	}
}
