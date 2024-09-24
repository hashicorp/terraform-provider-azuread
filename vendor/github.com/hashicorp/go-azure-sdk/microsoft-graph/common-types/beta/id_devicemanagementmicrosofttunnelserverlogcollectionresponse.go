package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementMicrosoftTunnelServerLogCollectionResponseId{}

// DeviceManagementMicrosoftTunnelServerLogCollectionResponseId is a struct representing the Resource ID for a Device Management Microsoft Tunnel Server Log Collection Response
type DeviceManagementMicrosoftTunnelServerLogCollectionResponseId struct {
	MicrosoftTunnelServerLogCollectionResponseId string
}

// NewDeviceManagementMicrosoftTunnelServerLogCollectionResponseID returns a new DeviceManagementMicrosoftTunnelServerLogCollectionResponseId struct
func NewDeviceManagementMicrosoftTunnelServerLogCollectionResponseID(microsoftTunnelServerLogCollectionResponseId string) DeviceManagementMicrosoftTunnelServerLogCollectionResponseId {
	return DeviceManagementMicrosoftTunnelServerLogCollectionResponseId{
		MicrosoftTunnelServerLogCollectionResponseId: microsoftTunnelServerLogCollectionResponseId,
	}
}

// ParseDeviceManagementMicrosoftTunnelServerLogCollectionResponseID parses 'input' into a DeviceManagementMicrosoftTunnelServerLogCollectionResponseId
func ParseDeviceManagementMicrosoftTunnelServerLogCollectionResponseID(input string) (*DeviceManagementMicrosoftTunnelServerLogCollectionResponseId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementMicrosoftTunnelServerLogCollectionResponseId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementMicrosoftTunnelServerLogCollectionResponseId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseDeviceManagementMicrosoftTunnelServerLogCollectionResponseIDInsensitively parses 'input' case-insensitively into a DeviceManagementMicrosoftTunnelServerLogCollectionResponseId
// note: this method should only be used for API response data and not user input
func ParseDeviceManagementMicrosoftTunnelServerLogCollectionResponseIDInsensitively(input string) (*DeviceManagementMicrosoftTunnelServerLogCollectionResponseId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementMicrosoftTunnelServerLogCollectionResponseId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementMicrosoftTunnelServerLogCollectionResponseId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *DeviceManagementMicrosoftTunnelServerLogCollectionResponseId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.MicrosoftTunnelServerLogCollectionResponseId, ok = input.Parsed["microsoftTunnelServerLogCollectionResponseId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "microsoftTunnelServerLogCollectionResponseId", input)
	}

	return nil
}

// ValidateDeviceManagementMicrosoftTunnelServerLogCollectionResponseID checks that 'input' can be parsed as a Device Management Microsoft Tunnel Server Log Collection Response ID
func ValidateDeviceManagementMicrosoftTunnelServerLogCollectionResponseID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDeviceManagementMicrosoftTunnelServerLogCollectionResponseID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Device Management Microsoft Tunnel Server Log Collection Response ID
func (id DeviceManagementMicrosoftTunnelServerLogCollectionResponseId) ID() string {
	fmtString := "/deviceManagement/microsoftTunnelServerLogCollectionResponses/%s"
	return fmt.Sprintf(fmtString, id.MicrosoftTunnelServerLogCollectionResponseId)
}

// Segments returns a slice of Resource ID Segments which comprise this Device Management Microsoft Tunnel Server Log Collection Response ID
func (id DeviceManagementMicrosoftTunnelServerLogCollectionResponseId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("deviceManagement", "deviceManagement", "deviceManagement"),
		resourceids.StaticSegment("microsoftTunnelServerLogCollectionResponses", "microsoftTunnelServerLogCollectionResponses", "microsoftTunnelServerLogCollectionResponses"),
		resourceids.UserSpecifiedSegment("microsoftTunnelServerLogCollectionResponseId", "microsoftTunnelServerLogCollectionResponseId"),
	}
}

// String returns a human-readable description of this Device Management Microsoft Tunnel Server Log Collection Response ID
func (id DeviceManagementMicrosoftTunnelServerLogCollectionResponseId) String() string {
	components := []string{
		fmt.Sprintf("Microsoft Tunnel Server Log Collection Response: %q", id.MicrosoftTunnelServerLogCollectionResponseId),
	}
	return fmt.Sprintf("Device Management Microsoft Tunnel Server Log Collection Response (%s)", strings.Join(components, "\n"))
}
