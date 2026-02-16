package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementGroupPolicyConfigurationIdAssignmentId{}

// DeviceManagementGroupPolicyConfigurationIdAssignmentId is a struct representing the Resource ID for a Device Management Group Policy Configuration Id Assignment
type DeviceManagementGroupPolicyConfigurationIdAssignmentId struct {
	GroupPolicyConfigurationId           string
	GroupPolicyConfigurationAssignmentId string
}

// NewDeviceManagementGroupPolicyConfigurationIdAssignmentID returns a new DeviceManagementGroupPolicyConfigurationIdAssignmentId struct
func NewDeviceManagementGroupPolicyConfigurationIdAssignmentID(groupPolicyConfigurationId string, groupPolicyConfigurationAssignmentId string) DeviceManagementGroupPolicyConfigurationIdAssignmentId {
	return DeviceManagementGroupPolicyConfigurationIdAssignmentId{
		GroupPolicyConfigurationId:           groupPolicyConfigurationId,
		GroupPolicyConfigurationAssignmentId: groupPolicyConfigurationAssignmentId,
	}
}

// ParseDeviceManagementGroupPolicyConfigurationIdAssignmentID parses 'input' into a DeviceManagementGroupPolicyConfigurationIdAssignmentId
func ParseDeviceManagementGroupPolicyConfigurationIdAssignmentID(input string) (*DeviceManagementGroupPolicyConfigurationIdAssignmentId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementGroupPolicyConfigurationIdAssignmentId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementGroupPolicyConfigurationIdAssignmentId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseDeviceManagementGroupPolicyConfigurationIdAssignmentIDInsensitively parses 'input' case-insensitively into a DeviceManagementGroupPolicyConfigurationIdAssignmentId
// note: this method should only be used for API response data and not user input
func ParseDeviceManagementGroupPolicyConfigurationIdAssignmentIDInsensitively(input string) (*DeviceManagementGroupPolicyConfigurationIdAssignmentId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementGroupPolicyConfigurationIdAssignmentId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementGroupPolicyConfigurationIdAssignmentId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *DeviceManagementGroupPolicyConfigurationIdAssignmentId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.GroupPolicyConfigurationId, ok = input.Parsed["groupPolicyConfigurationId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "groupPolicyConfigurationId", input)
	}

	if id.GroupPolicyConfigurationAssignmentId, ok = input.Parsed["groupPolicyConfigurationAssignmentId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "groupPolicyConfigurationAssignmentId", input)
	}

	return nil
}

// ValidateDeviceManagementGroupPolicyConfigurationIdAssignmentID checks that 'input' can be parsed as a Device Management Group Policy Configuration Id Assignment ID
func ValidateDeviceManagementGroupPolicyConfigurationIdAssignmentID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDeviceManagementGroupPolicyConfigurationIdAssignmentID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Device Management Group Policy Configuration Id Assignment ID
func (id DeviceManagementGroupPolicyConfigurationIdAssignmentId) ID() string {
	fmtString := "/deviceManagement/groupPolicyConfigurations/%s/assignments/%s"
	return fmt.Sprintf(fmtString, id.GroupPolicyConfigurationId, id.GroupPolicyConfigurationAssignmentId)
}

// Segments returns a slice of Resource ID Segments which comprise this Device Management Group Policy Configuration Id Assignment ID
func (id DeviceManagementGroupPolicyConfigurationIdAssignmentId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("deviceManagement", "deviceManagement", "deviceManagement"),
		resourceids.StaticSegment("groupPolicyConfigurations", "groupPolicyConfigurations", "groupPolicyConfigurations"),
		resourceids.UserSpecifiedSegment("groupPolicyConfigurationId", "groupPolicyConfigurationId"),
		resourceids.StaticSegment("assignments", "assignments", "assignments"),
		resourceids.UserSpecifiedSegment("groupPolicyConfigurationAssignmentId", "groupPolicyConfigurationAssignmentId"),
	}
}

// String returns a human-readable description of this Device Management Group Policy Configuration Id Assignment ID
func (id DeviceManagementGroupPolicyConfigurationIdAssignmentId) String() string {
	components := []string{
		fmt.Sprintf("Group Policy Configuration: %q", id.GroupPolicyConfigurationId),
		fmt.Sprintf("Group Policy Configuration Assignment: %q", id.GroupPolicyConfigurationAssignmentId),
	}
	return fmt.Sprintf("Device Management Group Policy Configuration Id Assignment (%s)", strings.Join(components, "\n"))
}
