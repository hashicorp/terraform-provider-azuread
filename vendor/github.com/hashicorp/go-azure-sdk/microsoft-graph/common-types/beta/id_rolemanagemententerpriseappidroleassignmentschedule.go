package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &RoleManagementEnterpriseAppIdRoleAssignmentScheduleId{}

// RoleManagementEnterpriseAppIdRoleAssignmentScheduleId is a struct representing the Resource ID for a Role Management Enterprise App Id Role Assignment Schedule
type RoleManagementEnterpriseAppIdRoleAssignmentScheduleId struct {
	RbacApplicationId               string
	UnifiedRoleAssignmentScheduleId string
}

// NewRoleManagementEnterpriseAppIdRoleAssignmentScheduleID returns a new RoleManagementEnterpriseAppIdRoleAssignmentScheduleId struct
func NewRoleManagementEnterpriseAppIdRoleAssignmentScheduleID(rbacApplicationId string, unifiedRoleAssignmentScheduleId string) RoleManagementEnterpriseAppIdRoleAssignmentScheduleId {
	return RoleManagementEnterpriseAppIdRoleAssignmentScheduleId{
		RbacApplicationId:               rbacApplicationId,
		UnifiedRoleAssignmentScheduleId: unifiedRoleAssignmentScheduleId,
	}
}

// ParseRoleManagementEnterpriseAppIdRoleAssignmentScheduleID parses 'input' into a RoleManagementEnterpriseAppIdRoleAssignmentScheduleId
func ParseRoleManagementEnterpriseAppIdRoleAssignmentScheduleID(input string) (*RoleManagementEnterpriseAppIdRoleAssignmentScheduleId, error) {
	parser := resourceids.NewParserFromResourceIdType(&RoleManagementEnterpriseAppIdRoleAssignmentScheduleId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := RoleManagementEnterpriseAppIdRoleAssignmentScheduleId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseRoleManagementEnterpriseAppIdRoleAssignmentScheduleIDInsensitively parses 'input' case-insensitively into a RoleManagementEnterpriseAppIdRoleAssignmentScheduleId
// note: this method should only be used for API response data and not user input
func ParseRoleManagementEnterpriseAppIdRoleAssignmentScheduleIDInsensitively(input string) (*RoleManagementEnterpriseAppIdRoleAssignmentScheduleId, error) {
	parser := resourceids.NewParserFromResourceIdType(&RoleManagementEnterpriseAppIdRoleAssignmentScheduleId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := RoleManagementEnterpriseAppIdRoleAssignmentScheduleId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *RoleManagementEnterpriseAppIdRoleAssignmentScheduleId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.RbacApplicationId, ok = input.Parsed["rbacApplicationId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "rbacApplicationId", input)
	}

	if id.UnifiedRoleAssignmentScheduleId, ok = input.Parsed["unifiedRoleAssignmentScheduleId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "unifiedRoleAssignmentScheduleId", input)
	}

	return nil
}

// ValidateRoleManagementEnterpriseAppIdRoleAssignmentScheduleID checks that 'input' can be parsed as a Role Management Enterprise App Id Role Assignment Schedule ID
func ValidateRoleManagementEnterpriseAppIdRoleAssignmentScheduleID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseRoleManagementEnterpriseAppIdRoleAssignmentScheduleID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Role Management Enterprise App Id Role Assignment Schedule ID
func (id RoleManagementEnterpriseAppIdRoleAssignmentScheduleId) ID() string {
	fmtString := "/roleManagement/enterpriseApps/%s/roleAssignmentSchedules/%s"
	return fmt.Sprintf(fmtString, id.RbacApplicationId, id.UnifiedRoleAssignmentScheduleId)
}

// Segments returns a slice of Resource ID Segments which comprise this Role Management Enterprise App Id Role Assignment Schedule ID
func (id RoleManagementEnterpriseAppIdRoleAssignmentScheduleId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("roleManagement", "roleManagement", "roleManagement"),
		resourceids.StaticSegment("enterpriseApps", "enterpriseApps", "enterpriseApps"),
		resourceids.UserSpecifiedSegment("rbacApplicationId", "rbacApplicationId"),
		resourceids.StaticSegment("roleAssignmentSchedules", "roleAssignmentSchedules", "roleAssignmentSchedules"),
		resourceids.UserSpecifiedSegment("unifiedRoleAssignmentScheduleId", "unifiedRoleAssignmentScheduleId"),
	}
}

// String returns a human-readable description of this Role Management Enterprise App Id Role Assignment Schedule ID
func (id RoleManagementEnterpriseAppIdRoleAssignmentScheduleId) String() string {
	components := []string{
		fmt.Sprintf("Rbac Application: %q", id.RbacApplicationId),
		fmt.Sprintf("Unified Role Assignment Schedule: %q", id.UnifiedRoleAssignmentScheduleId),
	}
	return fmt.Sprintf("Role Management Enterprise App Id Role Assignment Schedule (%s)", strings.Join(components, "\n"))
}
