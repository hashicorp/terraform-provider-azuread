package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementIosUpdateStatusId{}

// DeviceManagementIosUpdateStatusId is a struct representing the Resource ID for a Device Management Ios Update Status
type DeviceManagementIosUpdateStatusId struct {
	IosUpdateDeviceStatusId string
}

// NewDeviceManagementIosUpdateStatusID returns a new DeviceManagementIosUpdateStatusId struct
func NewDeviceManagementIosUpdateStatusID(iosUpdateDeviceStatusId string) DeviceManagementIosUpdateStatusId {
	return DeviceManagementIosUpdateStatusId{
		IosUpdateDeviceStatusId: iosUpdateDeviceStatusId,
	}
}

// ParseDeviceManagementIosUpdateStatusID parses 'input' into a DeviceManagementIosUpdateStatusId
func ParseDeviceManagementIosUpdateStatusID(input string) (*DeviceManagementIosUpdateStatusId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementIosUpdateStatusId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementIosUpdateStatusId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseDeviceManagementIosUpdateStatusIDInsensitively parses 'input' case-insensitively into a DeviceManagementIosUpdateStatusId
// note: this method should only be used for API response data and not user input
func ParseDeviceManagementIosUpdateStatusIDInsensitively(input string) (*DeviceManagementIosUpdateStatusId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementIosUpdateStatusId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementIosUpdateStatusId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *DeviceManagementIosUpdateStatusId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.IosUpdateDeviceStatusId, ok = input.Parsed["iosUpdateDeviceStatusId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "iosUpdateDeviceStatusId", input)
	}

	return nil
}

// ValidateDeviceManagementIosUpdateStatusID checks that 'input' can be parsed as a Device Management Ios Update Status ID
func ValidateDeviceManagementIosUpdateStatusID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDeviceManagementIosUpdateStatusID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Device Management Ios Update Status ID
func (id DeviceManagementIosUpdateStatusId) ID() string {
	fmtString := "/deviceManagement/iosUpdateStatuses/%s"
	return fmt.Sprintf(fmtString, id.IosUpdateDeviceStatusId)
}

// Segments returns a slice of Resource ID Segments which comprise this Device Management Ios Update Status ID
func (id DeviceManagementIosUpdateStatusId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("deviceManagement", "deviceManagement", "deviceManagement"),
		resourceids.StaticSegment("iosUpdateStatuses", "iosUpdateStatuses", "iosUpdateStatuses"),
		resourceids.UserSpecifiedSegment("iosUpdateDeviceStatusId", "iosUpdateDeviceStatusId"),
	}
}

// String returns a human-readable description of this Device Management Ios Update Status ID
func (id DeviceManagementIosUpdateStatusId) String() string {
	components := []string{
		fmt.Sprintf("Ios Update Device Status: %q", id.IosUpdateDeviceStatusId),
	}
	return fmt.Sprintf("Device Management Ios Update Status (%s)", strings.Join(components, "\n"))
}
