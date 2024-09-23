package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &IdentityGovernanceEntitlementManagementResourceEnvironmentIdResourceIdRoleId{}

// IdentityGovernanceEntitlementManagementResourceEnvironmentIdResourceIdRoleId is a struct representing the Resource ID for a Identity Governance Entitlement Management Resource Environment Id Resource Id Role
type IdentityGovernanceEntitlementManagementResourceEnvironmentIdResourceIdRoleId struct {
	AccessPackageResourceEnvironmentId string
	AccessPackageResourceId            string
	AccessPackageResourceRoleId        string
}

// NewIdentityGovernanceEntitlementManagementResourceEnvironmentIdResourceIdRoleID returns a new IdentityGovernanceEntitlementManagementResourceEnvironmentIdResourceIdRoleId struct
func NewIdentityGovernanceEntitlementManagementResourceEnvironmentIdResourceIdRoleID(accessPackageResourceEnvironmentId string, accessPackageResourceId string, accessPackageResourceRoleId string) IdentityGovernanceEntitlementManagementResourceEnvironmentIdResourceIdRoleId {
	return IdentityGovernanceEntitlementManagementResourceEnvironmentIdResourceIdRoleId{
		AccessPackageResourceEnvironmentId: accessPackageResourceEnvironmentId,
		AccessPackageResourceId:            accessPackageResourceId,
		AccessPackageResourceRoleId:        accessPackageResourceRoleId,
	}
}

// ParseIdentityGovernanceEntitlementManagementResourceEnvironmentIdResourceIdRoleID parses 'input' into a IdentityGovernanceEntitlementManagementResourceEnvironmentIdResourceIdRoleId
func ParseIdentityGovernanceEntitlementManagementResourceEnvironmentIdResourceIdRoleID(input string) (*IdentityGovernanceEntitlementManagementResourceEnvironmentIdResourceIdRoleId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityGovernanceEntitlementManagementResourceEnvironmentIdResourceIdRoleId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityGovernanceEntitlementManagementResourceEnvironmentIdResourceIdRoleId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseIdentityGovernanceEntitlementManagementResourceEnvironmentIdResourceIdRoleIDInsensitively parses 'input' case-insensitively into a IdentityGovernanceEntitlementManagementResourceEnvironmentIdResourceIdRoleId
// note: this method should only be used for API response data and not user input
func ParseIdentityGovernanceEntitlementManagementResourceEnvironmentIdResourceIdRoleIDInsensitively(input string) (*IdentityGovernanceEntitlementManagementResourceEnvironmentIdResourceIdRoleId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityGovernanceEntitlementManagementResourceEnvironmentIdResourceIdRoleId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityGovernanceEntitlementManagementResourceEnvironmentIdResourceIdRoleId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *IdentityGovernanceEntitlementManagementResourceEnvironmentIdResourceIdRoleId) FromParseResult(input resourceids.ParseResult) error {
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

	return nil
}

// ValidateIdentityGovernanceEntitlementManagementResourceEnvironmentIdResourceIdRoleID checks that 'input' can be parsed as a Identity Governance Entitlement Management Resource Environment Id Resource Id Role ID
func ValidateIdentityGovernanceEntitlementManagementResourceEnvironmentIdResourceIdRoleID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseIdentityGovernanceEntitlementManagementResourceEnvironmentIdResourceIdRoleID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Identity Governance Entitlement Management Resource Environment Id Resource Id Role ID
func (id IdentityGovernanceEntitlementManagementResourceEnvironmentIdResourceIdRoleId) ID() string {
	fmtString := "/identityGovernance/entitlementManagement/resourceEnvironments/%s/resources/%s/roles/%s"
	return fmt.Sprintf(fmtString, id.AccessPackageResourceEnvironmentId, id.AccessPackageResourceId, id.AccessPackageResourceRoleId)
}

// Segments returns a slice of Resource ID Segments which comprise this Identity Governance Entitlement Management Resource Environment Id Resource Id Role ID
func (id IdentityGovernanceEntitlementManagementResourceEnvironmentIdResourceIdRoleId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("identityGovernance", "identityGovernance", "identityGovernance"),
		resourceids.StaticSegment("entitlementManagement", "entitlementManagement", "entitlementManagement"),
		resourceids.StaticSegment("resourceEnvironments", "resourceEnvironments", "resourceEnvironments"),
		resourceids.UserSpecifiedSegment("accessPackageResourceEnvironmentId", "accessPackageResourceEnvironmentId"),
		resourceids.StaticSegment("resources", "resources", "resources"),
		resourceids.UserSpecifiedSegment("accessPackageResourceId", "accessPackageResourceId"),
		resourceids.StaticSegment("roles", "roles", "roles"),
		resourceids.UserSpecifiedSegment("accessPackageResourceRoleId", "accessPackageResourceRoleId"),
	}
}

// String returns a human-readable description of this Identity Governance Entitlement Management Resource Environment Id Resource Id Role ID
func (id IdentityGovernanceEntitlementManagementResourceEnvironmentIdResourceIdRoleId) String() string {
	components := []string{
		fmt.Sprintf("Access Package Resource Environment: %q", id.AccessPackageResourceEnvironmentId),
		fmt.Sprintf("Access Package Resource: %q", id.AccessPackageResourceId),
		fmt.Sprintf("Access Package Resource Role: %q", id.AccessPackageResourceRoleId),
	}
	return fmt.Sprintf("Identity Governance Entitlement Management Resource Environment Id Resource Id Role (%s)", strings.Join(components, "\n"))
}
