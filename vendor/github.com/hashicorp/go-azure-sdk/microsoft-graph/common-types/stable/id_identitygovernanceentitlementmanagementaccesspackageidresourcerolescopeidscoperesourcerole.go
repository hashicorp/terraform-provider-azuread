package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &IdentityGovernanceEntitlementManagementAccessPackageIdResourceRoleScopeIdScopeResourceRoleId{}

// IdentityGovernanceEntitlementManagementAccessPackageIdResourceRoleScopeIdScopeResourceRoleId is a struct representing the Resource ID for a Identity Governance Entitlement Management Access Package Id Resource Role Scope Id Scope Resource Role
type IdentityGovernanceEntitlementManagementAccessPackageIdResourceRoleScopeIdScopeResourceRoleId struct {
	AccessPackageId                  string
	AccessPackageResourceRoleScopeId string
	AccessPackageResourceRoleId      string
}

// NewIdentityGovernanceEntitlementManagementAccessPackageIdResourceRoleScopeIdScopeResourceRoleID returns a new IdentityGovernanceEntitlementManagementAccessPackageIdResourceRoleScopeIdScopeResourceRoleId struct
func NewIdentityGovernanceEntitlementManagementAccessPackageIdResourceRoleScopeIdScopeResourceRoleID(accessPackageId string, accessPackageResourceRoleScopeId string, accessPackageResourceRoleId string) IdentityGovernanceEntitlementManagementAccessPackageIdResourceRoleScopeIdScopeResourceRoleId {
	return IdentityGovernanceEntitlementManagementAccessPackageIdResourceRoleScopeIdScopeResourceRoleId{
		AccessPackageId:                  accessPackageId,
		AccessPackageResourceRoleScopeId: accessPackageResourceRoleScopeId,
		AccessPackageResourceRoleId:      accessPackageResourceRoleId,
	}
}

// ParseIdentityGovernanceEntitlementManagementAccessPackageIdResourceRoleScopeIdScopeResourceRoleID parses 'input' into a IdentityGovernanceEntitlementManagementAccessPackageIdResourceRoleScopeIdScopeResourceRoleId
func ParseIdentityGovernanceEntitlementManagementAccessPackageIdResourceRoleScopeIdScopeResourceRoleID(input string) (*IdentityGovernanceEntitlementManagementAccessPackageIdResourceRoleScopeIdScopeResourceRoleId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityGovernanceEntitlementManagementAccessPackageIdResourceRoleScopeIdScopeResourceRoleId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityGovernanceEntitlementManagementAccessPackageIdResourceRoleScopeIdScopeResourceRoleId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseIdentityGovernanceEntitlementManagementAccessPackageIdResourceRoleScopeIdScopeResourceRoleIDInsensitively parses 'input' case-insensitively into a IdentityGovernanceEntitlementManagementAccessPackageIdResourceRoleScopeIdScopeResourceRoleId
// note: this method should only be used for API response data and not user input
func ParseIdentityGovernanceEntitlementManagementAccessPackageIdResourceRoleScopeIdScopeResourceRoleIDInsensitively(input string) (*IdentityGovernanceEntitlementManagementAccessPackageIdResourceRoleScopeIdScopeResourceRoleId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityGovernanceEntitlementManagementAccessPackageIdResourceRoleScopeIdScopeResourceRoleId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityGovernanceEntitlementManagementAccessPackageIdResourceRoleScopeIdScopeResourceRoleId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *IdentityGovernanceEntitlementManagementAccessPackageIdResourceRoleScopeIdScopeResourceRoleId) FromParseResult(input resourceids.ParseResult) error {
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

// ValidateIdentityGovernanceEntitlementManagementAccessPackageIdResourceRoleScopeIdScopeResourceRoleID checks that 'input' can be parsed as a Identity Governance Entitlement Management Access Package Id Resource Role Scope Id Scope Resource Role ID
func ValidateIdentityGovernanceEntitlementManagementAccessPackageIdResourceRoleScopeIdScopeResourceRoleID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseIdentityGovernanceEntitlementManagementAccessPackageIdResourceRoleScopeIdScopeResourceRoleID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Identity Governance Entitlement Management Access Package Id Resource Role Scope Id Scope Resource Role ID
func (id IdentityGovernanceEntitlementManagementAccessPackageIdResourceRoleScopeIdScopeResourceRoleId) ID() string {
	fmtString := "/identityGovernance/entitlementManagement/accessPackages/%s/resourceRoleScopes/%s/scope/resource/roles/%s"
	return fmt.Sprintf(fmtString, id.AccessPackageId, id.AccessPackageResourceRoleScopeId, id.AccessPackageResourceRoleId)
}

// Segments returns a slice of Resource ID Segments which comprise this Identity Governance Entitlement Management Access Package Id Resource Role Scope Id Scope Resource Role ID
func (id IdentityGovernanceEntitlementManagementAccessPackageIdResourceRoleScopeIdScopeResourceRoleId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("identityGovernance", "identityGovernance", "identityGovernance"),
		resourceids.StaticSegment("entitlementManagement", "entitlementManagement", "entitlementManagement"),
		resourceids.StaticSegment("accessPackages", "accessPackages", "accessPackages"),
		resourceids.UserSpecifiedSegment("accessPackageId", "accessPackageId"),
		resourceids.StaticSegment("resourceRoleScopes", "resourceRoleScopes", "resourceRoleScopes"),
		resourceids.UserSpecifiedSegment("accessPackageResourceRoleScopeId", "accessPackageResourceRoleScopeId"),
		resourceids.StaticSegment("scope", "scope", "scope"),
		resourceids.StaticSegment("resource", "resource", "resource"),
		resourceids.StaticSegment("roles", "roles", "roles"),
		resourceids.UserSpecifiedSegment("accessPackageResourceRoleId", "accessPackageResourceRoleId"),
	}
}

// String returns a human-readable description of this Identity Governance Entitlement Management Access Package Id Resource Role Scope Id Scope Resource Role ID
func (id IdentityGovernanceEntitlementManagementAccessPackageIdResourceRoleScopeIdScopeResourceRoleId) String() string {
	components := []string{
		fmt.Sprintf("Access Package: %q", id.AccessPackageId),
		fmt.Sprintf("Access Package Resource Role Scope: %q", id.AccessPackageResourceRoleScopeId),
		fmt.Sprintf("Access Package Resource Role: %q", id.AccessPackageResourceRoleId),
	}
	return fmt.Sprintf("Identity Governance Entitlement Management Access Package Id Resource Role Scope Id Scope Resource Role (%s)", strings.Join(components, "\n"))
}
