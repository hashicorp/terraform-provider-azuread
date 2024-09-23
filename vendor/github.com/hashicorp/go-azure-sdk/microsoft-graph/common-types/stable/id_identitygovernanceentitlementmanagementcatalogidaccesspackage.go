package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &IdentityGovernanceEntitlementManagementCatalogIdAccessPackageId{}

// IdentityGovernanceEntitlementManagementCatalogIdAccessPackageId is a struct representing the Resource ID for a Identity Governance Entitlement Management Catalog Id Access Package
type IdentityGovernanceEntitlementManagementCatalogIdAccessPackageId struct {
	AccessPackageCatalogId string
	AccessPackageId        string
}

// NewIdentityGovernanceEntitlementManagementCatalogIdAccessPackageID returns a new IdentityGovernanceEntitlementManagementCatalogIdAccessPackageId struct
func NewIdentityGovernanceEntitlementManagementCatalogIdAccessPackageID(accessPackageCatalogId string, accessPackageId string) IdentityGovernanceEntitlementManagementCatalogIdAccessPackageId {
	return IdentityGovernanceEntitlementManagementCatalogIdAccessPackageId{
		AccessPackageCatalogId: accessPackageCatalogId,
		AccessPackageId:        accessPackageId,
	}
}

// ParseIdentityGovernanceEntitlementManagementCatalogIdAccessPackageID parses 'input' into a IdentityGovernanceEntitlementManagementCatalogIdAccessPackageId
func ParseIdentityGovernanceEntitlementManagementCatalogIdAccessPackageID(input string) (*IdentityGovernanceEntitlementManagementCatalogIdAccessPackageId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityGovernanceEntitlementManagementCatalogIdAccessPackageId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityGovernanceEntitlementManagementCatalogIdAccessPackageId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseIdentityGovernanceEntitlementManagementCatalogIdAccessPackageIDInsensitively parses 'input' case-insensitively into a IdentityGovernanceEntitlementManagementCatalogIdAccessPackageId
// note: this method should only be used for API response data and not user input
func ParseIdentityGovernanceEntitlementManagementCatalogIdAccessPackageIDInsensitively(input string) (*IdentityGovernanceEntitlementManagementCatalogIdAccessPackageId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityGovernanceEntitlementManagementCatalogIdAccessPackageId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityGovernanceEntitlementManagementCatalogIdAccessPackageId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *IdentityGovernanceEntitlementManagementCatalogIdAccessPackageId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.AccessPackageCatalogId, ok = input.Parsed["accessPackageCatalogId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "accessPackageCatalogId", input)
	}

	if id.AccessPackageId, ok = input.Parsed["accessPackageId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "accessPackageId", input)
	}

	return nil
}

// ValidateIdentityGovernanceEntitlementManagementCatalogIdAccessPackageID checks that 'input' can be parsed as a Identity Governance Entitlement Management Catalog Id Access Package ID
func ValidateIdentityGovernanceEntitlementManagementCatalogIdAccessPackageID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseIdentityGovernanceEntitlementManagementCatalogIdAccessPackageID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Identity Governance Entitlement Management Catalog Id Access Package ID
func (id IdentityGovernanceEntitlementManagementCatalogIdAccessPackageId) ID() string {
	fmtString := "/identityGovernance/entitlementManagement/catalogs/%s/accessPackages/%s"
	return fmt.Sprintf(fmtString, id.AccessPackageCatalogId, id.AccessPackageId)
}

// Segments returns a slice of Resource ID Segments which comprise this Identity Governance Entitlement Management Catalog Id Access Package ID
func (id IdentityGovernanceEntitlementManagementCatalogIdAccessPackageId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("identityGovernance", "identityGovernance", "identityGovernance"),
		resourceids.StaticSegment("entitlementManagement", "entitlementManagement", "entitlementManagement"),
		resourceids.StaticSegment("catalogs", "catalogs", "catalogs"),
		resourceids.UserSpecifiedSegment("accessPackageCatalogId", "accessPackageCatalogId"),
		resourceids.StaticSegment("accessPackages", "accessPackages", "accessPackages"),
		resourceids.UserSpecifiedSegment("accessPackageId", "accessPackageId"),
	}
}

// String returns a human-readable description of this Identity Governance Entitlement Management Catalog Id Access Package ID
func (id IdentityGovernanceEntitlementManagementCatalogIdAccessPackageId) String() string {
	components := []string{
		fmt.Sprintf("Access Package Catalog: %q", id.AccessPackageCatalogId),
		fmt.Sprintf("Access Package: %q", id.AccessPackageId),
	}
	return fmt.Sprintf("Identity Governance Entitlement Management Catalog Id Access Package (%s)", strings.Join(components, "\n"))
}
