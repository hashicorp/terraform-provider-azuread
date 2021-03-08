package serviceprincipals_test

import (
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	"github.com/hashicorp/terraform-provider-azuread/internal/acceptance"
	"github.com/hashicorp/terraform-provider-azuread/internal/acceptance/check"
)

type ClientConfigDataSource struct{}

func TestAccClientConfigDataSource_basic(t *testing.T) {
	data := acceptance.BuildTestData(t, "data.azuread_client_config", "test")
	clientId := os.Getenv("ARM_CLIENT_ID")
	tenantId := os.Getenv("ARM_TENANT_ID")

	data.DataSourceTest(t, []resource.TestStep{
		{
			Config: ClientConfigDataSource{}.basic(),
			Check: resource.ComposeTestCheckFunc(
				check.That(data.ResourceName).Key("client_id").HasValue(clientId),
				check.That(data.ResourceName).Key("tenant_id").HasValue(tenantId),
				check.That(data.ResourceName).Key("object_id").IsUuid(),
			),
		},
	})
}

func (ClientConfigDataSource) basic() string {
	return `data "azuread_client_config" "test" {}`
}
