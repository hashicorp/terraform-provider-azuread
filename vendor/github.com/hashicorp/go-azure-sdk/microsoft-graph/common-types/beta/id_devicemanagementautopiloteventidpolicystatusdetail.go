package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementAutopilotEventIdPolicyStatusDetailId{}

// DeviceManagementAutopilotEventIdPolicyStatusDetailId is a struct representing the Resource ID for a Device Management Autopilot Event Id Policy Status Detail
type DeviceManagementAutopilotEventIdPolicyStatusDetailId struct {
	DeviceManagementAutopilotEventId              string
	DeviceManagementAutopilotPolicyStatusDetailId string
}

// NewDeviceManagementAutopilotEventIdPolicyStatusDetailID returns a new DeviceManagementAutopilotEventIdPolicyStatusDetailId struct
func NewDeviceManagementAutopilotEventIdPolicyStatusDetailID(deviceManagementAutopilotEventId string, deviceManagementAutopilotPolicyStatusDetailId string) DeviceManagementAutopilotEventIdPolicyStatusDetailId {
	return DeviceManagementAutopilotEventIdPolicyStatusDetailId{
		DeviceManagementAutopilotEventId:              deviceManagementAutopilotEventId,
		DeviceManagementAutopilotPolicyStatusDetailId: deviceManagementAutopilotPolicyStatusDetailId,
	}
}

// ParseDeviceManagementAutopilotEventIdPolicyStatusDetailID parses 'input' into a DeviceManagementAutopilotEventIdPolicyStatusDetailId
func ParseDeviceManagementAutopilotEventIdPolicyStatusDetailID(input string) (*DeviceManagementAutopilotEventIdPolicyStatusDetailId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementAutopilotEventIdPolicyStatusDetailId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementAutopilotEventIdPolicyStatusDetailId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseDeviceManagementAutopilotEventIdPolicyStatusDetailIDInsensitively parses 'input' case-insensitively into a DeviceManagementAutopilotEventIdPolicyStatusDetailId
// note: this method should only be used for API response data and not user input
func ParseDeviceManagementAutopilotEventIdPolicyStatusDetailIDInsensitively(input string) (*DeviceManagementAutopilotEventIdPolicyStatusDetailId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementAutopilotEventIdPolicyStatusDetailId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementAutopilotEventIdPolicyStatusDetailId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *DeviceManagementAutopilotEventIdPolicyStatusDetailId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.DeviceManagementAutopilotEventId, ok = input.Parsed["deviceManagementAutopilotEventId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "deviceManagementAutopilotEventId", input)
	}

	if id.DeviceManagementAutopilotPolicyStatusDetailId, ok = input.Parsed["deviceManagementAutopilotPolicyStatusDetailId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "deviceManagementAutopilotPolicyStatusDetailId", input)
	}

	return nil
}

// ValidateDeviceManagementAutopilotEventIdPolicyStatusDetailID checks that 'input' can be parsed as a Device Management Autopilot Event Id Policy Status Detail ID
func ValidateDeviceManagementAutopilotEventIdPolicyStatusDetailID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDeviceManagementAutopilotEventIdPolicyStatusDetailID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Device Management Autopilot Event Id Policy Status Detail ID
func (id DeviceManagementAutopilotEventIdPolicyStatusDetailId) ID() string {
	fmtString := "/deviceManagement/autopilotEvents/%s/policyStatusDetails/%s"
	return fmt.Sprintf(fmtString, id.DeviceManagementAutopilotEventId, id.DeviceManagementAutopilotPolicyStatusDetailId)
}

// Segments returns a slice of Resource ID Segments which comprise this Device Management Autopilot Event Id Policy Status Detail ID
func (id DeviceManagementAutopilotEventIdPolicyStatusDetailId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("deviceManagement", "deviceManagement", "deviceManagement"),
		resourceids.StaticSegment("autopilotEvents", "autopilotEvents", "autopilotEvents"),
		resourceids.UserSpecifiedSegment("deviceManagementAutopilotEventId", "deviceManagementAutopilotEventId"),
		resourceids.StaticSegment("policyStatusDetails", "policyStatusDetails", "policyStatusDetails"),
		resourceids.UserSpecifiedSegment("deviceManagementAutopilotPolicyStatusDetailId", "deviceManagementAutopilotPolicyStatusDetailId"),
	}
}

// String returns a human-readable description of this Device Management Autopilot Event Id Policy Status Detail ID
func (id DeviceManagementAutopilotEventIdPolicyStatusDetailId) String() string {
	components := []string{
		fmt.Sprintf("Device Management Autopilot Event: %q", id.DeviceManagementAutopilotEventId),
		fmt.Sprintf("Device Management Autopilot Policy Status Detail: %q", id.DeviceManagementAutopilotPolicyStatusDetailId),
	}
	return fmt.Sprintf("Device Management Autopilot Event Id Policy Status Detail (%s)", strings.Join(components, "\n"))
}
