package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &IdentityGovernanceEntitlementManagementCatalogIdResourceIdScopeId{}

// IdentityGovernanceEntitlementManagementCatalogIdResourceIdScopeId is a struct representing the Resource ID for a Identity Governance Entitlement Management Catalog Id Resource Id Scope
type IdentityGovernanceEntitlementManagementCatalogIdResourceIdScopeId struct {
	AccessPackageCatalogId       string
	AccessPackageResourceId      string
	AccessPackageResourceScopeId string
}

// NewIdentityGovernanceEntitlementManagementCatalogIdResourceIdScopeID returns a new IdentityGovernanceEntitlementManagementCatalogIdResourceIdScopeId struct
func NewIdentityGovernanceEntitlementManagementCatalogIdResourceIdScopeID(accessPackageCatalogId string, accessPackageResourceId string, accessPackageResourceScopeId string) IdentityGovernanceEntitlementManagementCatalogIdResourceIdScopeId {
	return IdentityGovernanceEntitlementManagementCatalogIdResourceIdScopeId{
		AccessPackageCatalogId:       accessPackageCatalogId,
		AccessPackageResourceId:      accessPackageResourceId,
		AccessPackageResourceScopeId: accessPackageResourceScopeId,
	}
}

// ParseIdentityGovernanceEntitlementManagementCatalogIdResourceIdScopeID parses 'input' into a IdentityGovernanceEntitlementManagementCatalogIdResourceIdScopeId
func ParseIdentityGovernanceEntitlementManagementCatalogIdResourceIdScopeID(input string) (*IdentityGovernanceEntitlementManagementCatalogIdResourceIdScopeId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityGovernanceEntitlementManagementCatalogIdResourceIdScopeId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityGovernanceEntitlementManagementCatalogIdResourceIdScopeId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseIdentityGovernanceEntitlementManagementCatalogIdResourceIdScopeIDInsensitively parses 'input' case-insensitively into a IdentityGovernanceEntitlementManagementCatalogIdResourceIdScopeId
// note: this method should only be used for API response data and not user input
func ParseIdentityGovernanceEntitlementManagementCatalogIdResourceIdScopeIDInsensitively(input string) (*IdentityGovernanceEntitlementManagementCatalogIdResourceIdScopeId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityGovernanceEntitlementManagementCatalogIdResourceIdScopeId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityGovernanceEntitlementManagementCatalogIdResourceIdScopeId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *IdentityGovernanceEntitlementManagementCatalogIdResourceIdScopeId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.AccessPackageCatalogId, ok = input.Parsed["accessPackageCatalogId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "accessPackageCatalogId", input)
	}

	if id.AccessPackageResourceId, ok = input.Parsed["accessPackageResourceId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "accessPackageResourceId", input)
	}

	if id.AccessPackageResourceScopeId, ok = input.Parsed["accessPackageResourceScopeId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "accessPackageResourceScopeId", input)
	}

	return nil
}

// ValidateIdentityGovernanceEntitlementManagementCatalogIdResourceIdScopeID checks that 'input' can be parsed as a Identity Governance Entitlement Management Catalog Id Resource Id Scope ID
func ValidateIdentityGovernanceEntitlementManagementCatalogIdResourceIdScopeID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseIdentityGovernanceEntitlementManagementCatalogIdResourceIdScopeID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Identity Governance Entitlement Management Catalog Id Resource Id Scope ID
func (id IdentityGovernanceEntitlementManagementCatalogIdResourceIdScopeId) ID() string {
	fmtString := "/identityGovernance/entitlementManagement/catalogs/%s/resources/%s/scopes/%s"
	return fmt.Sprintf(fmtString, id.AccessPackageCatalogId, id.AccessPackageResourceId, id.AccessPackageResourceScopeId)
}

// Segments returns a slice of Resource ID Segments which comprise this Identity Governance Entitlement Management Catalog Id Resource Id Scope ID
func (id IdentityGovernanceEntitlementManagementCatalogIdResourceIdScopeId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("identityGovernance", "identityGovernance", "identityGovernance"),
		resourceids.StaticSegment("entitlementManagement", "entitlementManagement", "entitlementManagement"),
		resourceids.StaticSegment("catalogs", "catalogs", "catalogs"),
		resourceids.UserSpecifiedSegment("accessPackageCatalogId", "accessPackageCatalogId"),
		resourceids.StaticSegment("resources", "resources", "resources"),
		resourceids.UserSpecifiedSegment("accessPackageResourceId", "accessPackageResourceId"),
		resourceids.StaticSegment("scopes", "scopes", "scopes"),
		resourceids.UserSpecifiedSegment("accessPackageResourceScopeId", "accessPackageResourceScopeId"),
	}
}

// String returns a human-readable description of this Identity Governance Entitlement Management Catalog Id Resource Id Scope ID
func (id IdentityGovernanceEntitlementManagementCatalogIdResourceIdScopeId) String() string {
	components := []string{
		fmt.Sprintf("Access Package Catalog: %q", id.AccessPackageCatalogId),
		fmt.Sprintf("Access Package Resource: %q", id.AccessPackageResourceId),
		fmt.Sprintf("Access Package Resource Scope: %q", id.AccessPackageResourceScopeId),
	}
	return fmt.Sprintf("Identity Governance Entitlement Management Catalog Id Resource Id Scope (%s)", strings.Join(components, "\n"))
}
