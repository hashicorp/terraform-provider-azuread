package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &IdentityGovernanceEntitlementManagementResourceRequestIdCatalogResourceIdScopeId{}

// IdentityGovernanceEntitlementManagementResourceRequestIdCatalogResourceIdScopeId is a struct representing the Resource ID for a Identity Governance Entitlement Management Resource Request Id Catalog Resource Id Scope
type IdentityGovernanceEntitlementManagementResourceRequestIdCatalogResourceIdScopeId struct {
	AccessPackageResourceRequestId string
	AccessPackageResourceId        string
	AccessPackageResourceScopeId   string
}

// NewIdentityGovernanceEntitlementManagementResourceRequestIdCatalogResourceIdScopeID returns a new IdentityGovernanceEntitlementManagementResourceRequestIdCatalogResourceIdScopeId struct
func NewIdentityGovernanceEntitlementManagementResourceRequestIdCatalogResourceIdScopeID(accessPackageResourceRequestId string, accessPackageResourceId string, accessPackageResourceScopeId string) IdentityGovernanceEntitlementManagementResourceRequestIdCatalogResourceIdScopeId {
	return IdentityGovernanceEntitlementManagementResourceRequestIdCatalogResourceIdScopeId{
		AccessPackageResourceRequestId: accessPackageResourceRequestId,
		AccessPackageResourceId:        accessPackageResourceId,
		AccessPackageResourceScopeId:   accessPackageResourceScopeId,
	}
}

// ParseIdentityGovernanceEntitlementManagementResourceRequestIdCatalogResourceIdScopeID parses 'input' into a IdentityGovernanceEntitlementManagementResourceRequestIdCatalogResourceIdScopeId
func ParseIdentityGovernanceEntitlementManagementResourceRequestIdCatalogResourceIdScopeID(input string) (*IdentityGovernanceEntitlementManagementResourceRequestIdCatalogResourceIdScopeId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityGovernanceEntitlementManagementResourceRequestIdCatalogResourceIdScopeId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityGovernanceEntitlementManagementResourceRequestIdCatalogResourceIdScopeId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseIdentityGovernanceEntitlementManagementResourceRequestIdCatalogResourceIdScopeIDInsensitively parses 'input' case-insensitively into a IdentityGovernanceEntitlementManagementResourceRequestIdCatalogResourceIdScopeId
// note: this method should only be used for API response data and not user input
func ParseIdentityGovernanceEntitlementManagementResourceRequestIdCatalogResourceIdScopeIDInsensitively(input string) (*IdentityGovernanceEntitlementManagementResourceRequestIdCatalogResourceIdScopeId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityGovernanceEntitlementManagementResourceRequestIdCatalogResourceIdScopeId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityGovernanceEntitlementManagementResourceRequestIdCatalogResourceIdScopeId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *IdentityGovernanceEntitlementManagementResourceRequestIdCatalogResourceIdScopeId) FromParseResult(input resourceids.ParseResult) error {
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

	return nil
}

// ValidateIdentityGovernanceEntitlementManagementResourceRequestIdCatalogResourceIdScopeID checks that 'input' can be parsed as a Identity Governance Entitlement Management Resource Request Id Catalog Resource Id Scope ID
func ValidateIdentityGovernanceEntitlementManagementResourceRequestIdCatalogResourceIdScopeID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseIdentityGovernanceEntitlementManagementResourceRequestIdCatalogResourceIdScopeID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Identity Governance Entitlement Management Resource Request Id Catalog Resource Id Scope ID
func (id IdentityGovernanceEntitlementManagementResourceRequestIdCatalogResourceIdScopeId) ID() string {
	fmtString := "/identityGovernance/entitlementManagement/resourceRequests/%s/catalog/resources/%s/scopes/%s"
	return fmt.Sprintf(fmtString, id.AccessPackageResourceRequestId, id.AccessPackageResourceId, id.AccessPackageResourceScopeId)
}

// Segments returns a slice of Resource ID Segments which comprise this Identity Governance Entitlement Management Resource Request Id Catalog Resource Id Scope ID
func (id IdentityGovernanceEntitlementManagementResourceRequestIdCatalogResourceIdScopeId) Segments() []resourceids.Segment {
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
	}
}

// String returns a human-readable description of this Identity Governance Entitlement Management Resource Request Id Catalog Resource Id Scope ID
func (id IdentityGovernanceEntitlementManagementResourceRequestIdCatalogResourceIdScopeId) String() string {
	components := []string{
		fmt.Sprintf("Access Package Resource Request: %q", id.AccessPackageResourceRequestId),
		fmt.Sprintf("Access Package Resource: %q", id.AccessPackageResourceId),
		fmt.Sprintf("Access Package Resource Scope: %q", id.AccessPackageResourceScopeId),
	}
	return fmt.Sprintf("Identity Governance Entitlement Management Resource Request Id Catalog Resource Id Scope (%s)", strings.Join(components, "\n"))
}
