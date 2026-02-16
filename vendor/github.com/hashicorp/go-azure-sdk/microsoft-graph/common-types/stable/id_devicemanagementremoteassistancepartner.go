package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementRemoteAssistancePartnerId{}

// DeviceManagementRemoteAssistancePartnerId is a struct representing the Resource ID for a Device Management Remote Assistance Partner
type DeviceManagementRemoteAssistancePartnerId struct {
	RemoteAssistancePartnerId string
}

// NewDeviceManagementRemoteAssistancePartnerID returns a new DeviceManagementRemoteAssistancePartnerId struct
func NewDeviceManagementRemoteAssistancePartnerID(remoteAssistancePartnerId string) DeviceManagementRemoteAssistancePartnerId {
	return DeviceManagementRemoteAssistancePartnerId{
		RemoteAssistancePartnerId: remoteAssistancePartnerId,
	}
}

// ParseDeviceManagementRemoteAssistancePartnerID parses 'input' into a DeviceManagementRemoteAssistancePartnerId
func ParseDeviceManagementRemoteAssistancePartnerID(input string) (*DeviceManagementRemoteAssistancePartnerId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementRemoteAssistancePartnerId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementRemoteAssistancePartnerId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseDeviceManagementRemoteAssistancePartnerIDInsensitively parses 'input' case-insensitively into a DeviceManagementRemoteAssistancePartnerId
// note: this method should only be used for API response data and not user input
func ParseDeviceManagementRemoteAssistancePartnerIDInsensitively(input string) (*DeviceManagementRemoteAssistancePartnerId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementRemoteAssistancePartnerId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementRemoteAssistancePartnerId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *DeviceManagementRemoteAssistancePartnerId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.RemoteAssistancePartnerId, ok = input.Parsed["remoteAssistancePartnerId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "remoteAssistancePartnerId", input)
	}

	return nil
}

// ValidateDeviceManagementRemoteAssistancePartnerID checks that 'input' can be parsed as a Device Management Remote Assistance Partner ID
func ValidateDeviceManagementRemoteAssistancePartnerID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDeviceManagementRemoteAssistancePartnerID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Device Management Remote Assistance Partner ID
func (id DeviceManagementRemoteAssistancePartnerId) ID() string {
	fmtString := "/deviceManagement/remoteAssistancePartners/%s"
	return fmt.Sprintf(fmtString, id.RemoteAssistancePartnerId)
}

// Segments returns a slice of Resource ID Segments which comprise this Device Management Remote Assistance Partner ID
func (id DeviceManagementRemoteAssistancePartnerId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("deviceManagement", "deviceManagement", "deviceManagement"),
		resourceids.StaticSegment("remoteAssistancePartners", "remoteAssistancePartners", "remoteAssistancePartners"),
		resourceids.UserSpecifiedSegment("remoteAssistancePartnerId", "remoteAssistancePartnerId"),
	}
}

// String returns a human-readable description of this Device Management Remote Assistance Partner ID
func (id DeviceManagementRemoteAssistancePartnerId) String() string {
	components := []string{
		fmt.Sprintf("Remote Assistance Partner: %q", id.RemoteAssistancePartnerId),
	}
	return fmt.Sprintf("Device Management Remote Assistance Partner (%s)", strings.Join(components, "\n"))
}
