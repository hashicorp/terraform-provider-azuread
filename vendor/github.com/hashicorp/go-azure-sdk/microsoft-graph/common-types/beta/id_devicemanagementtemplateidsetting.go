package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementTemplateIdSettingId{}

// DeviceManagementTemplateIdSettingId is a struct representing the Resource ID for a Device Management Template Id Setting
type DeviceManagementTemplateIdSettingId struct {
	DeviceManagementTemplateId        string
	DeviceManagementSettingInstanceId string
}

// NewDeviceManagementTemplateIdSettingID returns a new DeviceManagementTemplateIdSettingId struct
func NewDeviceManagementTemplateIdSettingID(deviceManagementTemplateId string, deviceManagementSettingInstanceId string) DeviceManagementTemplateIdSettingId {
	return DeviceManagementTemplateIdSettingId{
		DeviceManagementTemplateId:        deviceManagementTemplateId,
		DeviceManagementSettingInstanceId: deviceManagementSettingInstanceId,
	}
}

// ParseDeviceManagementTemplateIdSettingID parses 'input' into a DeviceManagementTemplateIdSettingId
func ParseDeviceManagementTemplateIdSettingID(input string) (*DeviceManagementTemplateIdSettingId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementTemplateIdSettingId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementTemplateIdSettingId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseDeviceManagementTemplateIdSettingIDInsensitively parses 'input' case-insensitively into a DeviceManagementTemplateIdSettingId
// note: this method should only be used for API response data and not user input
func ParseDeviceManagementTemplateIdSettingIDInsensitively(input string) (*DeviceManagementTemplateIdSettingId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementTemplateIdSettingId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementTemplateIdSettingId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *DeviceManagementTemplateIdSettingId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.DeviceManagementTemplateId, ok = input.Parsed["deviceManagementTemplateId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "deviceManagementTemplateId", input)
	}

	if id.DeviceManagementSettingInstanceId, ok = input.Parsed["deviceManagementSettingInstanceId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "deviceManagementSettingInstanceId", input)
	}

	return nil
}

// ValidateDeviceManagementTemplateIdSettingID checks that 'input' can be parsed as a Device Management Template Id Setting ID
func ValidateDeviceManagementTemplateIdSettingID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDeviceManagementTemplateIdSettingID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Device Management Template Id Setting ID
func (id DeviceManagementTemplateIdSettingId) ID() string {
	fmtString := "/deviceManagement/templates/%s/settings/%s"
	return fmt.Sprintf(fmtString, id.DeviceManagementTemplateId, id.DeviceManagementSettingInstanceId)
}

// Segments returns a slice of Resource ID Segments which comprise this Device Management Template Id Setting ID
func (id DeviceManagementTemplateIdSettingId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("deviceManagement", "deviceManagement", "deviceManagement"),
		resourceids.StaticSegment("templates", "templates", "templates"),
		resourceids.UserSpecifiedSegment("deviceManagementTemplateId", "deviceManagementTemplateId"),
		resourceids.StaticSegment("settings", "settings", "settings"),
		resourceids.UserSpecifiedSegment("deviceManagementSettingInstanceId", "deviceManagementSettingInstanceId"),
	}
}

// String returns a human-readable description of this Device Management Template Id Setting ID
func (id DeviceManagementTemplateIdSettingId) String() string {
	components := []string{
		fmt.Sprintf("Device Management Template: %q", id.DeviceManagementTemplateId),
		fmt.Sprintf("Device Management Setting Instance: %q", id.DeviceManagementSettingInstanceId),
	}
	return fmt.Sprintf("Device Management Template Id Setting (%s)", strings.Join(components, "\n"))
}
