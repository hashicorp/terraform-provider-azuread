package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementReusablePolicySettingId{}

// DeviceManagementReusablePolicySettingId is a struct representing the Resource ID for a Device Management Reusable Policy Setting
type DeviceManagementReusablePolicySettingId struct {
	DeviceManagementReusablePolicySettingId string
}

// NewDeviceManagementReusablePolicySettingID returns a new DeviceManagementReusablePolicySettingId struct
func NewDeviceManagementReusablePolicySettingID(deviceManagementReusablePolicySettingId string) DeviceManagementReusablePolicySettingId {
	return DeviceManagementReusablePolicySettingId{
		DeviceManagementReusablePolicySettingId: deviceManagementReusablePolicySettingId,
	}
}

// ParseDeviceManagementReusablePolicySettingID parses 'input' into a DeviceManagementReusablePolicySettingId
func ParseDeviceManagementReusablePolicySettingID(input string) (*DeviceManagementReusablePolicySettingId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementReusablePolicySettingId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementReusablePolicySettingId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseDeviceManagementReusablePolicySettingIDInsensitively parses 'input' case-insensitively into a DeviceManagementReusablePolicySettingId
// note: this method should only be used for API response data and not user input
func ParseDeviceManagementReusablePolicySettingIDInsensitively(input string) (*DeviceManagementReusablePolicySettingId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementReusablePolicySettingId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementReusablePolicySettingId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *DeviceManagementReusablePolicySettingId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.DeviceManagementReusablePolicySettingId, ok = input.Parsed["deviceManagementReusablePolicySettingId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "deviceManagementReusablePolicySettingId", input)
	}

	return nil
}

// ValidateDeviceManagementReusablePolicySettingID checks that 'input' can be parsed as a Device Management Reusable Policy Setting ID
func ValidateDeviceManagementReusablePolicySettingID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDeviceManagementReusablePolicySettingID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Device Management Reusable Policy Setting ID
func (id DeviceManagementReusablePolicySettingId) ID() string {
	fmtString := "/deviceManagement/reusablePolicySettings/%s"
	return fmt.Sprintf(fmtString, id.DeviceManagementReusablePolicySettingId)
}

// Segments returns a slice of Resource ID Segments which comprise this Device Management Reusable Policy Setting ID
func (id DeviceManagementReusablePolicySettingId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("deviceManagement", "deviceManagement", "deviceManagement"),
		resourceids.StaticSegment("reusablePolicySettings", "reusablePolicySettings", "reusablePolicySettings"),
		resourceids.UserSpecifiedSegment("deviceManagementReusablePolicySettingId", "deviceManagementReusablePolicySettingId"),
	}
}

// String returns a human-readable description of this Device Management Reusable Policy Setting ID
func (id DeviceManagementReusablePolicySettingId) String() string {
	components := []string{
		fmt.Sprintf("Device Management Reusable Policy Setting: %q", id.DeviceManagementReusablePolicySettingId),
	}
	return fmt.Sprintf("Device Management Reusable Policy Setting (%s)", strings.Join(components, "\n"))
}
