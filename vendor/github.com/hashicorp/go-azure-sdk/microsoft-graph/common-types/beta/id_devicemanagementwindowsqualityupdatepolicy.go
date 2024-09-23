package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementWindowsQualityUpdatePolicyId{}

// DeviceManagementWindowsQualityUpdatePolicyId is a struct representing the Resource ID for a Device Management Windows Quality Update Policy
type DeviceManagementWindowsQualityUpdatePolicyId struct {
	WindowsQualityUpdatePolicyId string
}

// NewDeviceManagementWindowsQualityUpdatePolicyID returns a new DeviceManagementWindowsQualityUpdatePolicyId struct
func NewDeviceManagementWindowsQualityUpdatePolicyID(windowsQualityUpdatePolicyId string) DeviceManagementWindowsQualityUpdatePolicyId {
	return DeviceManagementWindowsQualityUpdatePolicyId{
		WindowsQualityUpdatePolicyId: windowsQualityUpdatePolicyId,
	}
}

// ParseDeviceManagementWindowsQualityUpdatePolicyID parses 'input' into a DeviceManagementWindowsQualityUpdatePolicyId
func ParseDeviceManagementWindowsQualityUpdatePolicyID(input string) (*DeviceManagementWindowsQualityUpdatePolicyId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementWindowsQualityUpdatePolicyId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementWindowsQualityUpdatePolicyId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseDeviceManagementWindowsQualityUpdatePolicyIDInsensitively parses 'input' case-insensitively into a DeviceManagementWindowsQualityUpdatePolicyId
// note: this method should only be used for API response data and not user input
func ParseDeviceManagementWindowsQualityUpdatePolicyIDInsensitively(input string) (*DeviceManagementWindowsQualityUpdatePolicyId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementWindowsQualityUpdatePolicyId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementWindowsQualityUpdatePolicyId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *DeviceManagementWindowsQualityUpdatePolicyId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.WindowsQualityUpdatePolicyId, ok = input.Parsed["windowsQualityUpdatePolicyId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "windowsQualityUpdatePolicyId", input)
	}

	return nil
}

// ValidateDeviceManagementWindowsQualityUpdatePolicyID checks that 'input' can be parsed as a Device Management Windows Quality Update Policy ID
func ValidateDeviceManagementWindowsQualityUpdatePolicyID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDeviceManagementWindowsQualityUpdatePolicyID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Device Management Windows Quality Update Policy ID
func (id DeviceManagementWindowsQualityUpdatePolicyId) ID() string {
	fmtString := "/deviceManagement/windowsQualityUpdatePolicies/%s"
	return fmt.Sprintf(fmtString, id.WindowsQualityUpdatePolicyId)
}

// Segments returns a slice of Resource ID Segments which comprise this Device Management Windows Quality Update Policy ID
func (id DeviceManagementWindowsQualityUpdatePolicyId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("deviceManagement", "deviceManagement", "deviceManagement"),
		resourceids.StaticSegment("windowsQualityUpdatePolicies", "windowsQualityUpdatePolicies", "windowsQualityUpdatePolicies"),
		resourceids.UserSpecifiedSegment("windowsQualityUpdatePolicyId", "windowsQualityUpdatePolicyId"),
	}
}

// String returns a human-readable description of this Device Management Windows Quality Update Policy ID
func (id DeviceManagementWindowsQualityUpdatePolicyId) String() string {
	components := []string{
		fmt.Sprintf("Windows Quality Update Policy: %q", id.WindowsQualityUpdatePolicyId),
	}
	return fmt.Sprintf("Device Management Windows Quality Update Policy (%s)", strings.Join(components, "\n"))
}
