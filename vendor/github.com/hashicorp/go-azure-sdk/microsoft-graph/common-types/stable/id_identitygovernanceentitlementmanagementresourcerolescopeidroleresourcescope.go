package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &IdentityGovernanceEntitlementManagementResourceRoleScopeIdRoleResourceScopeId{}

// IdentityGovernanceEntitlementManagementResourceRoleScopeIdRoleResourceScopeId is a struct representing the Resource ID for a Identity Governance Entitlement Management Resource Role Scope Id Role Resource Scope
type IdentityGovernanceEntitlementManagementResourceRoleScopeIdRoleResourceScopeId struct {
	AccessPackageResourceRoleScopeId string
	AccessPackageResourceScopeId     string
}

// NewIdentityGovernanceEntitlementManagementResourceRoleScopeIdRoleResourceScopeID returns a new IdentityGovernanceEntitlementManagementResourceRoleScopeIdRoleResourceScopeId struct
func NewIdentityGovernanceEntitlementManagementResourceRoleScopeIdRoleResourceScopeID(accessPackageResourceRoleScopeId string, accessPackageResourceScopeId string) IdentityGovernanceEntitlementManagementResourceRoleScopeIdRoleResourceScopeId {
	return IdentityGovernanceEntitlementManagementResourceRoleScopeIdRoleResourceScopeId{
		AccessPackageResourceRoleScopeId: accessPackageResourceRoleScopeId,
		AccessPackageResourceScopeId:     accessPackageResourceScopeId,
	}
}

// ParseIdentityGovernanceEntitlementManagementResourceRoleScopeIdRoleResourceScopeID parses 'input' into a IdentityGovernanceEntitlementManagementResourceRoleScopeIdRoleResourceScopeId
func ParseIdentityGovernanceEntitlementManagementResourceRoleScopeIdRoleResourceScopeID(input string) (*IdentityGovernanceEntitlementManagementResourceRoleScopeIdRoleResourceScopeId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityGovernanceEntitlementManagementResourceRoleScopeIdRoleResourceScopeId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityGovernanceEntitlementManagementResourceRoleScopeIdRoleResourceScopeId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseIdentityGovernanceEntitlementManagementResourceRoleScopeIdRoleResourceScopeIDInsensitively parses 'input' case-insensitively into a IdentityGovernanceEntitlementManagementResourceRoleScopeIdRoleResourceScopeId
// note: this method should only be used for API response data and not user input
func ParseIdentityGovernanceEntitlementManagementResourceRoleScopeIdRoleResourceScopeIDInsensitively(input string) (*IdentityGovernanceEntitlementManagementResourceRoleScopeIdRoleResourceScopeId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityGovernanceEntitlementManagementResourceRoleScopeIdRoleResourceScopeId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityGovernanceEntitlementManagementResourceRoleScopeIdRoleResourceScopeId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *IdentityGovernanceEntitlementManagementResourceRoleScopeIdRoleResourceScopeId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.AccessPackageResourceRoleScopeId, ok = input.Parsed["accessPackageResourceRoleScopeId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "accessPackageResourceRoleScopeId", input)
	}

	if id.AccessPackageResourceScopeId, ok = input.Parsed["accessPackageResourceScopeId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "accessPackageResourceScopeId", input)
	}

	return nil
}

// ValidateIdentityGovernanceEntitlementManagementResourceRoleScopeIdRoleResourceScopeID checks that 'input' can be parsed as a Identity Governance Entitlement Management Resource Role Scope Id Role Resource Scope ID
func ValidateIdentityGovernanceEntitlementManagementResourceRoleScopeIdRoleResourceScopeID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseIdentityGovernanceEntitlementManagementResourceRoleScopeIdRoleResourceScopeID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Identity Governance Entitlement Management Resource Role Scope Id Role Resource Scope ID
func (id IdentityGovernanceEntitlementManagementResourceRoleScopeIdRoleResourceScopeId) ID() string {
	fmtString := "/identityGovernance/entitlementManagement/resourceRoleScopes/%s/role/resource/scopes/%s"
	return fmt.Sprintf(fmtString, id.AccessPackageResourceRoleScopeId, id.AccessPackageResourceScopeId)
}

// Segments returns a slice of Resource ID Segments which comprise this Identity Governance Entitlement Management Resource Role Scope Id Role Resource Scope ID
func (id IdentityGovernanceEntitlementManagementResourceRoleScopeIdRoleResourceScopeId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("identityGovernance", "identityGovernance", "identityGovernance"),
		resourceids.StaticSegment("entitlementManagement", "entitlementManagement", "entitlementManagement"),
		resourceids.StaticSegment("resourceRoleScopes", "resourceRoleScopes", "resourceRoleScopes"),
		resourceids.UserSpecifiedSegment("accessPackageResourceRoleScopeId", "accessPackageResourceRoleScopeId"),
		resourceids.StaticSegment("role", "role", "role"),
		resourceids.StaticSegment("resource", "resource", "resource"),
		resourceids.StaticSegment("scopes", "scopes", "scopes"),
		resourceids.UserSpecifiedSegment("accessPackageResourceScopeId", "accessPackageResourceScopeId"),
	}
}

// String returns a human-readable description of this Identity Governance Entitlement Management Resource Role Scope Id Role Resource Scope ID
func (id IdentityGovernanceEntitlementManagementResourceRoleScopeIdRoleResourceScopeId) String() string {
	components := []string{
		fmt.Sprintf("Access Package Resource Role Scope: %q", id.AccessPackageResourceRoleScopeId),
		fmt.Sprintf("Access Package Resource Scope: %q", id.AccessPackageResourceScopeId),
	}
	return fmt.Sprintf("Identity Governance Entitlement Management Resource Role Scope Id Role Resource Scope (%s)", strings.Join(components, "\n"))
}
