package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementVirtualEndpointExternalPartnerSettingId{}

// DeviceManagementVirtualEndpointExternalPartnerSettingId is a struct representing the Resource ID for a Device Management Virtual Endpoint External Partner Setting
type DeviceManagementVirtualEndpointExternalPartnerSettingId struct {
	CloudPCExternalPartnerSettingId string
}

// NewDeviceManagementVirtualEndpointExternalPartnerSettingID returns a new DeviceManagementVirtualEndpointExternalPartnerSettingId struct
func NewDeviceManagementVirtualEndpointExternalPartnerSettingID(cloudPCExternalPartnerSettingId string) DeviceManagementVirtualEndpointExternalPartnerSettingId {
	return DeviceManagementVirtualEndpointExternalPartnerSettingId{
		CloudPCExternalPartnerSettingId: cloudPCExternalPartnerSettingId,
	}
}

// ParseDeviceManagementVirtualEndpointExternalPartnerSettingID parses 'input' into a DeviceManagementVirtualEndpointExternalPartnerSettingId
func ParseDeviceManagementVirtualEndpointExternalPartnerSettingID(input string) (*DeviceManagementVirtualEndpointExternalPartnerSettingId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementVirtualEndpointExternalPartnerSettingId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementVirtualEndpointExternalPartnerSettingId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseDeviceManagementVirtualEndpointExternalPartnerSettingIDInsensitively parses 'input' case-insensitively into a DeviceManagementVirtualEndpointExternalPartnerSettingId
// note: this method should only be used for API response data and not user input
func ParseDeviceManagementVirtualEndpointExternalPartnerSettingIDInsensitively(input string) (*DeviceManagementVirtualEndpointExternalPartnerSettingId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementVirtualEndpointExternalPartnerSettingId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementVirtualEndpointExternalPartnerSettingId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *DeviceManagementVirtualEndpointExternalPartnerSettingId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.CloudPCExternalPartnerSettingId, ok = input.Parsed["cloudPCExternalPartnerSettingId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "cloudPCExternalPartnerSettingId", input)
	}

	return nil
}

// ValidateDeviceManagementVirtualEndpointExternalPartnerSettingID checks that 'input' can be parsed as a Device Management Virtual Endpoint External Partner Setting ID
func ValidateDeviceManagementVirtualEndpointExternalPartnerSettingID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDeviceManagementVirtualEndpointExternalPartnerSettingID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Device Management Virtual Endpoint External Partner Setting ID
func (id DeviceManagementVirtualEndpointExternalPartnerSettingId) ID() string {
	fmtString := "/deviceManagement/virtualEndpoint/externalPartnerSettings/%s"
	return fmt.Sprintf(fmtString, id.CloudPCExternalPartnerSettingId)
}

// Segments returns a slice of Resource ID Segments which comprise this Device Management Virtual Endpoint External Partner Setting ID
func (id DeviceManagementVirtualEndpointExternalPartnerSettingId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("deviceManagement", "deviceManagement", "deviceManagement"),
		resourceids.StaticSegment("virtualEndpoint", "virtualEndpoint", "virtualEndpoint"),
		resourceids.StaticSegment("externalPartnerSettings", "externalPartnerSettings", "externalPartnerSettings"),
		resourceids.UserSpecifiedSegment("cloudPCExternalPartnerSettingId", "cloudPCExternalPartnerSettingId"),
	}
}

// String returns a human-readable description of this Device Management Virtual Endpoint External Partner Setting ID
func (id DeviceManagementVirtualEndpointExternalPartnerSettingId) String() string {
	components := []string{
		fmt.Sprintf("Cloud PC External Partner Setting: %q", id.CloudPCExternalPartnerSettingId),
	}
	return fmt.Sprintf("Device Management Virtual Endpoint External Partner Setting (%s)", strings.Join(components, "\n"))
}
