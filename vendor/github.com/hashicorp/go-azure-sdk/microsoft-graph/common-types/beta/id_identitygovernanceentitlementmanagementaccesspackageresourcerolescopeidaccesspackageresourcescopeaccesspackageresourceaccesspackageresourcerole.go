package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &IdentityGovernanceEntitlementManagementAccessPackageResourceRoleScopeIdAccessPackageResourceScopeAccessPackageResourceAccessPackageResourceRoleId{}

// IdentityGovernanceEntitlementManagementAccessPackageResourceRoleScopeIdAccessPackageResourceScopeAccessPackageResourceAccessPackageResourceRoleId is a struct representing the Resource ID for a Identity Governance Entitlement Management Access Package Resource Role Scope Id Access Package Resource Scope Access Package Resource Access Package Resource Role
type IdentityGovernanceEntitlementManagementAccessPackageResourceRoleScopeIdAccessPackageResourceScopeAccessPackageResourceAccessPackageResourceRoleId struct {
	AccessPackageResourceRoleScopeId string
	AccessPackageResourceRoleId      string
}

// NewIdentityGovernanceEntitlementManagementAccessPackageResourceRoleScopeIdAccessPackageResourceScopeAccessPackageResourceAccessPackageResourceRoleID returns a new IdentityGovernanceEntitlementManagementAccessPackageResourceRoleScopeIdAccessPackageResourceScopeAccessPackageResourceAccessPackageResourceRoleId struct
func NewIdentityGovernanceEntitlementManagementAccessPackageResourceRoleScopeIdAccessPackageResourceScopeAccessPackageResourceAccessPackageResourceRoleID(accessPackageResourceRoleScopeId string, accessPackageResourceRoleId string) IdentityGovernanceEntitlementManagementAccessPackageResourceRoleScopeIdAccessPackageResourceScopeAccessPackageResourceAccessPackageResourceRoleId {
	return IdentityGovernanceEntitlementManagementAccessPackageResourceRoleScopeIdAccessPackageResourceScopeAccessPackageResourceAccessPackageResourceRoleId{
		AccessPackageResourceRoleScopeId: accessPackageResourceRoleScopeId,
		AccessPackageResourceRoleId:      accessPackageResourceRoleId,
	}
}

// ParseIdentityGovernanceEntitlementManagementAccessPackageResourceRoleScopeIdAccessPackageResourceScopeAccessPackageResourceAccessPackageResourceRoleID parses 'input' into a IdentityGovernanceEntitlementManagementAccessPackageResourceRoleScopeIdAccessPackageResourceScopeAccessPackageResourceAccessPackageResourceRoleId
func ParseIdentityGovernanceEntitlementManagementAccessPackageResourceRoleScopeIdAccessPackageResourceScopeAccessPackageResourceAccessPackageResourceRoleID(input string) (*IdentityGovernanceEntitlementManagementAccessPackageResourceRoleScopeIdAccessPackageResourceScopeAccessPackageResourceAccessPackageResourceRoleId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityGovernanceEntitlementManagementAccessPackageResourceRoleScopeIdAccessPackageResourceScopeAccessPackageResourceAccessPackageResourceRoleId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityGovernanceEntitlementManagementAccessPackageResourceRoleScopeIdAccessPackageResourceScopeAccessPackageResourceAccessPackageResourceRoleId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseIdentityGovernanceEntitlementManagementAccessPackageResourceRoleScopeIdAccessPackageResourceScopeAccessPackageResourceAccessPackageResourceRoleIDInsensitively parses 'input' case-insensitively into a IdentityGovernanceEntitlementManagementAccessPackageResourceRoleScopeIdAccessPackageResourceScopeAccessPackageResourceAccessPackageResourceRoleId
// note: this method should only be used for API response data and not user input
func ParseIdentityGovernanceEntitlementManagementAccessPackageResourceRoleScopeIdAccessPackageResourceScopeAccessPackageResourceAccessPackageResourceRoleIDInsensitively(input string) (*IdentityGovernanceEntitlementManagementAccessPackageResourceRoleScopeIdAccessPackageResourceScopeAccessPackageResourceAccessPackageResourceRoleId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityGovernanceEntitlementManagementAccessPackageResourceRoleScopeIdAccessPackageResourceScopeAccessPackageResourceAccessPackageResourceRoleId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityGovernanceEntitlementManagementAccessPackageResourceRoleScopeIdAccessPackageResourceScopeAccessPackageResourceAccessPackageResourceRoleId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *IdentityGovernanceEntitlementManagementAccessPackageResourceRoleScopeIdAccessPackageResourceScopeAccessPackageResourceAccessPackageResourceRoleId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.AccessPackageResourceRoleScopeId, ok = input.Parsed["accessPackageResourceRoleScopeId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "accessPackageResourceRoleScopeId", input)
	}

	if id.AccessPackageResourceRoleId, ok = input.Parsed["accessPackageResourceRoleId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "accessPackageResourceRoleId", input)
	}

	return nil
}

// ValidateIdentityGovernanceEntitlementManagementAccessPackageResourceRoleScopeIdAccessPackageResourceScopeAccessPackageResourceAccessPackageResourceRoleID checks that 'input' can be parsed as a Identity Governance Entitlement Management Access Package Resource Role Scope Id Access Package Resource Scope Access Package Resource Access Package Resource Role ID
func ValidateIdentityGovernanceEntitlementManagementAccessPackageResourceRoleScopeIdAccessPackageResourceScopeAccessPackageResourceAccessPackageResourceRoleID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseIdentityGovernanceEntitlementManagementAccessPackageResourceRoleScopeIdAccessPackageResourceScopeAccessPackageResourceAccessPackageResourceRoleID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Identity Governance Entitlement Management Access Package Resource Role Scope Id Access Package Resource Scope Access Package Resource Access Package Resource Role ID
func (id IdentityGovernanceEntitlementManagementAccessPackageResourceRoleScopeIdAccessPackageResourceScopeAccessPackageResourceAccessPackageResourceRoleId) ID() string {
	fmtString := "/identityGovernance/entitlementManagement/accessPackageResourceRoleScopes/%s/accessPackageResourceScope/accessPackageResource/accessPackageResourceRoles/%s"
	return fmt.Sprintf(fmtString, id.AccessPackageResourceRoleScopeId, id.AccessPackageResourceRoleId)
}

// Segments returns a slice of Resource ID Segments which comprise this Identity Governance Entitlement Management Access Package Resource Role Scope Id Access Package Resource Scope Access Package Resource Access Package Resource Role ID
func (id IdentityGovernanceEntitlementManagementAccessPackageResourceRoleScopeIdAccessPackageResourceScopeAccessPackageResourceAccessPackageResourceRoleId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("identityGovernance", "identityGovernance", "identityGovernance"),
		resourceids.StaticSegment("entitlementManagement", "entitlementManagement", "entitlementManagement"),
		resourceids.StaticSegment("accessPackageResourceRoleScopes", "accessPackageResourceRoleScopes", "accessPackageResourceRoleScopes"),
		resourceids.UserSpecifiedSegment("accessPackageResourceRoleScopeId", "accessPackageResourceRoleScopeId"),
		resourceids.StaticSegment("accessPackageResourceScope", "accessPackageResourceScope", "accessPackageResourceScope"),
		resourceids.StaticSegment("accessPackageResource", "accessPackageResource", "accessPackageResource"),
		resourceids.StaticSegment("accessPackageResourceRoles", "accessPackageResourceRoles", "accessPackageResourceRoles"),
		resourceids.UserSpecifiedSegment("accessPackageResourceRoleId", "accessPackageResourceRoleId"),
	}
}

// String returns a human-readable description of this Identity Governance Entitlement Management Access Package Resource Role Scope Id Access Package Resource Scope Access Package Resource Access Package Resource Role ID
func (id IdentityGovernanceEntitlementManagementAccessPackageResourceRoleScopeIdAccessPackageResourceScopeAccessPackageResourceAccessPackageResourceRoleId) String() string {
	components := []string{
		fmt.Sprintf("Access Package Resource Role Scope: %q", id.AccessPackageResourceRoleScopeId),
		fmt.Sprintf("Access Package Resource Role: %q", id.AccessPackageResourceRoleId),
	}
	return fmt.Sprintf("Identity Governance Entitlement Management Access Package Resource Role Scope Id Access Package Resource Scope Access Package Resource Access Package Resource Role (%s)", strings.Join(components, "\n"))
}
