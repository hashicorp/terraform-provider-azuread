package client

import (
	"github.com/hashicorp/terraform-provider-azuread/internal/common"
	"github.com/manicminer/hamilton/msgraph"
)

type Client struct {
	DomainsClient *msgraph.DomainsClient
}

func NewClient(o *common.ClientOptions) *Client {
	msClient := msgraph.NewDomainsClient(o.TenantID)
	o.ConfigureClient(&msClient.BaseClient)

	return &Client{
		DomainsClient: msClient,
	}
}
