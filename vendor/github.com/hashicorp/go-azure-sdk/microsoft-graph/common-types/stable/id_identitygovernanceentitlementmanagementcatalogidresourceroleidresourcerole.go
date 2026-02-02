package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &IdentityGovernanceEntitlementManagementCatalogIdResourceRoleIdResourceRoleId{}

// IdentityGovernanceEntitlementManagementCatalogIdResourceRoleIdResourceRoleId is a struct representing the Resource ID for a Identity Governance Entitlement Management Catalog Id Resource Role Id Resource Role
type IdentityGovernanceEntitlementManagementCatalogIdResourceRoleIdResourceRoleId struct {
	AccessPackageCatalogId       string
	AccessPackageResourceRoleId  string
	AccessPackageResourceRoleId1 string
}

// NewIdentityGovernanceEntitlementManagementCatalogIdResourceRoleIdResourceRoleID returns a new IdentityGovernanceEntitlementManagementCatalogIdResourceRoleIdResourceRoleId struct
func NewIdentityGovernanceEntitlementManagementCatalogIdResourceRoleIdResourceRoleID(accessPackageCatalogId string, accessPackageResourceRoleId string, accessPackageResourceRoleId1 string) IdentityGovernanceEntitlementManagementCatalogIdResourceRoleIdResourceRoleId {
	return IdentityGovernanceEntitlementManagementCatalogIdResourceRoleIdResourceRoleId{
		AccessPackageCatalogId:       accessPackageCatalogId,
		AccessPackageResourceRoleId:  accessPackageResourceRoleId,
		AccessPackageResourceRoleId1: accessPackageResourceRoleId1,
	}
}

// ParseIdentityGovernanceEntitlementManagementCatalogIdResourceRoleIdResourceRoleID parses 'input' into a IdentityGovernanceEntitlementManagementCatalogIdResourceRoleIdResourceRoleId
func ParseIdentityGovernanceEntitlementManagementCatalogIdResourceRoleIdResourceRoleID(input string) (*IdentityGovernanceEntitlementManagementCatalogIdResourceRoleIdResourceRoleId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityGovernanceEntitlementManagementCatalogIdResourceRoleIdResourceRoleId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityGovernanceEntitlementManagementCatalogIdResourceRoleIdResourceRoleId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseIdentityGovernanceEntitlementManagementCatalogIdResourceRoleIdResourceRoleIDInsensitively parses 'input' case-insensitively into a IdentityGovernanceEntitlementManagementCatalogIdResourceRoleIdResourceRoleId
// note: this method should only be used for API response data and not user input
func ParseIdentityGovernanceEntitlementManagementCatalogIdResourceRoleIdResourceRoleIDInsensitively(input string) (*IdentityGovernanceEntitlementManagementCatalogIdResourceRoleIdResourceRoleId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityGovernanceEntitlementManagementCatalogIdResourceRoleIdResourceRoleId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityGovernanceEntitlementManagementCatalogIdResourceRoleIdResourceRoleId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *IdentityGovernanceEntitlementManagementCatalogIdResourceRoleIdResourceRoleId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.AccessPackageCatalogId, ok = input.Parsed["accessPackageCatalogId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "accessPackageCatalogId", input)
	}

	if id.AccessPackageResourceRoleId, ok = input.Parsed["accessPackageResourceRoleId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "accessPackageResourceRoleId", input)
	}

	if id.AccessPackageResourceRoleId1, ok = input.Parsed["accessPackageResourceRoleId1"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "accessPackageResourceRoleId1", input)
	}

	return nil
}

// ValidateIdentityGovernanceEntitlementManagementCatalogIdResourceRoleIdResourceRoleID checks that 'input' can be parsed as a Identity Governance Entitlement Management Catalog Id Resource Role Id Resource Role ID
func ValidateIdentityGovernanceEntitlementManagementCatalogIdResourceRoleIdResourceRoleID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseIdentityGovernanceEntitlementManagementCatalogIdResourceRoleIdResourceRoleID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Identity Governance Entitlement Management Catalog Id Resource Role Id Resource Role ID
func (id IdentityGovernanceEntitlementManagementCatalogIdResourceRoleIdResourceRoleId) ID() string {
	fmtString := "/identityGovernance/entitlementManagement/catalogs/%s/resourceRoles/%s/resource/roles/%s"
	return fmt.Sprintf(fmtString, id.AccessPackageCatalogId, id.AccessPackageResourceRoleId, id.AccessPackageResourceRoleId1)
}

// Segments returns a slice of Resource ID Segments which comprise this Identity Governance Entitlement Management Catalog Id Resource Role Id Resource Role ID
func (id IdentityGovernanceEntitlementManagementCatalogIdResourceRoleIdResourceRoleId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("identityGovernance", "identityGovernance", "identityGovernance"),
		resourceids.StaticSegment("entitlementManagement", "entitlementManagement", "entitlementManagement"),
		resourceids.StaticSegment("catalogs", "catalogs", "catalogs"),
		resourceids.UserSpecifiedSegment("accessPackageCatalogId", "accessPackageCatalogId"),
		resourceids.StaticSegment("resourceRoles", "resourceRoles", "resourceRoles"),
		resourceids.UserSpecifiedSegment("accessPackageResourceRoleId", "accessPackageResourceRoleId"),
		resourceids.StaticSegment("resource", "resource", "resource"),
		resourceids.StaticSegment("roles", "roles", "roles"),
		resourceids.UserSpecifiedSegment("accessPackageResourceRoleId1", "accessPackageResourceRoleId1"),
	}
}

// String returns a human-readable description of this Identity Governance Entitlement Management Catalog Id Resource Role Id Resource Role ID
func (id IdentityGovernanceEntitlementManagementCatalogIdResourceRoleIdResourceRoleId) String() string {
	components := []string{
		fmt.Sprintf("Access Package Catalog: %q", id.AccessPackageCatalogId),
		fmt.Sprintf("Access Package Resource Role: %q", id.AccessPackageResourceRoleId),
		fmt.Sprintf("Access Package Resource Role Id 1: %q", id.AccessPackageResourceRoleId1),
	}
	return fmt.Sprintf("Identity Governance Entitlement Management Catalog Id Resource Role Id Resource Role (%s)", strings.Join(components, "\n"))
}
