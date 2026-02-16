package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &IdentityGovernanceEntitlementManagementAccessPackageResourceRoleScopeId{}

// IdentityGovernanceEntitlementManagementAccessPackageResourceRoleScopeId is a struct representing the Resource ID for a Identity Governance Entitlement Management Access Package Resource Role Scope
type IdentityGovernanceEntitlementManagementAccessPackageResourceRoleScopeId struct {
	AccessPackageResourceRoleScopeId string
}

// NewIdentityGovernanceEntitlementManagementAccessPackageResourceRoleScopeID returns a new IdentityGovernanceEntitlementManagementAccessPackageResourceRoleScopeId struct
func NewIdentityGovernanceEntitlementManagementAccessPackageResourceRoleScopeID(accessPackageResourceRoleScopeId string) IdentityGovernanceEntitlementManagementAccessPackageResourceRoleScopeId {
	return IdentityGovernanceEntitlementManagementAccessPackageResourceRoleScopeId{
		AccessPackageResourceRoleScopeId: accessPackageResourceRoleScopeId,
	}
}

// ParseIdentityGovernanceEntitlementManagementAccessPackageResourceRoleScopeID parses 'input' into a IdentityGovernanceEntitlementManagementAccessPackageResourceRoleScopeId
func ParseIdentityGovernanceEntitlementManagementAccessPackageResourceRoleScopeID(input string) (*IdentityGovernanceEntitlementManagementAccessPackageResourceRoleScopeId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityGovernanceEntitlementManagementAccessPackageResourceRoleScopeId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityGovernanceEntitlementManagementAccessPackageResourceRoleScopeId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseIdentityGovernanceEntitlementManagementAccessPackageResourceRoleScopeIDInsensitively parses 'input' case-insensitively into a IdentityGovernanceEntitlementManagementAccessPackageResourceRoleScopeId
// note: this method should only be used for API response data and not user input
func ParseIdentityGovernanceEntitlementManagementAccessPackageResourceRoleScopeIDInsensitively(input string) (*IdentityGovernanceEntitlementManagementAccessPackageResourceRoleScopeId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityGovernanceEntitlementManagementAccessPackageResourceRoleScopeId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityGovernanceEntitlementManagementAccessPackageResourceRoleScopeId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *IdentityGovernanceEntitlementManagementAccessPackageResourceRoleScopeId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.AccessPackageResourceRoleScopeId, ok = input.Parsed["accessPackageResourceRoleScopeId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "accessPackageResourceRoleScopeId", input)
	}

	return nil
}

// ValidateIdentityGovernanceEntitlementManagementAccessPackageResourceRoleScopeID checks that 'input' can be parsed as a Identity Governance Entitlement Management Access Package Resource Role Scope ID
func ValidateIdentityGovernanceEntitlementManagementAccessPackageResourceRoleScopeID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseIdentityGovernanceEntitlementManagementAccessPackageResourceRoleScopeID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Identity Governance Entitlement Management Access Package Resource Role Scope ID
func (id IdentityGovernanceEntitlementManagementAccessPackageResourceRoleScopeId) ID() string {
	fmtString := "/identityGovernance/entitlementManagement/accessPackageResourceRoleScopes/%s"
	return fmt.Sprintf(fmtString, id.AccessPackageResourceRoleScopeId)
}

// Segments returns a slice of Resource ID Segments which comprise this Identity Governance Entitlement Management Access Package Resource Role Scope ID
func (id IdentityGovernanceEntitlementManagementAccessPackageResourceRoleScopeId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("identityGovernance", "identityGovernance", "identityGovernance"),
		resourceids.StaticSegment("entitlementManagement", "entitlementManagement", "entitlementManagement"),
		resourceids.StaticSegment("accessPackageResourceRoleScopes", "accessPackageResourceRoleScopes", "accessPackageResourceRoleScopes"),
		resourceids.UserSpecifiedSegment("accessPackageResourceRoleScopeId", "accessPackageResourceRoleScopeId"),
	}
}

// String returns a human-readable description of this Identity Governance Entitlement Management Access Package Resource Role Scope ID
func (id IdentityGovernanceEntitlementManagementAccessPackageResourceRoleScopeId) String() string {
	components := []string{
		fmt.Sprintf("Access Package Resource Role Scope: %q", id.AccessPackageResourceRoleScopeId),
	}
	return fmt.Sprintf("Identity Governance Entitlement Management Access Package Resource Role Scope (%s)", strings.Join(components, "\n"))
}
