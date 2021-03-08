package common

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/hashicorp/go-azure-helpers/sender"
	"github.com/hashicorp/terraform-plugin-sdk/v2/meta"

	"github.com/hashicorp/terraform-provider-azuread/version"
)

type ClientOptions struct {
	TenantID    string
	Environment azure.Environment

	PartnerID        string
	TerraformVersion string

	AadGraphAuthorizer autorest.Authorizer
	AadGraphEndpoint   string

	SkipProviderReg bool
}

func (o ClientOptions) ConfigureClient(c *autorest.Client, authorizer autorest.Authorizer) {
	setUserAgent(c, o.TerraformVersion, o.PartnerID)

	c.Authorizer = authorizer
	c.Sender = sender.BuildSender("AzureAD")
}

func setUserAgent(client *autorest.Client, tfVersion, partnerID string) {
	tfUserAgent := fmt.Sprintf("HashiCorp Terraform/%s (+https://www.terraform.io) Terraform Plugin SDK/%s", tfVersion, meta.SDKVersionString())

	providerUserAgent := fmt.Sprintf("%s terraform-provider-azuread/%s", tfUserAgent, version.ProviderVersion)
	client.UserAgent = strings.TrimSpace(fmt.Sprintf("%s %s", client.UserAgent, providerUserAgent))

	// append the CloudShell version to the user agent if it exists
	if azureAgent := os.Getenv("AZURE_HTTP_USER_AGENT"); azureAgent != "" {
		client.UserAgent = fmt.Sprintf("%s %s", client.UserAgent, azureAgent)
	}

	if partnerID != "" {
		client.UserAgent = fmt.Sprintf("%s pid-%s", client.UserAgent, partnerID)
	}

	log.Printf("[DEBUG] AzureAD Client User Agent: %s\n", client.UserAgent)
}
