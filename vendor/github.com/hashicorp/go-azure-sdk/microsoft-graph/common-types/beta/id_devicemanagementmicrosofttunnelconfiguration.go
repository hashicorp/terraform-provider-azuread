package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementMicrosoftTunnelConfigurationId{}

// DeviceManagementMicrosoftTunnelConfigurationId is a struct representing the Resource ID for a Device Management Microsoft Tunnel Configuration
type DeviceManagementMicrosoftTunnelConfigurationId struct {
	MicrosoftTunnelConfigurationId string
}

// NewDeviceManagementMicrosoftTunnelConfigurationID returns a new DeviceManagementMicrosoftTunnelConfigurationId struct
func NewDeviceManagementMicrosoftTunnelConfigurationID(microsoftTunnelConfigurationId string) DeviceManagementMicrosoftTunnelConfigurationId {
	return DeviceManagementMicrosoftTunnelConfigurationId{
		MicrosoftTunnelConfigurationId: microsoftTunnelConfigurationId,
	}
}

// ParseDeviceManagementMicrosoftTunnelConfigurationID parses 'input' into a DeviceManagementMicrosoftTunnelConfigurationId
func ParseDeviceManagementMicrosoftTunnelConfigurationID(input string) (*DeviceManagementMicrosoftTunnelConfigurationId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementMicrosoftTunnelConfigurationId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementMicrosoftTunnelConfigurationId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseDeviceManagementMicrosoftTunnelConfigurationIDInsensitively parses 'input' case-insensitively into a DeviceManagementMicrosoftTunnelConfigurationId
// note: this method should only be used for API response data and not user input
func ParseDeviceManagementMicrosoftTunnelConfigurationIDInsensitively(input string) (*DeviceManagementMicrosoftTunnelConfigurationId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementMicrosoftTunnelConfigurationId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementMicrosoftTunnelConfigurationId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *DeviceManagementMicrosoftTunnelConfigurationId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.MicrosoftTunnelConfigurationId, ok = input.Parsed["microsoftTunnelConfigurationId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "microsoftTunnelConfigurationId", input)
	}

	return nil
}

// ValidateDeviceManagementMicrosoftTunnelConfigurationID checks that 'input' can be parsed as a Device Management Microsoft Tunnel Configuration ID
func ValidateDeviceManagementMicrosoftTunnelConfigurationID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDeviceManagementMicrosoftTunnelConfigurationID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Device Management Microsoft Tunnel Configuration ID
func (id DeviceManagementMicrosoftTunnelConfigurationId) ID() string {
	fmtString := "/deviceManagement/microsoftTunnelConfigurations/%s"
	return fmt.Sprintf(fmtString, id.MicrosoftTunnelConfigurationId)
}

// Segments returns a slice of Resource ID Segments which comprise this Device Management Microsoft Tunnel Configuration ID
func (id DeviceManagementMicrosoftTunnelConfigurationId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("deviceManagement", "deviceManagement", "deviceManagement"),
		resourceids.StaticSegment("microsoftTunnelConfigurations", "microsoftTunnelConfigurations", "microsoftTunnelConfigurations"),
		resourceids.UserSpecifiedSegment("microsoftTunnelConfigurationId", "microsoftTunnelConfigurationId"),
	}
}

// String returns a human-readable description of this Device Management Microsoft Tunnel Configuration ID
func (id DeviceManagementMicrosoftTunnelConfigurationId) String() string {
	components := []string{
		fmt.Sprintf("Microsoft Tunnel Configuration: %q", id.MicrosoftTunnelConfigurationId),
	}
	return fmt.Sprintf("Device Management Microsoft Tunnel Configuration (%s)", strings.Join(components, "\n"))
}
