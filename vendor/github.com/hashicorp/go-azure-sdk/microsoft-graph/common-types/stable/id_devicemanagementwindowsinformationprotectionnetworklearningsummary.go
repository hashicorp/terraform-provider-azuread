package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementWindowsInformationProtectionNetworkLearningSummaryId{}

// DeviceManagementWindowsInformationProtectionNetworkLearningSummaryId is a struct representing the Resource ID for a Device Management Windows Information Protection Network Learning Summary
type DeviceManagementWindowsInformationProtectionNetworkLearningSummaryId struct {
	WindowsInformationProtectionNetworkLearningSummaryId string
}

// NewDeviceManagementWindowsInformationProtectionNetworkLearningSummaryID returns a new DeviceManagementWindowsInformationProtectionNetworkLearningSummaryId struct
func NewDeviceManagementWindowsInformationProtectionNetworkLearningSummaryID(windowsInformationProtectionNetworkLearningSummaryId string) DeviceManagementWindowsInformationProtectionNetworkLearningSummaryId {
	return DeviceManagementWindowsInformationProtectionNetworkLearningSummaryId{
		WindowsInformationProtectionNetworkLearningSummaryId: windowsInformationProtectionNetworkLearningSummaryId,
	}
}

// ParseDeviceManagementWindowsInformationProtectionNetworkLearningSummaryID parses 'input' into a DeviceManagementWindowsInformationProtectionNetworkLearningSummaryId
func ParseDeviceManagementWindowsInformationProtectionNetworkLearningSummaryID(input string) (*DeviceManagementWindowsInformationProtectionNetworkLearningSummaryId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementWindowsInformationProtectionNetworkLearningSummaryId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementWindowsInformationProtectionNetworkLearningSummaryId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseDeviceManagementWindowsInformationProtectionNetworkLearningSummaryIDInsensitively parses 'input' case-insensitively into a DeviceManagementWindowsInformationProtectionNetworkLearningSummaryId
// note: this method should only be used for API response data and not user input
func ParseDeviceManagementWindowsInformationProtectionNetworkLearningSummaryIDInsensitively(input string) (*DeviceManagementWindowsInformationProtectionNetworkLearningSummaryId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementWindowsInformationProtectionNetworkLearningSummaryId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementWindowsInformationProtectionNetworkLearningSummaryId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *DeviceManagementWindowsInformationProtectionNetworkLearningSummaryId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.WindowsInformationProtectionNetworkLearningSummaryId, ok = input.Parsed["windowsInformationProtectionNetworkLearningSummaryId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "windowsInformationProtectionNetworkLearningSummaryId", input)
	}

	return nil
}

// ValidateDeviceManagementWindowsInformationProtectionNetworkLearningSummaryID checks that 'input' can be parsed as a Device Management Windows Information Protection Network Learning Summary ID
func ValidateDeviceManagementWindowsInformationProtectionNetworkLearningSummaryID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDeviceManagementWindowsInformationProtectionNetworkLearningSummaryID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Device Management Windows Information Protection Network Learning Summary ID
func (id DeviceManagementWindowsInformationProtectionNetworkLearningSummaryId) ID() string {
	fmtString := "/deviceManagement/windowsInformationProtectionNetworkLearningSummaries/%s"
	return fmt.Sprintf(fmtString, id.WindowsInformationProtectionNetworkLearningSummaryId)
}

// Segments returns a slice of Resource ID Segments which comprise this Device Management Windows Information Protection Network Learning Summary ID
func (id DeviceManagementWindowsInformationProtectionNetworkLearningSummaryId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("deviceManagement", "deviceManagement", "deviceManagement"),
		resourceids.StaticSegment("windowsInformationProtectionNetworkLearningSummaries", "windowsInformationProtectionNetworkLearningSummaries", "windowsInformationProtectionNetworkLearningSummaries"),
		resourceids.UserSpecifiedSegment("windowsInformationProtectionNetworkLearningSummaryId", "windowsInformationProtectionNetworkLearningSummaryId"),
	}
}

// String returns a human-readable description of this Device Management Windows Information Protection Network Learning Summary ID
func (id DeviceManagementWindowsInformationProtectionNetworkLearningSummaryId) String() string {
	components := []string{
		fmt.Sprintf("Windows Information Protection Network Learning Summary: %q", id.WindowsInformationProtectionNetworkLearningSummaryId),
	}
	return fmt.Sprintf("Device Management Windows Information Protection Network Learning Summary (%s)", strings.Join(components, "\n"))
}
