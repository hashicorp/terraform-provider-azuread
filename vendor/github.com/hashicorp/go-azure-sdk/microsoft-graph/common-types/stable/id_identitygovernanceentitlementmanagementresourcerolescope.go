package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &IdentityGovernanceEntitlementManagementResourceRoleScopeId{}

// IdentityGovernanceEntitlementManagementResourceRoleScopeId is a struct representing the Resource ID for a Identity Governance Entitlement Management Resource Role Scope
type IdentityGovernanceEntitlementManagementResourceRoleScopeId struct {
	AccessPackageResourceRoleScopeId string
}

// NewIdentityGovernanceEntitlementManagementResourceRoleScopeID returns a new IdentityGovernanceEntitlementManagementResourceRoleScopeId struct
func NewIdentityGovernanceEntitlementManagementResourceRoleScopeID(accessPackageResourceRoleScopeId string) IdentityGovernanceEntitlementManagementResourceRoleScopeId {
	return IdentityGovernanceEntitlementManagementResourceRoleScopeId{
		AccessPackageResourceRoleScopeId: accessPackageResourceRoleScopeId,
	}
}

// ParseIdentityGovernanceEntitlementManagementResourceRoleScopeID parses 'input' into a IdentityGovernanceEntitlementManagementResourceRoleScopeId
func ParseIdentityGovernanceEntitlementManagementResourceRoleScopeID(input string) (*IdentityGovernanceEntitlementManagementResourceRoleScopeId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityGovernanceEntitlementManagementResourceRoleScopeId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityGovernanceEntitlementManagementResourceRoleScopeId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseIdentityGovernanceEntitlementManagementResourceRoleScopeIDInsensitively parses 'input' case-insensitively into a IdentityGovernanceEntitlementManagementResourceRoleScopeId
// note: this method should only be used for API response data and not user input
func ParseIdentityGovernanceEntitlementManagementResourceRoleScopeIDInsensitively(input string) (*IdentityGovernanceEntitlementManagementResourceRoleScopeId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityGovernanceEntitlementManagementResourceRoleScopeId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityGovernanceEntitlementManagementResourceRoleScopeId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *IdentityGovernanceEntitlementManagementResourceRoleScopeId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.AccessPackageResourceRoleScopeId, ok = input.Parsed["accessPackageResourceRoleScopeId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "accessPackageResourceRoleScopeId", input)
	}

	return nil
}

// ValidateIdentityGovernanceEntitlementManagementResourceRoleScopeID checks that 'input' can be parsed as a Identity Governance Entitlement Management Resource Role Scope ID
func ValidateIdentityGovernanceEntitlementManagementResourceRoleScopeID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseIdentityGovernanceEntitlementManagementResourceRoleScopeID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Identity Governance Entitlement Management Resource Role Scope ID
func (id IdentityGovernanceEntitlementManagementResourceRoleScopeId) ID() string {
	fmtString := "/identityGovernance/entitlementManagement/resourceRoleScopes/%s"
	return fmt.Sprintf(fmtString, id.AccessPackageResourceRoleScopeId)
}

// Segments returns a slice of Resource ID Segments which comprise this Identity Governance Entitlement Management Resource Role Scope ID
func (id IdentityGovernanceEntitlementManagementResourceRoleScopeId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("identityGovernance", "identityGovernance", "identityGovernance"),
		resourceids.StaticSegment("entitlementManagement", "entitlementManagement", "entitlementManagement"),
		resourceids.StaticSegment("resourceRoleScopes", "resourceRoleScopes", "resourceRoleScopes"),
		resourceids.UserSpecifiedSegment("accessPackageResourceRoleScopeId", "accessPackageResourceRoleScopeId"),
	}
}

// String returns a human-readable description of this Identity Governance Entitlement Management Resource Role Scope ID
func (id IdentityGovernanceEntitlementManagementResourceRoleScopeId) String() string {
	components := []string{
		fmt.Sprintf("Access Package Resource Role Scope: %q", id.AccessPackageResourceRoleScopeId),
	}
	return fmt.Sprintf("Identity Governance Entitlement Management Resource Role Scope (%s)", strings.Join(components, "\n"))
}
