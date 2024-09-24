package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &IdentityGovernanceEntitlementManagementAccessPackageAssignmentIdAccessPackageIncompatibleGroupId{}

// IdentityGovernanceEntitlementManagementAccessPackageAssignmentIdAccessPackageIncompatibleGroupId is a struct representing the Resource ID for a Identity Governance Entitlement Management Access Package Assignment Id Access Package Incompatible Group
type IdentityGovernanceEntitlementManagementAccessPackageAssignmentIdAccessPackageIncompatibleGroupId struct {
	AccessPackageAssignmentId string
	GroupId                   string
}

// NewIdentityGovernanceEntitlementManagementAccessPackageAssignmentIdAccessPackageIncompatibleGroupID returns a new IdentityGovernanceEntitlementManagementAccessPackageAssignmentIdAccessPackageIncompatibleGroupId struct
func NewIdentityGovernanceEntitlementManagementAccessPackageAssignmentIdAccessPackageIncompatibleGroupID(accessPackageAssignmentId string, groupId string) IdentityGovernanceEntitlementManagementAccessPackageAssignmentIdAccessPackageIncompatibleGroupId {
	return IdentityGovernanceEntitlementManagementAccessPackageAssignmentIdAccessPackageIncompatibleGroupId{
		AccessPackageAssignmentId: accessPackageAssignmentId,
		GroupId:                   groupId,
	}
}

// ParseIdentityGovernanceEntitlementManagementAccessPackageAssignmentIdAccessPackageIncompatibleGroupID parses 'input' into a IdentityGovernanceEntitlementManagementAccessPackageAssignmentIdAccessPackageIncompatibleGroupId
func ParseIdentityGovernanceEntitlementManagementAccessPackageAssignmentIdAccessPackageIncompatibleGroupID(input string) (*IdentityGovernanceEntitlementManagementAccessPackageAssignmentIdAccessPackageIncompatibleGroupId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityGovernanceEntitlementManagementAccessPackageAssignmentIdAccessPackageIncompatibleGroupId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityGovernanceEntitlementManagementAccessPackageAssignmentIdAccessPackageIncompatibleGroupId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseIdentityGovernanceEntitlementManagementAccessPackageAssignmentIdAccessPackageIncompatibleGroupIDInsensitively parses 'input' case-insensitively into a IdentityGovernanceEntitlementManagementAccessPackageAssignmentIdAccessPackageIncompatibleGroupId
// note: this method should only be used for API response data and not user input
func ParseIdentityGovernanceEntitlementManagementAccessPackageAssignmentIdAccessPackageIncompatibleGroupIDInsensitively(input string) (*IdentityGovernanceEntitlementManagementAccessPackageAssignmentIdAccessPackageIncompatibleGroupId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityGovernanceEntitlementManagementAccessPackageAssignmentIdAccessPackageIncompatibleGroupId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityGovernanceEntitlementManagementAccessPackageAssignmentIdAccessPackageIncompatibleGroupId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *IdentityGovernanceEntitlementManagementAccessPackageAssignmentIdAccessPackageIncompatibleGroupId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.AccessPackageAssignmentId, ok = input.Parsed["accessPackageAssignmentId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "accessPackageAssignmentId", input)
	}

	if id.GroupId, ok = input.Parsed["groupId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "groupId", input)
	}

	return nil
}

// ValidateIdentityGovernanceEntitlementManagementAccessPackageAssignmentIdAccessPackageIncompatibleGroupID checks that 'input' can be parsed as a Identity Governance Entitlement Management Access Package Assignment Id Access Package Incompatible Group ID
func ValidateIdentityGovernanceEntitlementManagementAccessPackageAssignmentIdAccessPackageIncompatibleGroupID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseIdentityGovernanceEntitlementManagementAccessPackageAssignmentIdAccessPackageIncompatibleGroupID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Identity Governance Entitlement Management Access Package Assignment Id Access Package Incompatible Group ID
func (id IdentityGovernanceEntitlementManagementAccessPackageAssignmentIdAccessPackageIncompatibleGroupId) ID() string {
	fmtString := "/identityGovernance/entitlementManagement/accessPackageAssignments/%s/accessPackage/incompatibleGroups/%s"
	return fmt.Sprintf(fmtString, id.AccessPackageAssignmentId, id.GroupId)
}

// Segments returns a slice of Resource ID Segments which comprise this Identity Governance Entitlement Management Access Package Assignment Id Access Package Incompatible Group ID
func (id IdentityGovernanceEntitlementManagementAccessPackageAssignmentIdAccessPackageIncompatibleGroupId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("identityGovernance", "identityGovernance", "identityGovernance"),
		resourceids.StaticSegment("entitlementManagement", "entitlementManagement", "entitlementManagement"),
		resourceids.StaticSegment("accessPackageAssignments", "accessPackageAssignments", "accessPackageAssignments"),
		resourceids.UserSpecifiedSegment("accessPackageAssignmentId", "accessPackageAssignmentId"),
		resourceids.StaticSegment("accessPackage", "accessPackage", "accessPackage"),
		resourceids.StaticSegment("incompatibleGroups", "incompatibleGroups", "incompatibleGroups"),
		resourceids.UserSpecifiedSegment("groupId", "groupId"),
	}
}

// String returns a human-readable description of this Identity Governance Entitlement Management Access Package Assignment Id Access Package Incompatible Group ID
func (id IdentityGovernanceEntitlementManagementAccessPackageAssignmentIdAccessPackageIncompatibleGroupId) String() string {
	components := []string{
		fmt.Sprintf("Access Package Assignment: %q", id.AccessPackageAssignmentId),
		fmt.Sprintf("Group: %q", id.GroupId),
	}
	return fmt.Sprintf("Identity Governance Entitlement Management Access Package Assignment Id Access Package Incompatible Group (%s)", strings.Join(components, "\n"))
}
