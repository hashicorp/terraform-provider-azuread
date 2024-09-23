package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &IdentityGovernanceEntitlementManagementResourceRoleScopeIdScopeResourceScopeId{}

// IdentityGovernanceEntitlementManagementResourceRoleScopeIdScopeResourceScopeId is a struct representing the Resource ID for a Identity Governance Entitlement Management Resource Role Scope Id Scope Resource Scope
type IdentityGovernanceEntitlementManagementResourceRoleScopeIdScopeResourceScopeId struct {
	AccessPackageResourceRoleScopeId string
	AccessPackageResourceScopeId     string
}

// NewIdentityGovernanceEntitlementManagementResourceRoleScopeIdScopeResourceScopeID returns a new IdentityGovernanceEntitlementManagementResourceRoleScopeIdScopeResourceScopeId struct
func NewIdentityGovernanceEntitlementManagementResourceRoleScopeIdScopeResourceScopeID(accessPackageResourceRoleScopeId string, accessPackageResourceScopeId string) IdentityGovernanceEntitlementManagementResourceRoleScopeIdScopeResourceScopeId {
	return IdentityGovernanceEntitlementManagementResourceRoleScopeIdScopeResourceScopeId{
		AccessPackageResourceRoleScopeId: accessPackageResourceRoleScopeId,
		AccessPackageResourceScopeId:     accessPackageResourceScopeId,
	}
}

// ParseIdentityGovernanceEntitlementManagementResourceRoleScopeIdScopeResourceScopeID parses 'input' into a IdentityGovernanceEntitlementManagementResourceRoleScopeIdScopeResourceScopeId
func ParseIdentityGovernanceEntitlementManagementResourceRoleScopeIdScopeResourceScopeID(input string) (*IdentityGovernanceEntitlementManagementResourceRoleScopeIdScopeResourceScopeId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityGovernanceEntitlementManagementResourceRoleScopeIdScopeResourceScopeId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityGovernanceEntitlementManagementResourceRoleScopeIdScopeResourceScopeId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseIdentityGovernanceEntitlementManagementResourceRoleScopeIdScopeResourceScopeIDInsensitively parses 'input' case-insensitively into a IdentityGovernanceEntitlementManagementResourceRoleScopeIdScopeResourceScopeId
// note: this method should only be used for API response data and not user input
func ParseIdentityGovernanceEntitlementManagementResourceRoleScopeIdScopeResourceScopeIDInsensitively(input string) (*IdentityGovernanceEntitlementManagementResourceRoleScopeIdScopeResourceScopeId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityGovernanceEntitlementManagementResourceRoleScopeIdScopeResourceScopeId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityGovernanceEntitlementManagementResourceRoleScopeIdScopeResourceScopeId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *IdentityGovernanceEntitlementManagementResourceRoleScopeIdScopeResourceScopeId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.AccessPackageResourceRoleScopeId, ok = input.Parsed["accessPackageResourceRoleScopeId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "accessPackageResourceRoleScopeId", input)
	}

	if id.AccessPackageResourceScopeId, ok = input.Parsed["accessPackageResourceScopeId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "accessPackageResourceScopeId", input)
	}

	return nil
}

// ValidateIdentityGovernanceEntitlementManagementResourceRoleScopeIdScopeResourceScopeID checks that 'input' can be parsed as a Identity Governance Entitlement Management Resource Role Scope Id Scope Resource Scope ID
func ValidateIdentityGovernanceEntitlementManagementResourceRoleScopeIdScopeResourceScopeID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseIdentityGovernanceEntitlementManagementResourceRoleScopeIdScopeResourceScopeID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Identity Governance Entitlement Management Resource Role Scope Id Scope Resource Scope ID
func (id IdentityGovernanceEntitlementManagementResourceRoleScopeIdScopeResourceScopeId) ID() string {
	fmtString := "/identityGovernance/entitlementManagement/resourceRoleScopes/%s/scope/resource/scopes/%s"
	return fmt.Sprintf(fmtString, id.AccessPackageResourceRoleScopeId, id.AccessPackageResourceScopeId)
}

// Segments returns a slice of Resource ID Segments which comprise this Identity Governance Entitlement Management Resource Role Scope Id Scope Resource Scope ID
func (id IdentityGovernanceEntitlementManagementResourceRoleScopeIdScopeResourceScopeId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("identityGovernance", "identityGovernance", "identityGovernance"),
		resourceids.StaticSegment("entitlementManagement", "entitlementManagement", "entitlementManagement"),
		resourceids.StaticSegment("resourceRoleScopes", "resourceRoleScopes", "resourceRoleScopes"),
		resourceids.UserSpecifiedSegment("accessPackageResourceRoleScopeId", "accessPackageResourceRoleScopeId"),
		resourceids.StaticSegment("scope", "scope", "scope"),
		resourceids.StaticSegment("resource", "resource", "resource"),
		resourceids.StaticSegment("scopes", "scopes", "scopes"),
		resourceids.UserSpecifiedSegment("accessPackageResourceScopeId", "accessPackageResourceScopeId"),
	}
}

// String returns a human-readable description of this Identity Governance Entitlement Management Resource Role Scope Id Scope Resource Scope ID
func (id IdentityGovernanceEntitlementManagementResourceRoleScopeIdScopeResourceScopeId) String() string {
	components := []string{
		fmt.Sprintf("Access Package Resource Role Scope: %q", id.AccessPackageResourceRoleScopeId),
		fmt.Sprintf("Access Package Resource Scope: %q", id.AccessPackageResourceScopeId),
	}
	return fmt.Sprintf("Identity Governance Entitlement Management Resource Role Scope Id Scope Resource Scope (%s)", strings.Join(components, "\n"))
}
