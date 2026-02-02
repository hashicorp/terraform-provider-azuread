package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &IdentityGovernanceEntitlementManagementAccessPackageCatalogIdAccessPackageResourceScopeIdAccessPackageResourceAccessPackageResourceRoleId{}

// IdentityGovernanceEntitlementManagementAccessPackageCatalogIdAccessPackageResourceScopeIdAccessPackageResourceAccessPackageResourceRoleId is a struct representing the Resource ID for a Identity Governance Entitlement Management Access Package Catalog Id Access Package Resource Scope Id Access Package Resource Access Package Resource Role
type IdentityGovernanceEntitlementManagementAccessPackageCatalogIdAccessPackageResourceScopeIdAccessPackageResourceAccessPackageResourceRoleId struct {
	AccessPackageCatalogId       string
	AccessPackageResourceScopeId string
	AccessPackageResourceRoleId  string
}

// NewIdentityGovernanceEntitlementManagementAccessPackageCatalogIdAccessPackageResourceScopeIdAccessPackageResourceAccessPackageResourceRoleID returns a new IdentityGovernanceEntitlementManagementAccessPackageCatalogIdAccessPackageResourceScopeIdAccessPackageResourceAccessPackageResourceRoleId struct
func NewIdentityGovernanceEntitlementManagementAccessPackageCatalogIdAccessPackageResourceScopeIdAccessPackageResourceAccessPackageResourceRoleID(accessPackageCatalogId string, accessPackageResourceScopeId string, accessPackageResourceRoleId string) IdentityGovernanceEntitlementManagementAccessPackageCatalogIdAccessPackageResourceScopeIdAccessPackageResourceAccessPackageResourceRoleId {
	return IdentityGovernanceEntitlementManagementAccessPackageCatalogIdAccessPackageResourceScopeIdAccessPackageResourceAccessPackageResourceRoleId{
		AccessPackageCatalogId:       accessPackageCatalogId,
		AccessPackageResourceScopeId: accessPackageResourceScopeId,
		AccessPackageResourceRoleId:  accessPackageResourceRoleId,
	}
}

// ParseIdentityGovernanceEntitlementManagementAccessPackageCatalogIdAccessPackageResourceScopeIdAccessPackageResourceAccessPackageResourceRoleID parses 'input' into a IdentityGovernanceEntitlementManagementAccessPackageCatalogIdAccessPackageResourceScopeIdAccessPackageResourceAccessPackageResourceRoleId
func ParseIdentityGovernanceEntitlementManagementAccessPackageCatalogIdAccessPackageResourceScopeIdAccessPackageResourceAccessPackageResourceRoleID(input string) (*IdentityGovernanceEntitlementManagementAccessPackageCatalogIdAccessPackageResourceScopeIdAccessPackageResourceAccessPackageResourceRoleId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityGovernanceEntitlementManagementAccessPackageCatalogIdAccessPackageResourceScopeIdAccessPackageResourceAccessPackageResourceRoleId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityGovernanceEntitlementManagementAccessPackageCatalogIdAccessPackageResourceScopeIdAccessPackageResourceAccessPackageResourceRoleId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseIdentityGovernanceEntitlementManagementAccessPackageCatalogIdAccessPackageResourceScopeIdAccessPackageResourceAccessPackageResourceRoleIDInsensitively parses 'input' case-insensitively into a IdentityGovernanceEntitlementManagementAccessPackageCatalogIdAccessPackageResourceScopeIdAccessPackageResourceAccessPackageResourceRoleId
// note: this method should only be used for API response data and not user input
func ParseIdentityGovernanceEntitlementManagementAccessPackageCatalogIdAccessPackageResourceScopeIdAccessPackageResourceAccessPackageResourceRoleIDInsensitively(input string) (*IdentityGovernanceEntitlementManagementAccessPackageCatalogIdAccessPackageResourceScopeIdAccessPackageResourceAccessPackageResourceRoleId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityGovernanceEntitlementManagementAccessPackageCatalogIdAccessPackageResourceScopeIdAccessPackageResourceAccessPackageResourceRoleId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityGovernanceEntitlementManagementAccessPackageCatalogIdAccessPackageResourceScopeIdAccessPackageResourceAccessPackageResourceRoleId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *IdentityGovernanceEntitlementManagementAccessPackageCatalogIdAccessPackageResourceScopeIdAccessPackageResourceAccessPackageResourceRoleId) FromParseResult(input resourceids.ParseResult) error {
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

// ValidateIdentityGovernanceEntitlementManagementAccessPackageCatalogIdAccessPackageResourceScopeIdAccessPackageResourceAccessPackageResourceRoleID checks that 'input' can be parsed as a Identity Governance Entitlement Management Access Package Catalog Id Access Package Resource Scope Id Access Package Resource Access Package Resource Role ID
func ValidateIdentityGovernanceEntitlementManagementAccessPackageCatalogIdAccessPackageResourceScopeIdAccessPackageResourceAccessPackageResourceRoleID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseIdentityGovernanceEntitlementManagementAccessPackageCatalogIdAccessPackageResourceScopeIdAccessPackageResourceAccessPackageResourceRoleID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Identity Governance Entitlement Management Access Package Catalog Id Access Package Resource Scope Id Access Package Resource Access Package Resource Role ID
func (id IdentityGovernanceEntitlementManagementAccessPackageCatalogIdAccessPackageResourceScopeIdAccessPackageResourceAccessPackageResourceRoleId) ID() string {
	fmtString := "/identityGovernance/entitlementManagement/accessPackageCatalogs/%s/accessPackageResourceScopes/%s/accessPackageResource/accessPackageResourceRoles/%s"
	return fmt.Sprintf(fmtString, id.AccessPackageCatalogId, id.AccessPackageResourceScopeId, id.AccessPackageResourceRoleId)
}

// Segments returns a slice of Resource ID Segments which comprise this Identity Governance Entitlement Management Access Package Catalog Id Access Package Resource Scope Id Access Package Resource Access Package Resource Role ID
func (id IdentityGovernanceEntitlementManagementAccessPackageCatalogIdAccessPackageResourceScopeIdAccessPackageResourceAccessPackageResourceRoleId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("identityGovernance", "identityGovernance", "identityGovernance"),
		resourceids.StaticSegment("entitlementManagement", "entitlementManagement", "entitlementManagement"),
		resourceids.StaticSegment("accessPackageCatalogs", "accessPackageCatalogs", "accessPackageCatalogs"),
		resourceids.UserSpecifiedSegment("accessPackageCatalogId", "accessPackageCatalogId"),
		resourceids.StaticSegment("accessPackageResourceScopes", "accessPackageResourceScopes", "accessPackageResourceScopes"),
		resourceids.UserSpecifiedSegment("accessPackageResourceScopeId", "accessPackageResourceScopeId"),
		resourceids.StaticSegment("accessPackageResource", "accessPackageResource", "accessPackageResource"),
		resourceids.StaticSegment("accessPackageResourceRoles", "accessPackageResourceRoles", "accessPackageResourceRoles"),
		resourceids.UserSpecifiedSegment("accessPackageResourceRoleId", "accessPackageResourceRoleId"),
	}
}

// String returns a human-readable description of this Identity Governance Entitlement Management Access Package Catalog Id Access Package Resource Scope Id Access Package Resource Access Package Resource Role ID
func (id IdentityGovernanceEntitlementManagementAccessPackageCatalogIdAccessPackageResourceScopeIdAccessPackageResourceAccessPackageResourceRoleId) String() string {
	components := []string{
		fmt.Sprintf("Access Package Catalog: %q", id.AccessPackageCatalogId),
		fmt.Sprintf("Access Package Resource Scope: %q", id.AccessPackageResourceScopeId),
		fmt.Sprintf("Access Package Resource Role: %q", id.AccessPackageResourceRoleId),
	}
	return fmt.Sprintf("Identity Governance Entitlement Management Access Package Catalog Id Access Package Resource Scope Id Access Package Resource Access Package Resource Role (%s)", strings.Join(components, "\n"))
}
