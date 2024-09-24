package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementConfigManagerCollectionId{}

// DeviceManagementConfigManagerCollectionId is a struct representing the Resource ID for a Device Management Config Manager Collection
type DeviceManagementConfigManagerCollectionId struct {
	ConfigManagerCollectionId string
}

// NewDeviceManagementConfigManagerCollectionID returns a new DeviceManagementConfigManagerCollectionId struct
func NewDeviceManagementConfigManagerCollectionID(configManagerCollectionId string) DeviceManagementConfigManagerCollectionId {
	return DeviceManagementConfigManagerCollectionId{
		ConfigManagerCollectionId: configManagerCollectionId,
	}
}

// ParseDeviceManagementConfigManagerCollectionID parses 'input' into a DeviceManagementConfigManagerCollectionId
func ParseDeviceManagementConfigManagerCollectionID(input string) (*DeviceManagementConfigManagerCollectionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementConfigManagerCollectionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementConfigManagerCollectionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseDeviceManagementConfigManagerCollectionIDInsensitively parses 'input' case-insensitively into a DeviceManagementConfigManagerCollectionId
// note: this method should only be used for API response data and not user input
func ParseDeviceManagementConfigManagerCollectionIDInsensitively(input string) (*DeviceManagementConfigManagerCollectionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementConfigManagerCollectionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementConfigManagerCollectionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *DeviceManagementConfigManagerCollectionId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.ConfigManagerCollectionId, ok = input.Parsed["configManagerCollectionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "configManagerCollectionId", input)
	}

	return nil
}

// ValidateDeviceManagementConfigManagerCollectionID checks that 'input' can be parsed as a Device Management Config Manager Collection ID
func ValidateDeviceManagementConfigManagerCollectionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDeviceManagementConfigManagerCollectionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Device Management Config Manager Collection ID
func (id DeviceManagementConfigManagerCollectionId) ID() string {
	fmtString := "/deviceManagement/configManagerCollections/%s"
	return fmt.Sprintf(fmtString, id.ConfigManagerCollectionId)
}

// Segments returns a slice of Resource ID Segments which comprise this Device Management Config Manager Collection ID
func (id DeviceManagementConfigManagerCollectionId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("deviceManagement", "deviceManagement", "deviceManagement"),
		resourceids.StaticSegment("configManagerCollections", "configManagerCollections", "configManagerCollections"),
		resourceids.UserSpecifiedSegment("configManagerCollectionId", "configManagerCollectionId"),
	}
}

// String returns a human-readable description of this Device Management Config Manager Collection ID
func (id DeviceManagementConfigManagerCollectionId) String() string {
	components := []string{
		fmt.Sprintf("Config Manager Collection: %q", id.ConfigManagerCollectionId),
	}
	return fmt.Sprintf("Device Management Config Manager Collection (%s)", strings.Join(components, "\n"))
}
