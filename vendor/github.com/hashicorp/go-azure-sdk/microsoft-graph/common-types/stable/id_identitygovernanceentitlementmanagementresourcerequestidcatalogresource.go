package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &IdentityGovernanceEntitlementManagementResourceRequestIdCatalogResourceId{}

// IdentityGovernanceEntitlementManagementResourceRequestIdCatalogResourceId is a struct representing the Resource ID for a Identity Governance Entitlement Management Resource Request Id Catalog Resource
type IdentityGovernanceEntitlementManagementResourceRequestIdCatalogResourceId struct {
	AccessPackageResourceRequestId string
	AccessPackageResourceId        string
}

// NewIdentityGovernanceEntitlementManagementResourceRequestIdCatalogResourceID returns a new IdentityGovernanceEntitlementManagementResourceRequestIdCatalogResourceId struct
func NewIdentityGovernanceEntitlementManagementResourceRequestIdCatalogResourceID(accessPackageResourceRequestId string, accessPackageResourceId string) IdentityGovernanceEntitlementManagementResourceRequestIdCatalogResourceId {
	return IdentityGovernanceEntitlementManagementResourceRequestIdCatalogResourceId{
		AccessPackageResourceRequestId: accessPackageResourceRequestId,
		AccessPackageResourceId:        accessPackageResourceId,
	}
}

// ParseIdentityGovernanceEntitlementManagementResourceRequestIdCatalogResourceID parses 'input' into a IdentityGovernanceEntitlementManagementResourceRequestIdCatalogResourceId
func ParseIdentityGovernanceEntitlementManagementResourceRequestIdCatalogResourceID(input string) (*IdentityGovernanceEntitlementManagementResourceRequestIdCatalogResourceId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityGovernanceEntitlementManagementResourceRequestIdCatalogResourceId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityGovernanceEntitlementManagementResourceRequestIdCatalogResourceId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseIdentityGovernanceEntitlementManagementResourceRequestIdCatalogResourceIDInsensitively parses 'input' case-insensitively into a IdentityGovernanceEntitlementManagementResourceRequestIdCatalogResourceId
// note: this method should only be used for API response data and not user input
func ParseIdentityGovernanceEntitlementManagementResourceRequestIdCatalogResourceIDInsensitively(input string) (*IdentityGovernanceEntitlementManagementResourceRequestIdCatalogResourceId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityGovernanceEntitlementManagementResourceRequestIdCatalogResourceId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityGovernanceEntitlementManagementResourceRequestIdCatalogResourceId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *IdentityGovernanceEntitlementManagementResourceRequestIdCatalogResourceId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.AccessPackageResourceRequestId, ok = input.Parsed["accessPackageResourceRequestId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "accessPackageResourceRequestId", input)
	}

	if id.AccessPackageResourceId, ok = input.Parsed["accessPackageResourceId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "accessPackageResourceId", input)
	}

	return nil
}

// ValidateIdentityGovernanceEntitlementManagementResourceRequestIdCatalogResourceID checks that 'input' can be parsed as a Identity Governance Entitlement Management Resource Request Id Catalog Resource ID
func ValidateIdentityGovernanceEntitlementManagementResourceRequestIdCatalogResourceID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseIdentityGovernanceEntitlementManagementResourceRequestIdCatalogResourceID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Identity Governance Entitlement Management Resource Request Id Catalog Resource ID
func (id IdentityGovernanceEntitlementManagementResourceRequestIdCatalogResourceId) ID() string {
	fmtString := "/identityGovernance/entitlementManagement/resourceRequests/%s/catalog/resources/%s"
	return fmt.Sprintf(fmtString, id.AccessPackageResourceRequestId, id.AccessPackageResourceId)
}

// Segments returns a slice of Resource ID Segments which comprise this Identity Governance Entitlement Management Resource Request Id Catalog Resource ID
func (id IdentityGovernanceEntitlementManagementResourceRequestIdCatalogResourceId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("identityGovernance", "identityGovernance", "identityGovernance"),
		resourceids.StaticSegment("entitlementManagement", "entitlementManagement", "entitlementManagement"),
		resourceids.StaticSegment("resourceRequests", "resourceRequests", "resourceRequests"),
		resourceids.UserSpecifiedSegment("accessPackageResourceRequestId", "accessPackageResourceRequestId"),
		resourceids.StaticSegment("catalog", "catalog", "catalog"),
		resourceids.StaticSegment("resources", "resources", "resources"),
		resourceids.UserSpecifiedSegment("accessPackageResourceId", "accessPackageResourceId"),
	}
}

// String returns a human-readable description of this Identity Governance Entitlement Management Resource Request Id Catalog Resource ID
func (id IdentityGovernanceEntitlementManagementResourceRequestIdCatalogResourceId) String() string {
	components := []string{
		fmt.Sprintf("Access Package Resource Request: %q", id.AccessPackageResourceRequestId),
		fmt.Sprintf("Access Package Resource: %q", id.AccessPackageResourceId),
	}
	return fmt.Sprintf("Identity Governance Entitlement Management Resource Request Id Catalog Resource (%s)", strings.Join(components, "\n"))
}
