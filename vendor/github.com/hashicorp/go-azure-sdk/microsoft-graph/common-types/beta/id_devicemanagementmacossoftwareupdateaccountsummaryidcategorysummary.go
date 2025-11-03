package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementMacOSSoftwareUpdateAccountSummaryIdCategorySummaryId{}

// DeviceManagementMacOSSoftwareUpdateAccountSummaryIdCategorySummaryId is a struct representing the Resource ID for a Device Management Mac OS Software Update Account Summary Id Category Summary
type DeviceManagementMacOSSoftwareUpdateAccountSummaryIdCategorySummaryId struct {
	MacOSSoftwareUpdateAccountSummaryId  string
	MacOSSoftwareUpdateCategorySummaryId string
}

// NewDeviceManagementMacOSSoftwareUpdateAccountSummaryIdCategorySummaryID returns a new DeviceManagementMacOSSoftwareUpdateAccountSummaryIdCategorySummaryId struct
func NewDeviceManagementMacOSSoftwareUpdateAccountSummaryIdCategorySummaryID(macOSSoftwareUpdateAccountSummaryId string, macOSSoftwareUpdateCategorySummaryId string) DeviceManagementMacOSSoftwareUpdateAccountSummaryIdCategorySummaryId {
	return DeviceManagementMacOSSoftwareUpdateAccountSummaryIdCategorySummaryId{
		MacOSSoftwareUpdateAccountSummaryId:  macOSSoftwareUpdateAccountSummaryId,
		MacOSSoftwareUpdateCategorySummaryId: macOSSoftwareUpdateCategorySummaryId,
	}
}

// ParseDeviceManagementMacOSSoftwareUpdateAccountSummaryIdCategorySummaryID parses 'input' into a DeviceManagementMacOSSoftwareUpdateAccountSummaryIdCategorySummaryId
func ParseDeviceManagementMacOSSoftwareUpdateAccountSummaryIdCategorySummaryID(input string) (*DeviceManagementMacOSSoftwareUpdateAccountSummaryIdCategorySummaryId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementMacOSSoftwareUpdateAccountSummaryIdCategorySummaryId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementMacOSSoftwareUpdateAccountSummaryIdCategorySummaryId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseDeviceManagementMacOSSoftwareUpdateAccountSummaryIdCategorySummaryIDInsensitively parses 'input' case-insensitively into a DeviceManagementMacOSSoftwareUpdateAccountSummaryIdCategorySummaryId
// note: this method should only be used for API response data and not user input
func ParseDeviceManagementMacOSSoftwareUpdateAccountSummaryIdCategorySummaryIDInsensitively(input string) (*DeviceManagementMacOSSoftwareUpdateAccountSummaryIdCategorySummaryId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementMacOSSoftwareUpdateAccountSummaryIdCategorySummaryId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementMacOSSoftwareUpdateAccountSummaryIdCategorySummaryId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *DeviceManagementMacOSSoftwareUpdateAccountSummaryIdCategorySummaryId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.MacOSSoftwareUpdateAccountSummaryId, ok = input.Parsed["macOSSoftwareUpdateAccountSummaryId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "macOSSoftwareUpdateAccountSummaryId", input)
	}

	if id.MacOSSoftwareUpdateCategorySummaryId, ok = input.Parsed["macOSSoftwareUpdateCategorySummaryId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "macOSSoftwareUpdateCategorySummaryId", input)
	}

	return nil
}

// ValidateDeviceManagementMacOSSoftwareUpdateAccountSummaryIdCategorySummaryID checks that 'input' can be parsed as a Device Management Mac OS Software Update Account Summary Id Category Summary ID
func ValidateDeviceManagementMacOSSoftwareUpdateAccountSummaryIdCategorySummaryID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDeviceManagementMacOSSoftwareUpdateAccountSummaryIdCategorySummaryID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Device Management Mac OS Software Update Account Summary Id Category Summary ID
func (id DeviceManagementMacOSSoftwareUpdateAccountSummaryIdCategorySummaryId) ID() string {
	fmtString := "/deviceManagement/macOSSoftwareUpdateAccountSummaries/%s/categorySummaries/%s"
	return fmt.Sprintf(fmtString, id.MacOSSoftwareUpdateAccountSummaryId, id.MacOSSoftwareUpdateCategorySummaryId)
}

// Segments returns a slice of Resource ID Segments which comprise this Device Management Mac OS Software Update Account Summary Id Category Summary ID
func (id DeviceManagementMacOSSoftwareUpdateAccountSummaryIdCategorySummaryId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("deviceManagement", "deviceManagement", "deviceManagement"),
		resourceids.StaticSegment("macOSSoftwareUpdateAccountSummaries", "macOSSoftwareUpdateAccountSummaries", "macOSSoftwareUpdateAccountSummaries"),
		resourceids.UserSpecifiedSegment("macOSSoftwareUpdateAccountSummaryId", "macOSSoftwareUpdateAccountSummaryId"),
		resourceids.StaticSegment("categorySummaries", "categorySummaries", "categorySummaries"),
		resourceids.UserSpecifiedSegment("macOSSoftwareUpdateCategorySummaryId", "macOSSoftwareUpdateCategorySummaryId"),
	}
}

// String returns a human-readable description of this Device Management Mac OS Software Update Account Summary Id Category Summary ID
func (id DeviceManagementMacOSSoftwareUpdateAccountSummaryIdCategorySummaryId) String() string {
	components := []string{
		fmt.Sprintf("Mac OS Software Update Account Summary: %q", id.MacOSSoftwareUpdateAccountSummaryId),
		fmt.Sprintf("Mac OS Software Update Category Summary: %q", id.MacOSSoftwareUpdateCategorySummaryId),
	}
	return fmt.Sprintf("Device Management Mac OS Software Update Account Summary Id Category Summary (%s)", strings.Join(components, "\n"))
}
