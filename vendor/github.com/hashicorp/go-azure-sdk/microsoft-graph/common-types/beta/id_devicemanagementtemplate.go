package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementTemplateId{}

// DeviceManagementTemplateId is a struct representing the Resource ID for a Device Management Template
type DeviceManagementTemplateId struct {
	DeviceManagementTemplateId string
}

// NewDeviceManagementTemplateID returns a new DeviceManagementTemplateId struct
func NewDeviceManagementTemplateID(deviceManagementTemplateId string) DeviceManagementTemplateId {
	return DeviceManagementTemplateId{
		DeviceManagementTemplateId: deviceManagementTemplateId,
	}
}

// ParseDeviceManagementTemplateID parses 'input' into a DeviceManagementTemplateId
func ParseDeviceManagementTemplateID(input string) (*DeviceManagementTemplateId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementTemplateId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementTemplateId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseDeviceManagementTemplateIDInsensitively parses 'input' case-insensitively into a DeviceManagementTemplateId
// note: this method should only be used for API response data and not user input
func ParseDeviceManagementTemplateIDInsensitively(input string) (*DeviceManagementTemplateId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementTemplateId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementTemplateId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *DeviceManagementTemplateId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.DeviceManagementTemplateId, ok = input.Parsed["deviceManagementTemplateId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "deviceManagementTemplateId", input)
	}

	return nil
}

// ValidateDeviceManagementTemplateID checks that 'input' can be parsed as a Device Management Template ID
func ValidateDeviceManagementTemplateID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDeviceManagementTemplateID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Device Management Template ID
func (id DeviceManagementTemplateId) ID() string {
	fmtString := "/deviceManagement/templates/%s"
	return fmt.Sprintf(fmtString, id.DeviceManagementTemplateId)
}

// Segments returns a slice of Resource ID Segments which comprise this Device Management Template ID
func (id DeviceManagementTemplateId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("deviceManagement", "deviceManagement", "deviceManagement"),
		resourceids.StaticSegment("templates", "templates", "templates"),
		resourceids.UserSpecifiedSegment("deviceManagementTemplateId", "deviceManagementTemplateId"),
	}
}

// String returns a human-readable description of this Device Management Template ID
func (id DeviceManagementTemplateId) String() string {
	components := []string{
		fmt.Sprintf("Device Management Template: %q", id.DeviceManagementTemplateId),
	}
	return fmt.Sprintf("Device Management Template (%s)", strings.Join(components, "\n"))
}
