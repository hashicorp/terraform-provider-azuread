package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &IdentityGovernanceEntitlementManagementResourceRequestIdCatalogResourceIdScopeIdResourceRoleId{}

// IdentityGovernanceEntitlementManagementResourceRequestIdCatalogResourceIdScopeIdResourceRoleId is a struct representing the Resource ID for a Identity Governance Entitlement Management Resource Request Id Catalog Resource Id Scope Id Resource Role
type IdentityGovernanceEntitlementManagementResourceRequestIdCatalogResourceIdScopeIdResourceRoleId struct {
	AccessPackageResourceRequestId string
	AccessPackageResourceId        string
	AccessPackageResourceScopeId   string
	AccessPackageResourceRoleId    string
}

// NewIdentityGovernanceEntitlementManagementResourceRequestIdCatalogResourceIdScopeIdResourceRoleID returns a new IdentityGovernanceEntitlementManagementResourceRequestIdCatalogResourceIdScopeIdResourceRoleId struct
func NewIdentityGovernanceEntitlementManagementResourceRequestIdCatalogResourceIdScopeIdResourceRoleID(accessPackageResourceRequestId string, accessPackageResourceId string, accessPackageResourceScopeId string, accessPackageResourceRoleId string) IdentityGovernanceEntitlementManagementResourceRequestIdCatalogResourceIdScopeIdResourceRoleId {
	return IdentityGovernanceEntitlementManagementResourceRequestIdCatalogResourceIdScopeIdResourceRoleId{
		AccessPackageResourceRequestId: accessPackageResourceRequestId,
		AccessPackageResourceId:        accessPackageResourceId,
		AccessPackageResourceScopeId:   accessPackageResourceScopeId,
		AccessPackageResourceRoleId:    accessPackageResourceRoleId,
	}
}

// ParseIdentityGovernanceEntitlementManagementResourceRequestIdCatalogResourceIdScopeIdResourceRoleID parses 'input' into a IdentityGovernanceEntitlementManagementResourceRequestIdCatalogResourceIdScopeIdResourceRoleId
func ParseIdentityGovernanceEntitlementManagementResourceRequestIdCatalogResourceIdScopeIdResourceRoleID(input string) (*IdentityGovernanceEntitlementManagementResourceRequestIdCatalogResourceIdScopeIdResourceRoleId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityGovernanceEntitlementManagementResourceRequestIdCatalogResourceIdScopeIdResourceRoleId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityGovernanceEntitlementManagementResourceRequestIdCatalogResourceIdScopeIdResourceRoleId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseIdentityGovernanceEntitlementManagementResourceRequestIdCatalogResourceIdScopeIdResourceRoleIDInsensitively parses 'input' case-insensitively into a IdentityGovernanceEntitlementManagementResourceRequestIdCatalogResourceIdScopeIdResourceRoleId
// note: this method should only be used for API response data and not user input
func ParseIdentityGovernanceEntitlementManagementResourceRequestIdCatalogResourceIdScopeIdResourceRoleIDInsensitively(input string) (*IdentityGovernanceEntitlementManagementResourceRequestIdCatalogResourceIdScopeIdResourceRoleId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityGovernanceEntitlementManagementResourceRequestIdCatalogResourceIdScopeIdResourceRoleId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityGovernanceEntitlementManagementResourceRequestIdCatalogResourceIdScopeIdResourceRoleId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *IdentityGovernanceEntitlementManagementResourceRequestIdCatalogResourceIdScopeIdResourceRoleId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.AccessPackageResourceRequestId, ok = input.Parsed["accessPackageResourceRequestId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "accessPackageResourceRequestId", input)
	}

	if id.AccessPackageResourceId, ok = input.Parsed["accessPackageResourceId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "accessPackageResourceId", input)
	}

	if id.AccessPackageResourceScopeId, ok = input.Parsed["accessPackageResourceScopeId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "accessPackageResourceScopeId", input)
	}

	if id.AccessPackageResourceRoleId, ok = input.Parsed["accessPackageResourceRoleId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "accessPackageResourceRoleId", input)
	}

	return nil
}

// ValidateIdentityGovernanceEntitlementManagementResourceRequestIdCatalogResourceIdScopeIdResourceRoleID checks that 'input' can be parsed as a Identity Governance Entitlement Management Resource Request Id Catalog Resource Id Scope Id Resource Role ID
func ValidateIdentityGovernanceEntitlementManagementResourceRequestIdCatalogResourceIdScopeIdResourceRoleID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseIdentityGovernanceEntitlementManagementResourceRequestIdCatalogResourceIdScopeIdResourceRoleID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Identity Governance Entitlement Management Resource Request Id Catalog Resource Id Scope Id Resource Role ID
func (id IdentityGovernanceEntitlementManagementResourceRequestIdCatalogResourceIdScopeIdResourceRoleId) ID() string {
	fmtString := "/identityGovernance/entitlementManagement/resourceRequests/%s/catalog/resources/%s/scopes/%s/resource/roles/%s"
	return fmt.Sprintf(fmtString, id.AccessPackageResourceRequestId, id.AccessPackageResourceId, id.AccessPackageResourceScopeId, id.AccessPackageResourceRoleId)
}

// Segments returns a slice of Resource ID Segments which comprise this Identity Governance Entitlement Management Resource Request Id Catalog Resource Id Scope Id Resource Role ID
func (id IdentityGovernanceEntitlementManagementResourceRequestIdCatalogResourceIdScopeIdResourceRoleId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("identityGovernance", "identityGovernance", "identityGovernance"),
		resourceids.StaticSegment("entitlementManagement", "entitlementManagement", "entitlementManagement"),
		resourceids.StaticSegment("resourceRequests", "resourceRequests", "resourceRequests"),
		resourceids.UserSpecifiedSegment("accessPackageResourceRequestId", "accessPackageResourceRequestId"),
		resourceids.StaticSegment("catalog", "catalog", "catalog"),
		resourceids.StaticSegment("resources", "resources", "resources"),
		resourceids.UserSpecifiedSegment("accessPackageResourceId", "accessPackageResourceId"),
		resourceids.StaticSegment("scopes", "scopes", "scopes"),
		resourceids.UserSpecifiedSegment("accessPackageResourceScopeId", "accessPackageResourceScopeId"),
		resourceids.StaticSegment("resource", "resource", "resource"),
		resourceids.StaticSegment("roles", "roles", "roles"),
		resourceids.UserSpecifiedSegment("accessPackageResourceRoleId", "accessPackageResourceRoleId"),
	}
}

// String returns a human-readable description of this Identity Governance Entitlement Management Resource Request Id Catalog Resource Id Scope Id Resource Role ID
func (id IdentityGovernanceEntitlementManagementResourceRequestIdCatalogResourceIdScopeIdResourceRoleId) String() string {
	components := []string{
		fmt.Sprintf("Access Package Resource Request: %q", id.AccessPackageResourceRequestId),
		fmt.Sprintf("Access Package Resource: %q", id.AccessPackageResourceId),
		fmt.Sprintf("Access Package Resource Scope: %q", id.AccessPackageResourceScopeId),
		fmt.Sprintf("Access Package Resource Role: %q", id.AccessPackageResourceRoleId),
	}
	return fmt.Sprintf("Identity Governance Entitlement Management Resource Request Id Catalog Resource Id Scope Id Resource Role (%s)", strings.Join(components, "\n"))
}
