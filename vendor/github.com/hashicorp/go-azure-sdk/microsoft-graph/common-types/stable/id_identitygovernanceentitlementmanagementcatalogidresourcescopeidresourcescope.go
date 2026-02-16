package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &IdentityGovernanceEntitlementManagementCatalogIdResourceScopeIdResourceScopeId{}

// IdentityGovernanceEntitlementManagementCatalogIdResourceScopeIdResourceScopeId is a struct representing the Resource ID for a Identity Governance Entitlement Management Catalog Id Resource Scope Id Resource Scope
type IdentityGovernanceEntitlementManagementCatalogIdResourceScopeIdResourceScopeId struct {
	AccessPackageCatalogId        string
	AccessPackageResourceScopeId  string
	AccessPackageResourceScopeId1 string
}

// NewIdentityGovernanceEntitlementManagementCatalogIdResourceScopeIdResourceScopeID returns a new IdentityGovernanceEntitlementManagementCatalogIdResourceScopeIdResourceScopeId struct
func NewIdentityGovernanceEntitlementManagementCatalogIdResourceScopeIdResourceScopeID(accessPackageCatalogId string, accessPackageResourceScopeId string, accessPackageResourceScopeId1 string) IdentityGovernanceEntitlementManagementCatalogIdResourceScopeIdResourceScopeId {
	return IdentityGovernanceEntitlementManagementCatalogIdResourceScopeIdResourceScopeId{
		AccessPackageCatalogId:        accessPackageCatalogId,
		AccessPackageResourceScopeId:  accessPackageResourceScopeId,
		AccessPackageResourceScopeId1: accessPackageResourceScopeId1,
	}
}

// ParseIdentityGovernanceEntitlementManagementCatalogIdResourceScopeIdResourceScopeID parses 'input' into a IdentityGovernanceEntitlementManagementCatalogIdResourceScopeIdResourceScopeId
func ParseIdentityGovernanceEntitlementManagementCatalogIdResourceScopeIdResourceScopeID(input string) (*IdentityGovernanceEntitlementManagementCatalogIdResourceScopeIdResourceScopeId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityGovernanceEntitlementManagementCatalogIdResourceScopeIdResourceScopeId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityGovernanceEntitlementManagementCatalogIdResourceScopeIdResourceScopeId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseIdentityGovernanceEntitlementManagementCatalogIdResourceScopeIdResourceScopeIDInsensitively parses 'input' case-insensitively into a IdentityGovernanceEntitlementManagementCatalogIdResourceScopeIdResourceScopeId
// note: this method should only be used for API response data and not user input
func ParseIdentityGovernanceEntitlementManagementCatalogIdResourceScopeIdResourceScopeIDInsensitively(input string) (*IdentityGovernanceEntitlementManagementCatalogIdResourceScopeIdResourceScopeId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityGovernanceEntitlementManagementCatalogIdResourceScopeIdResourceScopeId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityGovernanceEntitlementManagementCatalogIdResourceScopeIdResourceScopeId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *IdentityGovernanceEntitlementManagementCatalogIdResourceScopeIdResourceScopeId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.AccessPackageCatalogId, ok = input.Parsed["accessPackageCatalogId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "accessPackageCatalogId", input)
	}

	if id.AccessPackageResourceScopeId, ok = input.Parsed["accessPackageResourceScopeId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "accessPackageResourceScopeId", input)
	}

	if id.AccessPackageResourceScopeId1, ok = input.Parsed["accessPackageResourceScopeId1"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "accessPackageResourceScopeId1", input)
	}

	return nil
}

// ValidateIdentityGovernanceEntitlementManagementCatalogIdResourceScopeIdResourceScopeID checks that 'input' can be parsed as a Identity Governance Entitlement Management Catalog Id Resource Scope Id Resource Scope ID
func ValidateIdentityGovernanceEntitlementManagementCatalogIdResourceScopeIdResourceScopeID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseIdentityGovernanceEntitlementManagementCatalogIdResourceScopeIdResourceScopeID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Identity Governance Entitlement Management Catalog Id Resource Scope Id Resource Scope ID
func (id IdentityGovernanceEntitlementManagementCatalogIdResourceScopeIdResourceScopeId) ID() string {
	fmtString := "/identityGovernance/entitlementManagement/catalogs/%s/resourceScopes/%s/resource/scopes/%s"
	return fmt.Sprintf(fmtString, id.AccessPackageCatalogId, id.AccessPackageResourceScopeId, id.AccessPackageResourceScopeId1)
}

// Segments returns a slice of Resource ID Segments which comprise this Identity Governance Entitlement Management Catalog Id Resource Scope Id Resource Scope ID
func (id IdentityGovernanceEntitlementManagementCatalogIdResourceScopeIdResourceScopeId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("identityGovernance", "identityGovernance", "identityGovernance"),
		resourceids.StaticSegment("entitlementManagement", "entitlementManagement", "entitlementManagement"),
		resourceids.StaticSegment("catalogs", "catalogs", "catalogs"),
		resourceids.UserSpecifiedSegment("accessPackageCatalogId", "accessPackageCatalogId"),
		resourceids.StaticSegment("resourceScopes", "resourceScopes", "resourceScopes"),
		resourceids.UserSpecifiedSegment("accessPackageResourceScopeId", "accessPackageResourceScopeId"),
		resourceids.StaticSegment("resource", "resource", "resource"),
		resourceids.StaticSegment("scopes", "scopes", "scopes"),
		resourceids.UserSpecifiedSegment("accessPackageResourceScopeId1", "accessPackageResourceScopeId1"),
	}
}

// String returns a human-readable description of this Identity Governance Entitlement Management Catalog Id Resource Scope Id Resource Scope ID
func (id IdentityGovernanceEntitlementManagementCatalogIdResourceScopeIdResourceScopeId) String() string {
	components := []string{
		fmt.Sprintf("Access Package Catalog: %q", id.AccessPackageCatalogId),
		fmt.Sprintf("Access Package Resource Scope: %q", id.AccessPackageResourceScopeId),
		fmt.Sprintf("Access Package Resource Scope Id 1: %q", id.AccessPackageResourceScopeId1),
	}
	return fmt.Sprintf("Identity Governance Entitlement Management Catalog Id Resource Scope Id Resource Scope (%s)", strings.Join(components, "\n"))
}
