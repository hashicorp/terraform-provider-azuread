package check

import (
	"context"
	"fmt"
	"regexp"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	"github.com/hashicorp/terraform-provider-azuread/internal/acceptance"
	"github.com/hashicorp/terraform-provider-azuread/internal/acceptance/helpers"
	"github.com/hashicorp/terraform-provider-azuread/internal/acceptance/types"
	"github.com/hashicorp/terraform-provider-azuread/internal/clients"
)

type thatType struct {
	// resourceName being the full resource name e.g. azurerm_foo.bar
	resourceName string
}

// Key returns a type which can be used for more fluent assertions for a given Resource
func That(resourceName string) thatType {
	return thatType{
		resourceName: resourceName,
	}
}

// ExistsInAzure validates that the specified resource exists within Azure
func (t thatType) ExistsInAzure(testResource types.TestResource) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client := acceptance.AzureADProvider.Meta().(*clients.Client)
		return helpers.ExistsInAzure(client, testResource, t.resourceName)(s)
	}
}

// Key returns a type which can be used for more fluent assertions for a given Resource & Key combination
func (t thatType) Key(key string) thatWithKeyType {
	return thatWithKeyType{
		resourceName: t.resourceName,
		key:          key,
	}
}

type thatWithKeyType struct {
	// resourceName being the full resource name e.g. azurerm_foo.bar
	resourceName string

	// key being the specific field we're querying e.g. bar or a nested object ala foo.0.bar
	key string
}

// DoesNotExist returns a TestCheckFunc which validates that the specific key
// does not exist on the resource
func (t thatWithKeyType) DoesNotExist() resource.TestCheckFunc {
	return resource.TestCheckNoResourceAttr(t.resourceName, t.key)
}

// Exists returns a TestCheckFunc which validates that the specific key exists on the resource
func (t thatWithKeyType) Exists() resource.TestCheckFunc {
	return resource.TestCheckResourceAttrSet(t.resourceName, t.key)
}

// IsEmpty returns a TestCheckFunc which validates that the specific key is empty on the resource
func (t thatWithKeyType) IsEmpty() resource.TestCheckFunc {
	return resource.TestCheckResourceAttr(t.resourceName, t.key, "")
}

// IsUuid returns a TestCheckFunc which validates that the specific key value is a valid UUID
func (t thatWithKeyType) IsUuid() resource.TestCheckFunc {
	r, _ := regexp.Compile(`^[A-Fa-f0-9]{8}-[A-Fa-f0-9]{4}-[A-Fa-f0-9]{4}-[A-Fa-f0-9]{4}-[A-Fa-f0-9]{12}$`)
	return t.MatchesRegex(r)
}

// HasValue returns a TestCheckFunc which validates that the specific key has the
// specified value on the resource
func (t thatWithKeyType) HasValue(value string) resource.TestCheckFunc {
	return resource.TestCheckResourceAttr(t.resourceName, t.key, value)
}

// MatchesOtherKey returns a TestCheckFunc which validates that the key on this resource
// matches another other key on another resource
func (t thatWithKeyType) MatchesOtherKey(other thatWithKeyType) resource.TestCheckFunc {
	return resource.TestCheckResourceAttrPair(t.resourceName, t.key, other.resourceName, other.key)
}

// MatchesRegex returns a TestCheckFunc which validates that the key on this resource matches
// the given regular expression
func (t thatWithKeyType) MatchesRegex(r *regexp.Regexp) resource.TestCheckFunc {
	return resource.TestMatchResourceAttr(t.resourceName, t.key, r)
}

func (t thatWithKeyType) ValidatesWith(validationFunc KeyValidationFunc) resource.TestCheckFunc {
	return func(state *terraform.State) error {
		ms := state.RootModule()
		rs, ok := ms.Resources[t.resourceName]
		if !ok {
			return fmt.Errorf("Not found: %s in %s", t.resourceName, ms.Path)
		}
		is := rs.Primary
		if is == nil {
			return fmt.Errorf("No primary instance: %s in %s", t.resourceName, ms.Path)
		}

		var values []interface{}
		for attr, val := range is.Attributes {
			if attrParts := strings.Split(attr, "."); len(attrParts) == 2 && attrParts[0] == t.key && attrParts[1] != "#" && attrParts[1] != "%" {
				values = append(values, val)
			}
		}

		clients := acceptance.AzureADProvider.Meta().(*clients.Client)
		return validationFunc(clients.StopContext, clients, values)
	}
}

type KeyValidationFunc func(context.Context, *clients.Client, []interface{}) error
