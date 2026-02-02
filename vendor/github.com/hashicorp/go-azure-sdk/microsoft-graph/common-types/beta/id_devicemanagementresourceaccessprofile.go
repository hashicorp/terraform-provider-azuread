package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementResourceAccessProfileId{}

// DeviceManagementResourceAccessProfileId is a struct representing the Resource ID for a Device Management Resource Access Profile
type DeviceManagementResourceAccessProfileId struct {
	DeviceManagementResourceAccessProfileBaseId string
}

// NewDeviceManagementResourceAccessProfileID returns a new DeviceManagementResourceAccessProfileId struct
func NewDeviceManagementResourceAccessProfileID(deviceManagementResourceAccessProfileBaseId string) DeviceManagementResourceAccessProfileId {
	return DeviceManagementResourceAccessProfileId{
		DeviceManagementResourceAccessProfileBaseId: deviceManagementResourceAccessProfileBaseId,
	}
}

// ParseDeviceManagementResourceAccessProfileID parses 'input' into a DeviceManagementResourceAccessProfileId
func ParseDeviceManagementResourceAccessProfileID(input string) (*DeviceManagementResourceAccessProfileId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementResourceAccessProfileId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementResourceAccessProfileId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseDeviceManagementResourceAccessProfileIDInsensitively parses 'input' case-insensitively into a DeviceManagementResourceAccessProfileId
// note: this method should only be used for API response data and not user input
func ParseDeviceManagementResourceAccessProfileIDInsensitively(input string) (*DeviceManagementResourceAccessProfileId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementResourceAccessProfileId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementResourceAccessProfileId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *DeviceManagementResourceAccessProfileId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.DeviceManagementResourceAccessProfileBaseId, ok = input.Parsed["deviceManagementResourceAccessProfileBaseId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "deviceManagementResourceAccessProfileBaseId", input)
	}

	return nil
}

// ValidateDeviceManagementResourceAccessProfileID checks that 'input' can be parsed as a Device Management Resource Access Profile ID
func ValidateDeviceManagementResourceAccessProfileID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDeviceManagementResourceAccessProfileID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Device Management Resource Access Profile ID
func (id DeviceManagementResourceAccessProfileId) ID() string {
	fmtString := "/deviceManagement/resourceAccessProfiles/%s"
	return fmt.Sprintf(fmtString, id.DeviceManagementResourceAccessProfileBaseId)
}

// Segments returns a slice of Resource ID Segments which comprise this Device Management Resource Access Profile ID
func (id DeviceManagementResourceAccessProfileId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("deviceManagement", "deviceManagement", "deviceManagement"),
		resourceids.StaticSegment("resourceAccessProfiles", "resourceAccessProfiles", "resourceAccessProfiles"),
		resourceids.UserSpecifiedSegment("deviceManagementResourceAccessProfileBaseId", "deviceManagementResourceAccessProfileBaseId"),
	}
}

// String returns a human-readable description of this Device Management Resource Access Profile ID
func (id DeviceManagementResourceAccessProfileId) String() string {
	components := []string{
		fmt.Sprintf("Device Management Resource Access Profile Base: %q", id.DeviceManagementResourceAccessProfileBaseId),
	}
	return fmt.Sprintf("Device Management Resource Access Profile (%s)", strings.Join(components, "\n"))
}
