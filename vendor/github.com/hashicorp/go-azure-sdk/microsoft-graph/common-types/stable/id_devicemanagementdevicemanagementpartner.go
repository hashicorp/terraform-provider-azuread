package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementDeviceManagementPartnerId{}

// DeviceManagementDeviceManagementPartnerId is a struct representing the Resource ID for a Device Management Device Management Partner
type DeviceManagementDeviceManagementPartnerId struct {
	DeviceManagementPartnerId string
}

// NewDeviceManagementDeviceManagementPartnerID returns a new DeviceManagementDeviceManagementPartnerId struct
func NewDeviceManagementDeviceManagementPartnerID(deviceManagementPartnerId string) DeviceManagementDeviceManagementPartnerId {
	return DeviceManagementDeviceManagementPartnerId{
		DeviceManagementPartnerId: deviceManagementPartnerId,
	}
}

// ParseDeviceManagementDeviceManagementPartnerID parses 'input' into a DeviceManagementDeviceManagementPartnerId
func ParseDeviceManagementDeviceManagementPartnerID(input string) (*DeviceManagementDeviceManagementPartnerId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementDeviceManagementPartnerId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementDeviceManagementPartnerId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseDeviceManagementDeviceManagementPartnerIDInsensitively parses 'input' case-insensitively into a DeviceManagementDeviceManagementPartnerId
// note: this method should only be used for API response data and not user input
func ParseDeviceManagementDeviceManagementPartnerIDInsensitively(input string) (*DeviceManagementDeviceManagementPartnerId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementDeviceManagementPartnerId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementDeviceManagementPartnerId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *DeviceManagementDeviceManagementPartnerId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.DeviceManagementPartnerId, ok = input.Parsed["deviceManagementPartnerId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "deviceManagementPartnerId", input)
	}

	return nil
}

// ValidateDeviceManagementDeviceManagementPartnerID checks that 'input' can be parsed as a Device Management Device Management Partner ID
func ValidateDeviceManagementDeviceManagementPartnerID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDeviceManagementDeviceManagementPartnerID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Device Management Device Management Partner ID
func (id DeviceManagementDeviceManagementPartnerId) ID() string {
	fmtString := "/deviceManagement/deviceManagementPartners/%s"
	return fmt.Sprintf(fmtString, id.DeviceManagementPartnerId)
}

// Segments returns a slice of Resource ID Segments which comprise this Device Management Device Management Partner ID
func (id DeviceManagementDeviceManagementPartnerId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("deviceManagement", "deviceManagement", "deviceManagement"),
		resourceids.StaticSegment("deviceManagementPartners", "deviceManagementPartners", "deviceManagementPartners"),
		resourceids.UserSpecifiedSegment("deviceManagementPartnerId", "deviceManagementPartnerId"),
	}
}

// String returns a human-readable description of this Device Management Device Management Partner ID
func (id DeviceManagementDeviceManagementPartnerId) String() string {
	components := []string{
		fmt.Sprintf("Device Management Partner: %q", id.DeviceManagementPartnerId),
	}
	return fmt.Sprintf("Device Management Device Management Partner (%s)", strings.Join(components, "\n"))
}
