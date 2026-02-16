package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementHardwarePasswordInfoId{}

// DeviceManagementHardwarePasswordInfoId is a struct representing the Resource ID for a Device Management Hardware Password Info
type DeviceManagementHardwarePasswordInfoId struct {
	HardwarePasswordInfoId string
}

// NewDeviceManagementHardwarePasswordInfoID returns a new DeviceManagementHardwarePasswordInfoId struct
func NewDeviceManagementHardwarePasswordInfoID(hardwarePasswordInfoId string) DeviceManagementHardwarePasswordInfoId {
	return DeviceManagementHardwarePasswordInfoId{
		HardwarePasswordInfoId: hardwarePasswordInfoId,
	}
}

// ParseDeviceManagementHardwarePasswordInfoID parses 'input' into a DeviceManagementHardwarePasswordInfoId
func ParseDeviceManagementHardwarePasswordInfoID(input string) (*DeviceManagementHardwarePasswordInfoId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementHardwarePasswordInfoId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementHardwarePasswordInfoId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseDeviceManagementHardwarePasswordInfoIDInsensitively parses 'input' case-insensitively into a DeviceManagementHardwarePasswordInfoId
// note: this method should only be used for API response data and not user input
func ParseDeviceManagementHardwarePasswordInfoIDInsensitively(input string) (*DeviceManagementHardwarePasswordInfoId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementHardwarePasswordInfoId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementHardwarePasswordInfoId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *DeviceManagementHardwarePasswordInfoId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.HardwarePasswordInfoId, ok = input.Parsed["hardwarePasswordInfoId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "hardwarePasswordInfoId", input)
	}

	return nil
}

// ValidateDeviceManagementHardwarePasswordInfoID checks that 'input' can be parsed as a Device Management Hardware Password Info ID
func ValidateDeviceManagementHardwarePasswordInfoID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDeviceManagementHardwarePasswordInfoID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Device Management Hardware Password Info ID
func (id DeviceManagementHardwarePasswordInfoId) ID() string {
	fmtString := "/deviceManagement/hardwarePasswordInfo/%s"
	return fmt.Sprintf(fmtString, id.HardwarePasswordInfoId)
}

// Segments returns a slice of Resource ID Segments which comprise this Device Management Hardware Password Info ID
func (id DeviceManagementHardwarePasswordInfoId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("deviceManagement", "deviceManagement", "deviceManagement"),
		resourceids.StaticSegment("hardwarePasswordInfo", "hardwarePasswordInfo", "hardwarePasswordInfo"),
		resourceids.UserSpecifiedSegment("hardwarePasswordInfoId", "hardwarePasswordInfoId"),
	}
}

// String returns a human-readable description of this Device Management Hardware Password Info ID
func (id DeviceManagementHardwarePasswordInfoId) String() string {
	components := []string{
		fmt.Sprintf("Hardware Password Info: %q", id.HardwarePasswordInfoId),
	}
	return fmt.Sprintf("Device Management Hardware Password Info (%s)", strings.Join(components, "\n"))
}
