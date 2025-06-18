package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DirectoryTemplateDeviceTemplateIdDeviceInstanceId{}

// DirectoryTemplateDeviceTemplateIdDeviceInstanceId is a struct representing the Resource ID for a Directory Template Device Template Id Device Instance
type DirectoryTemplateDeviceTemplateIdDeviceInstanceId struct {
	DeviceTemplateId string
	DeviceId         string
}

// NewDirectoryTemplateDeviceTemplateIdDeviceInstanceID returns a new DirectoryTemplateDeviceTemplateIdDeviceInstanceId struct
func NewDirectoryTemplateDeviceTemplateIdDeviceInstanceID(deviceTemplateId string, deviceId string) DirectoryTemplateDeviceTemplateIdDeviceInstanceId {
	return DirectoryTemplateDeviceTemplateIdDeviceInstanceId{
		DeviceTemplateId: deviceTemplateId,
		DeviceId:         deviceId,
	}
}

// ParseDirectoryTemplateDeviceTemplateIdDeviceInstanceID parses 'input' into a DirectoryTemplateDeviceTemplateIdDeviceInstanceId
func ParseDirectoryTemplateDeviceTemplateIdDeviceInstanceID(input string) (*DirectoryTemplateDeviceTemplateIdDeviceInstanceId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DirectoryTemplateDeviceTemplateIdDeviceInstanceId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DirectoryTemplateDeviceTemplateIdDeviceInstanceId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseDirectoryTemplateDeviceTemplateIdDeviceInstanceIDInsensitively parses 'input' case-insensitively into a DirectoryTemplateDeviceTemplateIdDeviceInstanceId
// note: this method should only be used for API response data and not user input
func ParseDirectoryTemplateDeviceTemplateIdDeviceInstanceIDInsensitively(input string) (*DirectoryTemplateDeviceTemplateIdDeviceInstanceId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DirectoryTemplateDeviceTemplateIdDeviceInstanceId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DirectoryTemplateDeviceTemplateIdDeviceInstanceId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *DirectoryTemplateDeviceTemplateIdDeviceInstanceId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.DeviceTemplateId, ok = input.Parsed["deviceTemplateId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "deviceTemplateId", input)
	}

	if id.DeviceId, ok = input.Parsed["deviceId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "deviceId", input)
	}

	return nil
}

// ValidateDirectoryTemplateDeviceTemplateIdDeviceInstanceID checks that 'input' can be parsed as a Directory Template Device Template Id Device Instance ID
func ValidateDirectoryTemplateDeviceTemplateIdDeviceInstanceID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDirectoryTemplateDeviceTemplateIdDeviceInstanceID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Directory Template Device Template Id Device Instance ID
func (id DirectoryTemplateDeviceTemplateIdDeviceInstanceId) ID() string {
	fmtString := "/directory/templates/deviceTemplates/%s/deviceInstances/%s"
	return fmt.Sprintf(fmtString, id.DeviceTemplateId, id.DeviceId)
}

// Segments returns a slice of Resource ID Segments which comprise this Directory Template Device Template Id Device Instance ID
func (id DirectoryTemplateDeviceTemplateIdDeviceInstanceId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("directory", "directory", "directory"),
		resourceids.StaticSegment("templates", "templates", "templates"),
		resourceids.StaticSegment("deviceTemplates", "deviceTemplates", "deviceTemplates"),
		resourceids.UserSpecifiedSegment("deviceTemplateId", "deviceTemplateId"),
		resourceids.StaticSegment("deviceInstances", "deviceInstances", "deviceInstances"),
		resourceids.UserSpecifiedSegment("deviceId", "deviceId"),
	}
}

// String returns a human-readable description of this Directory Template Device Template Id Device Instance ID
func (id DirectoryTemplateDeviceTemplateIdDeviceInstanceId) String() string {
	components := []string{
		fmt.Sprintf("Device Template: %q", id.DeviceTemplateId),
		fmt.Sprintf("Device: %q", id.DeviceId),
	}
	return fmt.Sprintf("Directory Template Device Template Id Device Instance (%s)", strings.Join(components, "\n"))
}
