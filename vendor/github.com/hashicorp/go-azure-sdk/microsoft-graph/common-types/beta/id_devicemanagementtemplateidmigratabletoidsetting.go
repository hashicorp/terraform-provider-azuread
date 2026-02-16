package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementTemplateIdMigratableToIdSettingId{}

// DeviceManagementTemplateIdMigratableToIdSettingId is a struct representing the Resource ID for a Device Management Template Id Migratable To Id Setting
type DeviceManagementTemplateIdMigratableToIdSettingId struct {
	DeviceManagementTemplateId        string
	DeviceManagementTemplateId1       string
	DeviceManagementSettingInstanceId string
}

// NewDeviceManagementTemplateIdMigratableToIdSettingID returns a new DeviceManagementTemplateIdMigratableToIdSettingId struct
func NewDeviceManagementTemplateIdMigratableToIdSettingID(deviceManagementTemplateId string, deviceManagementTemplateId1 string, deviceManagementSettingInstanceId string) DeviceManagementTemplateIdMigratableToIdSettingId {
	return DeviceManagementTemplateIdMigratableToIdSettingId{
		DeviceManagementTemplateId:        deviceManagementTemplateId,
		DeviceManagementTemplateId1:       deviceManagementTemplateId1,
		DeviceManagementSettingInstanceId: deviceManagementSettingInstanceId,
	}
}

// ParseDeviceManagementTemplateIdMigratableToIdSettingID parses 'input' into a DeviceManagementTemplateIdMigratableToIdSettingId
func ParseDeviceManagementTemplateIdMigratableToIdSettingID(input string) (*DeviceManagementTemplateIdMigratableToIdSettingId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementTemplateIdMigratableToIdSettingId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementTemplateIdMigratableToIdSettingId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseDeviceManagementTemplateIdMigratableToIdSettingIDInsensitively parses 'input' case-insensitively into a DeviceManagementTemplateIdMigratableToIdSettingId
// note: this method should only be used for API response data and not user input
func ParseDeviceManagementTemplateIdMigratableToIdSettingIDInsensitively(input string) (*DeviceManagementTemplateIdMigratableToIdSettingId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementTemplateIdMigratableToIdSettingId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementTemplateIdMigratableToIdSettingId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *DeviceManagementTemplateIdMigratableToIdSettingId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.DeviceManagementTemplateId, ok = input.Parsed["deviceManagementTemplateId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "deviceManagementTemplateId", input)
	}

	if id.DeviceManagementTemplateId1, ok = input.Parsed["deviceManagementTemplateId1"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "deviceManagementTemplateId1", input)
	}

	if id.DeviceManagementSettingInstanceId, ok = input.Parsed["deviceManagementSettingInstanceId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "deviceManagementSettingInstanceId", input)
	}

	return nil
}

// ValidateDeviceManagementTemplateIdMigratableToIdSettingID checks that 'input' can be parsed as a Device Management Template Id Migratable To Id Setting ID
func ValidateDeviceManagementTemplateIdMigratableToIdSettingID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDeviceManagementTemplateIdMigratableToIdSettingID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Device Management Template Id Migratable To Id Setting ID
func (id DeviceManagementTemplateIdMigratableToIdSettingId) ID() string {
	fmtString := "/deviceManagement/templates/%s/migratableTo/%s/settings/%s"
	return fmt.Sprintf(fmtString, id.DeviceManagementTemplateId, id.DeviceManagementTemplateId1, id.DeviceManagementSettingInstanceId)
}

// Segments returns a slice of Resource ID Segments which comprise this Device Management Template Id Migratable To Id Setting ID
func (id DeviceManagementTemplateIdMigratableToIdSettingId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("deviceManagement", "deviceManagement", "deviceManagement"),
		resourceids.StaticSegment("templates", "templates", "templates"),
		resourceids.UserSpecifiedSegment("deviceManagementTemplateId", "deviceManagementTemplateId"),
		resourceids.StaticSegment("migratableTo", "migratableTo", "migratableTo"),
		resourceids.UserSpecifiedSegment("deviceManagementTemplateId1", "deviceManagementTemplateId1"),
		resourceids.StaticSegment("settings", "settings", "settings"),
		resourceids.UserSpecifiedSegment("deviceManagementSettingInstanceId", "deviceManagementSettingInstanceId"),
	}
}

// String returns a human-readable description of this Device Management Template Id Migratable To Id Setting ID
func (id DeviceManagementTemplateIdMigratableToIdSettingId) String() string {
	components := []string{
		fmt.Sprintf("Device Management Template: %q", id.DeviceManagementTemplateId),
		fmt.Sprintf("Device Management Template Id 1: %q", id.DeviceManagementTemplateId1),
		fmt.Sprintf("Device Management Setting Instance: %q", id.DeviceManagementSettingInstanceId),
	}
	return fmt.Sprintf("Device Management Template Id Migratable To Id Setting (%s)", strings.Join(components, "\n"))
}
