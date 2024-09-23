package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &IdentityGovernanceEntitlementManagementResourceRequestIdCatalogResourceIdRoleId{}

// IdentityGovernanceEntitlementManagementResourceRequestIdCatalogResourceIdRoleId is a struct representing the Resource ID for a Identity Governance Entitlement Management Resource Request Id Catalog Resource Id Role
type IdentityGovernanceEntitlementManagementResourceRequestIdCatalogResourceIdRoleId struct {
	AccessPackageResourceRequestId string
	AccessPackageResourceId        string
	AccessPackageResourceRoleId    string
}

// NewIdentityGovernanceEntitlementManagementResourceRequestIdCatalogResourceIdRoleID returns a new IdentityGovernanceEntitlementManagementResourceRequestIdCatalogResourceIdRoleId struct
func NewIdentityGovernanceEntitlementManagementResourceRequestIdCatalogResourceIdRoleID(accessPackageResourceRequestId string, accessPackageResourceId string, accessPackageResourceRoleId string) IdentityGovernanceEntitlementManagementResourceRequestIdCatalogResourceIdRoleId {
	return IdentityGovernanceEntitlementManagementResourceRequestIdCatalogResourceIdRoleId{
		AccessPackageResourceRequestId: accessPackageResourceRequestId,
		AccessPackageResourceId:        accessPackageResourceId,
		AccessPackageResourceRoleId:    accessPackageResourceRoleId,
	}
}

// ParseIdentityGovernanceEntitlementManagementResourceRequestIdCatalogResourceIdRoleID parses 'input' into a IdentityGovernanceEntitlementManagementResourceRequestIdCatalogResourceIdRoleId
func ParseIdentityGovernanceEntitlementManagementResourceRequestIdCatalogResourceIdRoleID(input string) (*IdentityGovernanceEntitlementManagementResourceRequestIdCatalogResourceIdRoleId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityGovernanceEntitlementManagementResourceRequestIdCatalogResourceIdRoleId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityGovernanceEntitlementManagementResourceRequestIdCatalogResourceIdRoleId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseIdentityGovernanceEntitlementManagementResourceRequestIdCatalogResourceIdRoleIDInsensitively parses 'input' case-insensitively into a IdentityGovernanceEntitlementManagementResourceRequestIdCatalogResourceIdRoleId
// note: this method should only be used for API response data and not user input
func ParseIdentityGovernanceEntitlementManagementResourceRequestIdCatalogResourceIdRoleIDInsensitively(input string) (*IdentityGovernanceEntitlementManagementResourceRequestIdCatalogResourceIdRoleId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityGovernanceEntitlementManagementResourceRequestIdCatalogResourceIdRoleId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityGovernanceEntitlementManagementResourceRequestIdCatalogResourceIdRoleId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *IdentityGovernanceEntitlementManagementResourceRequestIdCatalogResourceIdRoleId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.AccessPackageResourceRequestId, ok = input.Parsed["accessPackageResourceRequestId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "accessPackageResourceRequestId", input)
	}

	if id.AccessPackageResourceId, ok = input.Parsed["accessPackageResourceId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "accessPackageResourceId", input)
	}

	if id.AccessPackageResourceRoleId, ok = input.Parsed["accessPackageResourceRoleId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "accessPackageResourceRoleId", input)
	}

	return nil
}

// ValidateIdentityGovernanceEntitlementManagementResourceRequestIdCatalogResourceIdRoleID checks that 'input' can be parsed as a Identity Governance Entitlement Management Resource Request Id Catalog Resource Id Role ID
func ValidateIdentityGovernanceEntitlementManagementResourceRequestIdCatalogResourceIdRoleID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseIdentityGovernanceEntitlementManagementResourceRequestIdCatalogResourceIdRoleID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Identity Governance Entitlement Management Resource Request Id Catalog Resource Id Role ID
func (id IdentityGovernanceEntitlementManagementResourceRequestIdCatalogResourceIdRoleId) ID() string {
	fmtString := "/identityGovernance/entitlementManagement/resourceRequests/%s/catalog/resources/%s/roles/%s"
	return fmt.Sprintf(fmtString, id.AccessPackageResourceRequestId, id.AccessPackageResourceId, id.AccessPackageResourceRoleId)
}

// Segments returns a slice of Resource ID Segments which comprise this Identity Governance Entitlement Management Resource Request Id Catalog Resource Id Role ID
func (id IdentityGovernanceEntitlementManagementResourceRequestIdCatalogResourceIdRoleId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("identityGovernance", "identityGovernance", "identityGovernance"),
		resourceids.StaticSegment("entitlementManagement", "entitlementManagement", "entitlementManagement"),
		resourceids.StaticSegment("resourceRequests", "resourceRequests", "resourceRequests"),
		resourceids.UserSpecifiedSegment("accessPackageResourceRequestId", "accessPackageResourceRequestId"),
		resourceids.StaticSegment("catalog", "catalog", "catalog"),
		resourceids.StaticSegment("resources", "resources", "resources"),
		resourceids.UserSpecifiedSegment("accessPackageResourceId", "accessPackageResourceId"),
		resourceids.StaticSegment("roles", "roles", "roles"),
		resourceids.UserSpecifiedSegment("accessPackageResourceRoleId", "accessPackageResourceRoleId"),
	}
}

// String returns a human-readable description of this Identity Governance Entitlement Management Resource Request Id Catalog Resource Id Role ID
func (id IdentityGovernanceEntitlementManagementResourceRequestIdCatalogResourceIdRoleId) String() string {
	components := []string{
		fmt.Sprintf("Access Package Resource Request: %q", id.AccessPackageResourceRequestId),
		fmt.Sprintf("Access Package Resource: %q", id.AccessPackageResourceId),
		fmt.Sprintf("Access Package Resource Role: %q", id.AccessPackageResourceRoleId),
	}
	return fmt.Sprintf("Identity Governance Entitlement Management Resource Request Id Catalog Resource Id Role (%s)", strings.Join(components, "\n"))
}
