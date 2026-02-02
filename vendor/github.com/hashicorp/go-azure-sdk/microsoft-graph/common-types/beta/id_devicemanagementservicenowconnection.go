package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementServiceNowConnectionId{}

// DeviceManagementServiceNowConnectionId is a struct representing the Resource ID for a Device Management Service Now Connection
type DeviceManagementServiceNowConnectionId struct {
	ServiceNowConnectionId string
}

// NewDeviceManagementServiceNowConnectionID returns a new DeviceManagementServiceNowConnectionId struct
func NewDeviceManagementServiceNowConnectionID(serviceNowConnectionId string) DeviceManagementServiceNowConnectionId {
	return DeviceManagementServiceNowConnectionId{
		ServiceNowConnectionId: serviceNowConnectionId,
	}
}

// ParseDeviceManagementServiceNowConnectionID parses 'input' into a DeviceManagementServiceNowConnectionId
func ParseDeviceManagementServiceNowConnectionID(input string) (*DeviceManagementServiceNowConnectionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementServiceNowConnectionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementServiceNowConnectionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseDeviceManagementServiceNowConnectionIDInsensitively parses 'input' case-insensitively into a DeviceManagementServiceNowConnectionId
// note: this method should only be used for API response data and not user input
func ParseDeviceManagementServiceNowConnectionIDInsensitively(input string) (*DeviceManagementServiceNowConnectionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementServiceNowConnectionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementServiceNowConnectionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *DeviceManagementServiceNowConnectionId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.ServiceNowConnectionId, ok = input.Parsed["serviceNowConnectionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "serviceNowConnectionId", input)
	}

	return nil
}

// ValidateDeviceManagementServiceNowConnectionID checks that 'input' can be parsed as a Device Management Service Now Connection ID
func ValidateDeviceManagementServiceNowConnectionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDeviceManagementServiceNowConnectionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Device Management Service Now Connection ID
func (id DeviceManagementServiceNowConnectionId) ID() string {
	fmtString := "/deviceManagement/serviceNowConnections/%s"
	return fmt.Sprintf(fmtString, id.ServiceNowConnectionId)
}

// Segments returns a slice of Resource ID Segments which comprise this Device Management Service Now Connection ID
func (id DeviceManagementServiceNowConnectionId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("deviceManagement", "deviceManagement", "deviceManagement"),
		resourceids.StaticSegment("serviceNowConnections", "serviceNowConnections", "serviceNowConnections"),
		resourceids.UserSpecifiedSegment("serviceNowConnectionId", "serviceNowConnectionId"),
	}
}

// String returns a human-readable description of this Device Management Service Now Connection ID
func (id DeviceManagementServiceNowConnectionId) String() string {
	components := []string{
		fmt.Sprintf("Service Now Connection: %q", id.ServiceNowConnectionId),
	}
	return fmt.Sprintf("Device Management Service Now Connection (%s)", strings.Join(components, "\n"))
}
