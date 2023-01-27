package client

import (
	"github.com/hashicorp/terraform-provider-azuread/internal/common"
	"github.com/manicminer/hamilton/msgraph"
)

type Client struct {
	NamedLocationsClient *msgraph.NamedLocationsClient
	PoliciesClient       *msgraph.ConditionalAccessPoliciesClient
}

func NewClient(o *common.ClientOptions) *Client {
	namedLocationsClient := msgraph.NewNamedLocationsClient(o.TenantID)
	o.ConfigureClient(&namedLocationsClient.BaseClient)

	policiesClient := msgraph.NewConditionalAccessPoliciesClient(o.TenantID)
	o.ConfigureClient(&policiesClient.BaseClient)

	return &Client{
		NamedLocationsClient: namedLocationsClient,
		PoliciesClient:       policiesClient,
	}
}
