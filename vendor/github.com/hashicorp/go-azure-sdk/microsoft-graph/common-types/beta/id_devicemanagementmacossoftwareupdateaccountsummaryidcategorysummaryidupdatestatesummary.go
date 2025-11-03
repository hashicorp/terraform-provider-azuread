package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementMacOSSoftwareUpdateAccountSummaryIdCategorySummaryIdUpdateStateSummaryId{}

// DeviceManagementMacOSSoftwareUpdateAccountSummaryIdCategorySummaryIdUpdateStateSummaryId is a struct representing the Resource ID for a Device Management Mac OS Software Update Account Summary Id Category Summary Id Update State Summary
type DeviceManagementMacOSSoftwareUpdateAccountSummaryIdCategorySummaryIdUpdateStateSummaryId struct {
	MacOSSoftwareUpdateAccountSummaryId  string
	MacOSSoftwareUpdateCategorySummaryId string
	MacOSSoftwareUpdateStateSummaryId    string
}

// NewDeviceManagementMacOSSoftwareUpdateAccountSummaryIdCategorySummaryIdUpdateStateSummaryID returns a new DeviceManagementMacOSSoftwareUpdateAccountSummaryIdCategorySummaryIdUpdateStateSummaryId struct
func NewDeviceManagementMacOSSoftwareUpdateAccountSummaryIdCategorySummaryIdUpdateStateSummaryID(macOSSoftwareUpdateAccountSummaryId string, macOSSoftwareUpdateCategorySummaryId string, macOSSoftwareUpdateStateSummaryId string) DeviceManagementMacOSSoftwareUpdateAccountSummaryIdCategorySummaryIdUpdateStateSummaryId {
	return DeviceManagementMacOSSoftwareUpdateAccountSummaryIdCategorySummaryIdUpdateStateSummaryId{
		MacOSSoftwareUpdateAccountSummaryId:  macOSSoftwareUpdateAccountSummaryId,
		MacOSSoftwareUpdateCategorySummaryId: macOSSoftwareUpdateCategorySummaryId,
		MacOSSoftwareUpdateStateSummaryId:    macOSSoftwareUpdateStateSummaryId,
	}
}

// ParseDeviceManagementMacOSSoftwareUpdateAccountSummaryIdCategorySummaryIdUpdateStateSummaryID parses 'input' into a DeviceManagementMacOSSoftwareUpdateAccountSummaryIdCategorySummaryIdUpdateStateSummaryId
func ParseDeviceManagementMacOSSoftwareUpdateAccountSummaryIdCategorySummaryIdUpdateStateSummaryID(input string) (*DeviceManagementMacOSSoftwareUpdateAccountSummaryIdCategorySummaryIdUpdateStateSummaryId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementMacOSSoftwareUpdateAccountSummaryIdCategorySummaryIdUpdateStateSummaryId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementMacOSSoftwareUpdateAccountSummaryIdCategorySummaryIdUpdateStateSummaryId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseDeviceManagementMacOSSoftwareUpdateAccountSummaryIdCategorySummaryIdUpdateStateSummaryIDInsensitively parses 'input' case-insensitively into a DeviceManagementMacOSSoftwareUpdateAccountSummaryIdCategorySummaryIdUpdateStateSummaryId
// note: this method should only be used for API response data and not user input
func ParseDeviceManagementMacOSSoftwareUpdateAccountSummaryIdCategorySummaryIdUpdateStateSummaryIDInsensitively(input string) (*DeviceManagementMacOSSoftwareUpdateAccountSummaryIdCategorySummaryIdUpdateStateSummaryId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementMacOSSoftwareUpdateAccountSummaryIdCategorySummaryIdUpdateStateSummaryId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementMacOSSoftwareUpdateAccountSummaryIdCategorySummaryIdUpdateStateSummaryId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *DeviceManagementMacOSSoftwareUpdateAccountSummaryIdCategorySummaryIdUpdateStateSummaryId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.MacOSSoftwareUpdateAccountSummaryId, ok = input.Parsed["macOSSoftwareUpdateAccountSummaryId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "macOSSoftwareUpdateAccountSummaryId", input)
	}

	if id.MacOSSoftwareUpdateCategorySummaryId, ok = input.Parsed["macOSSoftwareUpdateCategorySummaryId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "macOSSoftwareUpdateCategorySummaryId", input)
	}

	if id.MacOSSoftwareUpdateStateSummaryId, ok = input.Parsed["macOSSoftwareUpdateStateSummaryId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "macOSSoftwareUpdateStateSummaryId", input)
	}

	return nil
}

// ValidateDeviceManagementMacOSSoftwareUpdateAccountSummaryIdCategorySummaryIdUpdateStateSummaryID checks that 'input' can be parsed as a Device Management Mac OS Software Update Account Summary Id Category Summary Id Update State Summary ID
func ValidateDeviceManagementMacOSSoftwareUpdateAccountSummaryIdCategorySummaryIdUpdateStateSummaryID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDeviceManagementMacOSSoftwareUpdateAccountSummaryIdCategorySummaryIdUpdateStateSummaryID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Device Management Mac OS Software Update Account Summary Id Category Summary Id Update State Summary ID
func (id DeviceManagementMacOSSoftwareUpdateAccountSummaryIdCategorySummaryIdUpdateStateSummaryId) ID() string {
	fmtString := "/deviceManagement/macOSSoftwareUpdateAccountSummaries/%s/categorySummaries/%s/updateStateSummaries/%s"
	return fmt.Sprintf(fmtString, id.MacOSSoftwareUpdateAccountSummaryId, id.MacOSSoftwareUpdateCategorySummaryId, id.MacOSSoftwareUpdateStateSummaryId)
}

// Segments returns a slice of Resource ID Segments which comprise this Device Management Mac OS Software Update Account Summary Id Category Summary Id Update State Summary ID
func (id DeviceManagementMacOSSoftwareUpdateAccountSummaryIdCategorySummaryIdUpdateStateSummaryId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("deviceManagement", "deviceManagement", "deviceManagement"),
		resourceids.StaticSegment("macOSSoftwareUpdateAccountSummaries", "macOSSoftwareUpdateAccountSummaries", "macOSSoftwareUpdateAccountSummaries"),
		resourceids.UserSpecifiedSegment("macOSSoftwareUpdateAccountSummaryId", "macOSSoftwareUpdateAccountSummaryId"),
		resourceids.StaticSegment("categorySummaries", "categorySummaries", "categorySummaries"),
		resourceids.UserSpecifiedSegment("macOSSoftwareUpdateCategorySummaryId", "macOSSoftwareUpdateCategorySummaryId"),
		resourceids.StaticSegment("updateStateSummaries", "updateStateSummaries", "updateStateSummaries"),
		resourceids.UserSpecifiedSegment("macOSSoftwareUpdateStateSummaryId", "macOSSoftwareUpdateStateSummaryId"),
	}
}

// String returns a human-readable description of this Device Management Mac OS Software Update Account Summary Id Category Summary Id Update State Summary ID
func (id DeviceManagementMacOSSoftwareUpdateAccountSummaryIdCategorySummaryIdUpdateStateSummaryId) String() string {
	components := []string{
		fmt.Sprintf("Mac OS Software Update Account Summary: %q", id.MacOSSoftwareUpdateAccountSummaryId),
		fmt.Sprintf("Mac OS Software Update Category Summary: %q", id.MacOSSoftwareUpdateCategorySummaryId),
		fmt.Sprintf("Mac OS Software Update State Summary: %q", id.MacOSSoftwareUpdateStateSummaryId),
	}
	return fmt.Sprintf("Device Management Mac OS Software Update Account Summary Id Category Summary Id Update State Summary (%s)", strings.Join(components, "\n"))
}
