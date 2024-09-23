package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &IdentityGovernanceEntitlementManagementCatalogIdResourceId{}

// IdentityGovernanceEntitlementManagementCatalogIdResourceId is a struct representing the Resource ID for a Identity Governance Entitlement Management Catalog Id Resource
type IdentityGovernanceEntitlementManagementCatalogIdResourceId struct {
	AccessPackageCatalogId  string
	AccessPackageResourceId string
}

// NewIdentityGovernanceEntitlementManagementCatalogIdResourceID returns a new IdentityGovernanceEntitlementManagementCatalogIdResourceId struct
func NewIdentityGovernanceEntitlementManagementCatalogIdResourceID(accessPackageCatalogId string, accessPackageResourceId string) IdentityGovernanceEntitlementManagementCatalogIdResourceId {
	return IdentityGovernanceEntitlementManagementCatalogIdResourceId{
		AccessPackageCatalogId:  accessPackageCatalogId,
		AccessPackageResourceId: accessPackageResourceId,
	}
}

// ParseIdentityGovernanceEntitlementManagementCatalogIdResourceID parses 'input' into a IdentityGovernanceEntitlementManagementCatalogIdResourceId
func ParseIdentityGovernanceEntitlementManagementCatalogIdResourceID(input string) (*IdentityGovernanceEntitlementManagementCatalogIdResourceId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityGovernanceEntitlementManagementCatalogIdResourceId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityGovernanceEntitlementManagementCatalogIdResourceId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseIdentityGovernanceEntitlementManagementCatalogIdResourceIDInsensitively parses 'input' case-insensitively into a IdentityGovernanceEntitlementManagementCatalogIdResourceId
// note: this method should only be used for API response data and not user input
func ParseIdentityGovernanceEntitlementManagementCatalogIdResourceIDInsensitively(input string) (*IdentityGovernanceEntitlementManagementCatalogIdResourceId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityGovernanceEntitlementManagementCatalogIdResourceId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityGovernanceEntitlementManagementCatalogIdResourceId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *IdentityGovernanceEntitlementManagementCatalogIdResourceId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.AccessPackageCatalogId, ok = input.Parsed["accessPackageCatalogId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "accessPackageCatalogId", input)
	}

	if id.AccessPackageResourceId, ok = input.Parsed["accessPackageResourceId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "accessPackageResourceId", input)
	}

	return nil
}

// ValidateIdentityGovernanceEntitlementManagementCatalogIdResourceID checks that 'input' can be parsed as a Identity Governance Entitlement Management Catalog Id Resource ID
func ValidateIdentityGovernanceEntitlementManagementCatalogIdResourceID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseIdentityGovernanceEntitlementManagementCatalogIdResourceID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Identity Governance Entitlement Management Catalog Id Resource ID
func (id IdentityGovernanceEntitlementManagementCatalogIdResourceId) ID() string {
	fmtString := "/identityGovernance/entitlementManagement/catalogs/%s/resources/%s"
	return fmt.Sprintf(fmtString, id.AccessPackageCatalogId, id.AccessPackageResourceId)
}

// Segments returns a slice of Resource ID Segments which comprise this Identity Governance Entitlement Management Catalog Id Resource ID
func (id IdentityGovernanceEntitlementManagementCatalogIdResourceId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("identityGovernance", "identityGovernance", "identityGovernance"),
		resourceids.StaticSegment("entitlementManagement", "entitlementManagement", "entitlementManagement"),
		resourceids.StaticSegment("catalogs", "catalogs", "catalogs"),
		resourceids.UserSpecifiedSegment("accessPackageCatalogId", "accessPackageCatalogId"),
		resourceids.StaticSegment("resources", "resources", "resources"),
		resourceids.UserSpecifiedSegment("accessPackageResourceId", "accessPackageResourceId"),
	}
}

// String returns a human-readable description of this Identity Governance Entitlement Management Catalog Id Resource ID
func (id IdentityGovernanceEntitlementManagementCatalogIdResourceId) String() string {
	components := []string{
		fmt.Sprintf("Access Package Catalog: %q", id.AccessPackageCatalogId),
		fmt.Sprintf("Access Package Resource: %q", id.AccessPackageResourceId),
	}
	return fmt.Sprintf("Identity Governance Entitlement Management Catalog Id Resource (%s)", strings.Join(components, "\n"))
}
