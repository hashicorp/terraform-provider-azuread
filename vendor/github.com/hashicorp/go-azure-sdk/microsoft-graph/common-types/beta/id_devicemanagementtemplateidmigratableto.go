package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementTemplateIdMigratableToId{}

// DeviceManagementTemplateIdMigratableToId is a struct representing the Resource ID for a Device Management Template Id Migratable To
type DeviceManagementTemplateIdMigratableToId struct {
	DeviceManagementTemplateId  string
	DeviceManagementTemplateId1 string
}

// NewDeviceManagementTemplateIdMigratableToID returns a new DeviceManagementTemplateIdMigratableToId struct
func NewDeviceManagementTemplateIdMigratableToID(deviceManagementTemplateId string, deviceManagementTemplateId1 string) DeviceManagementTemplateIdMigratableToId {
	return DeviceManagementTemplateIdMigratableToId{
		DeviceManagementTemplateId:  deviceManagementTemplateId,
		DeviceManagementTemplateId1: deviceManagementTemplateId1,
	}
}

// ParseDeviceManagementTemplateIdMigratableToID parses 'input' into a DeviceManagementTemplateIdMigratableToId
func ParseDeviceManagementTemplateIdMigratableToID(input string) (*DeviceManagementTemplateIdMigratableToId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementTemplateIdMigratableToId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementTemplateIdMigratableToId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseDeviceManagementTemplateIdMigratableToIDInsensitively parses 'input' case-insensitively into a DeviceManagementTemplateIdMigratableToId
// note: this method should only be used for API response data and not user input
func ParseDeviceManagementTemplateIdMigratableToIDInsensitively(input string) (*DeviceManagementTemplateIdMigratableToId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementTemplateIdMigratableToId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementTemplateIdMigratableToId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *DeviceManagementTemplateIdMigratableToId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.DeviceManagementTemplateId, ok = input.Parsed["deviceManagementTemplateId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "deviceManagementTemplateId", input)
	}

	if id.DeviceManagementTemplateId1, ok = input.Parsed["deviceManagementTemplateId1"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "deviceManagementTemplateId1", input)
	}

	return nil
}

// ValidateDeviceManagementTemplateIdMigratableToID checks that 'input' can be parsed as a Device Management Template Id Migratable To ID
func ValidateDeviceManagementTemplateIdMigratableToID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDeviceManagementTemplateIdMigratableToID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Device Management Template Id Migratable To ID
func (id DeviceManagementTemplateIdMigratableToId) ID() string {
	fmtString := "/deviceManagement/templates/%s/migratableTo/%s"
	return fmt.Sprintf(fmtString, id.DeviceManagementTemplateId, id.DeviceManagementTemplateId1)
}

// Segments returns a slice of Resource ID Segments which comprise this Device Management Template Id Migratable To ID
func (id DeviceManagementTemplateIdMigratableToId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("deviceManagement", "deviceManagement", "deviceManagement"),
		resourceids.StaticSegment("templates", "templates", "templates"),
		resourceids.UserSpecifiedSegment("deviceManagementTemplateId", "deviceManagementTemplateId"),
		resourceids.StaticSegment("migratableTo", "migratableTo", "migratableTo"),
		resourceids.UserSpecifiedSegment("deviceManagementTemplateId1", "deviceManagementTemplateId1"),
	}
}

// String returns a human-readable description of this Device Management Template Id Migratable To ID
func (id DeviceManagementTemplateIdMigratableToId) String() string {
	components := []string{
		fmt.Sprintf("Device Management Template: %q", id.DeviceManagementTemplateId),
		fmt.Sprintf("Device Management Template Id 1: %q", id.DeviceManagementTemplateId1),
	}
	return fmt.Sprintf("Device Management Template Id Migratable To (%s)", strings.Join(components, "\n"))
}
