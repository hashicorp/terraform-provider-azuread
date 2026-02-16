package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &IdentityGovernanceEntitlementManagementAccessPackageCatalogIdAccessPackageId{}

// IdentityGovernanceEntitlementManagementAccessPackageCatalogIdAccessPackageId is a struct representing the Resource ID for a Identity Governance Entitlement Management Access Package Catalog Id Access Package
type IdentityGovernanceEntitlementManagementAccessPackageCatalogIdAccessPackageId struct {
	AccessPackageCatalogId string
	AccessPackageId        string
}

// NewIdentityGovernanceEntitlementManagementAccessPackageCatalogIdAccessPackageID returns a new IdentityGovernanceEntitlementManagementAccessPackageCatalogIdAccessPackageId struct
func NewIdentityGovernanceEntitlementManagementAccessPackageCatalogIdAccessPackageID(accessPackageCatalogId string, accessPackageId string) IdentityGovernanceEntitlementManagementAccessPackageCatalogIdAccessPackageId {
	return IdentityGovernanceEntitlementManagementAccessPackageCatalogIdAccessPackageId{
		AccessPackageCatalogId: accessPackageCatalogId,
		AccessPackageId:        accessPackageId,
	}
}

// ParseIdentityGovernanceEntitlementManagementAccessPackageCatalogIdAccessPackageID parses 'input' into a IdentityGovernanceEntitlementManagementAccessPackageCatalogIdAccessPackageId
func ParseIdentityGovernanceEntitlementManagementAccessPackageCatalogIdAccessPackageID(input string) (*IdentityGovernanceEntitlementManagementAccessPackageCatalogIdAccessPackageId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityGovernanceEntitlementManagementAccessPackageCatalogIdAccessPackageId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityGovernanceEntitlementManagementAccessPackageCatalogIdAccessPackageId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseIdentityGovernanceEntitlementManagementAccessPackageCatalogIdAccessPackageIDInsensitively parses 'input' case-insensitively into a IdentityGovernanceEntitlementManagementAccessPackageCatalogIdAccessPackageId
// note: this method should only be used for API response data and not user input
func ParseIdentityGovernanceEntitlementManagementAccessPackageCatalogIdAccessPackageIDInsensitively(input string) (*IdentityGovernanceEntitlementManagementAccessPackageCatalogIdAccessPackageId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityGovernanceEntitlementManagementAccessPackageCatalogIdAccessPackageId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityGovernanceEntitlementManagementAccessPackageCatalogIdAccessPackageId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *IdentityGovernanceEntitlementManagementAccessPackageCatalogIdAccessPackageId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.AccessPackageCatalogId, ok = input.Parsed["accessPackageCatalogId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "accessPackageCatalogId", input)
	}

	if id.AccessPackageId, ok = input.Parsed["accessPackageId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "accessPackageId", input)
	}

	return nil
}

// ValidateIdentityGovernanceEntitlementManagementAccessPackageCatalogIdAccessPackageID checks that 'input' can be parsed as a Identity Governance Entitlement Management Access Package Catalog Id Access Package ID
func ValidateIdentityGovernanceEntitlementManagementAccessPackageCatalogIdAccessPackageID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseIdentityGovernanceEntitlementManagementAccessPackageCatalogIdAccessPackageID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Identity Governance Entitlement Management Access Package Catalog Id Access Package ID
func (id IdentityGovernanceEntitlementManagementAccessPackageCatalogIdAccessPackageId) ID() string {
	fmtString := "/identityGovernance/entitlementManagement/accessPackageCatalogs/%s/accessPackages/%s"
	return fmt.Sprintf(fmtString, id.AccessPackageCatalogId, id.AccessPackageId)
}

// Segments returns a slice of Resource ID Segments which comprise this Identity Governance Entitlement Management Access Package Catalog Id Access Package ID
func (id IdentityGovernanceEntitlementManagementAccessPackageCatalogIdAccessPackageId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("identityGovernance", "identityGovernance", "identityGovernance"),
		resourceids.StaticSegment("entitlementManagement", "entitlementManagement", "entitlementManagement"),
		resourceids.StaticSegment("accessPackageCatalogs", "accessPackageCatalogs", "accessPackageCatalogs"),
		resourceids.UserSpecifiedSegment("accessPackageCatalogId", "accessPackageCatalogId"),
		resourceids.StaticSegment("accessPackages", "accessPackages", "accessPackages"),
		resourceids.UserSpecifiedSegment("accessPackageId", "accessPackageId"),
	}
}

// String returns a human-readable description of this Identity Governance Entitlement Management Access Package Catalog Id Access Package ID
func (id IdentityGovernanceEntitlementManagementAccessPackageCatalogIdAccessPackageId) String() string {
	components := []string{
		fmt.Sprintf("Access Package Catalog: %q", id.AccessPackageCatalogId),
		fmt.Sprintf("Access Package: %q", id.AccessPackageId),
	}
	return fmt.Sprintf("Identity Governance Entitlement Management Access Package Catalog Id Access Package (%s)", strings.Join(components, "\n"))
}
