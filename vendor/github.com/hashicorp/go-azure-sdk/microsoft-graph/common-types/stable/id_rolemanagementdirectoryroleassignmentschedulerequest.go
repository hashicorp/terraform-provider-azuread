package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &RoleManagementDirectoryRoleAssignmentScheduleRequestId{}

// RoleManagementDirectoryRoleAssignmentScheduleRequestId is a struct representing the Resource ID for a Role Management Directory Role Assignment Schedule Request
type RoleManagementDirectoryRoleAssignmentScheduleRequestId struct {
	UnifiedRoleAssignmentScheduleRequestId string
}

// NewRoleManagementDirectoryRoleAssignmentScheduleRequestID returns a new RoleManagementDirectoryRoleAssignmentScheduleRequestId struct
func NewRoleManagementDirectoryRoleAssignmentScheduleRequestID(unifiedRoleAssignmentScheduleRequestId string) RoleManagementDirectoryRoleAssignmentScheduleRequestId {
	return RoleManagementDirectoryRoleAssignmentScheduleRequestId{
		UnifiedRoleAssignmentScheduleRequestId: unifiedRoleAssignmentScheduleRequestId,
	}
}

// ParseRoleManagementDirectoryRoleAssignmentScheduleRequestID parses 'input' into a RoleManagementDirectoryRoleAssignmentScheduleRequestId
func ParseRoleManagementDirectoryRoleAssignmentScheduleRequestID(input string) (*RoleManagementDirectoryRoleAssignmentScheduleRequestId, error) {
	parser := resourceids.NewParserFromResourceIdType(&RoleManagementDirectoryRoleAssignmentScheduleRequestId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := RoleManagementDirectoryRoleAssignmentScheduleRequestId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseRoleManagementDirectoryRoleAssignmentScheduleRequestIDInsensitively parses 'input' case-insensitively into a RoleManagementDirectoryRoleAssignmentScheduleRequestId
// note: this method should only be used for API response data and not user input
func ParseRoleManagementDirectoryRoleAssignmentScheduleRequestIDInsensitively(input string) (*RoleManagementDirectoryRoleAssignmentScheduleRequestId, error) {
	parser := resourceids.NewParserFromResourceIdType(&RoleManagementDirectoryRoleAssignmentScheduleRequestId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := RoleManagementDirectoryRoleAssignmentScheduleRequestId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *RoleManagementDirectoryRoleAssignmentScheduleRequestId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.UnifiedRoleAssignmentScheduleRequestId, ok = input.Parsed["unifiedRoleAssignmentScheduleRequestId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "unifiedRoleAssignmentScheduleRequestId", input)
	}

	return nil
}

// ValidateRoleManagementDirectoryRoleAssignmentScheduleRequestID checks that 'input' can be parsed as a Role Management Directory Role Assignment Schedule Request ID
func ValidateRoleManagementDirectoryRoleAssignmentScheduleRequestID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseRoleManagementDirectoryRoleAssignmentScheduleRequestID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Role Management Directory Role Assignment Schedule Request ID
func (id RoleManagementDirectoryRoleAssignmentScheduleRequestId) ID() string {
	fmtString := "/roleManagement/directory/roleAssignmentScheduleRequests/%s"
	return fmt.Sprintf(fmtString, id.UnifiedRoleAssignmentScheduleRequestId)
}

// Segments returns a slice of Resource ID Segments which comprise this Role Management Directory Role Assignment Schedule Request ID
func (id RoleManagementDirectoryRoleAssignmentScheduleRequestId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("roleManagement", "roleManagement", "roleManagement"),
		resourceids.StaticSegment("directory", "directory", "directory"),
		resourceids.StaticSegment("roleAssignmentScheduleRequests", "roleAssignmentScheduleRequests", "roleAssignmentScheduleRequests"),
		resourceids.UserSpecifiedSegment("unifiedRoleAssignmentScheduleRequestId", "unifiedRoleAssignmentScheduleRequestId"),
	}
}

// String returns a human-readable description of this Role Management Directory Role Assignment Schedule Request ID
func (id RoleManagementDirectoryRoleAssignmentScheduleRequestId) String() string {
	components := []string{
		fmt.Sprintf("Unified Role Assignment Schedule Request: %q", id.UnifiedRoleAssignmentScheduleRequestId),
	}
	return fmt.Sprintf("Role Management Directory Role Assignment Schedule Request (%s)", strings.Join(components, "\n"))
}
