package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementMobileAppTroubleshootingEventId{}

// DeviceManagementMobileAppTroubleshootingEventId is a struct representing the Resource ID for a Device Management Mobile App Troubleshooting Event
type DeviceManagementMobileAppTroubleshootingEventId struct {
	MobileAppTroubleshootingEventId string
}

// NewDeviceManagementMobileAppTroubleshootingEventID returns a new DeviceManagementMobileAppTroubleshootingEventId struct
func NewDeviceManagementMobileAppTroubleshootingEventID(mobileAppTroubleshootingEventId string) DeviceManagementMobileAppTroubleshootingEventId {
	return DeviceManagementMobileAppTroubleshootingEventId{
		MobileAppTroubleshootingEventId: mobileAppTroubleshootingEventId,
	}
}

// ParseDeviceManagementMobileAppTroubleshootingEventID parses 'input' into a DeviceManagementMobileAppTroubleshootingEventId
func ParseDeviceManagementMobileAppTroubleshootingEventID(input string) (*DeviceManagementMobileAppTroubleshootingEventId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementMobileAppTroubleshootingEventId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementMobileAppTroubleshootingEventId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseDeviceManagementMobileAppTroubleshootingEventIDInsensitively parses 'input' case-insensitively into a DeviceManagementMobileAppTroubleshootingEventId
// note: this method should only be used for API response data and not user input
func ParseDeviceManagementMobileAppTroubleshootingEventIDInsensitively(input string) (*DeviceManagementMobileAppTroubleshootingEventId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementMobileAppTroubleshootingEventId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementMobileAppTroubleshootingEventId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *DeviceManagementMobileAppTroubleshootingEventId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.MobileAppTroubleshootingEventId, ok = input.Parsed["mobileAppTroubleshootingEventId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "mobileAppTroubleshootingEventId", input)
	}

	return nil
}

// ValidateDeviceManagementMobileAppTroubleshootingEventID checks that 'input' can be parsed as a Device Management Mobile App Troubleshooting Event ID
func ValidateDeviceManagementMobileAppTroubleshootingEventID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDeviceManagementMobileAppTroubleshootingEventID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Device Management Mobile App Troubleshooting Event ID
func (id DeviceManagementMobileAppTroubleshootingEventId) ID() string {
	fmtString := "/deviceManagement/mobileAppTroubleshootingEvents/%s"
	return fmt.Sprintf(fmtString, id.MobileAppTroubleshootingEventId)
}

// Segments returns a slice of Resource ID Segments which comprise this Device Management Mobile App Troubleshooting Event ID
func (id DeviceManagementMobileAppTroubleshootingEventId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("deviceManagement", "deviceManagement", "deviceManagement"),
		resourceids.StaticSegment("mobileAppTroubleshootingEvents", "mobileAppTroubleshootingEvents", "mobileAppTroubleshootingEvents"),
		resourceids.UserSpecifiedSegment("mobileAppTroubleshootingEventId", "mobileAppTroubleshootingEventId"),
	}
}

// String returns a human-readable description of this Device Management Mobile App Troubleshooting Event ID
func (id DeviceManagementMobileAppTroubleshootingEventId) String() string {
	components := []string{
		fmt.Sprintf("Mobile App Troubleshooting Event: %q", id.MobileAppTroubleshootingEventId),
	}
	return fmt.Sprintf("Device Management Mobile App Troubleshooting Event (%s)", strings.Join(components, "\n"))
}
