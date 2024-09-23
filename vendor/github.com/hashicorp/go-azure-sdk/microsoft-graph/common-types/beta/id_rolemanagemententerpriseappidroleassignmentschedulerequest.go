package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &RoleManagementEnterpriseAppIdRoleAssignmentScheduleRequestId{}

// RoleManagementEnterpriseAppIdRoleAssignmentScheduleRequestId is a struct representing the Resource ID for a Role Management Enterprise App Id Role Assignment Schedule Request
type RoleManagementEnterpriseAppIdRoleAssignmentScheduleRequestId struct {
	RbacApplicationId                      string
	UnifiedRoleAssignmentScheduleRequestId string
}

// NewRoleManagementEnterpriseAppIdRoleAssignmentScheduleRequestID returns a new RoleManagementEnterpriseAppIdRoleAssignmentScheduleRequestId struct
func NewRoleManagementEnterpriseAppIdRoleAssignmentScheduleRequestID(rbacApplicationId string, unifiedRoleAssignmentScheduleRequestId string) RoleManagementEnterpriseAppIdRoleAssignmentScheduleRequestId {
	return RoleManagementEnterpriseAppIdRoleAssignmentScheduleRequestId{
		RbacApplicationId:                      rbacApplicationId,
		UnifiedRoleAssignmentScheduleRequestId: unifiedRoleAssignmentScheduleRequestId,
	}
}

// ParseRoleManagementEnterpriseAppIdRoleAssignmentScheduleRequestID parses 'input' into a RoleManagementEnterpriseAppIdRoleAssignmentScheduleRequestId
func ParseRoleManagementEnterpriseAppIdRoleAssignmentScheduleRequestID(input string) (*RoleManagementEnterpriseAppIdRoleAssignmentScheduleRequestId, error) {
	parser := resourceids.NewParserFromResourceIdType(&RoleManagementEnterpriseAppIdRoleAssignmentScheduleRequestId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := RoleManagementEnterpriseAppIdRoleAssignmentScheduleRequestId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseRoleManagementEnterpriseAppIdRoleAssignmentScheduleRequestIDInsensitively parses 'input' case-insensitively into a RoleManagementEnterpriseAppIdRoleAssignmentScheduleRequestId
// note: this method should only be used for API response data and not user input
func ParseRoleManagementEnterpriseAppIdRoleAssignmentScheduleRequestIDInsensitively(input string) (*RoleManagementEnterpriseAppIdRoleAssignmentScheduleRequestId, error) {
	parser := resourceids.NewParserFromResourceIdType(&RoleManagementEnterpriseAppIdRoleAssignmentScheduleRequestId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := RoleManagementEnterpriseAppIdRoleAssignmentScheduleRequestId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *RoleManagementEnterpriseAppIdRoleAssignmentScheduleRequestId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.RbacApplicationId, ok = input.Parsed["rbacApplicationId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "rbacApplicationId", input)
	}

	if id.UnifiedRoleAssignmentScheduleRequestId, ok = input.Parsed["unifiedRoleAssignmentScheduleRequestId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "unifiedRoleAssignmentScheduleRequestId", input)
	}

	return nil
}

// ValidateRoleManagementEnterpriseAppIdRoleAssignmentScheduleRequestID checks that 'input' can be parsed as a Role Management Enterprise App Id Role Assignment Schedule Request ID
func ValidateRoleManagementEnterpriseAppIdRoleAssignmentScheduleRequestID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseRoleManagementEnterpriseAppIdRoleAssignmentScheduleRequestID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Role Management Enterprise App Id Role Assignment Schedule Request ID
func (id RoleManagementEnterpriseAppIdRoleAssignmentScheduleRequestId) ID() string {
	fmtString := "/roleManagement/enterpriseApps/%s/roleAssignmentScheduleRequests/%s"
	return fmt.Sprintf(fmtString, id.RbacApplicationId, id.UnifiedRoleAssignmentScheduleRequestId)
}

// Segments returns a slice of Resource ID Segments which comprise this Role Management Enterprise App Id Role Assignment Schedule Request ID
func (id RoleManagementEnterpriseAppIdRoleAssignmentScheduleRequestId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("roleManagement", "roleManagement", "roleManagement"),
		resourceids.StaticSegment("enterpriseApps", "enterpriseApps", "enterpriseApps"),
		resourceids.UserSpecifiedSegment("rbacApplicationId", "rbacApplicationId"),
		resourceids.StaticSegment("roleAssignmentScheduleRequests", "roleAssignmentScheduleRequests", "roleAssignmentScheduleRequests"),
		resourceids.UserSpecifiedSegment("unifiedRoleAssignmentScheduleRequestId", "unifiedRoleAssignmentScheduleRequestId"),
	}
}

// String returns a human-readable description of this Role Management Enterprise App Id Role Assignment Schedule Request ID
func (id RoleManagementEnterpriseAppIdRoleAssignmentScheduleRequestId) String() string {
	components := []string{
		fmt.Sprintf("Rbac Application: %q", id.RbacApplicationId),
		fmt.Sprintf("Unified Role Assignment Schedule Request: %q", id.UnifiedRoleAssignmentScheduleRequestId),
	}
	return fmt.Sprintf("Role Management Enterprise App Id Role Assignment Schedule Request (%s)", strings.Join(components, "\n"))
}
