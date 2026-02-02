package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &IdentityGovernanceEntitlementManagementCatalogId{}

// IdentityGovernanceEntitlementManagementCatalogId is a struct representing the Resource ID for a Identity Governance Entitlement Management Catalog
type IdentityGovernanceEntitlementManagementCatalogId struct {
	AccessPackageCatalogId string
}

// NewIdentityGovernanceEntitlementManagementCatalogID returns a new IdentityGovernanceEntitlementManagementCatalogId struct
func NewIdentityGovernanceEntitlementManagementCatalogID(accessPackageCatalogId string) IdentityGovernanceEntitlementManagementCatalogId {
	return IdentityGovernanceEntitlementManagementCatalogId{
		AccessPackageCatalogId: accessPackageCatalogId,
	}
}

// ParseIdentityGovernanceEntitlementManagementCatalogID parses 'input' into a IdentityGovernanceEntitlementManagementCatalogId
func ParseIdentityGovernanceEntitlementManagementCatalogID(input string) (*IdentityGovernanceEntitlementManagementCatalogId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityGovernanceEntitlementManagementCatalogId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityGovernanceEntitlementManagementCatalogId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseIdentityGovernanceEntitlementManagementCatalogIDInsensitively parses 'input' case-insensitively into a IdentityGovernanceEntitlementManagementCatalogId
// note: this method should only be used for API response data and not user input
func ParseIdentityGovernanceEntitlementManagementCatalogIDInsensitively(input string) (*IdentityGovernanceEntitlementManagementCatalogId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityGovernanceEntitlementManagementCatalogId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityGovernanceEntitlementManagementCatalogId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *IdentityGovernanceEntitlementManagementCatalogId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.AccessPackageCatalogId, ok = input.Parsed["accessPackageCatalogId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "accessPackageCatalogId", input)
	}

	return nil
}

// ValidateIdentityGovernanceEntitlementManagementCatalogID checks that 'input' can be parsed as a Identity Governance Entitlement Management Catalog ID
func ValidateIdentityGovernanceEntitlementManagementCatalogID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseIdentityGovernanceEntitlementManagementCatalogID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Identity Governance Entitlement Management Catalog ID
func (id IdentityGovernanceEntitlementManagementCatalogId) ID() string {
	fmtString := "/identityGovernance/entitlementManagement/catalogs/%s"
	return fmt.Sprintf(fmtString, id.AccessPackageCatalogId)
}

// Segments returns a slice of Resource ID Segments which comprise this Identity Governance Entitlement Management Catalog ID
func (id IdentityGovernanceEntitlementManagementCatalogId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("identityGovernance", "identityGovernance", "identityGovernance"),
		resourceids.StaticSegment("entitlementManagement", "entitlementManagement", "entitlementManagement"),
		resourceids.StaticSegment("catalogs", "catalogs", "catalogs"),
		resourceids.UserSpecifiedSegment("accessPackageCatalogId", "accessPackageCatalogId"),
	}
}

// String returns a human-readable description of this Identity Governance Entitlement Management Catalog ID
func (id IdentityGovernanceEntitlementManagementCatalogId) String() string {
	components := []string{
		fmt.Sprintf("Access Package Catalog: %q", id.AccessPackageCatalogId),
	}
	return fmt.Sprintf("Identity Governance Entitlement Management Catalog (%s)", strings.Join(components, "\n"))
}
