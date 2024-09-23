package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &IdentityGovernanceEntitlementManagementCatalogIdResourceIdRoleId{}

// IdentityGovernanceEntitlementManagementCatalogIdResourceIdRoleId is a struct representing the Resource ID for a Identity Governance Entitlement Management Catalog Id Resource Id Role
type IdentityGovernanceEntitlementManagementCatalogIdResourceIdRoleId struct {
	AccessPackageCatalogId      string
	AccessPackageResourceId     string
	AccessPackageResourceRoleId string
}

// NewIdentityGovernanceEntitlementManagementCatalogIdResourceIdRoleID returns a new IdentityGovernanceEntitlementManagementCatalogIdResourceIdRoleId struct
func NewIdentityGovernanceEntitlementManagementCatalogIdResourceIdRoleID(accessPackageCatalogId string, accessPackageResourceId string, accessPackageResourceRoleId string) IdentityGovernanceEntitlementManagementCatalogIdResourceIdRoleId {
	return IdentityGovernanceEntitlementManagementCatalogIdResourceIdRoleId{
		AccessPackageCatalogId:      accessPackageCatalogId,
		AccessPackageResourceId:     accessPackageResourceId,
		AccessPackageResourceRoleId: accessPackageResourceRoleId,
	}
}

// ParseIdentityGovernanceEntitlementManagementCatalogIdResourceIdRoleID parses 'input' into a IdentityGovernanceEntitlementManagementCatalogIdResourceIdRoleId
func ParseIdentityGovernanceEntitlementManagementCatalogIdResourceIdRoleID(input string) (*IdentityGovernanceEntitlementManagementCatalogIdResourceIdRoleId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityGovernanceEntitlementManagementCatalogIdResourceIdRoleId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityGovernanceEntitlementManagementCatalogIdResourceIdRoleId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseIdentityGovernanceEntitlementManagementCatalogIdResourceIdRoleIDInsensitively parses 'input' case-insensitively into a IdentityGovernanceEntitlementManagementCatalogIdResourceIdRoleId
// note: this method should only be used for API response data and not user input
func ParseIdentityGovernanceEntitlementManagementCatalogIdResourceIdRoleIDInsensitively(input string) (*IdentityGovernanceEntitlementManagementCatalogIdResourceIdRoleId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityGovernanceEntitlementManagementCatalogIdResourceIdRoleId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityGovernanceEntitlementManagementCatalogIdResourceIdRoleId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *IdentityGovernanceEntitlementManagementCatalogIdResourceIdRoleId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.AccessPackageCatalogId, ok = input.Parsed["accessPackageCatalogId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "accessPackageCatalogId", input)
	}

	if id.AccessPackageResourceId, ok = input.Parsed["accessPackageResourceId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "accessPackageResourceId", input)
	}

	if id.AccessPackageResourceRoleId, ok = input.Parsed["accessPackageResourceRoleId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "accessPackageResourceRoleId", input)
	}

	return nil
}

// ValidateIdentityGovernanceEntitlementManagementCatalogIdResourceIdRoleID checks that 'input' can be parsed as a Identity Governance Entitlement Management Catalog Id Resource Id Role ID
func ValidateIdentityGovernanceEntitlementManagementCatalogIdResourceIdRoleID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseIdentityGovernanceEntitlementManagementCatalogIdResourceIdRoleID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Identity Governance Entitlement Management Catalog Id Resource Id Role ID
func (id IdentityGovernanceEntitlementManagementCatalogIdResourceIdRoleId) ID() string {
	fmtString := "/identityGovernance/entitlementManagement/catalogs/%s/resources/%s/roles/%s"
	return fmt.Sprintf(fmtString, id.AccessPackageCatalogId, id.AccessPackageResourceId, id.AccessPackageResourceRoleId)
}

// Segments returns a slice of Resource ID Segments which comprise this Identity Governance Entitlement Management Catalog Id Resource Id Role ID
func (id IdentityGovernanceEntitlementManagementCatalogIdResourceIdRoleId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("identityGovernance", "identityGovernance", "identityGovernance"),
		resourceids.StaticSegment("entitlementManagement", "entitlementManagement", "entitlementManagement"),
		resourceids.StaticSegment("catalogs", "catalogs", "catalogs"),
		resourceids.UserSpecifiedSegment("accessPackageCatalogId", "accessPackageCatalogId"),
		resourceids.StaticSegment("resources", "resources", "resources"),
		resourceids.UserSpecifiedSegment("accessPackageResourceId", "accessPackageResourceId"),
		resourceids.StaticSegment("roles", "roles", "roles"),
		resourceids.UserSpecifiedSegment("accessPackageResourceRoleId", "accessPackageResourceRoleId"),
	}
}

// String returns a human-readable description of this Identity Governance Entitlement Management Catalog Id Resource Id Role ID
func (id IdentityGovernanceEntitlementManagementCatalogIdResourceIdRoleId) String() string {
	components := []string{
		fmt.Sprintf("Access Package Catalog: %q", id.AccessPackageCatalogId),
		fmt.Sprintf("Access Package Resource: %q", id.AccessPackageResourceId),
		fmt.Sprintf("Access Package Resource Role: %q", id.AccessPackageResourceRoleId),
	}
	return fmt.Sprintf("Identity Governance Entitlement Management Catalog Id Resource Id Role (%s)", strings.Join(components, "\n"))
}
