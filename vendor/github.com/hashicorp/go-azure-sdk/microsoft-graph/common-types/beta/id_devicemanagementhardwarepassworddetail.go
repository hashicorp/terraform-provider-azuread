package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementHardwarePasswordDetailId{}

// DeviceManagementHardwarePasswordDetailId is a struct representing the Resource ID for a Device Management Hardware Password Detail
type DeviceManagementHardwarePasswordDetailId struct {
	HardwarePasswordDetailId string
}

// NewDeviceManagementHardwarePasswordDetailID returns a new DeviceManagementHardwarePasswordDetailId struct
func NewDeviceManagementHardwarePasswordDetailID(hardwarePasswordDetailId string) DeviceManagementHardwarePasswordDetailId {
	return DeviceManagementHardwarePasswordDetailId{
		HardwarePasswordDetailId: hardwarePasswordDetailId,
	}
}

// ParseDeviceManagementHardwarePasswordDetailID parses 'input' into a DeviceManagementHardwarePasswordDetailId
func ParseDeviceManagementHardwarePasswordDetailID(input string) (*DeviceManagementHardwarePasswordDetailId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementHardwarePasswordDetailId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementHardwarePasswordDetailId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseDeviceManagementHardwarePasswordDetailIDInsensitively parses 'input' case-insensitively into a DeviceManagementHardwarePasswordDetailId
// note: this method should only be used for API response data and not user input
func ParseDeviceManagementHardwarePasswordDetailIDInsensitively(input string) (*DeviceManagementHardwarePasswordDetailId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementHardwarePasswordDetailId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementHardwarePasswordDetailId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *DeviceManagementHardwarePasswordDetailId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.HardwarePasswordDetailId, ok = input.Parsed["hardwarePasswordDetailId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "hardwarePasswordDetailId", input)
	}

	return nil
}

// ValidateDeviceManagementHardwarePasswordDetailID checks that 'input' can be parsed as a Device Management Hardware Password Detail ID
func ValidateDeviceManagementHardwarePasswordDetailID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDeviceManagementHardwarePasswordDetailID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Device Management Hardware Password Detail ID
func (id DeviceManagementHardwarePasswordDetailId) ID() string {
	fmtString := "/deviceManagement/hardwarePasswordDetails/%s"
	return fmt.Sprintf(fmtString, id.HardwarePasswordDetailId)
}

// Segments returns a slice of Resource ID Segments which comprise this Device Management Hardware Password Detail ID
func (id DeviceManagementHardwarePasswordDetailId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("deviceManagement", "deviceManagement", "deviceManagement"),
		resourceids.StaticSegment("hardwarePasswordDetails", "hardwarePasswordDetails", "hardwarePasswordDetails"),
		resourceids.UserSpecifiedSegment("hardwarePasswordDetailId", "hardwarePasswordDetailId"),
	}
}

// String returns a human-readable description of this Device Management Hardware Password Detail ID
func (id DeviceManagementHardwarePasswordDetailId) String() string {
	components := []string{
		fmt.Sprintf("Hardware Password Detail: %q", id.HardwarePasswordDetailId),
	}
	return fmt.Sprintf("Device Management Hardware Password Detail (%s)", strings.Join(components, "\n"))
}
