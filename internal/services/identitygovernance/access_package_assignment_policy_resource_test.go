package identitygovernance_test

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/hashicorp/terraform-provider-azuread/internal/acceptance"
	"github.com/hashicorp/terraform-provider-azuread/internal/acceptance/check"
	"github.com/hashicorp/terraform-provider-azuread/internal/clients"
	"github.com/hashicorp/terraform-provider-azuread/internal/utils"
	"github.com/manicminer/hamilton/odata"
)

type AccessPackageAssignmentPolicyResource struct{}

func TestAccAccessPackageAssignmentPolicy_simple(t *testing.T) {
	data := acceptance.BuildTestData(t, "azuread_access_package_assignment_policy", "test")
	r := AccessPackageAssignmentPolicyResource{}

	data.ResourceTest(t, r, []resource.TestStep{
		{
			Config: r.simple(data),
			Check: resource.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		data.ImportStep("access_package_id"),
	})
}

func TestAccAccessPackageAssignmentPolicy_basic(t *testing.T) {
	data := acceptance.BuildTestData(t, "azuread_access_package_assignment_policy", "test")
	r := AccessPackageAssignmentPolicyResource{}

	data.ResourceTest(t, r, []resource.TestStep{
		{
			Config: r.basic(data),
			Check: resource.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		data.ImportStep("access_package_id"),
	})
}

func TestAccAccessPackageAssignmentPolicy_complete(t *testing.T) {
	data := acceptance.BuildTestData(t, "azuread_access_package_assignment_policy", "test")
	r := AccessPackageAssignmentPolicyResource{}

	data.ResourceTest(t, r, []resource.TestStep{
		{
			Config: r.complete(data),
			Check: resource.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		data.ImportStep("access_package_id"),
	})
}

func TestAccAccessPackageAssignmentPolicy_update(t *testing.T) {
	data := acceptance.BuildTestData(t, "azuread_access_package_assignment_policy", "test")
	r := AccessPackageAssignmentPolicyResource{}

	data.ResourceTest(t, r, []resource.TestStep{
		{
			Config: r.basic(data),
			Check: resource.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		data.ImportStep(),
		{
			Config: r.complete(data),
			Check: resource.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		data.ImportStep(),
		{
			Config: r.basic(data),
			Check: resource.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		data.ImportStep(),
	})
}

func (AccessPackageAssignmentPolicyResource) Exists(ctx context.Context, clients *clients.Client, state *terraform.InstanceState) (*bool, error) {
	client := clients.IdentityGovernance.AccessPackageAssignmentPolicyClient
	client.BaseClient.DisableRetries = true

	accessPackageAssignmentPolicy, status, err := client.Get(ctx, state.ID, odata.Query{})
	if err != nil {
		if status == http.StatusNotFound {
			return nil, fmt.Errorf("Access package assignment policy with object ID %q does not exist", state.ID)
		}
		return nil, fmt.Errorf("failed to retrieve Access package assignment policy with object ID %q: %+v", state.ID, err)
	}
	return utils.Bool(accessPackageAssignmentPolicy.ID != nil && *accessPackageAssignmentPolicy.ID == state.ID), nil
}

func (AccessPackageAssignmentPolicyResource) simple(data acceptance.TestData) string {
	return fmt.Sprintf(`
provider "azuread" {}

resource "azuread_access_package_catalog" "test_catalog" {
  display_name = "test-catalog-%[1]d"
  description  = "Test Catalog %[1]d"
}

resource "azuread_access_package" "test" {
  display_name = "access-package-%[1]d"
  description  = "Test Access Package %[1]d"
  catalog_id   = azuread_access_package_catalog.test_catalog.id
}

resource "azuread_access_package_assignment_policy" "test" {
  display_name      = "access-package-assignment-policy-%[1]d"
  description       = "Test Access Package Assignnment Policy %[1]d"
  access_package_id = azuread_access_package.test.id
}
`, data.RandomInteger)
}

func (AccessPackageAssignmentPolicyResource) basic(data acceptance.TestData) string {
	return fmt.Sprintf(`
provider "azuread" {}

resource "azuread_group" "test" {
  display_name     = "test-group-%[1]d"
  security_enabled = true
}

resource "azuread_access_package_catalog" "test_catalog" {
  display_name = "testacc-asscess-assignment-%[1]d"
  description  = "TestAcc Catalog %[1]d for access assignment policy"
}

resource "azuread_access_package" "test" {
  display_name = "testacc-asscess-assignment-%[1]d"
  description  = "TestAcc Access Package %[1]d for access assignment policy"
  catalog_id   = azuread_access_package_catalog.test_catalog.id
}

resource "azuread_access_package_assignment_policy" "test" {
  display_name      = "testacc-asscess-assignment-%[1]d"
  description       = "TestAcc Access Package Assignnment Policy %[1]d"
  duration_in_days  = 90
  access_package_id = azuread_access_package.test.id
  requestor_settings {
    scope_type = "AllExistingDirectoryMemberUsers"
  }
  approval_settings {
    is_approval_required     = true
    approval_stage {
	  approval_timeout_in_days = 14
        primary_approver {
          object_id                = azuread_group.test.object_id
		  subject_type             = "groupMembers"
      }
    }
  }
  assignment_review_settings {
    is_enabled                     = true
	review_frequency               = "weekly"
	duration_in_days               = 3
    review_type      			   = "Self"
	access_review_timeout_behavior = "keepAccess"	
  }
  question {
	text {
		default_text = "hello, how are you?"
	}	
  }
}
`, data.RandomInteger)
}

func (AccessPackageAssignmentPolicyResource) complete(data acceptance.TestData) string {
	return fmt.Sprintf(`
provider "azuread" {}

resource "azuread_group" "requestor" {
  display_name     = "test-requestor-%[1]d"
  security_enabled = true
}

resource "azuread_group" "first_approver" {
  display_name     = "test-approver-%[1]d"
  security_enabled = true
}

resource "azuread_group" "second_approver" {
  display_name     = "test-s-approver-%[1]d"
  security_enabled = true
}

resource "azuread_access_package_catalog" "test_catalog" {
  display_name = "testacc-asscess-assignment-%[1]d"
  description  = "TestAcc Catalog %[1]d for access assignment policy"
}

resource "azuread_access_package" "test" {
  display_name = "testacc-asscess-assignment-%[1]d"
  description  = "Test Access Package %[1]d for assignment policy"
  catalog_id   = azuread_access_package_catalog.test_catalog.id
}

resource "azuread_access_package_assignment_policy" "test" {
  display_name      = "access-package-assignment-policy-%[1]d"
  description       = "Test Access Package Assignnment Policy %[1]d"
  can_extend        = true
  expiration_date   = "2096-09-23T01:02:03Z"
  access_package_id = azuread_access_package.test.id
  requestor_settings {
    scope_type      = "SpecificDirectorySubjects"
    accept_requests = true
    requestor {
      object_id    = azuread_group.requestor.object_id
	  subject_type = "groupMembers"
    }
  }
  approval_settings {
    is_approval_required                = true
    is_approval_required_for_extension  = true
    is_requestor_justification_required = true
    approval_stage {
      approval_timeout_in_days            = 14
      is_approver_justification_required  = true
      is_alternative_approval_enabled     = true
      enable_alternative_approval_in_days = 8
      primary_approver {
		subject_type = "requestorManager"	
      }
      alternative_approver {
        object_id    = azuread_group.second_approver.object_id
		subject_type = "groupMembers"	
      }
    }

    approval_stage {
      approval_timeout_in_days            = 14
      primary_approver {
        object_id    = azuread_group.second_approver.object_id
		subject_type = "groupMembers"
      }
      primary_approver {
        object_id    = azuread_group.first_approver.object_id
		subject_type = "groupMembers"
		is_backup    = true
      }
    }
  }
  assignment_review_settings {
    is_enabled                         = true
    review_frequency                   = "annual"
    review_type                        = "Reviewers"
	duration_in_days                   = "10"
	is_access_recommendation_enabled   = true
	access_review_timeout_behavior     = "acceptAccessRecommendation"
	reviewer {
	  object_id    = azuread_group.first_approver.object_id
	  subject_type = "groupMembers"
    }
  }

  question {
    is_required             = true
    sequence                = 1
    text {
      default_text = "Hello Why"
      localized_text {
        language_code = "CN"
        content       = "Hello why CN?"
      }
      localized_text {
        language_code = "FR"
        content       = "Hello why BE?"
      }
    }
  }

  question {
    is_required             = false
    sequence                = 2
	choice {
		actual_value = "a"
		display_value {
			default_text = "AA"
			localized_text {
				language_code = "CN"
			    content       = "AAB"
			}
		}
	}
    text {
      default_text = "Hello Why again"
    }
  }
}

`, data.RandomInteger)
}
