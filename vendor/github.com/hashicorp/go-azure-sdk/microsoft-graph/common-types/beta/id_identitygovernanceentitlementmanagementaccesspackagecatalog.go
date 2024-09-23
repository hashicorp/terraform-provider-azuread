package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &IdentityGovernanceEntitlementManagementAccessPackageCatalogId{}

// IdentityGovernanceEntitlementManagementAccessPackageCatalogId is a struct representing the Resource ID for a Identity Governance Entitlement Management Access Package Catalog
type IdentityGovernanceEntitlementManagementAccessPackageCatalogId struct {
	AccessPackageCatalogId string
}

// NewIdentityGovernanceEntitlementManagementAccessPackageCatalogID returns a new IdentityGovernanceEntitlementManagementAccessPackageCatalogId struct
func NewIdentityGovernanceEntitlementManagementAccessPackageCatalogID(accessPackageCatalogId string) IdentityGovernanceEntitlementManagementAccessPackageCatalogId {
	return IdentityGovernanceEntitlementManagementAccessPackageCatalogId{
		AccessPackageCatalogId: accessPackageCatalogId,
	}
}

// ParseIdentityGovernanceEntitlementManagementAccessPackageCatalogID parses 'input' into a IdentityGovernanceEntitlementManagementAccessPackageCatalogId
func ParseIdentityGovernanceEntitlementManagementAccessPackageCatalogID(input string) (*IdentityGovernanceEntitlementManagementAccessPackageCatalogId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityGovernanceEntitlementManagementAccessPackageCatalogId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityGovernanceEntitlementManagementAccessPackageCatalogId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseIdentityGovernanceEntitlementManagementAccessPackageCatalogIDInsensitively parses 'input' case-insensitively into a IdentityGovernanceEntitlementManagementAccessPackageCatalogId
// note: this method should only be used for API response data and not user input
func ParseIdentityGovernanceEntitlementManagementAccessPackageCatalogIDInsensitively(input string) (*IdentityGovernanceEntitlementManagementAccessPackageCatalogId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityGovernanceEntitlementManagementAccessPackageCatalogId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityGovernanceEntitlementManagementAccessPackageCatalogId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *IdentityGovernanceEntitlementManagementAccessPackageCatalogId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.AccessPackageCatalogId, ok = input.Parsed["accessPackageCatalogId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "accessPackageCatalogId", input)
	}

	return nil
}

// ValidateIdentityGovernanceEntitlementManagementAccessPackageCatalogID checks that 'input' can be parsed as a Identity Governance Entitlement Management Access Package Catalog ID
func ValidateIdentityGovernanceEntitlementManagementAccessPackageCatalogID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseIdentityGovernanceEntitlementManagementAccessPackageCatalogID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Identity Governance Entitlement Management Access Package Catalog ID
func (id IdentityGovernanceEntitlementManagementAccessPackageCatalogId) ID() string {
	fmtString := "/identityGovernance/entitlementManagement/accessPackageCatalogs/%s"
	return fmt.Sprintf(fmtString, id.AccessPackageCatalogId)
}

// Segments returns a slice of Resource ID Segments which comprise this Identity Governance Entitlement Management Access Package Catalog ID
func (id IdentityGovernanceEntitlementManagementAccessPackageCatalogId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("identityGovernance", "identityGovernance", "identityGovernance"),
		resourceids.StaticSegment("entitlementManagement", "entitlementManagement", "entitlementManagement"),
		resourceids.StaticSegment("accessPackageCatalogs", "accessPackageCatalogs", "accessPackageCatalogs"),
		resourceids.UserSpecifiedSegment("accessPackageCatalogId", "accessPackageCatalogId"),
	}
}

// String returns a human-readable description of this Identity Governance Entitlement Management Access Package Catalog ID
func (id IdentityGovernanceEntitlementManagementAccessPackageCatalogId) String() string {
	components := []string{
		fmt.Sprintf("Access Package Catalog: %q", id.AccessPackageCatalogId),
	}
	return fmt.Sprintf("Identity Governance Entitlement Management Access Package Catalog (%s)", strings.Join(components, "\n"))
}
