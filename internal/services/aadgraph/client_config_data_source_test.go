package aadgraph_test

import (
	"github.com/terraform-providers/terraform-provider-azuread/internal/acceptance/check"
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/terraform-providers/terraform-provider-azuread/internal/acceptance"
)

type ClientConfigDataSource struct{}

func TestAccClientConfigDataSource_basic(t *testing.T) {
	data := acceptance.BuildTestData(t, "data.azuread_client_config", "test")
	data.AdditionalData["client_id"] = os.Getenv("ARM_CLIENT_ID")
	data.AdditionalData["tenant_id"] = os.Getenv("ARM_TENANT_ID")

	data.DataSourceTest(t, []resource.TestStep{
		{
			Config: ClientConfigDataSource{}.basic(),
			Check: resource.ComposeTestCheckFunc(
				check.That(data.ResourceName).Key("client_id").HasValue(data.AdditionalData["client_id"].(string)),
				check.That(data.ResourceName).Key("tenant_id").HasValue(data.AdditionalData["tenant_id"].(string)),
				check.That(data.ResourceName).Key("object_id").IsGuid(),
			),
		},
	})
}

func (ClientConfigDataSource) basic() string {
	return `data "azuread_client_config" "test" {}`
}
