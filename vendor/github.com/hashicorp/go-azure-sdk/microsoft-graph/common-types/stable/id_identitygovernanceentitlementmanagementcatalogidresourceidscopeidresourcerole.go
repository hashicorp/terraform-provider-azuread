package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &IdentityGovernanceEntitlementManagementCatalogIdResourceIdScopeIdResourceRoleId{}

// IdentityGovernanceEntitlementManagementCatalogIdResourceIdScopeIdResourceRoleId is a struct representing the Resource ID for a Identity Governance Entitlement Management Catalog Id Resource Id Scope Id Resource Role
type IdentityGovernanceEntitlementManagementCatalogIdResourceIdScopeIdResourceRoleId struct {
	AccessPackageCatalogId       string
	AccessPackageResourceId      string
	AccessPackageResourceScopeId string
	AccessPackageResourceRoleId  string
}

// NewIdentityGovernanceEntitlementManagementCatalogIdResourceIdScopeIdResourceRoleID returns a new IdentityGovernanceEntitlementManagementCatalogIdResourceIdScopeIdResourceRoleId struct
func NewIdentityGovernanceEntitlementManagementCatalogIdResourceIdScopeIdResourceRoleID(accessPackageCatalogId string, accessPackageResourceId string, accessPackageResourceScopeId string, accessPackageResourceRoleId string) IdentityGovernanceEntitlementManagementCatalogIdResourceIdScopeIdResourceRoleId {
	return IdentityGovernanceEntitlementManagementCatalogIdResourceIdScopeIdResourceRoleId{
		AccessPackageCatalogId:       accessPackageCatalogId,
		AccessPackageResourceId:      accessPackageResourceId,
		AccessPackageResourceScopeId: accessPackageResourceScopeId,
		AccessPackageResourceRoleId:  accessPackageResourceRoleId,
	}
}

// ParseIdentityGovernanceEntitlementManagementCatalogIdResourceIdScopeIdResourceRoleID parses 'input' into a IdentityGovernanceEntitlementManagementCatalogIdResourceIdScopeIdResourceRoleId
func ParseIdentityGovernanceEntitlementManagementCatalogIdResourceIdScopeIdResourceRoleID(input string) (*IdentityGovernanceEntitlementManagementCatalogIdResourceIdScopeIdResourceRoleId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityGovernanceEntitlementManagementCatalogIdResourceIdScopeIdResourceRoleId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityGovernanceEntitlementManagementCatalogIdResourceIdScopeIdResourceRoleId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseIdentityGovernanceEntitlementManagementCatalogIdResourceIdScopeIdResourceRoleIDInsensitively parses 'input' case-insensitively into a IdentityGovernanceEntitlementManagementCatalogIdResourceIdScopeIdResourceRoleId
// note: this method should only be used for API response data and not user input
func ParseIdentityGovernanceEntitlementManagementCatalogIdResourceIdScopeIdResourceRoleIDInsensitively(input string) (*IdentityGovernanceEntitlementManagementCatalogIdResourceIdScopeIdResourceRoleId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityGovernanceEntitlementManagementCatalogIdResourceIdScopeIdResourceRoleId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityGovernanceEntitlementManagementCatalogIdResourceIdScopeIdResourceRoleId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *IdentityGovernanceEntitlementManagementCatalogIdResourceIdScopeIdResourceRoleId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.AccessPackageCatalogId, ok = input.Parsed["accessPackageCatalogId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "accessPackageCatalogId", input)
	}

	if id.AccessPackageResourceId, ok = input.Parsed["accessPackageResourceId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "accessPackageResourceId", input)
	}

	if id.AccessPackageResourceScopeId, ok = input.Parsed["accessPackageResourceScopeId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "accessPackageResourceScopeId", input)
	}

	if id.AccessPackageResourceRoleId, ok = input.Parsed["accessPackageResourceRoleId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "accessPackageResourceRoleId", input)
	}

	return nil
}

// ValidateIdentityGovernanceEntitlementManagementCatalogIdResourceIdScopeIdResourceRoleID checks that 'input' can be parsed as a Identity Governance Entitlement Management Catalog Id Resource Id Scope Id Resource Role ID
func ValidateIdentityGovernanceEntitlementManagementCatalogIdResourceIdScopeIdResourceRoleID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseIdentityGovernanceEntitlementManagementCatalogIdResourceIdScopeIdResourceRoleID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Identity Governance Entitlement Management Catalog Id Resource Id Scope Id Resource Role ID
func (id IdentityGovernanceEntitlementManagementCatalogIdResourceIdScopeIdResourceRoleId) ID() string {
	fmtString := "/identityGovernance/entitlementManagement/catalogs/%s/resources/%s/scopes/%s/resource/roles/%s"
	return fmt.Sprintf(fmtString, id.AccessPackageCatalogId, id.AccessPackageResourceId, id.AccessPackageResourceScopeId, id.AccessPackageResourceRoleId)
}

// Segments returns a slice of Resource ID Segments which comprise this Identity Governance Entitlement Management Catalog Id Resource Id Scope Id Resource Role ID
func (id IdentityGovernanceEntitlementManagementCatalogIdResourceIdScopeIdResourceRoleId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("identityGovernance", "identityGovernance", "identityGovernance"),
		resourceids.StaticSegment("entitlementManagement", "entitlementManagement", "entitlementManagement"),
		resourceids.StaticSegment("catalogs", "catalogs", "catalogs"),
		resourceids.UserSpecifiedSegment("accessPackageCatalogId", "accessPackageCatalogId"),
		resourceids.StaticSegment("resources", "resources", "resources"),
		resourceids.UserSpecifiedSegment("accessPackageResourceId", "accessPackageResourceId"),
		resourceids.StaticSegment("scopes", "scopes", "scopes"),
		resourceids.UserSpecifiedSegment("accessPackageResourceScopeId", "accessPackageResourceScopeId"),
		resourceids.StaticSegment("resource", "resource", "resource"),
		resourceids.StaticSegment("roles", "roles", "roles"),
		resourceids.UserSpecifiedSegment("accessPackageResourceRoleId", "accessPackageResourceRoleId"),
	}
}

// String returns a human-readable description of this Identity Governance Entitlement Management Catalog Id Resource Id Scope Id Resource Role ID
func (id IdentityGovernanceEntitlementManagementCatalogIdResourceIdScopeIdResourceRoleId) String() string {
	components := []string{
		fmt.Sprintf("Access Package Catalog: %q", id.AccessPackageCatalogId),
		fmt.Sprintf("Access Package Resource: %q", id.AccessPackageResourceId),
		fmt.Sprintf("Access Package Resource Scope: %q", id.AccessPackageResourceScopeId),
		fmt.Sprintf("Access Package Resource Role: %q", id.AccessPackageResourceRoleId),
	}
	return fmt.Sprintf("Identity Governance Entitlement Management Catalog Id Resource Id Scope Id Resource Role (%s)", strings.Join(components, "\n"))
}
