package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &RoleManagementEntitlementManagementRoleAssignmentScheduleInstanceId{}

// RoleManagementEntitlementManagementRoleAssignmentScheduleInstanceId is a struct representing the Resource ID for a Role Management Entitlement Management Role Assignment Schedule Instance
type RoleManagementEntitlementManagementRoleAssignmentScheduleInstanceId struct {
	UnifiedRoleAssignmentScheduleInstanceId string
}

// NewRoleManagementEntitlementManagementRoleAssignmentScheduleInstanceID returns a new RoleManagementEntitlementManagementRoleAssignmentScheduleInstanceId struct
func NewRoleManagementEntitlementManagementRoleAssignmentScheduleInstanceID(unifiedRoleAssignmentScheduleInstanceId string) RoleManagementEntitlementManagementRoleAssignmentScheduleInstanceId {
	return RoleManagementEntitlementManagementRoleAssignmentScheduleInstanceId{
		UnifiedRoleAssignmentScheduleInstanceId: unifiedRoleAssignmentScheduleInstanceId,
	}
}

// ParseRoleManagementEntitlementManagementRoleAssignmentScheduleInstanceID parses 'input' into a RoleManagementEntitlementManagementRoleAssignmentScheduleInstanceId
func ParseRoleManagementEntitlementManagementRoleAssignmentScheduleInstanceID(input string) (*RoleManagementEntitlementManagementRoleAssignmentScheduleInstanceId, error) {
	parser := resourceids.NewParserFromResourceIdType(&RoleManagementEntitlementManagementRoleAssignmentScheduleInstanceId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := RoleManagementEntitlementManagementRoleAssignmentScheduleInstanceId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseRoleManagementEntitlementManagementRoleAssignmentScheduleInstanceIDInsensitively parses 'input' case-insensitively into a RoleManagementEntitlementManagementRoleAssignmentScheduleInstanceId
// note: this method should only be used for API response data and not user input
func ParseRoleManagementEntitlementManagementRoleAssignmentScheduleInstanceIDInsensitively(input string) (*RoleManagementEntitlementManagementRoleAssignmentScheduleInstanceId, error) {
	parser := resourceids.NewParserFromResourceIdType(&RoleManagementEntitlementManagementRoleAssignmentScheduleInstanceId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := RoleManagementEntitlementManagementRoleAssignmentScheduleInstanceId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *RoleManagementEntitlementManagementRoleAssignmentScheduleInstanceId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.UnifiedRoleAssignmentScheduleInstanceId, ok = input.Parsed["unifiedRoleAssignmentScheduleInstanceId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "unifiedRoleAssignmentScheduleInstanceId", input)
	}

	return nil
}

// ValidateRoleManagementEntitlementManagementRoleAssignmentScheduleInstanceID checks that 'input' can be parsed as a Role Management Entitlement Management Role Assignment Schedule Instance ID
func ValidateRoleManagementEntitlementManagementRoleAssignmentScheduleInstanceID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseRoleManagementEntitlementManagementRoleAssignmentScheduleInstanceID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Role Management Entitlement Management Role Assignment Schedule Instance ID
func (id RoleManagementEntitlementManagementRoleAssignmentScheduleInstanceId) ID() string {
	fmtString := "/roleManagement/entitlementManagement/roleAssignmentScheduleInstances/%s"
	return fmt.Sprintf(fmtString, id.UnifiedRoleAssignmentScheduleInstanceId)
}

// Segments returns a slice of Resource ID Segments which comprise this Role Management Entitlement Management Role Assignment Schedule Instance ID
func (id RoleManagementEntitlementManagementRoleAssignmentScheduleInstanceId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("roleManagement", "roleManagement", "roleManagement"),
		resourceids.StaticSegment("entitlementManagement", "entitlementManagement", "entitlementManagement"),
		resourceids.StaticSegment("roleAssignmentScheduleInstances", "roleAssignmentScheduleInstances", "roleAssignmentScheduleInstances"),
		resourceids.UserSpecifiedSegment("unifiedRoleAssignmentScheduleInstanceId", "unifiedRoleAssignmentScheduleInstanceId"),
	}
}

// String returns a human-readable description of this Role Management Entitlement Management Role Assignment Schedule Instance ID
func (id RoleManagementEntitlementManagementRoleAssignmentScheduleInstanceId) String() string {
	components := []string{
		fmt.Sprintf("Unified Role Assignment Schedule Instance: %q", id.UnifiedRoleAssignmentScheduleInstanceId),
	}
	return fmt.Sprintf("Role Management Entitlement Management Role Assignment Schedule Instance (%s)", strings.Join(components, "\n"))
}
