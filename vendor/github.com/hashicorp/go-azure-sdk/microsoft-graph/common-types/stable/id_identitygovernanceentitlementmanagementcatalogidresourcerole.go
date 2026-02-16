package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &IdentityGovernanceEntitlementManagementCatalogIdResourceRoleId{}

// IdentityGovernanceEntitlementManagementCatalogIdResourceRoleId is a struct representing the Resource ID for a Identity Governance Entitlement Management Catalog Id Resource Role
type IdentityGovernanceEntitlementManagementCatalogIdResourceRoleId struct {
	AccessPackageCatalogId      string
	AccessPackageResourceRoleId string
}

// NewIdentityGovernanceEntitlementManagementCatalogIdResourceRoleID returns a new IdentityGovernanceEntitlementManagementCatalogIdResourceRoleId struct
func NewIdentityGovernanceEntitlementManagementCatalogIdResourceRoleID(accessPackageCatalogId string, accessPackageResourceRoleId string) IdentityGovernanceEntitlementManagementCatalogIdResourceRoleId {
	return IdentityGovernanceEntitlementManagementCatalogIdResourceRoleId{
		AccessPackageCatalogId:      accessPackageCatalogId,
		AccessPackageResourceRoleId: accessPackageResourceRoleId,
	}
}

// ParseIdentityGovernanceEntitlementManagementCatalogIdResourceRoleID parses 'input' into a IdentityGovernanceEntitlementManagementCatalogIdResourceRoleId
func ParseIdentityGovernanceEntitlementManagementCatalogIdResourceRoleID(input string) (*IdentityGovernanceEntitlementManagementCatalogIdResourceRoleId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityGovernanceEntitlementManagementCatalogIdResourceRoleId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityGovernanceEntitlementManagementCatalogIdResourceRoleId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseIdentityGovernanceEntitlementManagementCatalogIdResourceRoleIDInsensitively parses 'input' case-insensitively into a IdentityGovernanceEntitlementManagementCatalogIdResourceRoleId
// note: this method should only be used for API response data and not user input
func ParseIdentityGovernanceEntitlementManagementCatalogIdResourceRoleIDInsensitively(input string) (*IdentityGovernanceEntitlementManagementCatalogIdResourceRoleId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityGovernanceEntitlementManagementCatalogIdResourceRoleId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityGovernanceEntitlementManagementCatalogIdResourceRoleId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *IdentityGovernanceEntitlementManagementCatalogIdResourceRoleId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.AccessPackageCatalogId, ok = input.Parsed["accessPackageCatalogId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "accessPackageCatalogId", input)
	}

	if id.AccessPackageResourceRoleId, ok = input.Parsed["accessPackageResourceRoleId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "accessPackageResourceRoleId", input)
	}

	return nil
}

// ValidateIdentityGovernanceEntitlementManagementCatalogIdResourceRoleID checks that 'input' can be parsed as a Identity Governance Entitlement Management Catalog Id Resource Role ID
func ValidateIdentityGovernanceEntitlementManagementCatalogIdResourceRoleID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseIdentityGovernanceEntitlementManagementCatalogIdResourceRoleID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Identity Governance Entitlement Management Catalog Id Resource Role ID
func (id IdentityGovernanceEntitlementManagementCatalogIdResourceRoleId) ID() string {
	fmtString := "/identityGovernance/entitlementManagement/catalogs/%s/resourceRoles/%s"
	return fmt.Sprintf(fmtString, id.AccessPackageCatalogId, id.AccessPackageResourceRoleId)
}

// Segments returns a slice of Resource ID Segments which comprise this Identity Governance Entitlement Management Catalog Id Resource Role ID
func (id IdentityGovernanceEntitlementManagementCatalogIdResourceRoleId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("identityGovernance", "identityGovernance", "identityGovernance"),
		resourceids.StaticSegment("entitlementManagement", "entitlementManagement", "entitlementManagement"),
		resourceids.StaticSegment("catalogs", "catalogs", "catalogs"),
		resourceids.UserSpecifiedSegment("accessPackageCatalogId", "accessPackageCatalogId"),
		resourceids.StaticSegment("resourceRoles", "resourceRoles", "resourceRoles"),
		resourceids.UserSpecifiedSegment("accessPackageResourceRoleId", "accessPackageResourceRoleId"),
	}
}

// String returns a human-readable description of this Identity Governance Entitlement Management Catalog Id Resource Role ID
func (id IdentityGovernanceEntitlementManagementCatalogIdResourceRoleId) String() string {
	components := []string{
		fmt.Sprintf("Access Package Catalog: %q", id.AccessPackageCatalogId),
		fmt.Sprintf("Access Package Resource Role: %q", id.AccessPackageResourceRoleId),
	}
	return fmt.Sprintf("Identity Governance Entitlement Management Catalog Id Resource Role (%s)", strings.Join(components, "\n"))
}
