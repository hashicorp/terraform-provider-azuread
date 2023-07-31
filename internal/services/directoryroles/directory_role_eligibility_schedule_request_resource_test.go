package directoryroles_test

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/hashicorp/go-azure-sdk/sdk/odata"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/hashicorp/terraform-provider-azuread/internal/acceptance"
	"github.com/hashicorp/terraform-provider-azuread/internal/acceptance/check"
	"github.com/hashicorp/terraform-provider-azuread/internal/clients"
	"github.com/hashicorp/terraform-provider-azuread/internal/utils"
)

type RoleEligibilityScheduleRequestResource struct{}

func TestAccRoleEligibilityScheduleRequest_basic(t *testing.T) {
	data := acceptance.BuildTestData(t, "azuread_directory_role_eligibility_schedule_request", "test")
	r := RoleEligibilityScheduleRequestResource{}

	data.ResourceTestIgnoreDangling(t, r, []resource.TestStep{
		{
			Config: r.basic(data),
			Check: resource.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
	})
}

func (r RoleEligibilityScheduleRequestResource) Exists(ctx context.Context, clients *clients.Client, state *terraform.InstanceState) (*bool, error) {
	client := clients.DirectoryRoles.RoleEligibilityScheduleRequestClient

	resr, status, err := client.Get(ctx, state.ID, odata.Query{})
	if err != nil {
		fmt.Printf("%s, %v\n", err.Error(), status)
		if status == http.StatusNotFound {
			return nil, fmt.Errorf("Role Eligibility Schedule Request with ID %q does not exist", state.ID)
		}
		return nil, fmt.Errorf("failed to retrieve Role Eligibility Schedule Request with object ID %q: %+v", state.ID, err)
	}

	return utils.Bool(resr.ID != nil && *resr.ID == state.ID), nil
}

func (r RoleEligibilityScheduleRequestResource) basic(data acceptance.TestData) string {
	return fmt.Sprintf(`
provider "azuread" {}

data "azuread_domains" "test" {
  only_initial = true
}

resource "azuread_user" "test" {
  user_principal_name = "acctestManager.%[1]d@${data.azuread_domains.test.domains.0.domain_name}"
  display_name        = "acctestManager-%[1]d"
  password            = "%[2]s"
}

resource "azuread_directory_role" "test" {
  display_name = "Application Administrator"
}

resource "azuread_directory_role_eligibility_schedule_request" "test" {
  role_definition_id = azuread_directory_role.test.template_id
  principal_id       = azuread_user.test.object_id
  directory_scope_id = "/"
  justification      = "abc"
}
`, data.RandomInteger, data.RandomPassword)
}
