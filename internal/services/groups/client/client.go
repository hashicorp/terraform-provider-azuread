package client

import (
	"github.com/Azure/azure-sdk-for-go/services/graphrbac/1.6/graphrbac"

	"github.com/hashicorp/terraform-provider-azuread/internal/common"
)

type Client struct {
	AadClient *graphrbac.GroupsClient
}

func NewClient(o *common.ClientOptions) *Client {
	aadClient := graphrbac.NewGroupsClientWithBaseURI(o.AadGraphEndpoint, o.TenantID)
	o.ConfigureClient(&aadClient.Client, o.AadGraphAuthorizer)

	return &Client{
		AadClient: &aadClient,
	}
}
