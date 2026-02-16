package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &IdentityGovernanceEntitlementManagementAccessPackageIdAccessPackageResourceRoleScopeId{}

// IdentityGovernanceEntitlementManagementAccessPackageIdAccessPackageResourceRoleScopeId is a struct representing the Resource ID for a Identity Governance Entitlement Management Access Package Id Access Package Resource Role Scope
type IdentityGovernanceEntitlementManagementAccessPackageIdAccessPackageResourceRoleScopeId struct {
	AccessPackageId                  string
	AccessPackageResourceRoleScopeId string
}

// NewIdentityGovernanceEntitlementManagementAccessPackageIdAccessPackageResourceRoleScopeID returns a new IdentityGovernanceEntitlementManagementAccessPackageIdAccessPackageResourceRoleScopeId struct
func NewIdentityGovernanceEntitlementManagementAccessPackageIdAccessPackageResourceRoleScopeID(accessPackageId string, accessPackageResourceRoleScopeId string) IdentityGovernanceEntitlementManagementAccessPackageIdAccessPackageResourceRoleScopeId {
	return IdentityGovernanceEntitlementManagementAccessPackageIdAccessPackageResourceRoleScopeId{
		AccessPackageId:                  accessPackageId,
		AccessPackageResourceRoleScopeId: accessPackageResourceRoleScopeId,
	}
}

// ParseIdentityGovernanceEntitlementManagementAccessPackageIdAccessPackageResourceRoleScopeID parses 'input' into a IdentityGovernanceEntitlementManagementAccessPackageIdAccessPackageResourceRoleScopeId
func ParseIdentityGovernanceEntitlementManagementAccessPackageIdAccessPackageResourceRoleScopeID(input string) (*IdentityGovernanceEntitlementManagementAccessPackageIdAccessPackageResourceRoleScopeId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityGovernanceEntitlementManagementAccessPackageIdAccessPackageResourceRoleScopeId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityGovernanceEntitlementManagementAccessPackageIdAccessPackageResourceRoleScopeId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseIdentityGovernanceEntitlementManagementAccessPackageIdAccessPackageResourceRoleScopeIDInsensitively parses 'input' case-insensitively into a IdentityGovernanceEntitlementManagementAccessPackageIdAccessPackageResourceRoleScopeId
// note: this method should only be used for API response data and not user input
func ParseIdentityGovernanceEntitlementManagementAccessPackageIdAccessPackageResourceRoleScopeIDInsensitively(input string) (*IdentityGovernanceEntitlementManagementAccessPackageIdAccessPackageResourceRoleScopeId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityGovernanceEntitlementManagementAccessPackageIdAccessPackageResourceRoleScopeId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityGovernanceEntitlementManagementAccessPackageIdAccessPackageResourceRoleScopeId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *IdentityGovernanceEntitlementManagementAccessPackageIdAccessPackageResourceRoleScopeId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.AccessPackageId, ok = input.Parsed["accessPackageId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "accessPackageId", input)
	}

	if id.AccessPackageResourceRoleScopeId, ok = input.Parsed["accessPackageResourceRoleScopeId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "accessPackageResourceRoleScopeId", input)
	}

	return nil
}

// ValidateIdentityGovernanceEntitlementManagementAccessPackageIdAccessPackageResourceRoleScopeID checks that 'input' can be parsed as a Identity Governance Entitlement Management Access Package Id Access Package Resource Role Scope ID
func ValidateIdentityGovernanceEntitlementManagementAccessPackageIdAccessPackageResourceRoleScopeID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseIdentityGovernanceEntitlementManagementAccessPackageIdAccessPackageResourceRoleScopeID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Identity Governance Entitlement Management Access Package Id Access Package Resource Role Scope ID
func (id IdentityGovernanceEntitlementManagementAccessPackageIdAccessPackageResourceRoleScopeId) ID() string {
	fmtString := "/identityGovernance/entitlementManagement/accessPackages/%s/accessPackageResourceRoleScopes/%s"
	return fmt.Sprintf(fmtString, id.AccessPackageId, id.AccessPackageResourceRoleScopeId)
}

// Segments returns a slice of Resource ID Segments which comprise this Identity Governance Entitlement Management Access Package Id Access Package Resource Role Scope ID
func (id IdentityGovernanceEntitlementManagementAccessPackageIdAccessPackageResourceRoleScopeId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("identityGovernance", "identityGovernance", "identityGovernance"),
		resourceids.StaticSegment("entitlementManagement", "entitlementManagement", "entitlementManagement"),
		resourceids.StaticSegment("accessPackages", "accessPackages", "accessPackages"),
		resourceids.UserSpecifiedSegment("accessPackageId", "accessPackageId"),
		resourceids.StaticSegment("accessPackageResourceRoleScopes", "accessPackageResourceRoleScopes", "accessPackageResourceRoleScopes"),
		resourceids.UserSpecifiedSegment("accessPackageResourceRoleScopeId", "accessPackageResourceRoleScopeId"),
	}
}

// String returns a human-readable description of this Identity Governance Entitlement Management Access Package Id Access Package Resource Role Scope ID
func (id IdentityGovernanceEntitlementManagementAccessPackageIdAccessPackageResourceRoleScopeId) String() string {
	components := []string{
		fmt.Sprintf("Access Package: %q", id.AccessPackageId),
		fmt.Sprintf("Access Package Resource Role Scope: %q", id.AccessPackageResourceRoleScopeId),
	}
	return fmt.Sprintf("Identity Governance Entitlement Management Access Package Id Access Package Resource Role Scope (%s)", strings.Join(components, "\n"))
}
