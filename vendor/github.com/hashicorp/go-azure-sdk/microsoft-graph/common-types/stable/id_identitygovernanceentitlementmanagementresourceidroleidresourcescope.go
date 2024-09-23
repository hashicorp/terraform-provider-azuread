package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &IdentityGovernanceEntitlementManagementResourceIdRoleIdResourceScopeId{}

// IdentityGovernanceEntitlementManagementResourceIdRoleIdResourceScopeId is a struct representing the Resource ID for a Identity Governance Entitlement Management Resource Id Role Id Resource Scope
type IdentityGovernanceEntitlementManagementResourceIdRoleIdResourceScopeId struct {
	AccessPackageResourceId      string
	AccessPackageResourceRoleId  string
	AccessPackageResourceScopeId string
}

// NewIdentityGovernanceEntitlementManagementResourceIdRoleIdResourceScopeID returns a new IdentityGovernanceEntitlementManagementResourceIdRoleIdResourceScopeId struct
func NewIdentityGovernanceEntitlementManagementResourceIdRoleIdResourceScopeID(accessPackageResourceId string, accessPackageResourceRoleId string, accessPackageResourceScopeId string) IdentityGovernanceEntitlementManagementResourceIdRoleIdResourceScopeId {
	return IdentityGovernanceEntitlementManagementResourceIdRoleIdResourceScopeId{
		AccessPackageResourceId:      accessPackageResourceId,
		AccessPackageResourceRoleId:  accessPackageResourceRoleId,
		AccessPackageResourceScopeId: accessPackageResourceScopeId,
	}
}

// ParseIdentityGovernanceEntitlementManagementResourceIdRoleIdResourceScopeID parses 'input' into a IdentityGovernanceEntitlementManagementResourceIdRoleIdResourceScopeId
func ParseIdentityGovernanceEntitlementManagementResourceIdRoleIdResourceScopeID(input string) (*IdentityGovernanceEntitlementManagementResourceIdRoleIdResourceScopeId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityGovernanceEntitlementManagementResourceIdRoleIdResourceScopeId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityGovernanceEntitlementManagementResourceIdRoleIdResourceScopeId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseIdentityGovernanceEntitlementManagementResourceIdRoleIdResourceScopeIDInsensitively parses 'input' case-insensitively into a IdentityGovernanceEntitlementManagementResourceIdRoleIdResourceScopeId
// note: this method should only be used for API response data and not user input
func ParseIdentityGovernanceEntitlementManagementResourceIdRoleIdResourceScopeIDInsensitively(input string) (*IdentityGovernanceEntitlementManagementResourceIdRoleIdResourceScopeId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityGovernanceEntitlementManagementResourceIdRoleIdResourceScopeId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityGovernanceEntitlementManagementResourceIdRoleIdResourceScopeId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *IdentityGovernanceEntitlementManagementResourceIdRoleIdResourceScopeId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.AccessPackageResourceId, ok = input.Parsed["accessPackageResourceId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "accessPackageResourceId", input)
	}

	if id.AccessPackageResourceRoleId, ok = input.Parsed["accessPackageResourceRoleId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "accessPackageResourceRoleId", input)
	}

	if id.AccessPackageResourceScopeId, ok = input.Parsed["accessPackageResourceScopeId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "accessPackageResourceScopeId", input)
	}

	return nil
}

// ValidateIdentityGovernanceEntitlementManagementResourceIdRoleIdResourceScopeID checks that 'input' can be parsed as a Identity Governance Entitlement Management Resource Id Role Id Resource Scope ID
func ValidateIdentityGovernanceEntitlementManagementResourceIdRoleIdResourceScopeID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseIdentityGovernanceEntitlementManagementResourceIdRoleIdResourceScopeID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Identity Governance Entitlement Management Resource Id Role Id Resource Scope ID
func (id IdentityGovernanceEntitlementManagementResourceIdRoleIdResourceScopeId) ID() string {
	fmtString := "/identityGovernance/entitlementManagement/resources/%s/roles/%s/resource/scopes/%s"
	return fmt.Sprintf(fmtString, id.AccessPackageResourceId, id.AccessPackageResourceRoleId, id.AccessPackageResourceScopeId)
}

// Segments returns a slice of Resource ID Segments which comprise this Identity Governance Entitlement Management Resource Id Role Id Resource Scope ID
func (id IdentityGovernanceEntitlementManagementResourceIdRoleIdResourceScopeId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("identityGovernance", "identityGovernance", "identityGovernance"),
		resourceids.StaticSegment("entitlementManagement", "entitlementManagement", "entitlementManagement"),
		resourceids.StaticSegment("resources", "resources", "resources"),
		resourceids.UserSpecifiedSegment("accessPackageResourceId", "accessPackageResourceId"),
		resourceids.StaticSegment("roles", "roles", "roles"),
		resourceids.UserSpecifiedSegment("accessPackageResourceRoleId", "accessPackageResourceRoleId"),
		resourceids.StaticSegment("resource", "resource", "resource"),
		resourceids.StaticSegment("scopes", "scopes", "scopes"),
		resourceids.UserSpecifiedSegment("accessPackageResourceScopeId", "accessPackageResourceScopeId"),
	}
}

// String returns a human-readable description of this Identity Governance Entitlement Management Resource Id Role Id Resource Scope ID
func (id IdentityGovernanceEntitlementManagementResourceIdRoleIdResourceScopeId) String() string {
	components := []string{
		fmt.Sprintf("Access Package Resource: %q", id.AccessPackageResourceId),
		fmt.Sprintf("Access Package Resource Role: %q", id.AccessPackageResourceRoleId),
		fmt.Sprintf("Access Package Resource Scope: %q", id.AccessPackageResourceScopeId),
	}
	return fmt.Sprintf("Identity Governance Entitlement Management Resource Id Role Id Resource Scope (%s)", strings.Join(components, "\n"))
}
