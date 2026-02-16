package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementMicrosoftTunnelSiteIdMicrosoftTunnelServerId{}

// DeviceManagementMicrosoftTunnelSiteIdMicrosoftTunnelServerId is a struct representing the Resource ID for a Device Management Microsoft Tunnel Site Id Microsoft Tunnel Server
type DeviceManagementMicrosoftTunnelSiteIdMicrosoftTunnelServerId struct {
	MicrosoftTunnelSiteId   string
	MicrosoftTunnelServerId string
}

// NewDeviceManagementMicrosoftTunnelSiteIdMicrosoftTunnelServerID returns a new DeviceManagementMicrosoftTunnelSiteIdMicrosoftTunnelServerId struct
func NewDeviceManagementMicrosoftTunnelSiteIdMicrosoftTunnelServerID(microsoftTunnelSiteId string, microsoftTunnelServerId string) DeviceManagementMicrosoftTunnelSiteIdMicrosoftTunnelServerId {
	return DeviceManagementMicrosoftTunnelSiteIdMicrosoftTunnelServerId{
		MicrosoftTunnelSiteId:   microsoftTunnelSiteId,
		MicrosoftTunnelServerId: microsoftTunnelServerId,
	}
}

// ParseDeviceManagementMicrosoftTunnelSiteIdMicrosoftTunnelServerID parses 'input' into a DeviceManagementMicrosoftTunnelSiteIdMicrosoftTunnelServerId
func ParseDeviceManagementMicrosoftTunnelSiteIdMicrosoftTunnelServerID(input string) (*DeviceManagementMicrosoftTunnelSiteIdMicrosoftTunnelServerId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementMicrosoftTunnelSiteIdMicrosoftTunnelServerId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementMicrosoftTunnelSiteIdMicrosoftTunnelServerId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseDeviceManagementMicrosoftTunnelSiteIdMicrosoftTunnelServerIDInsensitively parses 'input' case-insensitively into a DeviceManagementMicrosoftTunnelSiteIdMicrosoftTunnelServerId
// note: this method should only be used for API response data and not user input
func ParseDeviceManagementMicrosoftTunnelSiteIdMicrosoftTunnelServerIDInsensitively(input string) (*DeviceManagementMicrosoftTunnelSiteIdMicrosoftTunnelServerId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementMicrosoftTunnelSiteIdMicrosoftTunnelServerId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementMicrosoftTunnelSiteIdMicrosoftTunnelServerId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *DeviceManagementMicrosoftTunnelSiteIdMicrosoftTunnelServerId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.MicrosoftTunnelSiteId, ok = input.Parsed["microsoftTunnelSiteId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "microsoftTunnelSiteId", input)
	}

	if id.MicrosoftTunnelServerId, ok = input.Parsed["microsoftTunnelServerId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "microsoftTunnelServerId", input)
	}

	return nil
}

// ValidateDeviceManagementMicrosoftTunnelSiteIdMicrosoftTunnelServerID checks that 'input' can be parsed as a Device Management Microsoft Tunnel Site Id Microsoft Tunnel Server ID
func ValidateDeviceManagementMicrosoftTunnelSiteIdMicrosoftTunnelServerID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDeviceManagementMicrosoftTunnelSiteIdMicrosoftTunnelServerID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Device Management Microsoft Tunnel Site Id Microsoft Tunnel Server ID
func (id DeviceManagementMicrosoftTunnelSiteIdMicrosoftTunnelServerId) ID() string {
	fmtString := "/deviceManagement/microsoftTunnelSites/%s/microsoftTunnelServers/%s"
	return fmt.Sprintf(fmtString, id.MicrosoftTunnelSiteId, id.MicrosoftTunnelServerId)
}

// Segments returns a slice of Resource ID Segments which comprise this Device Management Microsoft Tunnel Site Id Microsoft Tunnel Server ID
func (id DeviceManagementMicrosoftTunnelSiteIdMicrosoftTunnelServerId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("deviceManagement", "deviceManagement", "deviceManagement"),
		resourceids.StaticSegment("microsoftTunnelSites", "microsoftTunnelSites", "microsoftTunnelSites"),
		resourceids.UserSpecifiedSegment("microsoftTunnelSiteId", "microsoftTunnelSiteId"),
		resourceids.StaticSegment("microsoftTunnelServers", "microsoftTunnelServers", "microsoftTunnelServers"),
		resourceids.UserSpecifiedSegment("microsoftTunnelServerId", "microsoftTunnelServerId"),
	}
}

// String returns a human-readable description of this Device Management Microsoft Tunnel Site Id Microsoft Tunnel Server ID
func (id DeviceManagementMicrosoftTunnelSiteIdMicrosoftTunnelServerId) String() string {
	components := []string{
		fmt.Sprintf("Microsoft Tunnel Site: %q", id.MicrosoftTunnelSiteId),
		fmt.Sprintf("Microsoft Tunnel Server: %q", id.MicrosoftTunnelServerId),
	}
	return fmt.Sprintf("Device Management Microsoft Tunnel Site Id Microsoft Tunnel Server (%s)", strings.Join(components, "\n"))
}
