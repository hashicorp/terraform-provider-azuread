package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &IdentityGovernanceEntitlementManagementResourceRequestIdCatalogResourceScopeIdResourceRoleId{}

// IdentityGovernanceEntitlementManagementResourceRequestIdCatalogResourceScopeIdResourceRoleId is a struct representing the Resource ID for a Identity Governance Entitlement Management Resource Request Id Catalog Resource Scope Id Resource Role
type IdentityGovernanceEntitlementManagementResourceRequestIdCatalogResourceScopeIdResourceRoleId struct {
	AccessPackageResourceRequestId string
	AccessPackageResourceScopeId   string
	AccessPackageResourceRoleId    string
}

// NewIdentityGovernanceEntitlementManagementResourceRequestIdCatalogResourceScopeIdResourceRoleID returns a new IdentityGovernanceEntitlementManagementResourceRequestIdCatalogResourceScopeIdResourceRoleId struct
func NewIdentityGovernanceEntitlementManagementResourceRequestIdCatalogResourceScopeIdResourceRoleID(accessPackageResourceRequestId string, accessPackageResourceScopeId string, accessPackageResourceRoleId string) IdentityGovernanceEntitlementManagementResourceRequestIdCatalogResourceScopeIdResourceRoleId {
	return IdentityGovernanceEntitlementManagementResourceRequestIdCatalogResourceScopeIdResourceRoleId{
		AccessPackageResourceRequestId: accessPackageResourceRequestId,
		AccessPackageResourceScopeId:   accessPackageResourceScopeId,
		AccessPackageResourceRoleId:    accessPackageResourceRoleId,
	}
}

// ParseIdentityGovernanceEntitlementManagementResourceRequestIdCatalogResourceScopeIdResourceRoleID parses 'input' into a IdentityGovernanceEntitlementManagementResourceRequestIdCatalogResourceScopeIdResourceRoleId
func ParseIdentityGovernanceEntitlementManagementResourceRequestIdCatalogResourceScopeIdResourceRoleID(input string) (*IdentityGovernanceEntitlementManagementResourceRequestIdCatalogResourceScopeIdResourceRoleId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityGovernanceEntitlementManagementResourceRequestIdCatalogResourceScopeIdResourceRoleId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityGovernanceEntitlementManagementResourceRequestIdCatalogResourceScopeIdResourceRoleId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseIdentityGovernanceEntitlementManagementResourceRequestIdCatalogResourceScopeIdResourceRoleIDInsensitively parses 'input' case-insensitively into a IdentityGovernanceEntitlementManagementResourceRequestIdCatalogResourceScopeIdResourceRoleId
// note: this method should only be used for API response data and not user input
func ParseIdentityGovernanceEntitlementManagementResourceRequestIdCatalogResourceScopeIdResourceRoleIDInsensitively(input string) (*IdentityGovernanceEntitlementManagementResourceRequestIdCatalogResourceScopeIdResourceRoleId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityGovernanceEntitlementManagementResourceRequestIdCatalogResourceScopeIdResourceRoleId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityGovernanceEntitlementManagementResourceRequestIdCatalogResourceScopeIdResourceRoleId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *IdentityGovernanceEntitlementManagementResourceRequestIdCatalogResourceScopeIdResourceRoleId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.AccessPackageResourceRequestId, ok = input.Parsed["accessPackageResourceRequestId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "accessPackageResourceRequestId", input)
	}

	if id.AccessPackageResourceScopeId, ok = input.Parsed["accessPackageResourceScopeId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "accessPackageResourceScopeId", input)
	}

	if id.AccessPackageResourceRoleId, ok = input.Parsed["accessPackageResourceRoleId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "accessPackageResourceRoleId", input)
	}

	return nil
}

// ValidateIdentityGovernanceEntitlementManagementResourceRequestIdCatalogResourceScopeIdResourceRoleID checks that 'input' can be parsed as a Identity Governance Entitlement Management Resource Request Id Catalog Resource Scope Id Resource Role ID
func ValidateIdentityGovernanceEntitlementManagementResourceRequestIdCatalogResourceScopeIdResourceRoleID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseIdentityGovernanceEntitlementManagementResourceRequestIdCatalogResourceScopeIdResourceRoleID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Identity Governance Entitlement Management Resource Request Id Catalog Resource Scope Id Resource Role ID
func (id IdentityGovernanceEntitlementManagementResourceRequestIdCatalogResourceScopeIdResourceRoleId) ID() string {
	fmtString := "/identityGovernance/entitlementManagement/resourceRequests/%s/catalog/resourceScopes/%s/resource/roles/%s"
	return fmt.Sprintf(fmtString, id.AccessPackageResourceRequestId, id.AccessPackageResourceScopeId, id.AccessPackageResourceRoleId)
}

// Segments returns a slice of Resource ID Segments which comprise this Identity Governance Entitlement Management Resource Request Id Catalog Resource Scope Id Resource Role ID
func (id IdentityGovernanceEntitlementManagementResourceRequestIdCatalogResourceScopeIdResourceRoleId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("identityGovernance", "identityGovernance", "identityGovernance"),
		resourceids.StaticSegment("entitlementManagement", "entitlementManagement", "entitlementManagement"),
		resourceids.StaticSegment("resourceRequests", "resourceRequests", "resourceRequests"),
		resourceids.UserSpecifiedSegment("accessPackageResourceRequestId", "accessPackageResourceRequestId"),
		resourceids.StaticSegment("catalog", "catalog", "catalog"),
		resourceids.StaticSegment("resourceScopes", "resourceScopes", "resourceScopes"),
		resourceids.UserSpecifiedSegment("accessPackageResourceScopeId", "accessPackageResourceScopeId"),
		resourceids.StaticSegment("resource", "resource", "resource"),
		resourceids.StaticSegment("roles", "roles", "roles"),
		resourceids.UserSpecifiedSegment("accessPackageResourceRoleId", "accessPackageResourceRoleId"),
	}
}

// String returns a human-readable description of this Identity Governance Entitlement Management Resource Request Id Catalog Resource Scope Id Resource Role ID
func (id IdentityGovernanceEntitlementManagementResourceRequestIdCatalogResourceScopeIdResourceRoleId) String() string {
	components := []string{
		fmt.Sprintf("Access Package Resource Request: %q", id.AccessPackageResourceRequestId),
		fmt.Sprintf("Access Package Resource Scope: %q", id.AccessPackageResourceScopeId),
		fmt.Sprintf("Access Package Resource Role: %q", id.AccessPackageResourceRoleId),
	}
	return fmt.Sprintf("Identity Governance Entitlement Management Resource Request Id Catalog Resource Scope Id Resource Role (%s)", strings.Join(components, "\n"))
}
