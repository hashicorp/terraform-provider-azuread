package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &RoleManagementEntitlementManagementRoleAssignmentId{}

// RoleManagementEntitlementManagementRoleAssignmentId is a struct representing the Resource ID for a Role Management Entitlement Management Role Assignment
type RoleManagementEntitlementManagementRoleAssignmentId struct {
	UnifiedRoleAssignmentId string
}

// NewRoleManagementEntitlementManagementRoleAssignmentID returns a new RoleManagementEntitlementManagementRoleAssignmentId struct
func NewRoleManagementEntitlementManagementRoleAssignmentID(unifiedRoleAssignmentId string) RoleManagementEntitlementManagementRoleAssignmentId {
	return RoleManagementEntitlementManagementRoleAssignmentId{
		UnifiedRoleAssignmentId: unifiedRoleAssignmentId,
	}
}

// ParseRoleManagementEntitlementManagementRoleAssignmentID parses 'input' into a RoleManagementEntitlementManagementRoleAssignmentId
func ParseRoleManagementEntitlementManagementRoleAssignmentID(input string) (*RoleManagementEntitlementManagementRoleAssignmentId, error) {
	parser := resourceids.NewParserFromResourceIdType(&RoleManagementEntitlementManagementRoleAssignmentId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := RoleManagementEntitlementManagementRoleAssignmentId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseRoleManagementEntitlementManagementRoleAssignmentIDInsensitively parses 'input' case-insensitively into a RoleManagementEntitlementManagementRoleAssignmentId
// note: this method should only be used for API response data and not user input
func ParseRoleManagementEntitlementManagementRoleAssignmentIDInsensitively(input string) (*RoleManagementEntitlementManagementRoleAssignmentId, error) {
	parser := resourceids.NewParserFromResourceIdType(&RoleManagementEntitlementManagementRoleAssignmentId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := RoleManagementEntitlementManagementRoleAssignmentId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *RoleManagementEntitlementManagementRoleAssignmentId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.UnifiedRoleAssignmentId, ok = input.Parsed["unifiedRoleAssignmentId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "unifiedRoleAssignmentId", input)
	}

	return nil
}

// ValidateRoleManagementEntitlementManagementRoleAssignmentID checks that 'input' can be parsed as a Role Management Entitlement Management Role Assignment ID
func ValidateRoleManagementEntitlementManagementRoleAssignmentID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseRoleManagementEntitlementManagementRoleAssignmentID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Role Management Entitlement Management Role Assignment ID
func (id RoleManagementEntitlementManagementRoleAssignmentId) ID() string {
	fmtString := "/roleManagement/entitlementManagement/roleAssignments/%s"
	return fmt.Sprintf(fmtString, id.UnifiedRoleAssignmentId)
}

// Segments returns a slice of Resource ID Segments which comprise this Role Management Entitlement Management Role Assignment ID
func (id RoleManagementEntitlementManagementRoleAssignmentId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("roleManagement", "roleManagement", "roleManagement"),
		resourceids.StaticSegment("entitlementManagement", "entitlementManagement", "entitlementManagement"),
		resourceids.StaticSegment("roleAssignments", "roleAssignments", "roleAssignments"),
		resourceids.UserSpecifiedSegment("unifiedRoleAssignmentId", "unifiedRoleAssignmentId"),
	}
}

// String returns a human-readable description of this Role Management Entitlement Management Role Assignment ID
func (id RoleManagementEntitlementManagementRoleAssignmentId) String() string {
	components := []string{
		fmt.Sprintf("Unified Role Assignment: %q", id.UnifiedRoleAssignmentId),
	}
	return fmt.Sprintf("Role Management Entitlement Management Role Assignment (%s)", strings.Join(components, "\n"))
}
