package client

import (
	"github.com/manicminer/hamilton/msgraph"

	"github.com/hashicorp/terraform-provider-azuread/internal/common"
)

type Client struct {
	AccessPackageClient                 *msgraph.AccessPackageClient
	AccessPackageAssignmentPolicyClient *msgraph.AccessPackageAssignmentPolicyClient
}

func NewClient(o *common.ClientOptions) *Client {
	accessPackageClient := msgraph.NewAccessPackageClient(o.TenantID)
	o.ConfigureClient(&accessPackageClient.BaseClient)

	accessPackageAssignmentPolicyClient := msgraph.NewAccessPackageAssignmentPolicyClient(o.TenantID)
	o.ConfigureClient(&accessPackageAssignmentPolicyClient.BaseClient)

	return &Client{
		AccessPackageClient:                 accessPackageClient,
		AccessPackageAssignmentPolicyClient: nil,
	}
}
