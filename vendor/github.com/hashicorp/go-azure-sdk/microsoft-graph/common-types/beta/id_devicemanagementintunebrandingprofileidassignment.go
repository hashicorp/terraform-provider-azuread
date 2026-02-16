package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementIntuneBrandingProfileIdAssignmentId{}

// DeviceManagementIntuneBrandingProfileIdAssignmentId is a struct representing the Resource ID for a Device Management Intune Branding Profile Id Assignment
type DeviceManagementIntuneBrandingProfileIdAssignmentId struct {
	IntuneBrandingProfileId           string
	IntuneBrandingProfileAssignmentId string
}

// NewDeviceManagementIntuneBrandingProfileIdAssignmentID returns a new DeviceManagementIntuneBrandingProfileIdAssignmentId struct
func NewDeviceManagementIntuneBrandingProfileIdAssignmentID(intuneBrandingProfileId string, intuneBrandingProfileAssignmentId string) DeviceManagementIntuneBrandingProfileIdAssignmentId {
	return DeviceManagementIntuneBrandingProfileIdAssignmentId{
		IntuneBrandingProfileId:           intuneBrandingProfileId,
		IntuneBrandingProfileAssignmentId: intuneBrandingProfileAssignmentId,
	}
}

// ParseDeviceManagementIntuneBrandingProfileIdAssignmentID parses 'input' into a DeviceManagementIntuneBrandingProfileIdAssignmentId
func ParseDeviceManagementIntuneBrandingProfileIdAssignmentID(input string) (*DeviceManagementIntuneBrandingProfileIdAssignmentId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementIntuneBrandingProfileIdAssignmentId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementIntuneBrandingProfileIdAssignmentId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseDeviceManagementIntuneBrandingProfileIdAssignmentIDInsensitively parses 'input' case-insensitively into a DeviceManagementIntuneBrandingProfileIdAssignmentId
// note: this method should only be used for API response data and not user input
func ParseDeviceManagementIntuneBrandingProfileIdAssignmentIDInsensitively(input string) (*DeviceManagementIntuneBrandingProfileIdAssignmentId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementIntuneBrandingProfileIdAssignmentId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementIntuneBrandingProfileIdAssignmentId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *DeviceManagementIntuneBrandingProfileIdAssignmentId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.IntuneBrandingProfileId, ok = input.Parsed["intuneBrandingProfileId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "intuneBrandingProfileId", input)
	}

	if id.IntuneBrandingProfileAssignmentId, ok = input.Parsed["intuneBrandingProfileAssignmentId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "intuneBrandingProfileAssignmentId", input)
	}

	return nil
}

// ValidateDeviceManagementIntuneBrandingProfileIdAssignmentID checks that 'input' can be parsed as a Device Management Intune Branding Profile Id Assignment ID
func ValidateDeviceManagementIntuneBrandingProfileIdAssignmentID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDeviceManagementIntuneBrandingProfileIdAssignmentID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Device Management Intune Branding Profile Id Assignment ID
func (id DeviceManagementIntuneBrandingProfileIdAssignmentId) ID() string {
	fmtString := "/deviceManagement/intuneBrandingProfiles/%s/assignments/%s"
	return fmt.Sprintf(fmtString, id.IntuneBrandingProfileId, id.IntuneBrandingProfileAssignmentId)
}

// Segments returns a slice of Resource ID Segments which comprise this Device Management Intune Branding Profile Id Assignment ID
func (id DeviceManagementIntuneBrandingProfileIdAssignmentId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("deviceManagement", "deviceManagement", "deviceManagement"),
		resourceids.StaticSegment("intuneBrandingProfiles", "intuneBrandingProfiles", "intuneBrandingProfiles"),
		resourceids.UserSpecifiedSegment("intuneBrandingProfileId", "intuneBrandingProfileId"),
		resourceids.StaticSegment("assignments", "assignments", "assignments"),
		resourceids.UserSpecifiedSegment("intuneBrandingProfileAssignmentId", "intuneBrandingProfileAssignmentId"),
	}
}

// String returns a human-readable description of this Device Management Intune Branding Profile Id Assignment ID
func (id DeviceManagementIntuneBrandingProfileIdAssignmentId) String() string {
	components := []string{
		fmt.Sprintf("Intune Branding Profile: %q", id.IntuneBrandingProfileId),
		fmt.Sprintf("Intune Branding Profile Assignment: %q", id.IntuneBrandingProfileAssignmentId),
	}
	return fmt.Sprintf("Device Management Intune Branding Profile Id Assignment (%s)", strings.Join(components, "\n"))
}
