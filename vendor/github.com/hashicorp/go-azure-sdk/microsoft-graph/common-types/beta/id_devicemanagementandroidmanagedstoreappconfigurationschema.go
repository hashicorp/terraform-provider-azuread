package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementAndroidManagedStoreAppConfigurationSchemaId{}

// DeviceManagementAndroidManagedStoreAppConfigurationSchemaId is a struct representing the Resource ID for a Device Management Android Managed Store App Configuration Schema
type DeviceManagementAndroidManagedStoreAppConfigurationSchemaId struct {
	AndroidManagedStoreAppConfigurationSchemaId string
}

// NewDeviceManagementAndroidManagedStoreAppConfigurationSchemaID returns a new DeviceManagementAndroidManagedStoreAppConfigurationSchemaId struct
func NewDeviceManagementAndroidManagedStoreAppConfigurationSchemaID(androidManagedStoreAppConfigurationSchemaId string) DeviceManagementAndroidManagedStoreAppConfigurationSchemaId {
	return DeviceManagementAndroidManagedStoreAppConfigurationSchemaId{
		AndroidManagedStoreAppConfigurationSchemaId: androidManagedStoreAppConfigurationSchemaId,
	}
}

// ParseDeviceManagementAndroidManagedStoreAppConfigurationSchemaID parses 'input' into a DeviceManagementAndroidManagedStoreAppConfigurationSchemaId
func ParseDeviceManagementAndroidManagedStoreAppConfigurationSchemaID(input string) (*DeviceManagementAndroidManagedStoreAppConfigurationSchemaId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementAndroidManagedStoreAppConfigurationSchemaId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementAndroidManagedStoreAppConfigurationSchemaId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseDeviceManagementAndroidManagedStoreAppConfigurationSchemaIDInsensitively parses 'input' case-insensitively into a DeviceManagementAndroidManagedStoreAppConfigurationSchemaId
// note: this method should only be used for API response data and not user input
func ParseDeviceManagementAndroidManagedStoreAppConfigurationSchemaIDInsensitively(input string) (*DeviceManagementAndroidManagedStoreAppConfigurationSchemaId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementAndroidManagedStoreAppConfigurationSchemaId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementAndroidManagedStoreAppConfigurationSchemaId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *DeviceManagementAndroidManagedStoreAppConfigurationSchemaId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.AndroidManagedStoreAppConfigurationSchemaId, ok = input.Parsed["androidManagedStoreAppConfigurationSchemaId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "androidManagedStoreAppConfigurationSchemaId", input)
	}

	return nil
}

// ValidateDeviceManagementAndroidManagedStoreAppConfigurationSchemaID checks that 'input' can be parsed as a Device Management Android Managed Store App Configuration Schema ID
func ValidateDeviceManagementAndroidManagedStoreAppConfigurationSchemaID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDeviceManagementAndroidManagedStoreAppConfigurationSchemaID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Device Management Android Managed Store App Configuration Schema ID
func (id DeviceManagementAndroidManagedStoreAppConfigurationSchemaId) ID() string {
	fmtString := "/deviceManagement/androidManagedStoreAppConfigurationSchemas/%s"
	return fmt.Sprintf(fmtString, id.AndroidManagedStoreAppConfigurationSchemaId)
}

// Segments returns a slice of Resource ID Segments which comprise this Device Management Android Managed Store App Configuration Schema ID
func (id DeviceManagementAndroidManagedStoreAppConfigurationSchemaId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("deviceManagement", "deviceManagement", "deviceManagement"),
		resourceids.StaticSegment("androidManagedStoreAppConfigurationSchemas", "androidManagedStoreAppConfigurationSchemas", "androidManagedStoreAppConfigurationSchemas"),
		resourceids.UserSpecifiedSegment("androidManagedStoreAppConfigurationSchemaId", "androidManagedStoreAppConfigurationSchemaId"),
	}
}

// String returns a human-readable description of this Device Management Android Managed Store App Configuration Schema ID
func (id DeviceManagementAndroidManagedStoreAppConfigurationSchemaId) String() string {
	components := []string{
		fmt.Sprintf("Android Managed Store App Configuration Schema: %q", id.AndroidManagedStoreAppConfigurationSchemaId),
	}
	return fmt.Sprintf("Device Management Android Managed Store App Configuration Schema (%s)", strings.Join(components, "\n"))
}
