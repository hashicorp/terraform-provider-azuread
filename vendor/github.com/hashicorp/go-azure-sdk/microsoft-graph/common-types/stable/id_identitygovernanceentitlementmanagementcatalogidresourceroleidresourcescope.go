package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &IdentityGovernanceEntitlementManagementCatalogIdResourceRoleIdResourceScopeId{}

// IdentityGovernanceEntitlementManagementCatalogIdResourceRoleIdResourceScopeId is a struct representing the Resource ID for a Identity Governance Entitlement Management Catalog Id Resource Role Id Resource Scope
type IdentityGovernanceEntitlementManagementCatalogIdResourceRoleIdResourceScopeId struct {
	AccessPackageCatalogId       string
	AccessPackageResourceRoleId  string
	AccessPackageResourceScopeId string
}

// NewIdentityGovernanceEntitlementManagementCatalogIdResourceRoleIdResourceScopeID returns a new IdentityGovernanceEntitlementManagementCatalogIdResourceRoleIdResourceScopeId struct
func NewIdentityGovernanceEntitlementManagementCatalogIdResourceRoleIdResourceScopeID(accessPackageCatalogId string, accessPackageResourceRoleId string, accessPackageResourceScopeId string) IdentityGovernanceEntitlementManagementCatalogIdResourceRoleIdResourceScopeId {
	return IdentityGovernanceEntitlementManagementCatalogIdResourceRoleIdResourceScopeId{
		AccessPackageCatalogId:       accessPackageCatalogId,
		AccessPackageResourceRoleId:  accessPackageResourceRoleId,
		AccessPackageResourceScopeId: accessPackageResourceScopeId,
	}
}

// ParseIdentityGovernanceEntitlementManagementCatalogIdResourceRoleIdResourceScopeID parses 'input' into a IdentityGovernanceEntitlementManagementCatalogIdResourceRoleIdResourceScopeId
func ParseIdentityGovernanceEntitlementManagementCatalogIdResourceRoleIdResourceScopeID(input string) (*IdentityGovernanceEntitlementManagementCatalogIdResourceRoleIdResourceScopeId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityGovernanceEntitlementManagementCatalogIdResourceRoleIdResourceScopeId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityGovernanceEntitlementManagementCatalogIdResourceRoleIdResourceScopeId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseIdentityGovernanceEntitlementManagementCatalogIdResourceRoleIdResourceScopeIDInsensitively parses 'input' case-insensitively into a IdentityGovernanceEntitlementManagementCatalogIdResourceRoleIdResourceScopeId
// note: this method should only be used for API response data and not user input
func ParseIdentityGovernanceEntitlementManagementCatalogIdResourceRoleIdResourceScopeIDInsensitively(input string) (*IdentityGovernanceEntitlementManagementCatalogIdResourceRoleIdResourceScopeId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityGovernanceEntitlementManagementCatalogIdResourceRoleIdResourceScopeId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityGovernanceEntitlementManagementCatalogIdResourceRoleIdResourceScopeId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *IdentityGovernanceEntitlementManagementCatalogIdResourceRoleIdResourceScopeId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.AccessPackageCatalogId, ok = input.Parsed["accessPackageCatalogId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "accessPackageCatalogId", input)
	}

	if id.AccessPackageResourceRoleId, ok = input.Parsed["accessPackageResourceRoleId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "accessPackageResourceRoleId", input)
	}

	if id.AccessPackageResourceScopeId, ok = input.Parsed["accessPackageResourceScopeId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "accessPackageResourceScopeId", input)
	}

	return nil
}

// ValidateIdentityGovernanceEntitlementManagementCatalogIdResourceRoleIdResourceScopeID checks that 'input' can be parsed as a Identity Governance Entitlement Management Catalog Id Resource Role Id Resource Scope ID
func ValidateIdentityGovernanceEntitlementManagementCatalogIdResourceRoleIdResourceScopeID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseIdentityGovernanceEntitlementManagementCatalogIdResourceRoleIdResourceScopeID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Identity Governance Entitlement Management Catalog Id Resource Role Id Resource Scope ID
func (id IdentityGovernanceEntitlementManagementCatalogIdResourceRoleIdResourceScopeId) ID() string {
	fmtString := "/identityGovernance/entitlementManagement/catalogs/%s/resourceRoles/%s/resource/scopes/%s"
	return fmt.Sprintf(fmtString, id.AccessPackageCatalogId, id.AccessPackageResourceRoleId, id.AccessPackageResourceScopeId)
}

// Segments returns a slice of Resource ID Segments which comprise this Identity Governance Entitlement Management Catalog Id Resource Role Id Resource Scope ID
func (id IdentityGovernanceEntitlementManagementCatalogIdResourceRoleIdResourceScopeId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("identityGovernance", "identityGovernance", "identityGovernance"),
		resourceids.StaticSegment("entitlementManagement", "entitlementManagement", "entitlementManagement"),
		resourceids.StaticSegment("catalogs", "catalogs", "catalogs"),
		resourceids.UserSpecifiedSegment("accessPackageCatalogId", "accessPackageCatalogId"),
		resourceids.StaticSegment("resourceRoles", "resourceRoles", "resourceRoles"),
		resourceids.UserSpecifiedSegment("accessPackageResourceRoleId", "accessPackageResourceRoleId"),
		resourceids.StaticSegment("resource", "resource", "resource"),
		resourceids.StaticSegment("scopes", "scopes", "scopes"),
		resourceids.UserSpecifiedSegment("accessPackageResourceScopeId", "accessPackageResourceScopeId"),
	}
}

// String returns a human-readable description of this Identity Governance Entitlement Management Catalog Id Resource Role Id Resource Scope ID
func (id IdentityGovernanceEntitlementManagementCatalogIdResourceRoleIdResourceScopeId) String() string {
	components := []string{
		fmt.Sprintf("Access Package Catalog: %q", id.AccessPackageCatalogId),
		fmt.Sprintf("Access Package Resource Role: %q", id.AccessPackageResourceRoleId),
		fmt.Sprintf("Access Package Resource Scope: %q", id.AccessPackageResourceScopeId),
	}
	return fmt.Sprintf("Identity Governance Entitlement Management Catalog Id Resource Role Id Resource Scope (%s)", strings.Join(components, "\n"))
}
