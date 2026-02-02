package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &RoleManagementDirectoryRoleEligibilityScheduleId{}

// RoleManagementDirectoryRoleEligibilityScheduleId is a struct representing the Resource ID for a Role Management Directory Role Eligibility Schedule
type RoleManagementDirectoryRoleEligibilityScheduleId struct {
	UnifiedRoleEligibilityScheduleId string
}

// NewRoleManagementDirectoryRoleEligibilityScheduleID returns a new RoleManagementDirectoryRoleEligibilityScheduleId struct
func NewRoleManagementDirectoryRoleEligibilityScheduleID(unifiedRoleEligibilityScheduleId string) RoleManagementDirectoryRoleEligibilityScheduleId {
	return RoleManagementDirectoryRoleEligibilityScheduleId{
		UnifiedRoleEligibilityScheduleId: unifiedRoleEligibilityScheduleId,
	}
}

// ParseRoleManagementDirectoryRoleEligibilityScheduleID parses 'input' into a RoleManagementDirectoryRoleEligibilityScheduleId
func ParseRoleManagementDirectoryRoleEligibilityScheduleID(input string) (*RoleManagementDirectoryRoleEligibilityScheduleId, error) {
	parser := resourceids.NewParserFromResourceIdType(&RoleManagementDirectoryRoleEligibilityScheduleId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := RoleManagementDirectoryRoleEligibilityScheduleId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseRoleManagementDirectoryRoleEligibilityScheduleIDInsensitively parses 'input' case-insensitively into a RoleManagementDirectoryRoleEligibilityScheduleId
// note: this method should only be used for API response data and not user input
func ParseRoleManagementDirectoryRoleEligibilityScheduleIDInsensitively(input string) (*RoleManagementDirectoryRoleEligibilityScheduleId, error) {
	parser := resourceids.NewParserFromResourceIdType(&RoleManagementDirectoryRoleEligibilityScheduleId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := RoleManagementDirectoryRoleEligibilityScheduleId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *RoleManagementDirectoryRoleEligibilityScheduleId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.UnifiedRoleEligibilityScheduleId, ok = input.Parsed["unifiedRoleEligibilityScheduleId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "unifiedRoleEligibilityScheduleId", input)
	}

	return nil
}

// ValidateRoleManagementDirectoryRoleEligibilityScheduleID checks that 'input' can be parsed as a Role Management Directory Role Eligibility Schedule ID
func ValidateRoleManagementDirectoryRoleEligibilityScheduleID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseRoleManagementDirectoryRoleEligibilityScheduleID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Role Management Directory Role Eligibility Schedule ID
func (id RoleManagementDirectoryRoleEligibilityScheduleId) ID() string {
	fmtString := "/roleManagement/directory/roleEligibilitySchedules/%s"
	return fmt.Sprintf(fmtString, id.UnifiedRoleEligibilityScheduleId)
}

// Segments returns a slice of Resource ID Segments which comprise this Role Management Directory Role Eligibility Schedule ID
func (id RoleManagementDirectoryRoleEligibilityScheduleId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("roleManagement", "roleManagement", "roleManagement"),
		resourceids.StaticSegment("directory", "directory", "directory"),
		resourceids.StaticSegment("roleEligibilitySchedules", "roleEligibilitySchedules", "roleEligibilitySchedules"),
		resourceids.UserSpecifiedSegment("unifiedRoleEligibilityScheduleId", "unifiedRoleEligibilityScheduleId"),
	}
}

// String returns a human-readable description of this Role Management Directory Role Eligibility Schedule ID
func (id RoleManagementDirectoryRoleEligibilityScheduleId) String() string {
	components := []string{
		fmt.Sprintf("Unified Role Eligibility Schedule: %q", id.UnifiedRoleEligibilityScheduleId),
	}
	return fmt.Sprintf("Role Management Directory Role Eligibility Schedule (%s)", strings.Join(components, "\n"))
}
