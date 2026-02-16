package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &IdentityGovernanceEntitlementManagementCatalogIdResourceScopeId{}

// IdentityGovernanceEntitlementManagementCatalogIdResourceScopeId is a struct representing the Resource ID for a Identity Governance Entitlement Management Catalog Id Resource Scope
type IdentityGovernanceEntitlementManagementCatalogIdResourceScopeId struct {
	AccessPackageCatalogId       string
	AccessPackageResourceScopeId string
}

// NewIdentityGovernanceEntitlementManagementCatalogIdResourceScopeID returns a new IdentityGovernanceEntitlementManagementCatalogIdResourceScopeId struct
func NewIdentityGovernanceEntitlementManagementCatalogIdResourceScopeID(accessPackageCatalogId string, accessPackageResourceScopeId string) IdentityGovernanceEntitlementManagementCatalogIdResourceScopeId {
	return IdentityGovernanceEntitlementManagementCatalogIdResourceScopeId{
		AccessPackageCatalogId:       accessPackageCatalogId,
		AccessPackageResourceScopeId: accessPackageResourceScopeId,
	}
}

// ParseIdentityGovernanceEntitlementManagementCatalogIdResourceScopeID parses 'input' into a IdentityGovernanceEntitlementManagementCatalogIdResourceScopeId
func ParseIdentityGovernanceEntitlementManagementCatalogIdResourceScopeID(input string) (*IdentityGovernanceEntitlementManagementCatalogIdResourceScopeId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityGovernanceEntitlementManagementCatalogIdResourceScopeId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityGovernanceEntitlementManagementCatalogIdResourceScopeId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseIdentityGovernanceEntitlementManagementCatalogIdResourceScopeIDInsensitively parses 'input' case-insensitively into a IdentityGovernanceEntitlementManagementCatalogIdResourceScopeId
// note: this method should only be used for API response data and not user input
func ParseIdentityGovernanceEntitlementManagementCatalogIdResourceScopeIDInsensitively(input string) (*IdentityGovernanceEntitlementManagementCatalogIdResourceScopeId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityGovernanceEntitlementManagementCatalogIdResourceScopeId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityGovernanceEntitlementManagementCatalogIdResourceScopeId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *IdentityGovernanceEntitlementManagementCatalogIdResourceScopeId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.AccessPackageCatalogId, ok = input.Parsed["accessPackageCatalogId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "accessPackageCatalogId", input)
	}

	if id.AccessPackageResourceScopeId, ok = input.Parsed["accessPackageResourceScopeId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "accessPackageResourceScopeId", input)
	}

	return nil
}

// ValidateIdentityGovernanceEntitlementManagementCatalogIdResourceScopeID checks that 'input' can be parsed as a Identity Governance Entitlement Management Catalog Id Resource Scope ID
func ValidateIdentityGovernanceEntitlementManagementCatalogIdResourceScopeID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseIdentityGovernanceEntitlementManagementCatalogIdResourceScopeID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Identity Governance Entitlement Management Catalog Id Resource Scope ID
func (id IdentityGovernanceEntitlementManagementCatalogIdResourceScopeId) ID() string {
	fmtString := "/identityGovernance/entitlementManagement/catalogs/%s/resourceScopes/%s"
	return fmt.Sprintf(fmtString, id.AccessPackageCatalogId, id.AccessPackageResourceScopeId)
}

// Segments returns a slice of Resource ID Segments which comprise this Identity Governance Entitlement Management Catalog Id Resource Scope ID
func (id IdentityGovernanceEntitlementManagementCatalogIdResourceScopeId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("identityGovernance", "identityGovernance", "identityGovernance"),
		resourceids.StaticSegment("entitlementManagement", "entitlementManagement", "entitlementManagement"),
		resourceids.StaticSegment("catalogs", "catalogs", "catalogs"),
		resourceids.UserSpecifiedSegment("accessPackageCatalogId", "accessPackageCatalogId"),
		resourceids.StaticSegment("resourceScopes", "resourceScopes", "resourceScopes"),
		resourceids.UserSpecifiedSegment("accessPackageResourceScopeId", "accessPackageResourceScopeId"),
	}
}

// String returns a human-readable description of this Identity Governance Entitlement Management Catalog Id Resource Scope ID
func (id IdentityGovernanceEntitlementManagementCatalogIdResourceScopeId) String() string {
	components := []string{
		fmt.Sprintf("Access Package Catalog: %q", id.AccessPackageCatalogId),
		fmt.Sprintf("Access Package Resource Scope: %q", id.AccessPackageResourceScopeId),
	}
	return fmt.Sprintf("Identity Governance Entitlement Management Catalog Id Resource Scope (%s)", strings.Join(components, "\n"))
}
