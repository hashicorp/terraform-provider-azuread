package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementAndroidForWorkAppConfigurationSchemaId{}

// DeviceManagementAndroidForWorkAppConfigurationSchemaId is a struct representing the Resource ID for a Device Management Android For Work App Configuration Schema
type DeviceManagementAndroidForWorkAppConfigurationSchemaId struct {
	AndroidForWorkAppConfigurationSchemaId string
}

// NewDeviceManagementAndroidForWorkAppConfigurationSchemaID returns a new DeviceManagementAndroidForWorkAppConfigurationSchemaId struct
func NewDeviceManagementAndroidForWorkAppConfigurationSchemaID(androidForWorkAppConfigurationSchemaId string) DeviceManagementAndroidForWorkAppConfigurationSchemaId {
	return DeviceManagementAndroidForWorkAppConfigurationSchemaId{
		AndroidForWorkAppConfigurationSchemaId: androidForWorkAppConfigurationSchemaId,
	}
}

// ParseDeviceManagementAndroidForWorkAppConfigurationSchemaID parses 'input' into a DeviceManagementAndroidForWorkAppConfigurationSchemaId
func ParseDeviceManagementAndroidForWorkAppConfigurationSchemaID(input string) (*DeviceManagementAndroidForWorkAppConfigurationSchemaId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementAndroidForWorkAppConfigurationSchemaId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementAndroidForWorkAppConfigurationSchemaId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseDeviceManagementAndroidForWorkAppConfigurationSchemaIDInsensitively parses 'input' case-insensitively into a DeviceManagementAndroidForWorkAppConfigurationSchemaId
// note: this method should only be used for API response data and not user input
func ParseDeviceManagementAndroidForWorkAppConfigurationSchemaIDInsensitively(input string) (*DeviceManagementAndroidForWorkAppConfigurationSchemaId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementAndroidForWorkAppConfigurationSchemaId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementAndroidForWorkAppConfigurationSchemaId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *DeviceManagementAndroidForWorkAppConfigurationSchemaId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.AndroidForWorkAppConfigurationSchemaId, ok = input.Parsed["androidForWorkAppConfigurationSchemaId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "androidForWorkAppConfigurationSchemaId", input)
	}

	return nil
}

// ValidateDeviceManagementAndroidForWorkAppConfigurationSchemaID checks that 'input' can be parsed as a Device Management Android For Work App Configuration Schema ID
func ValidateDeviceManagementAndroidForWorkAppConfigurationSchemaID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDeviceManagementAndroidForWorkAppConfigurationSchemaID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Device Management Android For Work App Configuration Schema ID
func (id DeviceManagementAndroidForWorkAppConfigurationSchemaId) ID() string {
	fmtString := "/deviceManagement/androidForWorkAppConfigurationSchemas/%s"
	return fmt.Sprintf(fmtString, id.AndroidForWorkAppConfigurationSchemaId)
}

// Segments returns a slice of Resource ID Segments which comprise this Device Management Android For Work App Configuration Schema ID
func (id DeviceManagementAndroidForWorkAppConfigurationSchemaId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("deviceManagement", "deviceManagement", "deviceManagement"),
		resourceids.StaticSegment("androidForWorkAppConfigurationSchemas", "androidForWorkAppConfigurationSchemas", "androidForWorkAppConfigurationSchemas"),
		resourceids.UserSpecifiedSegment("androidForWorkAppConfigurationSchemaId", "androidForWorkAppConfigurationSchemaId"),
	}
}

// String returns a human-readable description of this Device Management Android For Work App Configuration Schema ID
func (id DeviceManagementAndroidForWorkAppConfigurationSchemaId) String() string {
	components := []string{
		fmt.Sprintf("Android For Work App Configuration Schema: %q", id.AndroidForWorkAppConfigurationSchemaId),
	}
	return fmt.Sprintf("Device Management Android For Work App Configuration Schema (%s)", strings.Join(components, "\n"))
}
