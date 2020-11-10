package services

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/hashicorp/go-azure-helpers/sender"
	"github.com/hashicorp/terraform-plugin-sdk/meta"

	"github.com/terraform-providers/terraform-provider-azuread/version"
)

const TerraformPartnerID = "222c6c49-1b0a-5959-a213-6608f9eb8820"

type ClientOptions struct {
	TenantID    string
	Environment azure.Environment

	PartnerID        string
	TerraformVersion string

	SkipProviderReg           bool
	DisableTerraformPartnerID bool
}

func (o ClientOptions) ConfigureClient(c *autorest.Client, authorizer autorest.Authorizer) {
	setUserAgent(c, o.TerraformVersion, o.PartnerID, o.DisableTerraformPartnerID)

	c.Authorizer = authorizer
	c.Sender = sender.BuildSender("AzureAD")
}

func setUserAgent(client *autorest.Client, tfVersion, partnerID string, disableTerraformPartnerID bool) {
	tfUserAgent := fmt.Sprintf("HashiCorp Terraform/%s (+https://www.terraform.io) Terraform Plugin SDK/%s", tfVersion, meta.SDKVersionString())

	providerUserAgent := fmt.Sprintf("%s terraform-provider-azuread/%s", tfUserAgent, version.ProviderVersion)
	client.UserAgent = strings.TrimSpace(fmt.Sprintf("%s %s", client.UserAgent, providerUserAgent))

	// append the CloudShell version to the user agent if it exists
	if azureAgent := os.Getenv("AZURE_HTTP_USER_AGENT"); azureAgent != "" {
		client.UserAgent = fmt.Sprintf("%s %s", client.UserAgent, azureAgent)
	}

	// only one pid can be interpreted currently
	// hence, send partner ID if present, otherwise send Terraform GUID
	// unless users have opted out
	if partnerID == "" && !disableTerraformPartnerID {
		// Microsoft’s Terraform Partner ID is this specific GUID
		partnerID = TerraformPartnerID
	}

	if partnerID != "" {
		client.UserAgent = fmt.Sprintf("%s pid-%s", client.UserAgent, partnerID)
	}

	log.Printf("[DEBUG] AzureAD Client User Agent: %s\n", client.UserAgent)
}
