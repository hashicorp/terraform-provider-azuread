package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &IdentityGovernanceEntitlementManagementCatalogIdResourceScopeIdResourceRoleId{}

// IdentityGovernanceEntitlementManagementCatalogIdResourceScopeIdResourceRoleId is a struct representing the Resource ID for a Identity Governance Entitlement Management Catalog Id Resource Scope Id Resource Role
type IdentityGovernanceEntitlementManagementCatalogIdResourceScopeIdResourceRoleId struct {
	AccessPackageCatalogId       string
	AccessPackageResourceScopeId string
	AccessPackageResourceRoleId  string
}

// NewIdentityGovernanceEntitlementManagementCatalogIdResourceScopeIdResourceRoleID returns a new IdentityGovernanceEntitlementManagementCatalogIdResourceScopeIdResourceRoleId struct
func NewIdentityGovernanceEntitlementManagementCatalogIdResourceScopeIdResourceRoleID(accessPackageCatalogId string, accessPackageResourceScopeId string, accessPackageResourceRoleId string) IdentityGovernanceEntitlementManagementCatalogIdResourceScopeIdResourceRoleId {
	return IdentityGovernanceEntitlementManagementCatalogIdResourceScopeIdResourceRoleId{
		AccessPackageCatalogId:       accessPackageCatalogId,
		AccessPackageResourceScopeId: accessPackageResourceScopeId,
		AccessPackageResourceRoleId:  accessPackageResourceRoleId,
	}
}

// ParseIdentityGovernanceEntitlementManagementCatalogIdResourceScopeIdResourceRoleID parses 'input' into a IdentityGovernanceEntitlementManagementCatalogIdResourceScopeIdResourceRoleId
func ParseIdentityGovernanceEntitlementManagementCatalogIdResourceScopeIdResourceRoleID(input string) (*IdentityGovernanceEntitlementManagementCatalogIdResourceScopeIdResourceRoleId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityGovernanceEntitlementManagementCatalogIdResourceScopeIdResourceRoleId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityGovernanceEntitlementManagementCatalogIdResourceScopeIdResourceRoleId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseIdentityGovernanceEntitlementManagementCatalogIdResourceScopeIdResourceRoleIDInsensitively parses 'input' case-insensitively into a IdentityGovernanceEntitlementManagementCatalogIdResourceScopeIdResourceRoleId
// note: this method should only be used for API response data and not user input
func ParseIdentityGovernanceEntitlementManagementCatalogIdResourceScopeIdResourceRoleIDInsensitively(input string) (*IdentityGovernanceEntitlementManagementCatalogIdResourceScopeIdResourceRoleId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityGovernanceEntitlementManagementCatalogIdResourceScopeIdResourceRoleId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityGovernanceEntitlementManagementCatalogIdResourceScopeIdResourceRoleId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *IdentityGovernanceEntitlementManagementCatalogIdResourceScopeIdResourceRoleId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.AccessPackageCatalogId, ok = input.Parsed["accessPackageCatalogId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "accessPackageCatalogId", input)
	}

	if id.AccessPackageResourceScopeId, ok = input.Parsed["accessPackageResourceScopeId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "accessPackageResourceScopeId", input)
	}

	if id.AccessPackageResourceRoleId, ok = input.Parsed["accessPackageResourceRoleId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "accessPackageResourceRoleId", input)
	}

	return nil
}

// ValidateIdentityGovernanceEntitlementManagementCatalogIdResourceScopeIdResourceRoleID checks that 'input' can be parsed as a Identity Governance Entitlement Management Catalog Id Resource Scope Id Resource Role ID
func ValidateIdentityGovernanceEntitlementManagementCatalogIdResourceScopeIdResourceRoleID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseIdentityGovernanceEntitlementManagementCatalogIdResourceScopeIdResourceRoleID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Identity Governance Entitlement Management Catalog Id Resource Scope Id Resource Role ID
func (id IdentityGovernanceEntitlementManagementCatalogIdResourceScopeIdResourceRoleId) ID() string {
	fmtString := "/identityGovernance/entitlementManagement/catalogs/%s/resourceScopes/%s/resource/roles/%s"
	return fmt.Sprintf(fmtString, id.AccessPackageCatalogId, id.AccessPackageResourceScopeId, id.AccessPackageResourceRoleId)
}

// Segments returns a slice of Resource ID Segments which comprise this Identity Governance Entitlement Management Catalog Id Resource Scope Id Resource Role ID
func (id IdentityGovernanceEntitlementManagementCatalogIdResourceScopeIdResourceRoleId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("identityGovernance", "identityGovernance", "identityGovernance"),
		resourceids.StaticSegment("entitlementManagement", "entitlementManagement", "entitlementManagement"),
		resourceids.StaticSegment("catalogs", "catalogs", "catalogs"),
		resourceids.UserSpecifiedSegment("accessPackageCatalogId", "accessPackageCatalogId"),
		resourceids.StaticSegment("resourceScopes", "resourceScopes", "resourceScopes"),
		resourceids.UserSpecifiedSegment("accessPackageResourceScopeId", "accessPackageResourceScopeId"),
		resourceids.StaticSegment("resource", "resource", "resource"),
		resourceids.StaticSegment("roles", "roles", "roles"),
		resourceids.UserSpecifiedSegment("accessPackageResourceRoleId", "accessPackageResourceRoleId"),
	}
}

// String returns a human-readable description of this Identity Governance Entitlement Management Catalog Id Resource Scope Id Resource Role ID
func (id IdentityGovernanceEntitlementManagementCatalogIdResourceScopeIdResourceRoleId) String() string {
	components := []string{
		fmt.Sprintf("Access Package Catalog: %q", id.AccessPackageCatalogId),
		fmt.Sprintf("Access Package Resource Scope: %q", id.AccessPackageResourceScopeId),
		fmt.Sprintf("Access Package Resource Role: %q", id.AccessPackageResourceRoleId),
	}
	return fmt.Sprintf("Identity Governance Entitlement Management Catalog Id Resource Scope Id Resource Role (%s)", strings.Join(components, "\n"))
}
