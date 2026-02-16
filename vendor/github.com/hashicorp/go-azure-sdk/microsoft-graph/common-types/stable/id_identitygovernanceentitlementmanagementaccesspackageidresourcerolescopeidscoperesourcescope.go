package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &IdentityGovernanceEntitlementManagementAccessPackageIdResourceRoleScopeIdScopeResourceScopeId{}

// IdentityGovernanceEntitlementManagementAccessPackageIdResourceRoleScopeIdScopeResourceScopeId is a struct representing the Resource ID for a Identity Governance Entitlement Management Access Package Id Resource Role Scope Id Scope Resource Scope
type IdentityGovernanceEntitlementManagementAccessPackageIdResourceRoleScopeIdScopeResourceScopeId struct {
	AccessPackageId                  string
	AccessPackageResourceRoleScopeId string
	AccessPackageResourceScopeId     string
}

// NewIdentityGovernanceEntitlementManagementAccessPackageIdResourceRoleScopeIdScopeResourceScopeID returns a new IdentityGovernanceEntitlementManagementAccessPackageIdResourceRoleScopeIdScopeResourceScopeId struct
func NewIdentityGovernanceEntitlementManagementAccessPackageIdResourceRoleScopeIdScopeResourceScopeID(accessPackageId string, accessPackageResourceRoleScopeId string, accessPackageResourceScopeId string) IdentityGovernanceEntitlementManagementAccessPackageIdResourceRoleScopeIdScopeResourceScopeId {
	return IdentityGovernanceEntitlementManagementAccessPackageIdResourceRoleScopeIdScopeResourceScopeId{
		AccessPackageId:                  accessPackageId,
		AccessPackageResourceRoleScopeId: accessPackageResourceRoleScopeId,
		AccessPackageResourceScopeId:     accessPackageResourceScopeId,
	}
}

// ParseIdentityGovernanceEntitlementManagementAccessPackageIdResourceRoleScopeIdScopeResourceScopeID parses 'input' into a IdentityGovernanceEntitlementManagementAccessPackageIdResourceRoleScopeIdScopeResourceScopeId
func ParseIdentityGovernanceEntitlementManagementAccessPackageIdResourceRoleScopeIdScopeResourceScopeID(input string) (*IdentityGovernanceEntitlementManagementAccessPackageIdResourceRoleScopeIdScopeResourceScopeId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityGovernanceEntitlementManagementAccessPackageIdResourceRoleScopeIdScopeResourceScopeId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityGovernanceEntitlementManagementAccessPackageIdResourceRoleScopeIdScopeResourceScopeId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseIdentityGovernanceEntitlementManagementAccessPackageIdResourceRoleScopeIdScopeResourceScopeIDInsensitively parses 'input' case-insensitively into a IdentityGovernanceEntitlementManagementAccessPackageIdResourceRoleScopeIdScopeResourceScopeId
// note: this method should only be used for API response data and not user input
func ParseIdentityGovernanceEntitlementManagementAccessPackageIdResourceRoleScopeIdScopeResourceScopeIDInsensitively(input string) (*IdentityGovernanceEntitlementManagementAccessPackageIdResourceRoleScopeIdScopeResourceScopeId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityGovernanceEntitlementManagementAccessPackageIdResourceRoleScopeIdScopeResourceScopeId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityGovernanceEntitlementManagementAccessPackageIdResourceRoleScopeIdScopeResourceScopeId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *IdentityGovernanceEntitlementManagementAccessPackageIdResourceRoleScopeIdScopeResourceScopeId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.AccessPackageId, ok = input.Parsed["accessPackageId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "accessPackageId", input)
	}

	if id.AccessPackageResourceRoleScopeId, ok = input.Parsed["accessPackageResourceRoleScopeId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "accessPackageResourceRoleScopeId", input)
	}

	if id.AccessPackageResourceScopeId, ok = input.Parsed["accessPackageResourceScopeId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "accessPackageResourceScopeId", input)
	}

	return nil
}

// ValidateIdentityGovernanceEntitlementManagementAccessPackageIdResourceRoleScopeIdScopeResourceScopeID checks that 'input' can be parsed as a Identity Governance Entitlement Management Access Package Id Resource Role Scope Id Scope Resource Scope ID
func ValidateIdentityGovernanceEntitlementManagementAccessPackageIdResourceRoleScopeIdScopeResourceScopeID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseIdentityGovernanceEntitlementManagementAccessPackageIdResourceRoleScopeIdScopeResourceScopeID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Identity Governance Entitlement Management Access Package Id Resource Role Scope Id Scope Resource Scope ID
func (id IdentityGovernanceEntitlementManagementAccessPackageIdResourceRoleScopeIdScopeResourceScopeId) ID() string {
	fmtString := "/identityGovernance/entitlementManagement/accessPackages/%s/resourceRoleScopes/%s/scope/resource/scopes/%s"
	return fmt.Sprintf(fmtString, id.AccessPackageId, id.AccessPackageResourceRoleScopeId, id.AccessPackageResourceScopeId)
}

// Segments returns a slice of Resource ID Segments which comprise this Identity Governance Entitlement Management Access Package Id Resource Role Scope Id Scope Resource Scope ID
func (id IdentityGovernanceEntitlementManagementAccessPackageIdResourceRoleScopeIdScopeResourceScopeId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("identityGovernance", "identityGovernance", "identityGovernance"),
		resourceids.StaticSegment("entitlementManagement", "entitlementManagement", "entitlementManagement"),
		resourceids.StaticSegment("accessPackages", "accessPackages", "accessPackages"),
		resourceids.UserSpecifiedSegment("accessPackageId", "accessPackageId"),
		resourceids.StaticSegment("resourceRoleScopes", "resourceRoleScopes", "resourceRoleScopes"),
		resourceids.UserSpecifiedSegment("accessPackageResourceRoleScopeId", "accessPackageResourceRoleScopeId"),
		resourceids.StaticSegment("scope", "scope", "scope"),
		resourceids.StaticSegment("resource", "resource", "resource"),
		resourceids.StaticSegment("scopes", "scopes", "scopes"),
		resourceids.UserSpecifiedSegment("accessPackageResourceScopeId", "accessPackageResourceScopeId"),
	}
}

// String returns a human-readable description of this Identity Governance Entitlement Management Access Package Id Resource Role Scope Id Scope Resource Scope ID
func (id IdentityGovernanceEntitlementManagementAccessPackageIdResourceRoleScopeIdScopeResourceScopeId) String() string {
	components := []string{
		fmt.Sprintf("Access Package: %q", id.AccessPackageId),
		fmt.Sprintf("Access Package Resource Role Scope: %q", id.AccessPackageResourceRoleScopeId),
		fmt.Sprintf("Access Package Resource Scope: %q", id.AccessPackageResourceScopeId),
	}
	return fmt.Sprintf("Identity Governance Entitlement Management Access Package Id Resource Role Scope Id Scope Resource Scope (%s)", strings.Join(components, "\n"))
}
