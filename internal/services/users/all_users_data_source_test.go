package users_test

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-provider-azuread/internal/acceptance"
	"github.com/hashicorp/terraform-provider-azuread/internal/acceptance/check"
	"testing"
)

type AllUsersDataSource struct{}

func TestAccAllUsersDataSource(t *testing.T) {
	data := acceptance.BuildTestData(t, "data.azuread_all_users", "test")
	r := AllUsersDataSource{}

	data.DataSourceTest(t, []resource.TestStep{{
		Check: r.testCheckFunc(data),
	}})
}

func (AllUsersDataSource) testCheckFunc(data acceptance.TestData) resource.TestCheckFunc {
	return resource.ComposeTestCheckFunc(
		check.That(data.ResourceName).Key("user_principal_names.#").Exists(),
		check.That(data.ResourceName).Key("object_ids.#").Exists(),
		check.That(data.ResourceName).Key("users.#").Exists(),
	)
}

