package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &IdentityGovernanceEntitlementManagementAccessPackageIdResourceRoleScopeId{}

// IdentityGovernanceEntitlementManagementAccessPackageIdResourceRoleScopeId is a struct representing the Resource ID for a Identity Governance Entitlement Management Access Package Id Resource Role Scope
type IdentityGovernanceEntitlementManagementAccessPackageIdResourceRoleScopeId struct {
	AccessPackageId                  string
	AccessPackageResourceRoleScopeId string
}

// NewIdentityGovernanceEntitlementManagementAccessPackageIdResourceRoleScopeID returns a new IdentityGovernanceEntitlementManagementAccessPackageIdResourceRoleScopeId struct
func NewIdentityGovernanceEntitlementManagementAccessPackageIdResourceRoleScopeID(accessPackageId string, accessPackageResourceRoleScopeId string) IdentityGovernanceEntitlementManagementAccessPackageIdResourceRoleScopeId {
	return IdentityGovernanceEntitlementManagementAccessPackageIdResourceRoleScopeId{
		AccessPackageId:                  accessPackageId,
		AccessPackageResourceRoleScopeId: accessPackageResourceRoleScopeId,
	}
}

// ParseIdentityGovernanceEntitlementManagementAccessPackageIdResourceRoleScopeID parses 'input' into a IdentityGovernanceEntitlementManagementAccessPackageIdResourceRoleScopeId
func ParseIdentityGovernanceEntitlementManagementAccessPackageIdResourceRoleScopeID(input string) (*IdentityGovernanceEntitlementManagementAccessPackageIdResourceRoleScopeId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityGovernanceEntitlementManagementAccessPackageIdResourceRoleScopeId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityGovernanceEntitlementManagementAccessPackageIdResourceRoleScopeId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseIdentityGovernanceEntitlementManagementAccessPackageIdResourceRoleScopeIDInsensitively parses 'input' case-insensitively into a IdentityGovernanceEntitlementManagementAccessPackageIdResourceRoleScopeId
// note: this method should only be used for API response data and not user input
func ParseIdentityGovernanceEntitlementManagementAccessPackageIdResourceRoleScopeIDInsensitively(input string) (*IdentityGovernanceEntitlementManagementAccessPackageIdResourceRoleScopeId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityGovernanceEntitlementManagementAccessPackageIdResourceRoleScopeId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityGovernanceEntitlementManagementAccessPackageIdResourceRoleScopeId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *IdentityGovernanceEntitlementManagementAccessPackageIdResourceRoleScopeId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.AccessPackageId, ok = input.Parsed["accessPackageId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "accessPackageId", input)
	}

	if id.AccessPackageResourceRoleScopeId, ok = input.Parsed["accessPackageResourceRoleScopeId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "accessPackageResourceRoleScopeId", input)
	}

	return nil
}

// ValidateIdentityGovernanceEntitlementManagementAccessPackageIdResourceRoleScopeID checks that 'input' can be parsed as a Identity Governance Entitlement Management Access Package Id Resource Role Scope ID
func ValidateIdentityGovernanceEntitlementManagementAccessPackageIdResourceRoleScopeID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseIdentityGovernanceEntitlementManagementAccessPackageIdResourceRoleScopeID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Identity Governance Entitlement Management Access Package Id Resource Role Scope ID
func (id IdentityGovernanceEntitlementManagementAccessPackageIdResourceRoleScopeId) ID() string {
	fmtString := "/identityGovernance/entitlementManagement/accessPackages/%s/resourceRoleScopes/%s"
	return fmt.Sprintf(fmtString, id.AccessPackageId, id.AccessPackageResourceRoleScopeId)
}

// Segments returns a slice of Resource ID Segments which comprise this Identity Governance Entitlement Management Access Package Id Resource Role Scope ID
func (id IdentityGovernanceEntitlementManagementAccessPackageIdResourceRoleScopeId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("identityGovernance", "identityGovernance", "identityGovernance"),
		resourceids.StaticSegment("entitlementManagement", "entitlementManagement", "entitlementManagement"),
		resourceids.StaticSegment("accessPackages", "accessPackages", "accessPackages"),
		resourceids.UserSpecifiedSegment("accessPackageId", "accessPackageId"),
		resourceids.StaticSegment("resourceRoleScopes", "resourceRoleScopes", "resourceRoleScopes"),
		resourceids.UserSpecifiedSegment("accessPackageResourceRoleScopeId", "accessPackageResourceRoleScopeId"),
	}
}

// String returns a human-readable description of this Identity Governance Entitlement Management Access Package Id Resource Role Scope ID
func (id IdentityGovernanceEntitlementManagementAccessPackageIdResourceRoleScopeId) String() string {
	components := []string{
		fmt.Sprintf("Access Package: %q", id.AccessPackageId),
		fmt.Sprintf("Access Package Resource Role Scope: %q", id.AccessPackageResourceRoleScopeId),
	}
	return fmt.Sprintf("Identity Governance Entitlement Management Access Package Id Resource Role Scope (%s)", strings.Join(components, "\n"))
}
