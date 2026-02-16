package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &RoleManagementDirectoryRoleEligibilityScheduleInstanceId{}

// RoleManagementDirectoryRoleEligibilityScheduleInstanceId is a struct representing the Resource ID for a Role Management Directory Role Eligibility Schedule Instance
type RoleManagementDirectoryRoleEligibilityScheduleInstanceId struct {
	UnifiedRoleEligibilityScheduleInstanceId string
}

// NewRoleManagementDirectoryRoleEligibilityScheduleInstanceID returns a new RoleManagementDirectoryRoleEligibilityScheduleInstanceId struct
func NewRoleManagementDirectoryRoleEligibilityScheduleInstanceID(unifiedRoleEligibilityScheduleInstanceId string) RoleManagementDirectoryRoleEligibilityScheduleInstanceId {
	return RoleManagementDirectoryRoleEligibilityScheduleInstanceId{
		UnifiedRoleEligibilityScheduleInstanceId: unifiedRoleEligibilityScheduleInstanceId,
	}
}

// ParseRoleManagementDirectoryRoleEligibilityScheduleInstanceID parses 'input' into a RoleManagementDirectoryRoleEligibilityScheduleInstanceId
func ParseRoleManagementDirectoryRoleEligibilityScheduleInstanceID(input string) (*RoleManagementDirectoryRoleEligibilityScheduleInstanceId, error) {
	parser := resourceids.NewParserFromResourceIdType(&RoleManagementDirectoryRoleEligibilityScheduleInstanceId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := RoleManagementDirectoryRoleEligibilityScheduleInstanceId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseRoleManagementDirectoryRoleEligibilityScheduleInstanceIDInsensitively parses 'input' case-insensitively into a RoleManagementDirectoryRoleEligibilityScheduleInstanceId
// note: this method should only be used for API response data and not user input
func ParseRoleManagementDirectoryRoleEligibilityScheduleInstanceIDInsensitively(input string) (*RoleManagementDirectoryRoleEligibilityScheduleInstanceId, error) {
	parser := resourceids.NewParserFromResourceIdType(&RoleManagementDirectoryRoleEligibilityScheduleInstanceId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := RoleManagementDirectoryRoleEligibilityScheduleInstanceId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *RoleManagementDirectoryRoleEligibilityScheduleInstanceId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.UnifiedRoleEligibilityScheduleInstanceId, ok = input.Parsed["unifiedRoleEligibilityScheduleInstanceId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "unifiedRoleEligibilityScheduleInstanceId", input)
	}

	return nil
}

// ValidateRoleManagementDirectoryRoleEligibilityScheduleInstanceID checks that 'input' can be parsed as a Role Management Directory Role Eligibility Schedule Instance ID
func ValidateRoleManagementDirectoryRoleEligibilityScheduleInstanceID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseRoleManagementDirectoryRoleEligibilityScheduleInstanceID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Role Management Directory Role Eligibility Schedule Instance ID
func (id RoleManagementDirectoryRoleEligibilityScheduleInstanceId) ID() string {
	fmtString := "/roleManagement/directory/roleEligibilityScheduleInstances/%s"
	return fmt.Sprintf(fmtString, id.UnifiedRoleEligibilityScheduleInstanceId)
}

// Segments returns a slice of Resource ID Segments which comprise this Role Management Directory Role Eligibility Schedule Instance ID
func (id RoleManagementDirectoryRoleEligibilityScheduleInstanceId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("roleManagement", "roleManagement", "roleManagement"),
		resourceids.StaticSegment("directory", "directory", "directory"),
		resourceids.StaticSegment("roleEligibilityScheduleInstances", "roleEligibilityScheduleInstances", "roleEligibilityScheduleInstances"),
		resourceids.UserSpecifiedSegment("unifiedRoleEligibilityScheduleInstanceId", "unifiedRoleEligibilityScheduleInstanceId"),
	}
}

// String returns a human-readable description of this Role Management Directory Role Eligibility Schedule Instance ID
func (id RoleManagementDirectoryRoleEligibilityScheduleInstanceId) String() string {
	components := []string{
		fmt.Sprintf("Unified Role Eligibility Schedule Instance: %q", id.UnifiedRoleEligibilityScheduleInstanceId),
	}
	return fmt.Sprintf("Role Management Directory Role Eligibility Schedule Instance (%s)", strings.Join(components, "\n"))
}
