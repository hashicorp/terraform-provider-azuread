package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &RoleManagementEnterpriseAppIdRoleAssignmentScheduleInstanceId{}

// RoleManagementEnterpriseAppIdRoleAssignmentScheduleInstanceId is a struct representing the Resource ID for a Role Management Enterprise App Id Role Assignment Schedule Instance
type RoleManagementEnterpriseAppIdRoleAssignmentScheduleInstanceId struct {
	RbacApplicationId                       string
	UnifiedRoleAssignmentScheduleInstanceId string
}

// NewRoleManagementEnterpriseAppIdRoleAssignmentScheduleInstanceID returns a new RoleManagementEnterpriseAppIdRoleAssignmentScheduleInstanceId struct
func NewRoleManagementEnterpriseAppIdRoleAssignmentScheduleInstanceID(rbacApplicationId string, unifiedRoleAssignmentScheduleInstanceId string) RoleManagementEnterpriseAppIdRoleAssignmentScheduleInstanceId {
	return RoleManagementEnterpriseAppIdRoleAssignmentScheduleInstanceId{
		RbacApplicationId:                       rbacApplicationId,
		UnifiedRoleAssignmentScheduleInstanceId: unifiedRoleAssignmentScheduleInstanceId,
	}
}

// ParseRoleManagementEnterpriseAppIdRoleAssignmentScheduleInstanceID parses 'input' into a RoleManagementEnterpriseAppIdRoleAssignmentScheduleInstanceId
func ParseRoleManagementEnterpriseAppIdRoleAssignmentScheduleInstanceID(input string) (*RoleManagementEnterpriseAppIdRoleAssignmentScheduleInstanceId, error) {
	parser := resourceids.NewParserFromResourceIdType(&RoleManagementEnterpriseAppIdRoleAssignmentScheduleInstanceId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := RoleManagementEnterpriseAppIdRoleAssignmentScheduleInstanceId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseRoleManagementEnterpriseAppIdRoleAssignmentScheduleInstanceIDInsensitively parses 'input' case-insensitively into a RoleManagementEnterpriseAppIdRoleAssignmentScheduleInstanceId
// note: this method should only be used for API response data and not user input
func ParseRoleManagementEnterpriseAppIdRoleAssignmentScheduleInstanceIDInsensitively(input string) (*RoleManagementEnterpriseAppIdRoleAssignmentScheduleInstanceId, error) {
	parser := resourceids.NewParserFromResourceIdType(&RoleManagementEnterpriseAppIdRoleAssignmentScheduleInstanceId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := RoleManagementEnterpriseAppIdRoleAssignmentScheduleInstanceId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *RoleManagementEnterpriseAppIdRoleAssignmentScheduleInstanceId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.RbacApplicationId, ok = input.Parsed["rbacApplicationId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "rbacApplicationId", input)
	}

	if id.UnifiedRoleAssignmentScheduleInstanceId, ok = input.Parsed["unifiedRoleAssignmentScheduleInstanceId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "unifiedRoleAssignmentScheduleInstanceId", input)
	}

	return nil
}

// ValidateRoleManagementEnterpriseAppIdRoleAssignmentScheduleInstanceID checks that 'input' can be parsed as a Role Management Enterprise App Id Role Assignment Schedule Instance ID
func ValidateRoleManagementEnterpriseAppIdRoleAssignmentScheduleInstanceID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseRoleManagementEnterpriseAppIdRoleAssignmentScheduleInstanceID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Role Management Enterprise App Id Role Assignment Schedule Instance ID
func (id RoleManagementEnterpriseAppIdRoleAssignmentScheduleInstanceId) ID() string {
	fmtString := "/roleManagement/enterpriseApps/%s/roleAssignmentScheduleInstances/%s"
	return fmt.Sprintf(fmtString, id.RbacApplicationId, id.UnifiedRoleAssignmentScheduleInstanceId)
}

// Segments returns a slice of Resource ID Segments which comprise this Role Management Enterprise App Id Role Assignment Schedule Instance ID
func (id RoleManagementEnterpriseAppIdRoleAssignmentScheduleInstanceId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("roleManagement", "roleManagement", "roleManagement"),
		resourceids.StaticSegment("enterpriseApps", "enterpriseApps", "enterpriseApps"),
		resourceids.UserSpecifiedSegment("rbacApplicationId", "rbacApplicationId"),
		resourceids.StaticSegment("roleAssignmentScheduleInstances", "roleAssignmentScheduleInstances", "roleAssignmentScheduleInstances"),
		resourceids.UserSpecifiedSegment("unifiedRoleAssignmentScheduleInstanceId", "unifiedRoleAssignmentScheduleInstanceId"),
	}
}

// String returns a human-readable description of this Role Management Enterprise App Id Role Assignment Schedule Instance ID
func (id RoleManagementEnterpriseAppIdRoleAssignmentScheduleInstanceId) String() string {
	components := []string{
		fmt.Sprintf("Rbac Application: %q", id.RbacApplicationId),
		fmt.Sprintf("Unified Role Assignment Schedule Instance: %q", id.UnifiedRoleAssignmentScheduleInstanceId),
	}
	return fmt.Sprintf("Role Management Enterprise App Id Role Assignment Schedule Instance (%s)", strings.Join(components, "\n"))
}
