package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementMicrosoftTunnelSiteId{}

// DeviceManagementMicrosoftTunnelSiteId is a struct representing the Resource ID for a Device Management Microsoft Tunnel Site
type DeviceManagementMicrosoftTunnelSiteId struct {
	MicrosoftTunnelSiteId string
}

// NewDeviceManagementMicrosoftTunnelSiteID returns a new DeviceManagementMicrosoftTunnelSiteId struct
func NewDeviceManagementMicrosoftTunnelSiteID(microsoftTunnelSiteId string) DeviceManagementMicrosoftTunnelSiteId {
	return DeviceManagementMicrosoftTunnelSiteId{
		MicrosoftTunnelSiteId: microsoftTunnelSiteId,
	}
}

// ParseDeviceManagementMicrosoftTunnelSiteID parses 'input' into a DeviceManagementMicrosoftTunnelSiteId
func ParseDeviceManagementMicrosoftTunnelSiteID(input string) (*DeviceManagementMicrosoftTunnelSiteId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementMicrosoftTunnelSiteId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementMicrosoftTunnelSiteId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseDeviceManagementMicrosoftTunnelSiteIDInsensitively parses 'input' case-insensitively into a DeviceManagementMicrosoftTunnelSiteId
// note: this method should only be used for API response data and not user input
func ParseDeviceManagementMicrosoftTunnelSiteIDInsensitively(input string) (*DeviceManagementMicrosoftTunnelSiteId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementMicrosoftTunnelSiteId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementMicrosoftTunnelSiteId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *DeviceManagementMicrosoftTunnelSiteId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.MicrosoftTunnelSiteId, ok = input.Parsed["microsoftTunnelSiteId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "microsoftTunnelSiteId", input)
	}

	return nil
}

// ValidateDeviceManagementMicrosoftTunnelSiteID checks that 'input' can be parsed as a Device Management Microsoft Tunnel Site ID
func ValidateDeviceManagementMicrosoftTunnelSiteID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDeviceManagementMicrosoftTunnelSiteID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Device Management Microsoft Tunnel Site ID
func (id DeviceManagementMicrosoftTunnelSiteId) ID() string {
	fmtString := "/deviceManagement/microsoftTunnelSites/%s"
	return fmt.Sprintf(fmtString, id.MicrosoftTunnelSiteId)
}

// Segments returns a slice of Resource ID Segments which comprise this Device Management Microsoft Tunnel Site ID
func (id DeviceManagementMicrosoftTunnelSiteId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("deviceManagement", "deviceManagement", "deviceManagement"),
		resourceids.StaticSegment("microsoftTunnelSites", "microsoftTunnelSites", "microsoftTunnelSites"),
		resourceids.UserSpecifiedSegment("microsoftTunnelSiteId", "microsoftTunnelSiteId"),
	}
}

// String returns a human-readable description of this Device Management Microsoft Tunnel Site ID
func (id DeviceManagementMicrosoftTunnelSiteId) String() string {
	components := []string{
		fmt.Sprintf("Microsoft Tunnel Site: %q", id.MicrosoftTunnelSiteId),
	}
	return fmt.Sprintf("Device Management Microsoft Tunnel Site (%s)", strings.Join(components, "\n"))
}
