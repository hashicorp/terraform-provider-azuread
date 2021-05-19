package client

import (
	"github.com/Azure/azure-sdk-for-go/services/graphrbac/1.6/graphrbac"
	"github.com/manicminer/hamilton/msgraph"

	"github.com/hashicorp/terraform-provider-azuread/internal/common"
)

type Client struct {
	AadClient *graphrbac.ApplicationsClient
	MsClient  *msgraph.ApplicationsClient
}

func NewClient(o *common.ClientOptions) *Client {
	aadClient := graphrbac.NewApplicationsClientWithBaseURI(o.AadGraphEndpoint, o.TenantID)
	msClient := msgraph.NewApplicationsClient(o.TenantID)
	o.ConfigureClient(&msClient.BaseClient, &aadClient.Client)

	return &Client{
		AadClient: &aadClient,
		MsClient:  msClient,
	}
}
