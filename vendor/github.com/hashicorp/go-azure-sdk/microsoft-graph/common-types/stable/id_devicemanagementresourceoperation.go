package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementResourceOperationId{}

// DeviceManagementResourceOperationId is a struct representing the Resource ID for a Device Management Resource Operation
type DeviceManagementResourceOperationId struct {
	ResourceOperationId string
}

// NewDeviceManagementResourceOperationID returns a new DeviceManagementResourceOperationId struct
func NewDeviceManagementResourceOperationID(resourceOperationId string) DeviceManagementResourceOperationId {
	return DeviceManagementResourceOperationId{
		ResourceOperationId: resourceOperationId,
	}
}

// ParseDeviceManagementResourceOperationID parses 'input' into a DeviceManagementResourceOperationId
func ParseDeviceManagementResourceOperationID(input string) (*DeviceManagementResourceOperationId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementResourceOperationId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementResourceOperationId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseDeviceManagementResourceOperationIDInsensitively parses 'input' case-insensitively into a DeviceManagementResourceOperationId
// note: this method should only be used for API response data and not user input
func ParseDeviceManagementResourceOperationIDInsensitively(input string) (*DeviceManagementResourceOperationId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementResourceOperationId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementResourceOperationId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *DeviceManagementResourceOperationId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.ResourceOperationId, ok = input.Parsed["resourceOperationId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "resourceOperationId", input)
	}

	return nil
}

// ValidateDeviceManagementResourceOperationID checks that 'input' can be parsed as a Device Management Resource Operation ID
func ValidateDeviceManagementResourceOperationID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDeviceManagementResourceOperationID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Device Management Resource Operation ID
func (id DeviceManagementResourceOperationId) ID() string {
	fmtString := "/deviceManagement/resourceOperations/%s"
	return fmt.Sprintf(fmtString, id.ResourceOperationId)
}

// Segments returns a slice of Resource ID Segments which comprise this Device Management Resource Operation ID
func (id DeviceManagementResourceOperationId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("deviceManagement", "deviceManagement", "deviceManagement"),
		resourceids.StaticSegment("resourceOperations", "resourceOperations", "resourceOperations"),
		resourceids.UserSpecifiedSegment("resourceOperationId", "resourceOperationId"),
	}
}

// String returns a human-readable description of this Device Management Resource Operation ID
func (id DeviceManagementResourceOperationId) String() string {
	components := []string{
		fmt.Sprintf("Resource Operation: %q", id.ResourceOperationId),
	}
	return fmt.Sprintf("Device Management Resource Operation (%s)", strings.Join(components, "\n"))
}
