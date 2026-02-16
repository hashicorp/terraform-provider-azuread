package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementPrivilegeManagementElevationId{}

// DeviceManagementPrivilegeManagementElevationId is a struct representing the Resource ID for a Device Management Privilege Management Elevation
type DeviceManagementPrivilegeManagementElevationId struct {
	PrivilegeManagementElevationId string
}

// NewDeviceManagementPrivilegeManagementElevationID returns a new DeviceManagementPrivilegeManagementElevationId struct
func NewDeviceManagementPrivilegeManagementElevationID(privilegeManagementElevationId string) DeviceManagementPrivilegeManagementElevationId {
	return DeviceManagementPrivilegeManagementElevationId{
		PrivilegeManagementElevationId: privilegeManagementElevationId,
	}
}

// ParseDeviceManagementPrivilegeManagementElevationID parses 'input' into a DeviceManagementPrivilegeManagementElevationId
func ParseDeviceManagementPrivilegeManagementElevationID(input string) (*DeviceManagementPrivilegeManagementElevationId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementPrivilegeManagementElevationId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementPrivilegeManagementElevationId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseDeviceManagementPrivilegeManagementElevationIDInsensitively parses 'input' case-insensitively into a DeviceManagementPrivilegeManagementElevationId
// note: this method should only be used for API response data and not user input
func ParseDeviceManagementPrivilegeManagementElevationIDInsensitively(input string) (*DeviceManagementPrivilegeManagementElevationId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementPrivilegeManagementElevationId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementPrivilegeManagementElevationId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *DeviceManagementPrivilegeManagementElevationId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.PrivilegeManagementElevationId, ok = input.Parsed["privilegeManagementElevationId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "privilegeManagementElevationId", input)
	}

	return nil
}

// ValidateDeviceManagementPrivilegeManagementElevationID checks that 'input' can be parsed as a Device Management Privilege Management Elevation ID
func ValidateDeviceManagementPrivilegeManagementElevationID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDeviceManagementPrivilegeManagementElevationID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Device Management Privilege Management Elevation ID
func (id DeviceManagementPrivilegeManagementElevationId) ID() string {
	fmtString := "/deviceManagement/privilegeManagementElevations/%s"
	return fmt.Sprintf(fmtString, id.PrivilegeManagementElevationId)
}

// Segments returns a slice of Resource ID Segments which comprise this Device Management Privilege Management Elevation ID
func (id DeviceManagementPrivilegeManagementElevationId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("deviceManagement", "deviceManagement", "deviceManagement"),
		resourceids.StaticSegment("privilegeManagementElevations", "privilegeManagementElevations", "privilegeManagementElevations"),
		resourceids.UserSpecifiedSegment("privilegeManagementElevationId", "privilegeManagementElevationId"),
	}
}

// String returns a human-readable description of this Device Management Privilege Management Elevation ID
func (id DeviceManagementPrivilegeManagementElevationId) String() string {
	components := []string{
		fmt.Sprintf("Privilege Management Elevation: %q", id.PrivilegeManagementElevationId),
	}
	return fmt.Sprintf("Device Management Privilege Management Elevation (%s)", strings.Join(components, "\n"))
}
