package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &IdentityGovernanceEntitlementManagementResourceEnvironmentIdResourceIdScopeIdResourceRoleId{}

// IdentityGovernanceEntitlementManagementResourceEnvironmentIdResourceIdScopeIdResourceRoleId is a struct representing the Resource ID for a Identity Governance Entitlement Management Resource Environment Id Resource Id Scope Id Resource Role
type IdentityGovernanceEntitlementManagementResourceEnvironmentIdResourceIdScopeIdResourceRoleId struct {
	AccessPackageResourceEnvironmentId string
	AccessPackageResourceId            string
	AccessPackageResourceScopeId       string
	AccessPackageResourceRoleId        string
}

// NewIdentityGovernanceEntitlementManagementResourceEnvironmentIdResourceIdScopeIdResourceRoleID returns a new IdentityGovernanceEntitlementManagementResourceEnvironmentIdResourceIdScopeIdResourceRoleId struct
func NewIdentityGovernanceEntitlementManagementResourceEnvironmentIdResourceIdScopeIdResourceRoleID(accessPackageResourceEnvironmentId string, accessPackageResourceId string, accessPackageResourceScopeId string, accessPackageResourceRoleId string) IdentityGovernanceEntitlementManagementResourceEnvironmentIdResourceIdScopeIdResourceRoleId {
	return IdentityGovernanceEntitlementManagementResourceEnvironmentIdResourceIdScopeIdResourceRoleId{
		AccessPackageResourceEnvironmentId: accessPackageResourceEnvironmentId,
		AccessPackageResourceId:            accessPackageResourceId,
		AccessPackageResourceScopeId:       accessPackageResourceScopeId,
		AccessPackageResourceRoleId:        accessPackageResourceRoleId,
	}
}

// ParseIdentityGovernanceEntitlementManagementResourceEnvironmentIdResourceIdScopeIdResourceRoleID parses 'input' into a IdentityGovernanceEntitlementManagementResourceEnvironmentIdResourceIdScopeIdResourceRoleId
func ParseIdentityGovernanceEntitlementManagementResourceEnvironmentIdResourceIdScopeIdResourceRoleID(input string) (*IdentityGovernanceEntitlementManagementResourceEnvironmentIdResourceIdScopeIdResourceRoleId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityGovernanceEntitlementManagementResourceEnvironmentIdResourceIdScopeIdResourceRoleId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityGovernanceEntitlementManagementResourceEnvironmentIdResourceIdScopeIdResourceRoleId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseIdentityGovernanceEntitlementManagementResourceEnvironmentIdResourceIdScopeIdResourceRoleIDInsensitively parses 'input' case-insensitively into a IdentityGovernanceEntitlementManagementResourceEnvironmentIdResourceIdScopeIdResourceRoleId
// note: this method should only be used for API response data and not user input
func ParseIdentityGovernanceEntitlementManagementResourceEnvironmentIdResourceIdScopeIdResourceRoleIDInsensitively(input string) (*IdentityGovernanceEntitlementManagementResourceEnvironmentIdResourceIdScopeIdResourceRoleId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityGovernanceEntitlementManagementResourceEnvironmentIdResourceIdScopeIdResourceRoleId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityGovernanceEntitlementManagementResourceEnvironmentIdResourceIdScopeIdResourceRoleId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *IdentityGovernanceEntitlementManagementResourceEnvironmentIdResourceIdScopeIdResourceRoleId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.AccessPackageResourceEnvironmentId, ok = input.Parsed["accessPackageResourceEnvironmentId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "accessPackageResourceEnvironmentId", input)
	}

	if id.AccessPackageResourceId, ok = input.Parsed["accessPackageResourceId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "accessPackageResourceId", input)
	}

	if id.AccessPackageResourceScopeId, ok = input.Parsed["accessPackageResourceScopeId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "accessPackageResourceScopeId", input)
	}

	if id.AccessPackageResourceRoleId, ok = input.Parsed["accessPackageResourceRoleId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "accessPackageResourceRoleId", input)
	}

	return nil
}

// ValidateIdentityGovernanceEntitlementManagementResourceEnvironmentIdResourceIdScopeIdResourceRoleID checks that 'input' can be parsed as a Identity Governance Entitlement Management Resource Environment Id Resource Id Scope Id Resource Role ID
func ValidateIdentityGovernanceEntitlementManagementResourceEnvironmentIdResourceIdScopeIdResourceRoleID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseIdentityGovernanceEntitlementManagementResourceEnvironmentIdResourceIdScopeIdResourceRoleID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Identity Governance Entitlement Management Resource Environment Id Resource Id Scope Id Resource Role ID
func (id IdentityGovernanceEntitlementManagementResourceEnvironmentIdResourceIdScopeIdResourceRoleId) ID() string {
	fmtString := "/identityGovernance/entitlementManagement/resourceEnvironments/%s/resources/%s/scopes/%s/resource/roles/%s"
	return fmt.Sprintf(fmtString, id.AccessPackageResourceEnvironmentId, id.AccessPackageResourceId, id.AccessPackageResourceScopeId, id.AccessPackageResourceRoleId)
}

// Segments returns a slice of Resource ID Segments which comprise this Identity Governance Entitlement Management Resource Environment Id Resource Id Scope Id Resource Role ID
func (id IdentityGovernanceEntitlementManagementResourceEnvironmentIdResourceIdScopeIdResourceRoleId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("identityGovernance", "identityGovernance", "identityGovernance"),
		resourceids.StaticSegment("entitlementManagement", "entitlementManagement", "entitlementManagement"),
		resourceids.StaticSegment("resourceEnvironments", "resourceEnvironments", "resourceEnvironments"),
		resourceids.UserSpecifiedSegment("accessPackageResourceEnvironmentId", "accessPackageResourceEnvironmentId"),
		resourceids.StaticSegment("resources", "resources", "resources"),
		resourceids.UserSpecifiedSegment("accessPackageResourceId", "accessPackageResourceId"),
		resourceids.StaticSegment("scopes", "scopes", "scopes"),
		resourceids.UserSpecifiedSegment("accessPackageResourceScopeId", "accessPackageResourceScopeId"),
		resourceids.StaticSegment("resource", "resource", "resource"),
		resourceids.StaticSegment("roles", "roles", "roles"),
		resourceids.UserSpecifiedSegment("accessPackageResourceRoleId", "accessPackageResourceRoleId"),
	}
}

// String returns a human-readable description of this Identity Governance Entitlement Management Resource Environment Id Resource Id Scope Id Resource Role ID
func (id IdentityGovernanceEntitlementManagementResourceEnvironmentIdResourceIdScopeIdResourceRoleId) String() string {
	components := []string{
		fmt.Sprintf("Access Package Resource Environment: %q", id.AccessPackageResourceEnvironmentId),
		fmt.Sprintf("Access Package Resource: %q", id.AccessPackageResourceId),
		fmt.Sprintf("Access Package Resource Scope: %q", id.AccessPackageResourceScopeId),
		fmt.Sprintf("Access Package Resource Role: %q", id.AccessPackageResourceRoleId),
	}
	return fmt.Sprintf("Identity Governance Entitlement Management Resource Environment Id Resource Id Scope Id Resource Role (%s)", strings.Join(components, "\n"))
}
