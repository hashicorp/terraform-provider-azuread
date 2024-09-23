package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementWindowsInformationProtectionAppLearningSummaryId{}

// DeviceManagementWindowsInformationProtectionAppLearningSummaryId is a struct representing the Resource ID for a Device Management Windows Information Protection App Learning Summary
type DeviceManagementWindowsInformationProtectionAppLearningSummaryId struct {
	WindowsInformationProtectionAppLearningSummaryId string
}

// NewDeviceManagementWindowsInformationProtectionAppLearningSummaryID returns a new DeviceManagementWindowsInformationProtectionAppLearningSummaryId struct
func NewDeviceManagementWindowsInformationProtectionAppLearningSummaryID(windowsInformationProtectionAppLearningSummaryId string) DeviceManagementWindowsInformationProtectionAppLearningSummaryId {
	return DeviceManagementWindowsInformationProtectionAppLearningSummaryId{
		WindowsInformationProtectionAppLearningSummaryId: windowsInformationProtectionAppLearningSummaryId,
	}
}

// ParseDeviceManagementWindowsInformationProtectionAppLearningSummaryID parses 'input' into a DeviceManagementWindowsInformationProtectionAppLearningSummaryId
func ParseDeviceManagementWindowsInformationProtectionAppLearningSummaryID(input string) (*DeviceManagementWindowsInformationProtectionAppLearningSummaryId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementWindowsInformationProtectionAppLearningSummaryId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementWindowsInformationProtectionAppLearningSummaryId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseDeviceManagementWindowsInformationProtectionAppLearningSummaryIDInsensitively parses 'input' case-insensitively into a DeviceManagementWindowsInformationProtectionAppLearningSummaryId
// note: this method should only be used for API response data and not user input
func ParseDeviceManagementWindowsInformationProtectionAppLearningSummaryIDInsensitively(input string) (*DeviceManagementWindowsInformationProtectionAppLearningSummaryId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementWindowsInformationProtectionAppLearningSummaryId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementWindowsInformationProtectionAppLearningSummaryId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *DeviceManagementWindowsInformationProtectionAppLearningSummaryId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.WindowsInformationProtectionAppLearningSummaryId, ok = input.Parsed["windowsInformationProtectionAppLearningSummaryId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "windowsInformationProtectionAppLearningSummaryId", input)
	}

	return nil
}

// ValidateDeviceManagementWindowsInformationProtectionAppLearningSummaryID checks that 'input' can be parsed as a Device Management Windows Information Protection App Learning Summary ID
func ValidateDeviceManagementWindowsInformationProtectionAppLearningSummaryID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDeviceManagementWindowsInformationProtectionAppLearningSummaryID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Device Management Windows Information Protection App Learning Summary ID
func (id DeviceManagementWindowsInformationProtectionAppLearningSummaryId) ID() string {
	fmtString := "/deviceManagement/windowsInformationProtectionAppLearningSummaries/%s"
	return fmt.Sprintf(fmtString, id.WindowsInformationProtectionAppLearningSummaryId)
}

// Segments returns a slice of Resource ID Segments which comprise this Device Management Windows Information Protection App Learning Summary ID
func (id DeviceManagementWindowsInformationProtectionAppLearningSummaryId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("deviceManagement", "deviceManagement", "deviceManagement"),
		resourceids.StaticSegment("windowsInformationProtectionAppLearningSummaries", "windowsInformationProtectionAppLearningSummaries", "windowsInformationProtectionAppLearningSummaries"),
		resourceids.UserSpecifiedSegment("windowsInformationProtectionAppLearningSummaryId", "windowsInformationProtectionAppLearningSummaryId"),
	}
}

// String returns a human-readable description of this Device Management Windows Information Protection App Learning Summary ID
func (id DeviceManagementWindowsInformationProtectionAppLearningSummaryId) String() string {
	components := []string{
		fmt.Sprintf("Windows Information Protection App Learning Summary: %q", id.WindowsInformationProtectionAppLearningSummaryId),
	}
	return fmt.Sprintf("Device Management Windows Information Protection App Learning Summary (%s)", strings.Join(components, "\n"))
}
