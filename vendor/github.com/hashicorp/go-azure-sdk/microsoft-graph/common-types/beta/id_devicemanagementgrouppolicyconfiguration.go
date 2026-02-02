package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementGroupPolicyConfigurationId{}

// DeviceManagementGroupPolicyConfigurationId is a struct representing the Resource ID for a Device Management Group Policy Configuration
type DeviceManagementGroupPolicyConfigurationId struct {
	GroupPolicyConfigurationId string
}

// NewDeviceManagementGroupPolicyConfigurationID returns a new DeviceManagementGroupPolicyConfigurationId struct
func NewDeviceManagementGroupPolicyConfigurationID(groupPolicyConfigurationId string) DeviceManagementGroupPolicyConfigurationId {
	return DeviceManagementGroupPolicyConfigurationId{
		GroupPolicyConfigurationId: groupPolicyConfigurationId,
	}
}

// ParseDeviceManagementGroupPolicyConfigurationID parses 'input' into a DeviceManagementGroupPolicyConfigurationId
func ParseDeviceManagementGroupPolicyConfigurationID(input string) (*DeviceManagementGroupPolicyConfigurationId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementGroupPolicyConfigurationId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementGroupPolicyConfigurationId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseDeviceManagementGroupPolicyConfigurationIDInsensitively parses 'input' case-insensitively into a DeviceManagementGroupPolicyConfigurationId
// note: this method should only be used for API response data and not user input
func ParseDeviceManagementGroupPolicyConfigurationIDInsensitively(input string) (*DeviceManagementGroupPolicyConfigurationId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementGroupPolicyConfigurationId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementGroupPolicyConfigurationId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *DeviceManagementGroupPolicyConfigurationId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.GroupPolicyConfigurationId, ok = input.Parsed["groupPolicyConfigurationId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "groupPolicyConfigurationId", input)
	}

	return nil
}

// ValidateDeviceManagementGroupPolicyConfigurationID checks that 'input' can be parsed as a Device Management Group Policy Configuration ID
func ValidateDeviceManagementGroupPolicyConfigurationID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDeviceManagementGroupPolicyConfigurationID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Device Management Group Policy Configuration ID
func (id DeviceManagementGroupPolicyConfigurationId) ID() string {
	fmtString := "/deviceManagement/groupPolicyConfigurations/%s"
	return fmt.Sprintf(fmtString, id.GroupPolicyConfigurationId)
}

// Segments returns a slice of Resource ID Segments which comprise this Device Management Group Policy Configuration ID
func (id DeviceManagementGroupPolicyConfigurationId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("deviceManagement", "deviceManagement", "deviceManagement"),
		resourceids.StaticSegment("groupPolicyConfigurations", "groupPolicyConfigurations", "groupPolicyConfigurations"),
		resourceids.UserSpecifiedSegment("groupPolicyConfigurationId", "groupPolicyConfigurationId"),
	}
}

// String returns a human-readable description of this Device Management Group Policy Configuration ID
func (id DeviceManagementGroupPolicyConfigurationId) String() string {
	components := []string{
		fmt.Sprintf("Group Policy Configuration: %q", id.GroupPolicyConfigurationId),
	}
	return fmt.Sprintf("Device Management Group Policy Configuration (%s)", strings.Join(components, "\n"))
}
