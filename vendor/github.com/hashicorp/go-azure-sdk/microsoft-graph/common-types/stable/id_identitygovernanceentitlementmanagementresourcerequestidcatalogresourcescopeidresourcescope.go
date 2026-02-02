package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &IdentityGovernanceEntitlementManagementResourceRequestIdCatalogResourceScopeIdResourceScopeId{}

// IdentityGovernanceEntitlementManagementResourceRequestIdCatalogResourceScopeIdResourceScopeId is a struct representing the Resource ID for a Identity Governance Entitlement Management Resource Request Id Catalog Resource Scope Id Resource Scope
type IdentityGovernanceEntitlementManagementResourceRequestIdCatalogResourceScopeIdResourceScopeId struct {
	AccessPackageResourceRequestId string
	AccessPackageResourceScopeId   string
	AccessPackageResourceScopeId1  string
}

// NewIdentityGovernanceEntitlementManagementResourceRequestIdCatalogResourceScopeIdResourceScopeID returns a new IdentityGovernanceEntitlementManagementResourceRequestIdCatalogResourceScopeIdResourceScopeId struct
func NewIdentityGovernanceEntitlementManagementResourceRequestIdCatalogResourceScopeIdResourceScopeID(accessPackageResourceRequestId string, accessPackageResourceScopeId string, accessPackageResourceScopeId1 string) IdentityGovernanceEntitlementManagementResourceRequestIdCatalogResourceScopeIdResourceScopeId {
	return IdentityGovernanceEntitlementManagementResourceRequestIdCatalogResourceScopeIdResourceScopeId{
		AccessPackageResourceRequestId: accessPackageResourceRequestId,
		AccessPackageResourceScopeId:   accessPackageResourceScopeId,
		AccessPackageResourceScopeId1:  accessPackageResourceScopeId1,
	}
}

// ParseIdentityGovernanceEntitlementManagementResourceRequestIdCatalogResourceScopeIdResourceScopeID parses 'input' into a IdentityGovernanceEntitlementManagementResourceRequestIdCatalogResourceScopeIdResourceScopeId
func ParseIdentityGovernanceEntitlementManagementResourceRequestIdCatalogResourceScopeIdResourceScopeID(input string) (*IdentityGovernanceEntitlementManagementResourceRequestIdCatalogResourceScopeIdResourceScopeId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityGovernanceEntitlementManagementResourceRequestIdCatalogResourceScopeIdResourceScopeId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityGovernanceEntitlementManagementResourceRequestIdCatalogResourceScopeIdResourceScopeId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseIdentityGovernanceEntitlementManagementResourceRequestIdCatalogResourceScopeIdResourceScopeIDInsensitively parses 'input' case-insensitively into a IdentityGovernanceEntitlementManagementResourceRequestIdCatalogResourceScopeIdResourceScopeId
// note: this method should only be used for API response data and not user input
func ParseIdentityGovernanceEntitlementManagementResourceRequestIdCatalogResourceScopeIdResourceScopeIDInsensitively(input string) (*IdentityGovernanceEntitlementManagementResourceRequestIdCatalogResourceScopeIdResourceScopeId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityGovernanceEntitlementManagementResourceRequestIdCatalogResourceScopeIdResourceScopeId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityGovernanceEntitlementManagementResourceRequestIdCatalogResourceScopeIdResourceScopeId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *IdentityGovernanceEntitlementManagementResourceRequestIdCatalogResourceScopeIdResourceScopeId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.AccessPackageResourceRequestId, ok = input.Parsed["accessPackageResourceRequestId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "accessPackageResourceRequestId", input)
	}

	if id.AccessPackageResourceScopeId, ok = input.Parsed["accessPackageResourceScopeId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "accessPackageResourceScopeId", input)
	}

	if id.AccessPackageResourceScopeId1, ok = input.Parsed["accessPackageResourceScopeId1"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "accessPackageResourceScopeId1", input)
	}

	return nil
}

// ValidateIdentityGovernanceEntitlementManagementResourceRequestIdCatalogResourceScopeIdResourceScopeID checks that 'input' can be parsed as a Identity Governance Entitlement Management Resource Request Id Catalog Resource Scope Id Resource Scope ID
func ValidateIdentityGovernanceEntitlementManagementResourceRequestIdCatalogResourceScopeIdResourceScopeID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseIdentityGovernanceEntitlementManagementResourceRequestIdCatalogResourceScopeIdResourceScopeID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Identity Governance Entitlement Management Resource Request Id Catalog Resource Scope Id Resource Scope ID
func (id IdentityGovernanceEntitlementManagementResourceRequestIdCatalogResourceScopeIdResourceScopeId) ID() string {
	fmtString := "/identityGovernance/entitlementManagement/resourceRequests/%s/catalog/resourceScopes/%s/resource/scopes/%s"
	return fmt.Sprintf(fmtString, id.AccessPackageResourceRequestId, id.AccessPackageResourceScopeId, id.AccessPackageResourceScopeId1)
}

// Segments returns a slice of Resource ID Segments which comprise this Identity Governance Entitlement Management Resource Request Id Catalog Resource Scope Id Resource Scope ID
func (id IdentityGovernanceEntitlementManagementResourceRequestIdCatalogResourceScopeIdResourceScopeId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("identityGovernance", "identityGovernance", "identityGovernance"),
		resourceids.StaticSegment("entitlementManagement", "entitlementManagement", "entitlementManagement"),
		resourceids.StaticSegment("resourceRequests", "resourceRequests", "resourceRequests"),
		resourceids.UserSpecifiedSegment("accessPackageResourceRequestId", "accessPackageResourceRequestId"),
		resourceids.StaticSegment("catalog", "catalog", "catalog"),
		resourceids.StaticSegment("resourceScopes", "resourceScopes", "resourceScopes"),
		resourceids.UserSpecifiedSegment("accessPackageResourceScopeId", "accessPackageResourceScopeId"),
		resourceids.StaticSegment("resource", "resource", "resource"),
		resourceids.StaticSegment("scopes", "scopes", "scopes"),
		resourceids.UserSpecifiedSegment("accessPackageResourceScopeId1", "accessPackageResourceScopeId1"),
	}
}

// String returns a human-readable description of this Identity Governance Entitlement Management Resource Request Id Catalog Resource Scope Id Resource Scope ID
func (id IdentityGovernanceEntitlementManagementResourceRequestIdCatalogResourceScopeIdResourceScopeId) String() string {
	components := []string{
		fmt.Sprintf("Access Package Resource Request: %q", id.AccessPackageResourceRequestId),
		fmt.Sprintf("Access Package Resource Scope: %q", id.AccessPackageResourceScopeId),
		fmt.Sprintf("Access Package Resource Scope Id 1: %q", id.AccessPackageResourceScopeId1),
	}
	return fmt.Sprintf("Identity Governance Entitlement Management Resource Request Id Catalog Resource Scope Id Resource Scope (%s)", strings.Join(components, "\n"))
}
