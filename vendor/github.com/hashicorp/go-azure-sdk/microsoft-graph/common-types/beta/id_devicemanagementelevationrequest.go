package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementElevationRequestId{}

// DeviceManagementElevationRequestId is a struct representing the Resource ID for a Device Management Elevation Request
type DeviceManagementElevationRequestId struct {
	PrivilegeManagementElevationRequestId string
}

// NewDeviceManagementElevationRequestID returns a new DeviceManagementElevationRequestId struct
func NewDeviceManagementElevationRequestID(privilegeManagementElevationRequestId string) DeviceManagementElevationRequestId {
	return DeviceManagementElevationRequestId{
		PrivilegeManagementElevationRequestId: privilegeManagementElevationRequestId,
	}
}

// ParseDeviceManagementElevationRequestID parses 'input' into a DeviceManagementElevationRequestId
func ParseDeviceManagementElevationRequestID(input string) (*DeviceManagementElevationRequestId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementElevationRequestId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementElevationRequestId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseDeviceManagementElevationRequestIDInsensitively parses 'input' case-insensitively into a DeviceManagementElevationRequestId
// note: this method should only be used for API response data and not user input
func ParseDeviceManagementElevationRequestIDInsensitively(input string) (*DeviceManagementElevationRequestId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementElevationRequestId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementElevationRequestId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *DeviceManagementElevationRequestId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.PrivilegeManagementElevationRequestId, ok = input.Parsed["privilegeManagementElevationRequestId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "privilegeManagementElevationRequestId", input)
	}

	return nil
}

// ValidateDeviceManagementElevationRequestID checks that 'input' can be parsed as a Device Management Elevation Request ID
func ValidateDeviceManagementElevationRequestID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDeviceManagementElevationRequestID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Device Management Elevation Request ID
func (id DeviceManagementElevationRequestId) ID() string {
	fmtString := "/deviceManagement/elevationRequests/%s"
	return fmt.Sprintf(fmtString, id.PrivilegeManagementElevationRequestId)
}

// Segments returns a slice of Resource ID Segments which comprise this Device Management Elevation Request ID
func (id DeviceManagementElevationRequestId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("deviceManagement", "deviceManagement", "deviceManagement"),
		resourceids.StaticSegment("elevationRequests", "elevationRequests", "elevationRequests"),
		resourceids.UserSpecifiedSegment("privilegeManagementElevationRequestId", "privilegeManagementElevationRequestId"),
	}
}

// String returns a human-readable description of this Device Management Elevation Request ID
func (id DeviceManagementElevationRequestId) String() string {
	components := []string{
		fmt.Sprintf("Privilege Management Elevation Request: %q", id.PrivilegeManagementElevationRequestId),
	}
	return fmt.Sprintf("Device Management Elevation Request (%s)", strings.Join(components, "\n"))
}
