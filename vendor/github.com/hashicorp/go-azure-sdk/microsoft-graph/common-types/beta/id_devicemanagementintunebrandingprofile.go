package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementIntuneBrandingProfileId{}

// DeviceManagementIntuneBrandingProfileId is a struct representing the Resource ID for a Device Management Intune Branding Profile
type DeviceManagementIntuneBrandingProfileId struct {
	IntuneBrandingProfileId string
}

// NewDeviceManagementIntuneBrandingProfileID returns a new DeviceManagementIntuneBrandingProfileId struct
func NewDeviceManagementIntuneBrandingProfileID(intuneBrandingProfileId string) DeviceManagementIntuneBrandingProfileId {
	return DeviceManagementIntuneBrandingProfileId{
		IntuneBrandingProfileId: intuneBrandingProfileId,
	}
}

// ParseDeviceManagementIntuneBrandingProfileID parses 'input' into a DeviceManagementIntuneBrandingProfileId
func ParseDeviceManagementIntuneBrandingProfileID(input string) (*DeviceManagementIntuneBrandingProfileId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementIntuneBrandingProfileId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementIntuneBrandingProfileId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseDeviceManagementIntuneBrandingProfileIDInsensitively parses 'input' case-insensitively into a DeviceManagementIntuneBrandingProfileId
// note: this method should only be used for API response data and not user input
func ParseDeviceManagementIntuneBrandingProfileIDInsensitively(input string) (*DeviceManagementIntuneBrandingProfileId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementIntuneBrandingProfileId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementIntuneBrandingProfileId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *DeviceManagementIntuneBrandingProfileId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.IntuneBrandingProfileId, ok = input.Parsed["intuneBrandingProfileId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "intuneBrandingProfileId", input)
	}

	return nil
}

// ValidateDeviceManagementIntuneBrandingProfileID checks that 'input' can be parsed as a Device Management Intune Branding Profile ID
func ValidateDeviceManagementIntuneBrandingProfileID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDeviceManagementIntuneBrandingProfileID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Device Management Intune Branding Profile ID
func (id DeviceManagementIntuneBrandingProfileId) ID() string {
	fmtString := "/deviceManagement/intuneBrandingProfiles/%s"
	return fmt.Sprintf(fmtString, id.IntuneBrandingProfileId)
}

// Segments returns a slice of Resource ID Segments which comprise this Device Management Intune Branding Profile ID
func (id DeviceManagementIntuneBrandingProfileId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("deviceManagement", "deviceManagement", "deviceManagement"),
		resourceids.StaticSegment("intuneBrandingProfiles", "intuneBrandingProfiles", "intuneBrandingProfiles"),
		resourceids.UserSpecifiedSegment("intuneBrandingProfileId", "intuneBrandingProfileId"),
	}
}

// String returns a human-readable description of this Device Management Intune Branding Profile ID
func (id DeviceManagementIntuneBrandingProfileId) String() string {
	components := []string{
		fmt.Sprintf("Intune Branding Profile: %q", id.IntuneBrandingProfileId),
	}
	return fmt.Sprintf("Device Management Intune Branding Profile (%s)", strings.Join(components, "\n"))
}
