// Copyright IBM Corp. 2014, 2025
// SPDX-License-Identifier: MPL-2.0

package check

import (
	"context"
	"fmt"
	"regexp"
	"strings"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/hashicorp/terraform-provider-azuread/internal/acceptance/helpers"
	"github.com/hashicorp/terraform-provider-azuread/internal/acceptance/testclient"
	"github.com/hashicorp/terraform-provider-azuread/internal/acceptance/types"
	"github.com/hashicorp/terraform-provider-azuread/internal/clients"
	"github.com/hashicorp/terraform-provider-azuread/internal/helpers/tf/pluginsdk"
)

type withTenantType struct {
	tenantId string
}

func (w withTenantType) That(resourceName string) thatType {
	that := That(resourceName)
	that.tenantId = w.tenantId
	return that
}

func WithTenant(tenantId string) withTenantType {
	return withTenantType{
		tenantId: tenantId,
	}
}

type thatType struct {
	// resourceName being the full resource name e.g. azurerm_foo.bar
	resourceName string

	// tenantId is the tenant to use when building the test client
	tenantId string
}

func That(resourceName string) thatType {
	return thatType{
		resourceName: resourceName,
	}
}

// DoesNotExistInAzure validates that the specified resource does not exist within Azure
func (t thatType) DoesNotExistInAzure(testResource types.TestResource) pluginsdk.TestCheckFunc {
	return func(s *terraform.State) error {
		client, err := testclient.Build(t.tenantId)
		if err != nil {
			return fmt.Errorf("building client: %+v", err)
		}
		return helpers.DoesNotExistInAzure(client, testResource, t.resourceName)(s)
	}
}

// ExistsInAzure validates that the specified resource exists within Azure
func (t thatType) ExistsInAzure(testResource types.TestResource) pluginsdk.TestCheckFunc {
	return func(s *terraform.State) error {
		client, err := testclient.Build(t.tenantId)
		if err != nil {
			return fmt.Errorf("building client: %+v", err)
		}
		return helpers.ExistsInAzure(client, testResource, t.resourceName)(s)
	}
}

// Key returns a type which can be used for more fluent assertions for a given Resource & Key combination
func (t thatType) Key(key string) thatWithKeyType {
	return thatWithKeyType{
		resourceName: t.resourceName,
		key:          key,
		tenantId:     t.tenantId,
	}
}

type thatWithKeyType struct {
	// resourceName being the full resource name e.g. azurerm_foo.bar
	resourceName string

	// key being the specific field we're querying e.g. bar or a nested object ala foo.0.bar
	key string

	// tenantId is the tenant to use when building the test client. When blank, the env var ARM_TENANT_ID is used.
	tenantId string
}

// DoesNotExist returns a TestCheckFunc which validates that the specific key
// does not exist on the resource
func (t thatWithKeyType) DoesNotExist() pluginsdk.TestCheckFunc {
	return resource.TestCheckNoResourceAttr(t.resourceName, t.key)
}

// Exists returns a TestCheckFunc which validates that the specific key exists on the resource
func (t thatWithKeyType) Exists() pluginsdk.TestCheckFunc {
	return resource.TestCheckResourceAttrSet(t.resourceName, t.key)
}

// IsEmpty returns a TestCheckFunc which validates that the specific key is empty on the resource
func (t thatWithKeyType) IsEmpty() pluginsdk.TestCheckFunc {
	return resource.TestCheckResourceAttr(t.resourceName, t.key, "")
}

// IsUuid returns a TestCheckFunc which validates that the specific key value is a valid UUID
func (t thatWithKeyType) IsUuid() pluginsdk.TestCheckFunc {
	return t.MatchesRegex(regexp.MustCompile(`^[A-Fa-f0-9]{8}-[A-Fa-f0-9]{4}-[A-Fa-f0-9]{4}-[A-Fa-f0-9]{4}-[A-Fa-f0-9]{12}$`))
}

// HasValue returns a TestCheckFunc which validates that the specific key has the
// specified value on the resource
func (t thatWithKeyType) HasValue(value string) pluginsdk.TestCheckFunc {
	return resource.TestCheckResourceAttr(t.resourceName, t.key, value)
}

// MatchesOtherKey returns a TestCheckFunc which validates that the key on this resource
// matches another other key on another resource
func (t thatWithKeyType) MatchesOtherKey(other thatWithKeyType) pluginsdk.TestCheckFunc {
	return resource.TestCheckResourceAttrPair(t.resourceName, t.key, other.resourceName, other.key)
}

// MatchesRegex returns a TestCheckFunc which validates that the key on this resource matches
// the given regular expression
func (t thatWithKeyType) MatchesRegex(r *regexp.Regexp) pluginsdk.TestCheckFunc {
	return resource.TestMatchResourceAttr(t.resourceName, t.key, r)
}

func (t thatWithKeyType) ValidatesWith(validationFunc KeyValidationFunc) pluginsdk.TestCheckFunc {
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

		client, err := testclient.Build(t.tenantId)
		if err != nil {
			return fmt.Errorf("building client: %+v", err)
		}

		return validationFunc(client.StopContext, client, values)
	}
}

type KeyValidationFunc func(context.Context, *clients.Client, []interface{}) error
