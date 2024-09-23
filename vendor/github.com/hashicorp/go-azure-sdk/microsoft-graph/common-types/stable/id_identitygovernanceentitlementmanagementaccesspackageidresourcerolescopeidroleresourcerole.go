package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &IdentityGovernanceEntitlementManagementAccessPackageIdResourceRoleScopeIdRoleResourceRoleId{}

// IdentityGovernanceEntitlementManagementAccessPackageIdResourceRoleScopeIdRoleResourceRoleId is a struct representing the Resource ID for a Identity Governance Entitlement Management Access Package Id Resource Role Scope Id Role Resource Role
type IdentityGovernanceEntitlementManagementAccessPackageIdResourceRoleScopeIdRoleResourceRoleId struct {
	AccessPackageId                  string
	AccessPackageResourceRoleScopeId string
	AccessPackageResourceRoleId      string
}

// NewIdentityGovernanceEntitlementManagementAccessPackageIdResourceRoleScopeIdRoleResourceRoleID returns a new IdentityGovernanceEntitlementManagementAccessPackageIdResourceRoleScopeIdRoleResourceRoleId struct
func NewIdentityGovernanceEntitlementManagementAccessPackageIdResourceRoleScopeIdRoleResourceRoleID(accessPackageId string, accessPackageResourceRoleScopeId string, accessPackageResourceRoleId string) IdentityGovernanceEntitlementManagementAccessPackageIdResourceRoleScopeIdRoleResourceRoleId {
	return IdentityGovernanceEntitlementManagementAccessPackageIdResourceRoleScopeIdRoleResourceRoleId{
		AccessPackageId:                  accessPackageId,
		AccessPackageResourceRoleScopeId: accessPackageResourceRoleScopeId,
		AccessPackageResourceRoleId:      accessPackageResourceRoleId,
	}
}

// ParseIdentityGovernanceEntitlementManagementAccessPackageIdResourceRoleScopeIdRoleResourceRoleID parses 'input' into a IdentityGovernanceEntitlementManagementAccessPackageIdResourceRoleScopeIdRoleResourceRoleId
func ParseIdentityGovernanceEntitlementManagementAccessPackageIdResourceRoleScopeIdRoleResourceRoleID(input string) (*IdentityGovernanceEntitlementManagementAccessPackageIdResourceRoleScopeIdRoleResourceRoleId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityGovernanceEntitlementManagementAccessPackageIdResourceRoleScopeIdRoleResourceRoleId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityGovernanceEntitlementManagementAccessPackageIdResourceRoleScopeIdRoleResourceRoleId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseIdentityGovernanceEntitlementManagementAccessPackageIdResourceRoleScopeIdRoleResourceRoleIDInsensitively parses 'input' case-insensitively into a IdentityGovernanceEntitlementManagementAccessPackageIdResourceRoleScopeIdRoleResourceRoleId
// note: this method should only be used for API response data and not user input
func ParseIdentityGovernanceEntitlementManagementAccessPackageIdResourceRoleScopeIdRoleResourceRoleIDInsensitively(input string) (*IdentityGovernanceEntitlementManagementAccessPackageIdResourceRoleScopeIdRoleResourceRoleId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityGovernanceEntitlementManagementAccessPackageIdResourceRoleScopeIdRoleResourceRoleId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityGovernanceEntitlementManagementAccessPackageIdResourceRoleScopeIdRoleResourceRoleId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *IdentityGovernanceEntitlementManagementAccessPackageIdResourceRoleScopeIdRoleResourceRoleId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.AccessPackageId, ok = input.Parsed["accessPackageId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "accessPackageId", input)
	}

	if id.AccessPackageResourceRoleScopeId, ok = input.Parsed["accessPackageResourceRoleScopeId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "accessPackageResourceRoleScopeId", input)
	}

	if id.AccessPackageResourceRoleId, ok = input.Parsed["accessPackageResourceRoleId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "accessPackageResourceRoleId", input)
	}

	return nil
}

// ValidateIdentityGovernanceEntitlementManagementAccessPackageIdResourceRoleScopeIdRoleResourceRoleID checks that 'input' can be parsed as a Identity Governance Entitlement Management Access Package Id Resource Role Scope Id Role Resource Role ID
func ValidateIdentityGovernanceEntitlementManagementAccessPackageIdResourceRoleScopeIdRoleResourceRoleID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseIdentityGovernanceEntitlementManagementAccessPackageIdResourceRoleScopeIdRoleResourceRoleID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Identity Governance Entitlement Management Access Package Id Resource Role Scope Id Role Resource Role ID
func (id IdentityGovernanceEntitlementManagementAccessPackageIdResourceRoleScopeIdRoleResourceRoleId) ID() string {
	fmtString := "/identityGovernance/entitlementManagement/accessPackages/%s/resourceRoleScopes/%s/role/resource/roles/%s"
	return fmt.Sprintf(fmtString, id.AccessPackageId, id.AccessPackageResourceRoleScopeId, id.AccessPackageResourceRoleId)
}

// Segments returns a slice of Resource ID Segments which comprise this Identity Governance Entitlement Management Access Package Id Resource Role Scope Id Role Resource Role ID
func (id IdentityGovernanceEntitlementManagementAccessPackageIdResourceRoleScopeIdRoleResourceRoleId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("identityGovernance", "identityGovernance", "identityGovernance"),
		resourceids.StaticSegment("entitlementManagement", "entitlementManagement", "entitlementManagement"),
		resourceids.StaticSegment("accessPackages", "accessPackages", "accessPackages"),
		resourceids.UserSpecifiedSegment("accessPackageId", "accessPackageId"),
		resourceids.StaticSegment("resourceRoleScopes", "resourceRoleScopes", "resourceRoleScopes"),
		resourceids.UserSpecifiedSegment("accessPackageResourceRoleScopeId", "accessPackageResourceRoleScopeId"),
		resourceids.StaticSegment("role", "role", "role"),
		resourceids.StaticSegment("resource", "resource", "resource"),
		resourceids.StaticSegment("roles", "roles", "roles"),
		resourceids.UserSpecifiedSegment("accessPackageResourceRoleId", "accessPackageResourceRoleId"),
	}
}

// String returns a human-readable description of this Identity Governance Entitlement Management Access Package Id Resource Role Scope Id Role Resource Role ID
func (id IdentityGovernanceEntitlementManagementAccessPackageIdResourceRoleScopeIdRoleResourceRoleId) String() string {
	components := []string{
		fmt.Sprintf("Access Package: %q", id.AccessPackageId),
		fmt.Sprintf("Access Package Resource Role Scope: %q", id.AccessPackageResourceRoleScopeId),
		fmt.Sprintf("Access Package Resource Role: %q", id.AccessPackageResourceRoleId),
	}
	return fmt.Sprintf("Identity Governance Entitlement Management Access Package Id Resource Role Scope Id Role Resource Role (%s)", strings.Join(components, "\n"))
}
