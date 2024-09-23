package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &RoleManagementDirectoryRoleEligibilityScheduleRequestId{}

// RoleManagementDirectoryRoleEligibilityScheduleRequestId is a struct representing the Resource ID for a Role Management Directory Role Eligibility Schedule Request
type RoleManagementDirectoryRoleEligibilityScheduleRequestId struct {
	UnifiedRoleEligibilityScheduleRequestId string
}

// NewRoleManagementDirectoryRoleEligibilityScheduleRequestID returns a new RoleManagementDirectoryRoleEligibilityScheduleRequestId struct
func NewRoleManagementDirectoryRoleEligibilityScheduleRequestID(unifiedRoleEligibilityScheduleRequestId string) RoleManagementDirectoryRoleEligibilityScheduleRequestId {
	return RoleManagementDirectoryRoleEligibilityScheduleRequestId{
		UnifiedRoleEligibilityScheduleRequestId: unifiedRoleEligibilityScheduleRequestId,
	}
}

// ParseRoleManagementDirectoryRoleEligibilityScheduleRequestID parses 'input' into a RoleManagementDirectoryRoleEligibilityScheduleRequestId
func ParseRoleManagementDirectoryRoleEligibilityScheduleRequestID(input string) (*RoleManagementDirectoryRoleEligibilityScheduleRequestId, error) {
	parser := resourceids.NewParserFromResourceIdType(&RoleManagementDirectoryRoleEligibilityScheduleRequestId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := RoleManagementDirectoryRoleEligibilityScheduleRequestId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseRoleManagementDirectoryRoleEligibilityScheduleRequestIDInsensitively parses 'input' case-insensitively into a RoleManagementDirectoryRoleEligibilityScheduleRequestId
// note: this method should only be used for API response data and not user input
func ParseRoleManagementDirectoryRoleEligibilityScheduleRequestIDInsensitively(input string) (*RoleManagementDirectoryRoleEligibilityScheduleRequestId, error) {
	parser := resourceids.NewParserFromResourceIdType(&RoleManagementDirectoryRoleEligibilityScheduleRequestId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := RoleManagementDirectoryRoleEligibilityScheduleRequestId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *RoleManagementDirectoryRoleEligibilityScheduleRequestId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.UnifiedRoleEligibilityScheduleRequestId, ok = input.Parsed["unifiedRoleEligibilityScheduleRequestId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "unifiedRoleEligibilityScheduleRequestId", input)
	}

	return nil
}

// ValidateRoleManagementDirectoryRoleEligibilityScheduleRequestID checks that 'input' can be parsed as a Role Management Directory Role Eligibility Schedule Request ID
func ValidateRoleManagementDirectoryRoleEligibilityScheduleRequestID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseRoleManagementDirectoryRoleEligibilityScheduleRequestID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Role Management Directory Role Eligibility Schedule Request ID
func (id RoleManagementDirectoryRoleEligibilityScheduleRequestId) ID() string {
	fmtString := "/roleManagement/directory/roleEligibilityScheduleRequests/%s"
	return fmt.Sprintf(fmtString, id.UnifiedRoleEligibilityScheduleRequestId)
}

// Segments returns a slice of Resource ID Segments which comprise this Role Management Directory Role Eligibility Schedule Request ID
func (id RoleManagementDirectoryRoleEligibilityScheduleRequestId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("roleManagement", "roleManagement", "roleManagement"),
		resourceids.StaticSegment("directory", "directory", "directory"),
		resourceids.StaticSegment("roleEligibilityScheduleRequests", "roleEligibilityScheduleRequests", "roleEligibilityScheduleRequests"),
		resourceids.UserSpecifiedSegment("unifiedRoleEligibilityScheduleRequestId", "unifiedRoleEligibilityScheduleRequestId"),
	}
}

// String returns a human-readable description of this Role Management Directory Role Eligibility Schedule Request ID
func (id RoleManagementDirectoryRoleEligibilityScheduleRequestId) String() string {
	components := []string{
		fmt.Sprintf("Unified Role Eligibility Schedule Request: %q", id.UnifiedRoleEligibilityScheduleRequestId),
	}
	return fmt.Sprintf("Role Management Directory Role Eligibility Schedule Request (%s)", strings.Join(components, "\n"))
}
