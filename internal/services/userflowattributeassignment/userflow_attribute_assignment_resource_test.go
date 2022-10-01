package userflowattributeassignment_test

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	_ "net/http/pprof"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/hashicorp/terraform-provider-azuread/internal/acceptance"
	"github.com/hashicorp/terraform-provider-azuread/internal/acceptance/check"
	"github.com/hashicorp/terraform-provider-azuread/internal/clients"
	"github.com/hashicorp/terraform-provider-azuread/internal/utils"
	"github.com/manicminer/hamilton/odata"
)

type AttributeAssignmentResouce struct{}

func TestAccAttributeAssignment_basic(t *testing.T) {

	go func() {
		if err := http.ListenAndServe(":8080", nil); err != nil {
			panic(err)
		}
	}()

	data := acceptance.BuildTestData(t, "azuread_userflow_attribute_assignment", "test")
	r := AttributeAssignmentResouce{}

	userflowId := "B2C_1_test2"
	attrId := "extension_95e6fbf115424a97af20e2237b0d7d84_ibrahimAttribute"
	data.ResourceTest(t, r, []resource.TestStep{
		{
			Config: r.basic(data, userflowId, attrId),
			Check: resource.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
	})
}

// func TestAccAttributeAssignment_complete(t *testing.T) {
// 	data := acceptance.BuildTestData(t, "azuread_userflow_attribute_assignment", "test")
// 	r := AttributeAssignmentResouce{}

// 	data.ResourceTest(t, r, []resource.TestStep{
// 		{
// 			Config: r.complete(data),
// 			Check: resource.ComposeTestCheckFunc(
// 				check.That(data.ResourceName).ExistsInAzure(r),
// 			),
// 		},
// 	})
// }

// func TestAccAttributeAssignment_update(t *testing.T) {
// 	data := acceptance.BuildTestData(t, "azuread_userflow_attribute_assignment", "test")
// 	r := AttributeAssignmentResouce{}

// 	data.ResourceTest(t, r, []resource.TestStep{
// 		{
// 			Config: r.basic(data),
// 			Check: resource.ComposeTestCheckFunc(
// 				check.That(data.ResourceName).ExistsInAzure(r),
// 			),
// 		},
// 		{
// 			Config: r.complete(data),
// 			Check: resource.ComposeTestCheckFunc(
// 				check.That(data.ResourceName).ExistsInAzure(r),
// 			),
// 		},
// 	})
// }

func (r AttributeAssignmentResouce) Exists(ctx context.Context, clients *clients.Client, state *terraform.InstanceState) (*bool, error) {
	client := clients.B2CUserFlow.UserFlowClient
	client.BaseClient.DisableRetries = true

	userflowID := state.Attributes["userflow_id"]
	attrAssignment, status, err := client.GetAssignedAttribute(ctx, userflowID, state.ID, odata.Query{})
	if err != nil {
		if status == http.StatusNotFound {
			return nil, fmt.Errorf("Assignment with object ID %q does not exist", state.ID)
		}
		return nil, fmt.Errorf("failed to retrieve Assignment with object ID %q: %+v", state.ID, err)
	}
	return utils.Bool(attrAssignment.ID != nil && *attrAssignment.ID == state.ID), nil
}

func (r AttributeAssignmentResouce) basic(data acceptance.TestData, userflowId, attrId string) string {
	return fmt.Sprintf(`
provider "azuread" {}

resource "azuread_userflow_attribute_assignment" "test" {
  display_name = "acctest-attribute-assignment-%[1]s"
  user_input_type = "textBox"
  userflow_attribute_id = %q
  userflow_id = %q
}`, data.RandomString, attrId, userflowId)
}

func (r AttributeAssignmentResouce) complete(data acceptance.TestData, userflowId, attrId string) string {
	return fmt.Sprintf(`
provider "azuread" {}

resource "azuread_userflow_attribute_assignment" "test" {
  display_name = "%sacctest-userflow"
  is_optional = true
  requires_verification = false
  user_attribute_values = ["values1", "values2"]
  user_input_type = "radioSingleSelect"
  userflow_attribute_id = %q
  userflow_id = %q
}`, data.RandomString, attrId, userflowId)
}
