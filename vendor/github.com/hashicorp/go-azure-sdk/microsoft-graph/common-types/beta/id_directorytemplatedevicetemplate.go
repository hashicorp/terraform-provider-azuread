package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DirectoryTemplateDeviceTemplateId{}

// DirectoryTemplateDeviceTemplateId is a struct representing the Resource ID for a Directory Template Device Template
type DirectoryTemplateDeviceTemplateId struct {
	DeviceTemplateId string
}

// NewDirectoryTemplateDeviceTemplateID returns a new DirectoryTemplateDeviceTemplateId struct
func NewDirectoryTemplateDeviceTemplateID(deviceTemplateId string) DirectoryTemplateDeviceTemplateId {
	return DirectoryTemplateDeviceTemplateId{
		DeviceTemplateId: deviceTemplateId,
	}
}

// ParseDirectoryTemplateDeviceTemplateID parses 'input' into a DirectoryTemplateDeviceTemplateId
func ParseDirectoryTemplateDeviceTemplateID(input string) (*DirectoryTemplateDeviceTemplateId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DirectoryTemplateDeviceTemplateId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DirectoryTemplateDeviceTemplateId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseDirectoryTemplateDeviceTemplateIDInsensitively parses 'input' case-insensitively into a DirectoryTemplateDeviceTemplateId
// note: this method should only be used for API response data and not user input
func ParseDirectoryTemplateDeviceTemplateIDInsensitively(input string) (*DirectoryTemplateDeviceTemplateId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DirectoryTemplateDeviceTemplateId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DirectoryTemplateDeviceTemplateId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *DirectoryTemplateDeviceTemplateId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.DeviceTemplateId, ok = input.Parsed["deviceTemplateId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "deviceTemplateId", input)
	}

	return nil
}

// ValidateDirectoryTemplateDeviceTemplateID checks that 'input' can be parsed as a Directory Template Device Template ID
func ValidateDirectoryTemplateDeviceTemplateID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDirectoryTemplateDeviceTemplateID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Directory Template Device Template ID
func (id DirectoryTemplateDeviceTemplateId) ID() string {
	fmtString := "/directory/templates/deviceTemplates/%s"
	return fmt.Sprintf(fmtString, id.DeviceTemplateId)
}

// Segments returns a slice of Resource ID Segments which comprise this Directory Template Device Template ID
func (id DirectoryTemplateDeviceTemplateId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("directory", "directory", "directory"),
		resourceids.StaticSegment("templates", "templates", "templates"),
		resourceids.StaticSegment("deviceTemplates", "deviceTemplates", "deviceTemplates"),
		resourceids.UserSpecifiedSegment("deviceTemplateId", "deviceTemplateId"),
	}
}

// String returns a human-readable description of this Directory Template Device Template ID
func (id DirectoryTemplateDeviceTemplateId) String() string {
	components := []string{
		fmt.Sprintf("Device Template: %q", id.DeviceTemplateId),
	}
	return fmt.Sprintf("Directory Template Device Template (%s)", strings.Join(components, "\n"))
}
