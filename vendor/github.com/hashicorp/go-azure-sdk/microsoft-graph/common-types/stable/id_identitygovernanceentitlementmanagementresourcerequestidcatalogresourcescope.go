package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &IdentityGovernanceEntitlementManagementResourceRequestIdCatalogResourceScopeId{}

// IdentityGovernanceEntitlementManagementResourceRequestIdCatalogResourceScopeId is a struct representing the Resource ID for a Identity Governance Entitlement Management Resource Request Id Catalog Resource Scope
type IdentityGovernanceEntitlementManagementResourceRequestIdCatalogResourceScopeId struct {
	AccessPackageResourceRequestId string
	AccessPackageResourceScopeId   string
}

// NewIdentityGovernanceEntitlementManagementResourceRequestIdCatalogResourceScopeID returns a new IdentityGovernanceEntitlementManagementResourceRequestIdCatalogResourceScopeId struct
func NewIdentityGovernanceEntitlementManagementResourceRequestIdCatalogResourceScopeID(accessPackageResourceRequestId string, accessPackageResourceScopeId string) IdentityGovernanceEntitlementManagementResourceRequestIdCatalogResourceScopeId {
	return IdentityGovernanceEntitlementManagementResourceRequestIdCatalogResourceScopeId{
		AccessPackageResourceRequestId: accessPackageResourceRequestId,
		AccessPackageResourceScopeId:   accessPackageResourceScopeId,
	}
}

// ParseIdentityGovernanceEntitlementManagementResourceRequestIdCatalogResourceScopeID parses 'input' into a IdentityGovernanceEntitlementManagementResourceRequestIdCatalogResourceScopeId
func ParseIdentityGovernanceEntitlementManagementResourceRequestIdCatalogResourceScopeID(input string) (*IdentityGovernanceEntitlementManagementResourceRequestIdCatalogResourceScopeId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityGovernanceEntitlementManagementResourceRequestIdCatalogResourceScopeId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityGovernanceEntitlementManagementResourceRequestIdCatalogResourceScopeId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseIdentityGovernanceEntitlementManagementResourceRequestIdCatalogResourceScopeIDInsensitively parses 'input' case-insensitively into a IdentityGovernanceEntitlementManagementResourceRequestIdCatalogResourceScopeId
// note: this method should only be used for API response data and not user input
func ParseIdentityGovernanceEntitlementManagementResourceRequestIdCatalogResourceScopeIDInsensitively(input string) (*IdentityGovernanceEntitlementManagementResourceRequestIdCatalogResourceScopeId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityGovernanceEntitlementManagementResourceRequestIdCatalogResourceScopeId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityGovernanceEntitlementManagementResourceRequestIdCatalogResourceScopeId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *IdentityGovernanceEntitlementManagementResourceRequestIdCatalogResourceScopeId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.AccessPackageResourceRequestId, ok = input.Parsed["accessPackageResourceRequestId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "accessPackageResourceRequestId", input)
	}

	if id.AccessPackageResourceScopeId, ok = input.Parsed["accessPackageResourceScopeId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "accessPackageResourceScopeId", input)
	}

	return nil
}

// ValidateIdentityGovernanceEntitlementManagementResourceRequestIdCatalogResourceScopeID checks that 'input' can be parsed as a Identity Governance Entitlement Management Resource Request Id Catalog Resource Scope ID
func ValidateIdentityGovernanceEntitlementManagementResourceRequestIdCatalogResourceScopeID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseIdentityGovernanceEntitlementManagementResourceRequestIdCatalogResourceScopeID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Identity Governance Entitlement Management Resource Request Id Catalog Resource Scope ID
func (id IdentityGovernanceEntitlementManagementResourceRequestIdCatalogResourceScopeId) ID() string {
	fmtString := "/identityGovernance/entitlementManagement/resourceRequests/%s/catalog/resourceScopes/%s"
	return fmt.Sprintf(fmtString, id.AccessPackageResourceRequestId, id.AccessPackageResourceScopeId)
}

// Segments returns a slice of Resource ID Segments which comprise this Identity Governance Entitlement Management Resource Request Id Catalog Resource Scope ID
func (id IdentityGovernanceEntitlementManagementResourceRequestIdCatalogResourceScopeId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("identityGovernance", "identityGovernance", "identityGovernance"),
		resourceids.StaticSegment("entitlementManagement", "entitlementManagement", "entitlementManagement"),
		resourceids.StaticSegment("resourceRequests", "resourceRequests", "resourceRequests"),
		resourceids.UserSpecifiedSegment("accessPackageResourceRequestId", "accessPackageResourceRequestId"),
		resourceids.StaticSegment("catalog", "catalog", "catalog"),
		resourceids.StaticSegment("resourceScopes", "resourceScopes", "resourceScopes"),
		resourceids.UserSpecifiedSegment("accessPackageResourceScopeId", "accessPackageResourceScopeId"),
	}
}

// String returns a human-readable description of this Identity Governance Entitlement Management Resource Request Id Catalog Resource Scope ID
func (id IdentityGovernanceEntitlementManagementResourceRequestIdCatalogResourceScopeId) String() string {
	components := []string{
		fmt.Sprintf("Access Package Resource Request: %q", id.AccessPackageResourceRequestId),
		fmt.Sprintf("Access Package Resource Scope: %q", id.AccessPackageResourceScopeId),
	}
	return fmt.Sprintf("Identity Governance Entitlement Management Resource Request Id Catalog Resource Scope (%s)", strings.Join(components, "\n"))
}
