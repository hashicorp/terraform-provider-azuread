package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &IdentityGovernanceEntitlementManagementResourceRequestIdCatalogAccessPackageId{}

// IdentityGovernanceEntitlementManagementResourceRequestIdCatalogAccessPackageId is a struct representing the Resource ID for a Identity Governance Entitlement Management Resource Request Id Catalog Access Package
type IdentityGovernanceEntitlementManagementResourceRequestIdCatalogAccessPackageId struct {
	AccessPackageResourceRequestId string
	AccessPackageId                string
}

// NewIdentityGovernanceEntitlementManagementResourceRequestIdCatalogAccessPackageID returns a new IdentityGovernanceEntitlementManagementResourceRequestIdCatalogAccessPackageId struct
func NewIdentityGovernanceEntitlementManagementResourceRequestIdCatalogAccessPackageID(accessPackageResourceRequestId string, accessPackageId string) IdentityGovernanceEntitlementManagementResourceRequestIdCatalogAccessPackageId {
	return IdentityGovernanceEntitlementManagementResourceRequestIdCatalogAccessPackageId{
		AccessPackageResourceRequestId: accessPackageResourceRequestId,
		AccessPackageId:                accessPackageId,
	}
}

// ParseIdentityGovernanceEntitlementManagementResourceRequestIdCatalogAccessPackageID parses 'input' into a IdentityGovernanceEntitlementManagementResourceRequestIdCatalogAccessPackageId
func ParseIdentityGovernanceEntitlementManagementResourceRequestIdCatalogAccessPackageID(input string) (*IdentityGovernanceEntitlementManagementResourceRequestIdCatalogAccessPackageId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityGovernanceEntitlementManagementResourceRequestIdCatalogAccessPackageId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityGovernanceEntitlementManagementResourceRequestIdCatalogAccessPackageId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseIdentityGovernanceEntitlementManagementResourceRequestIdCatalogAccessPackageIDInsensitively parses 'input' case-insensitively into a IdentityGovernanceEntitlementManagementResourceRequestIdCatalogAccessPackageId
// note: this method should only be used for API response data and not user input
func ParseIdentityGovernanceEntitlementManagementResourceRequestIdCatalogAccessPackageIDInsensitively(input string) (*IdentityGovernanceEntitlementManagementResourceRequestIdCatalogAccessPackageId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityGovernanceEntitlementManagementResourceRequestIdCatalogAccessPackageId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityGovernanceEntitlementManagementResourceRequestIdCatalogAccessPackageId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *IdentityGovernanceEntitlementManagementResourceRequestIdCatalogAccessPackageId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.AccessPackageResourceRequestId, ok = input.Parsed["accessPackageResourceRequestId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "accessPackageResourceRequestId", input)
	}

	if id.AccessPackageId, ok = input.Parsed["accessPackageId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "accessPackageId", input)
	}

	return nil
}

// ValidateIdentityGovernanceEntitlementManagementResourceRequestIdCatalogAccessPackageID checks that 'input' can be parsed as a Identity Governance Entitlement Management Resource Request Id Catalog Access Package ID
func ValidateIdentityGovernanceEntitlementManagementResourceRequestIdCatalogAccessPackageID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseIdentityGovernanceEntitlementManagementResourceRequestIdCatalogAccessPackageID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Identity Governance Entitlement Management Resource Request Id Catalog Access Package ID
func (id IdentityGovernanceEntitlementManagementResourceRequestIdCatalogAccessPackageId) ID() string {
	fmtString := "/identityGovernance/entitlementManagement/resourceRequests/%s/catalog/accessPackages/%s"
	return fmt.Sprintf(fmtString, id.AccessPackageResourceRequestId, id.AccessPackageId)
}

// Segments returns a slice of Resource ID Segments which comprise this Identity Governance Entitlement Management Resource Request Id Catalog Access Package ID
func (id IdentityGovernanceEntitlementManagementResourceRequestIdCatalogAccessPackageId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("identityGovernance", "identityGovernance", "identityGovernance"),
		resourceids.StaticSegment("entitlementManagement", "entitlementManagement", "entitlementManagement"),
		resourceids.StaticSegment("resourceRequests", "resourceRequests", "resourceRequests"),
		resourceids.UserSpecifiedSegment("accessPackageResourceRequestId", "accessPackageResourceRequestId"),
		resourceids.StaticSegment("catalog", "catalog", "catalog"),
		resourceids.StaticSegment("accessPackages", "accessPackages", "accessPackages"),
		resourceids.UserSpecifiedSegment("accessPackageId", "accessPackageId"),
	}
}

// String returns a human-readable description of this Identity Governance Entitlement Management Resource Request Id Catalog Access Package ID
func (id IdentityGovernanceEntitlementManagementResourceRequestIdCatalogAccessPackageId) String() string {
	components := []string{
		fmt.Sprintf("Access Package Resource Request: %q", id.AccessPackageResourceRequestId),
		fmt.Sprintf("Access Package: %q", id.AccessPackageId),
	}
	return fmt.Sprintf("Identity Governance Entitlement Management Resource Request Id Catalog Access Package (%s)", strings.Join(components, "\n"))
}
