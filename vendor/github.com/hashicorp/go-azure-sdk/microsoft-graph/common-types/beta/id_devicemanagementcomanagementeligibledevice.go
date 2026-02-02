package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementComanagementEligibleDeviceId{}

// DeviceManagementComanagementEligibleDeviceId is a struct representing the Resource ID for a Device Management Comanagement Eligible Device
type DeviceManagementComanagementEligibleDeviceId struct {
	ComanagementEligibleDeviceId string
}

// NewDeviceManagementComanagementEligibleDeviceID returns a new DeviceManagementComanagementEligibleDeviceId struct
func NewDeviceManagementComanagementEligibleDeviceID(comanagementEligibleDeviceId string) DeviceManagementComanagementEligibleDeviceId {
	return DeviceManagementComanagementEligibleDeviceId{
		ComanagementEligibleDeviceId: comanagementEligibleDeviceId,
	}
}

// ParseDeviceManagementComanagementEligibleDeviceID parses 'input' into a DeviceManagementComanagementEligibleDeviceId
func ParseDeviceManagementComanagementEligibleDeviceID(input string) (*DeviceManagementComanagementEligibleDeviceId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementComanagementEligibleDeviceId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementComanagementEligibleDeviceId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseDeviceManagementComanagementEligibleDeviceIDInsensitively parses 'input' case-insensitively into a DeviceManagementComanagementEligibleDeviceId
// note: this method should only be used for API response data and not user input
func ParseDeviceManagementComanagementEligibleDeviceIDInsensitively(input string) (*DeviceManagementComanagementEligibleDeviceId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementComanagementEligibleDeviceId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementComanagementEligibleDeviceId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *DeviceManagementComanagementEligibleDeviceId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.ComanagementEligibleDeviceId, ok = input.Parsed["comanagementEligibleDeviceId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "comanagementEligibleDeviceId", input)
	}

	return nil
}

// ValidateDeviceManagementComanagementEligibleDeviceID checks that 'input' can be parsed as a Device Management Comanagement Eligible Device ID
func ValidateDeviceManagementComanagementEligibleDeviceID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDeviceManagementComanagementEligibleDeviceID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Device Management Comanagement Eligible Device ID
func (id DeviceManagementComanagementEligibleDeviceId) ID() string {
	fmtString := "/deviceManagement/comanagementEligibleDevices/%s"
	return fmt.Sprintf(fmtString, id.ComanagementEligibleDeviceId)
}

// Segments returns a slice of Resource ID Segments which comprise this Device Management Comanagement Eligible Device ID
func (id DeviceManagementComanagementEligibleDeviceId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("deviceManagement", "deviceManagement", "deviceManagement"),
		resourceids.StaticSegment("comanagementEligibleDevices", "comanagementEligibleDevices", "comanagementEligibleDevices"),
		resourceids.UserSpecifiedSegment("comanagementEligibleDeviceId", "comanagementEligibleDeviceId"),
	}
}

// String returns a human-readable description of this Device Management Comanagement Eligible Device ID
func (id DeviceManagementComanagementEligibleDeviceId) String() string {
	components := []string{
		fmt.Sprintf("Comanagement Eligible Device: %q", id.ComanagementEligibleDeviceId),
	}
	return fmt.Sprintf("Device Management Comanagement Eligible Device (%s)", strings.Join(components, "\n"))
}
