package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DirectoryTemplateDeviceTemplateIdOwnerId{}

// DirectoryTemplateDeviceTemplateIdOwnerId is a struct representing the Resource ID for a Directory Template Device Template Id Owner
type DirectoryTemplateDeviceTemplateIdOwnerId struct {
	DeviceTemplateId  string
	DirectoryObjectId string
}

// NewDirectoryTemplateDeviceTemplateIdOwnerID returns a new DirectoryTemplateDeviceTemplateIdOwnerId struct
func NewDirectoryTemplateDeviceTemplateIdOwnerID(deviceTemplateId string, directoryObjectId string) DirectoryTemplateDeviceTemplateIdOwnerId {
	return DirectoryTemplateDeviceTemplateIdOwnerId{
		DeviceTemplateId:  deviceTemplateId,
		DirectoryObjectId: directoryObjectId,
	}
}

// ParseDirectoryTemplateDeviceTemplateIdOwnerID parses 'input' into a DirectoryTemplateDeviceTemplateIdOwnerId
func ParseDirectoryTemplateDeviceTemplateIdOwnerID(input string) (*DirectoryTemplateDeviceTemplateIdOwnerId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DirectoryTemplateDeviceTemplateIdOwnerId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DirectoryTemplateDeviceTemplateIdOwnerId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseDirectoryTemplateDeviceTemplateIdOwnerIDInsensitively parses 'input' case-insensitively into a DirectoryTemplateDeviceTemplateIdOwnerId
// note: this method should only be used for API response data and not user input
func ParseDirectoryTemplateDeviceTemplateIdOwnerIDInsensitively(input string) (*DirectoryTemplateDeviceTemplateIdOwnerId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DirectoryTemplateDeviceTemplateIdOwnerId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DirectoryTemplateDeviceTemplateIdOwnerId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *DirectoryTemplateDeviceTemplateIdOwnerId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.DeviceTemplateId, ok = input.Parsed["deviceTemplateId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "deviceTemplateId", input)
	}

	if id.DirectoryObjectId, ok = input.Parsed["directoryObjectId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "directoryObjectId", input)
	}

	return nil
}

// ValidateDirectoryTemplateDeviceTemplateIdOwnerID checks that 'input' can be parsed as a Directory Template Device Template Id Owner ID
func ValidateDirectoryTemplateDeviceTemplateIdOwnerID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDirectoryTemplateDeviceTemplateIdOwnerID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Directory Template Device Template Id Owner ID
func (id DirectoryTemplateDeviceTemplateIdOwnerId) ID() string {
	fmtString := "/directory/templates/deviceTemplates/%s/owners/%s"
	return fmt.Sprintf(fmtString, id.DeviceTemplateId, id.DirectoryObjectId)
}

// Segments returns a slice of Resource ID Segments which comprise this Directory Template Device Template Id Owner ID
func (id DirectoryTemplateDeviceTemplateIdOwnerId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("directory", "directory", "directory"),
		resourceids.StaticSegment("templates", "templates", "templates"),
		resourceids.StaticSegment("deviceTemplates", "deviceTemplates", "deviceTemplates"),
		resourceids.UserSpecifiedSegment("deviceTemplateId", "deviceTemplateId"),
		resourceids.StaticSegment("owners", "owners", "owners"),
		resourceids.UserSpecifiedSegment("directoryObjectId", "directoryObjectId"),
	}
}

// String returns a human-readable description of this Directory Template Device Template Id Owner ID
func (id DirectoryTemplateDeviceTemplateIdOwnerId) String() string {
	components := []string{
		fmt.Sprintf("Device Template: %q", id.DeviceTemplateId),
		fmt.Sprintf("Directory Object: %q", id.DirectoryObjectId),
	}
	return fmt.Sprintf("Directory Template Device Template Id Owner (%s)", strings.Join(components, "\n"))
}
