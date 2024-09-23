package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &IdentityGovernanceEntitlementManagementResourceRequestIdCatalogResourceRoleIdResourceScopeId{}

// IdentityGovernanceEntitlementManagementResourceRequestIdCatalogResourceRoleIdResourceScopeId is a struct representing the Resource ID for a Identity Governance Entitlement Management Resource Request Id Catalog Resource Role Id Resource Scope
type IdentityGovernanceEntitlementManagementResourceRequestIdCatalogResourceRoleIdResourceScopeId struct {
	AccessPackageResourceRequestId string
	AccessPackageResourceRoleId    string
	AccessPackageResourceScopeId   string
}

// NewIdentityGovernanceEntitlementManagementResourceRequestIdCatalogResourceRoleIdResourceScopeID returns a new IdentityGovernanceEntitlementManagementResourceRequestIdCatalogResourceRoleIdResourceScopeId struct
func NewIdentityGovernanceEntitlementManagementResourceRequestIdCatalogResourceRoleIdResourceScopeID(accessPackageResourceRequestId string, accessPackageResourceRoleId string, accessPackageResourceScopeId string) IdentityGovernanceEntitlementManagementResourceRequestIdCatalogResourceRoleIdResourceScopeId {
	return IdentityGovernanceEntitlementManagementResourceRequestIdCatalogResourceRoleIdResourceScopeId{
		AccessPackageResourceRequestId: accessPackageResourceRequestId,
		AccessPackageResourceRoleId:    accessPackageResourceRoleId,
		AccessPackageResourceScopeId:   accessPackageResourceScopeId,
	}
}

// ParseIdentityGovernanceEntitlementManagementResourceRequestIdCatalogResourceRoleIdResourceScopeID parses 'input' into a IdentityGovernanceEntitlementManagementResourceRequestIdCatalogResourceRoleIdResourceScopeId
func ParseIdentityGovernanceEntitlementManagementResourceRequestIdCatalogResourceRoleIdResourceScopeID(input string) (*IdentityGovernanceEntitlementManagementResourceRequestIdCatalogResourceRoleIdResourceScopeId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityGovernanceEntitlementManagementResourceRequestIdCatalogResourceRoleIdResourceScopeId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityGovernanceEntitlementManagementResourceRequestIdCatalogResourceRoleIdResourceScopeId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseIdentityGovernanceEntitlementManagementResourceRequestIdCatalogResourceRoleIdResourceScopeIDInsensitively parses 'input' case-insensitively into a IdentityGovernanceEntitlementManagementResourceRequestIdCatalogResourceRoleIdResourceScopeId
// note: this method should only be used for API response data and not user input
func ParseIdentityGovernanceEntitlementManagementResourceRequestIdCatalogResourceRoleIdResourceScopeIDInsensitively(input string) (*IdentityGovernanceEntitlementManagementResourceRequestIdCatalogResourceRoleIdResourceScopeId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityGovernanceEntitlementManagementResourceRequestIdCatalogResourceRoleIdResourceScopeId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityGovernanceEntitlementManagementResourceRequestIdCatalogResourceRoleIdResourceScopeId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *IdentityGovernanceEntitlementManagementResourceRequestIdCatalogResourceRoleIdResourceScopeId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.AccessPackageResourceRequestId, ok = input.Parsed["accessPackageResourceRequestId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "accessPackageResourceRequestId", input)
	}

	if id.AccessPackageResourceRoleId, ok = input.Parsed["accessPackageResourceRoleId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "accessPackageResourceRoleId", input)
	}

	if id.AccessPackageResourceScopeId, ok = input.Parsed["accessPackageResourceScopeId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "accessPackageResourceScopeId", input)
	}

	return nil
}

// ValidateIdentityGovernanceEntitlementManagementResourceRequestIdCatalogResourceRoleIdResourceScopeID checks that 'input' can be parsed as a Identity Governance Entitlement Management Resource Request Id Catalog Resource Role Id Resource Scope ID
func ValidateIdentityGovernanceEntitlementManagementResourceRequestIdCatalogResourceRoleIdResourceScopeID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseIdentityGovernanceEntitlementManagementResourceRequestIdCatalogResourceRoleIdResourceScopeID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Identity Governance Entitlement Management Resource Request Id Catalog Resource Role Id Resource Scope ID
func (id IdentityGovernanceEntitlementManagementResourceRequestIdCatalogResourceRoleIdResourceScopeId) ID() string {
	fmtString := "/identityGovernance/entitlementManagement/resourceRequests/%s/catalog/resourceRoles/%s/resource/scopes/%s"
	return fmt.Sprintf(fmtString, id.AccessPackageResourceRequestId, id.AccessPackageResourceRoleId, id.AccessPackageResourceScopeId)
}

// Segments returns a slice of Resource ID Segments which comprise this Identity Governance Entitlement Management Resource Request Id Catalog Resource Role Id Resource Scope ID
func (id IdentityGovernanceEntitlementManagementResourceRequestIdCatalogResourceRoleIdResourceScopeId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("identityGovernance", "identityGovernance", "identityGovernance"),
		resourceids.StaticSegment("entitlementManagement", "entitlementManagement", "entitlementManagement"),
		resourceids.StaticSegment("resourceRequests", "resourceRequests", "resourceRequests"),
		resourceids.UserSpecifiedSegment("accessPackageResourceRequestId", "accessPackageResourceRequestId"),
		resourceids.StaticSegment("catalog", "catalog", "catalog"),
		resourceids.StaticSegment("resourceRoles", "resourceRoles", "resourceRoles"),
		resourceids.UserSpecifiedSegment("accessPackageResourceRoleId", "accessPackageResourceRoleId"),
		resourceids.StaticSegment("resource", "resource", "resource"),
		resourceids.StaticSegment("scopes", "scopes", "scopes"),
		resourceids.UserSpecifiedSegment("accessPackageResourceScopeId", "accessPackageResourceScopeId"),
	}
}

// String returns a human-readable description of this Identity Governance Entitlement Management Resource Request Id Catalog Resource Role Id Resource Scope ID
func (id IdentityGovernanceEntitlementManagementResourceRequestIdCatalogResourceRoleIdResourceScopeId) String() string {
	components := []string{
		fmt.Sprintf("Access Package Resource Request: %q", id.AccessPackageResourceRequestId),
		fmt.Sprintf("Access Package Resource Role: %q", id.AccessPackageResourceRoleId),
		fmt.Sprintf("Access Package Resource Scope: %q", id.AccessPackageResourceScopeId),
	}
	return fmt.Sprintf("Identity Governance Entitlement Management Resource Request Id Catalog Resource Role Id Resource Scope (%s)", strings.Join(components, "\n"))
}
