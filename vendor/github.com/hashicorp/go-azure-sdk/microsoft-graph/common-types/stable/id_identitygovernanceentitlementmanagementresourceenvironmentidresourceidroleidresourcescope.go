package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &IdentityGovernanceEntitlementManagementResourceEnvironmentIdResourceIdRoleIdResourceScopeId{}

// IdentityGovernanceEntitlementManagementResourceEnvironmentIdResourceIdRoleIdResourceScopeId is a struct representing the Resource ID for a Identity Governance Entitlement Management Resource Environment Id Resource Id Role Id Resource Scope
type IdentityGovernanceEntitlementManagementResourceEnvironmentIdResourceIdRoleIdResourceScopeId struct {
	AccessPackageResourceEnvironmentId string
	AccessPackageResourceId            string
	AccessPackageResourceRoleId        string
	AccessPackageResourceScopeId       string
}

// NewIdentityGovernanceEntitlementManagementResourceEnvironmentIdResourceIdRoleIdResourceScopeID returns a new IdentityGovernanceEntitlementManagementResourceEnvironmentIdResourceIdRoleIdResourceScopeId struct
func NewIdentityGovernanceEntitlementManagementResourceEnvironmentIdResourceIdRoleIdResourceScopeID(accessPackageResourceEnvironmentId string, accessPackageResourceId string, accessPackageResourceRoleId string, accessPackageResourceScopeId string) IdentityGovernanceEntitlementManagementResourceEnvironmentIdResourceIdRoleIdResourceScopeId {
	return IdentityGovernanceEntitlementManagementResourceEnvironmentIdResourceIdRoleIdResourceScopeId{
		AccessPackageResourceEnvironmentId: accessPackageResourceEnvironmentId,
		AccessPackageResourceId:            accessPackageResourceId,
		AccessPackageResourceRoleId:        accessPackageResourceRoleId,
		AccessPackageResourceScopeId:       accessPackageResourceScopeId,
	}
}

// ParseIdentityGovernanceEntitlementManagementResourceEnvironmentIdResourceIdRoleIdResourceScopeID parses 'input' into a IdentityGovernanceEntitlementManagementResourceEnvironmentIdResourceIdRoleIdResourceScopeId
func ParseIdentityGovernanceEntitlementManagementResourceEnvironmentIdResourceIdRoleIdResourceScopeID(input string) (*IdentityGovernanceEntitlementManagementResourceEnvironmentIdResourceIdRoleIdResourceScopeId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityGovernanceEntitlementManagementResourceEnvironmentIdResourceIdRoleIdResourceScopeId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityGovernanceEntitlementManagementResourceEnvironmentIdResourceIdRoleIdResourceScopeId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseIdentityGovernanceEntitlementManagementResourceEnvironmentIdResourceIdRoleIdResourceScopeIDInsensitively parses 'input' case-insensitively into a IdentityGovernanceEntitlementManagementResourceEnvironmentIdResourceIdRoleIdResourceScopeId
// note: this method should only be used for API response data and not user input
func ParseIdentityGovernanceEntitlementManagementResourceEnvironmentIdResourceIdRoleIdResourceScopeIDInsensitively(input string) (*IdentityGovernanceEntitlementManagementResourceEnvironmentIdResourceIdRoleIdResourceScopeId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityGovernanceEntitlementManagementResourceEnvironmentIdResourceIdRoleIdResourceScopeId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityGovernanceEntitlementManagementResourceEnvironmentIdResourceIdRoleIdResourceScopeId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *IdentityGovernanceEntitlementManagementResourceEnvironmentIdResourceIdRoleIdResourceScopeId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.AccessPackageResourceEnvironmentId, ok = input.Parsed["accessPackageResourceEnvironmentId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "accessPackageResourceEnvironmentId", input)
	}

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

// ValidateIdentityGovernanceEntitlementManagementResourceEnvironmentIdResourceIdRoleIdResourceScopeID checks that 'input' can be parsed as a Identity Governance Entitlement Management Resource Environment Id Resource Id Role Id Resource Scope ID
func ValidateIdentityGovernanceEntitlementManagementResourceEnvironmentIdResourceIdRoleIdResourceScopeID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseIdentityGovernanceEntitlementManagementResourceEnvironmentIdResourceIdRoleIdResourceScopeID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Identity Governance Entitlement Management Resource Environment Id Resource Id Role Id Resource Scope ID
func (id IdentityGovernanceEntitlementManagementResourceEnvironmentIdResourceIdRoleIdResourceScopeId) ID() string {
	fmtString := "/identityGovernance/entitlementManagement/resourceEnvironments/%s/resources/%s/roles/%s/resource/scopes/%s"
	return fmt.Sprintf(fmtString, id.AccessPackageResourceEnvironmentId, id.AccessPackageResourceId, id.AccessPackageResourceRoleId, id.AccessPackageResourceScopeId)
}

// Segments returns a slice of Resource ID Segments which comprise this Identity Governance Entitlement Management Resource Environment Id Resource Id Role Id Resource Scope ID
func (id IdentityGovernanceEntitlementManagementResourceEnvironmentIdResourceIdRoleIdResourceScopeId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("identityGovernance", "identityGovernance", "identityGovernance"),
		resourceids.StaticSegment("entitlementManagement", "entitlementManagement", "entitlementManagement"),
		resourceids.StaticSegment("resourceEnvironments", "resourceEnvironments", "resourceEnvironments"),
		resourceids.UserSpecifiedSegment("accessPackageResourceEnvironmentId", "accessPackageResourceEnvironmentId"),
		resourceids.StaticSegment("resources", "resources", "resources"),
		resourceids.UserSpecifiedSegment("accessPackageResourceId", "accessPackageResourceId"),
		resourceids.StaticSegment("roles", "roles", "roles"),
		resourceids.UserSpecifiedSegment("accessPackageResourceRoleId", "accessPackageResourceRoleId"),
		resourceids.StaticSegment("resource", "resource", "resource"),
		resourceids.StaticSegment("scopes", "scopes", "scopes"),
		resourceids.UserSpecifiedSegment("accessPackageResourceScopeId", "accessPackageResourceScopeId"),
	}
}

// String returns a human-readable description of this Identity Governance Entitlement Management Resource Environment Id Resource Id Role Id Resource Scope ID
func (id IdentityGovernanceEntitlementManagementResourceEnvironmentIdResourceIdRoleIdResourceScopeId) String() string {
	components := []string{
		fmt.Sprintf("Access Package Resource Environment: %q", id.AccessPackageResourceEnvironmentId),
		fmt.Sprintf("Access Package Resource: %q", id.AccessPackageResourceId),
		fmt.Sprintf("Access Package Resource Role: %q", id.AccessPackageResourceRoleId),
		fmt.Sprintf("Access Package Resource Scope: %q", id.AccessPackageResourceScopeId),
	}
	return fmt.Sprintf("Identity Governance Entitlement Management Resource Environment Id Resource Id Role Id Resource Scope (%s)", strings.Join(components, "\n"))
}
