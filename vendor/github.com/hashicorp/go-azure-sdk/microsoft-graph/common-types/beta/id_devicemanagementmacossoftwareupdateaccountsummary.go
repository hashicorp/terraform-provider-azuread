package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementMacOSSoftwareUpdateAccountSummaryId{}

// DeviceManagementMacOSSoftwareUpdateAccountSummaryId is a struct representing the Resource ID for a Device Management Mac OS Software Update Account Summary
type DeviceManagementMacOSSoftwareUpdateAccountSummaryId struct {
	MacOSSoftwareUpdateAccountSummaryId string
}

// NewDeviceManagementMacOSSoftwareUpdateAccountSummaryID returns a new DeviceManagementMacOSSoftwareUpdateAccountSummaryId struct
func NewDeviceManagementMacOSSoftwareUpdateAccountSummaryID(macOSSoftwareUpdateAccountSummaryId string) DeviceManagementMacOSSoftwareUpdateAccountSummaryId {
	return DeviceManagementMacOSSoftwareUpdateAccountSummaryId{
		MacOSSoftwareUpdateAccountSummaryId: macOSSoftwareUpdateAccountSummaryId,
	}
}

// ParseDeviceManagementMacOSSoftwareUpdateAccountSummaryID parses 'input' into a DeviceManagementMacOSSoftwareUpdateAccountSummaryId
func ParseDeviceManagementMacOSSoftwareUpdateAccountSummaryID(input string) (*DeviceManagementMacOSSoftwareUpdateAccountSummaryId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementMacOSSoftwareUpdateAccountSummaryId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementMacOSSoftwareUpdateAccountSummaryId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseDeviceManagementMacOSSoftwareUpdateAccountSummaryIDInsensitively parses 'input' case-insensitively into a DeviceManagementMacOSSoftwareUpdateAccountSummaryId
// note: this method should only be used for API response data and not user input
func ParseDeviceManagementMacOSSoftwareUpdateAccountSummaryIDInsensitively(input string) (*DeviceManagementMacOSSoftwareUpdateAccountSummaryId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementMacOSSoftwareUpdateAccountSummaryId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementMacOSSoftwareUpdateAccountSummaryId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *DeviceManagementMacOSSoftwareUpdateAccountSummaryId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.MacOSSoftwareUpdateAccountSummaryId, ok = input.Parsed["macOSSoftwareUpdateAccountSummaryId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "macOSSoftwareUpdateAccountSummaryId", input)
	}

	return nil
}

// ValidateDeviceManagementMacOSSoftwareUpdateAccountSummaryID checks that 'input' can be parsed as a Device Management Mac OS Software Update Account Summary ID
func ValidateDeviceManagementMacOSSoftwareUpdateAccountSummaryID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDeviceManagementMacOSSoftwareUpdateAccountSummaryID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Device Management Mac OS Software Update Account Summary ID
func (id DeviceManagementMacOSSoftwareUpdateAccountSummaryId) ID() string {
	fmtString := "/deviceManagement/macOSSoftwareUpdateAccountSummaries/%s"
	return fmt.Sprintf(fmtString, id.MacOSSoftwareUpdateAccountSummaryId)
}

// Segments returns a slice of Resource ID Segments which comprise this Device Management Mac OS Software Update Account Summary ID
func (id DeviceManagementMacOSSoftwareUpdateAccountSummaryId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("deviceManagement", "deviceManagement", "deviceManagement"),
		resourceids.StaticSegment("macOSSoftwareUpdateAccountSummaries", "macOSSoftwareUpdateAccountSummaries", "macOSSoftwareUpdateAccountSummaries"),
		resourceids.UserSpecifiedSegment("macOSSoftwareUpdateAccountSummaryId", "macOSSoftwareUpdateAccountSummaryId"),
	}
}

// String returns a human-readable description of this Device Management Mac OS Software Update Account Summary ID
func (id DeviceManagementMacOSSoftwareUpdateAccountSummaryId) String() string {
	components := []string{
		fmt.Sprintf("Mac OS Software Update Account Summary: %q", id.MacOSSoftwareUpdateAccountSummaryId),
	}
	return fmt.Sprintf("Device Management Mac OS Software Update Account Summary (%s)", strings.Join(components, "\n"))
}
