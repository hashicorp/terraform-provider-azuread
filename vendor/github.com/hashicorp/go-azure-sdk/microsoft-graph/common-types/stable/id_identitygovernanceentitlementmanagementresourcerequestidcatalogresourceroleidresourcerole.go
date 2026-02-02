package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &IdentityGovernanceEntitlementManagementResourceRequestIdCatalogResourceRoleIdResourceRoleId{}

// IdentityGovernanceEntitlementManagementResourceRequestIdCatalogResourceRoleIdResourceRoleId is a struct representing the Resource ID for a Identity Governance Entitlement Management Resource Request Id Catalog Resource Role Id Resource Role
type IdentityGovernanceEntitlementManagementResourceRequestIdCatalogResourceRoleIdResourceRoleId struct {
	AccessPackageResourceRequestId string
	AccessPackageResourceRoleId    string
	AccessPackageResourceRoleId1   string
}

// NewIdentityGovernanceEntitlementManagementResourceRequestIdCatalogResourceRoleIdResourceRoleID returns a new IdentityGovernanceEntitlementManagementResourceRequestIdCatalogResourceRoleIdResourceRoleId struct
func NewIdentityGovernanceEntitlementManagementResourceRequestIdCatalogResourceRoleIdResourceRoleID(accessPackageResourceRequestId string, accessPackageResourceRoleId string, accessPackageResourceRoleId1 string) IdentityGovernanceEntitlementManagementResourceRequestIdCatalogResourceRoleIdResourceRoleId {
	return IdentityGovernanceEntitlementManagementResourceRequestIdCatalogResourceRoleIdResourceRoleId{
		AccessPackageResourceRequestId: accessPackageResourceRequestId,
		AccessPackageResourceRoleId:    accessPackageResourceRoleId,
		AccessPackageResourceRoleId1:   accessPackageResourceRoleId1,
	}
}

// ParseIdentityGovernanceEntitlementManagementResourceRequestIdCatalogResourceRoleIdResourceRoleID parses 'input' into a IdentityGovernanceEntitlementManagementResourceRequestIdCatalogResourceRoleIdResourceRoleId
func ParseIdentityGovernanceEntitlementManagementResourceRequestIdCatalogResourceRoleIdResourceRoleID(input string) (*IdentityGovernanceEntitlementManagementResourceRequestIdCatalogResourceRoleIdResourceRoleId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityGovernanceEntitlementManagementResourceRequestIdCatalogResourceRoleIdResourceRoleId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityGovernanceEntitlementManagementResourceRequestIdCatalogResourceRoleIdResourceRoleId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseIdentityGovernanceEntitlementManagementResourceRequestIdCatalogResourceRoleIdResourceRoleIDInsensitively parses 'input' case-insensitively into a IdentityGovernanceEntitlementManagementResourceRequestIdCatalogResourceRoleIdResourceRoleId
// note: this method should only be used for API response data and not user input
func ParseIdentityGovernanceEntitlementManagementResourceRequestIdCatalogResourceRoleIdResourceRoleIDInsensitively(input string) (*IdentityGovernanceEntitlementManagementResourceRequestIdCatalogResourceRoleIdResourceRoleId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityGovernanceEntitlementManagementResourceRequestIdCatalogResourceRoleIdResourceRoleId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityGovernanceEntitlementManagementResourceRequestIdCatalogResourceRoleIdResourceRoleId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *IdentityGovernanceEntitlementManagementResourceRequestIdCatalogResourceRoleIdResourceRoleId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.AccessPackageResourceRequestId, ok = input.Parsed["accessPackageResourceRequestId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "accessPackageResourceRequestId", input)
	}

	if id.AccessPackageResourceRoleId, ok = input.Parsed["accessPackageResourceRoleId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "accessPackageResourceRoleId", input)
	}

	if id.AccessPackageResourceRoleId1, ok = input.Parsed["accessPackageResourceRoleId1"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "accessPackageResourceRoleId1", input)
	}

	return nil
}

// ValidateIdentityGovernanceEntitlementManagementResourceRequestIdCatalogResourceRoleIdResourceRoleID checks that 'input' can be parsed as a Identity Governance Entitlement Management Resource Request Id Catalog Resource Role Id Resource Role ID
func ValidateIdentityGovernanceEntitlementManagementResourceRequestIdCatalogResourceRoleIdResourceRoleID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseIdentityGovernanceEntitlementManagementResourceRequestIdCatalogResourceRoleIdResourceRoleID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Identity Governance Entitlement Management Resource Request Id Catalog Resource Role Id Resource Role ID
func (id IdentityGovernanceEntitlementManagementResourceRequestIdCatalogResourceRoleIdResourceRoleId) ID() string {
	fmtString := "/identityGovernance/entitlementManagement/resourceRequests/%s/catalog/resourceRoles/%s/resource/roles/%s"
	return fmt.Sprintf(fmtString, id.AccessPackageResourceRequestId, id.AccessPackageResourceRoleId, id.AccessPackageResourceRoleId1)
}

// Segments returns a slice of Resource ID Segments which comprise this Identity Governance Entitlement Management Resource Request Id Catalog Resource Role Id Resource Role ID
func (id IdentityGovernanceEntitlementManagementResourceRequestIdCatalogResourceRoleIdResourceRoleId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("identityGovernance", "identityGovernance", "identityGovernance"),
		resourceids.StaticSegment("entitlementManagement", "entitlementManagement", "entitlementManagement"),
		resourceids.StaticSegment("resourceRequests", "resourceRequests", "resourceRequests"),
		resourceids.UserSpecifiedSegment("accessPackageResourceRequestId", "accessPackageResourceRequestId"),
		resourceids.StaticSegment("catalog", "catalog", "catalog"),
		resourceids.StaticSegment("resourceRoles", "resourceRoles", "resourceRoles"),
		resourceids.UserSpecifiedSegment("accessPackageResourceRoleId", "accessPackageResourceRoleId"),
		resourceids.StaticSegment("resource", "resource", "resource"),
		resourceids.StaticSegment("roles", "roles", "roles"),
		resourceids.UserSpecifiedSegment("accessPackageResourceRoleId1", "accessPackageResourceRoleId1"),
	}
}

// String returns a human-readable description of this Identity Governance Entitlement Management Resource Request Id Catalog Resource Role Id Resource Role ID
func (id IdentityGovernanceEntitlementManagementResourceRequestIdCatalogResourceRoleIdResourceRoleId) String() string {
	components := []string{
		fmt.Sprintf("Access Package Resource Request: %q", id.AccessPackageResourceRequestId),
		fmt.Sprintf("Access Package Resource Role: %q", id.AccessPackageResourceRoleId),
		fmt.Sprintf("Access Package Resource Role Id 1: %q", id.AccessPackageResourceRoleId1),
	}
	return fmt.Sprintf("Identity Governance Entitlement Management Resource Request Id Catalog Resource Role Id Resource Role (%s)", strings.Join(components, "\n"))
}
