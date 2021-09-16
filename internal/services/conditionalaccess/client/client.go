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

	policiesClient := msgraph.NewConditionalAccessPolicyClient(o.TenantID)
	o.ConfigureClient(&policiesClient.BaseClient)
	policiesClient.BaseClient.ApiVersion = msgraph.Version10

	return &Client{
		NamedLocationsClient: namedLocationsClient,
		PoliciesClient:       policiesClient,
	}
}
