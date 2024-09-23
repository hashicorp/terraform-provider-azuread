package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementDomainJoinConnectorId{}

// DeviceManagementDomainJoinConnectorId is a struct representing the Resource ID for a Device Management Domain Join Connector
type DeviceManagementDomainJoinConnectorId struct {
	DeviceManagementDomainJoinConnectorId string
}

// NewDeviceManagementDomainJoinConnectorID returns a new DeviceManagementDomainJoinConnectorId struct
func NewDeviceManagementDomainJoinConnectorID(deviceManagementDomainJoinConnectorId string) DeviceManagementDomainJoinConnectorId {
	return DeviceManagementDomainJoinConnectorId{
		DeviceManagementDomainJoinConnectorId: deviceManagementDomainJoinConnectorId,
	}
}

// ParseDeviceManagementDomainJoinConnectorID parses 'input' into a DeviceManagementDomainJoinConnectorId
func ParseDeviceManagementDomainJoinConnectorID(input string) (*DeviceManagementDomainJoinConnectorId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementDomainJoinConnectorId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementDomainJoinConnectorId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseDeviceManagementDomainJoinConnectorIDInsensitively parses 'input' case-insensitively into a DeviceManagementDomainJoinConnectorId
// note: this method should only be used for API response data and not user input
func ParseDeviceManagementDomainJoinConnectorIDInsensitively(input string) (*DeviceManagementDomainJoinConnectorId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementDomainJoinConnectorId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementDomainJoinConnectorId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *DeviceManagementDomainJoinConnectorId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.DeviceManagementDomainJoinConnectorId, ok = input.Parsed["deviceManagementDomainJoinConnectorId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "deviceManagementDomainJoinConnectorId", input)
	}

	return nil
}

// ValidateDeviceManagementDomainJoinConnectorID checks that 'input' can be parsed as a Device Management Domain Join Connector ID
func ValidateDeviceManagementDomainJoinConnectorID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDeviceManagementDomainJoinConnectorID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Device Management Domain Join Connector ID
func (id DeviceManagementDomainJoinConnectorId) ID() string {
	fmtString := "/deviceManagement/domainJoinConnectors/%s"
	return fmt.Sprintf(fmtString, id.DeviceManagementDomainJoinConnectorId)
}

// Segments returns a slice of Resource ID Segments which comprise this Device Management Domain Join Connector ID
func (id DeviceManagementDomainJoinConnectorId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("deviceManagement", "deviceManagement", "deviceManagement"),
		resourceids.StaticSegment("domainJoinConnectors", "domainJoinConnectors", "domainJoinConnectors"),
		resourceids.UserSpecifiedSegment("deviceManagementDomainJoinConnectorId", "deviceManagementDomainJoinConnectorId"),
	}
}

// String returns a human-readable description of this Device Management Domain Join Connector ID
func (id DeviceManagementDomainJoinConnectorId) String() string {
	components := []string{
		fmt.Sprintf("Device Management Domain Join Connector: %q", id.DeviceManagementDomainJoinConnectorId),
	}
	return fmt.Sprintf("Device Management Domain Join Connector (%s)", strings.Join(components, "\n"))
}
