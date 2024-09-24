package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementMobileThreatDefenseConnectorId{}

// DeviceManagementMobileThreatDefenseConnectorId is a struct representing the Resource ID for a Device Management Mobile Threat Defense Connector
type DeviceManagementMobileThreatDefenseConnectorId struct {
	MobileThreatDefenseConnectorId string
}

// NewDeviceManagementMobileThreatDefenseConnectorID returns a new DeviceManagementMobileThreatDefenseConnectorId struct
func NewDeviceManagementMobileThreatDefenseConnectorID(mobileThreatDefenseConnectorId string) DeviceManagementMobileThreatDefenseConnectorId {
	return DeviceManagementMobileThreatDefenseConnectorId{
		MobileThreatDefenseConnectorId: mobileThreatDefenseConnectorId,
	}
}

// ParseDeviceManagementMobileThreatDefenseConnectorID parses 'input' into a DeviceManagementMobileThreatDefenseConnectorId
func ParseDeviceManagementMobileThreatDefenseConnectorID(input string) (*DeviceManagementMobileThreatDefenseConnectorId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementMobileThreatDefenseConnectorId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementMobileThreatDefenseConnectorId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseDeviceManagementMobileThreatDefenseConnectorIDInsensitively parses 'input' case-insensitively into a DeviceManagementMobileThreatDefenseConnectorId
// note: this method should only be used for API response data and not user input
func ParseDeviceManagementMobileThreatDefenseConnectorIDInsensitively(input string) (*DeviceManagementMobileThreatDefenseConnectorId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementMobileThreatDefenseConnectorId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementMobileThreatDefenseConnectorId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *DeviceManagementMobileThreatDefenseConnectorId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.MobileThreatDefenseConnectorId, ok = input.Parsed["mobileThreatDefenseConnectorId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "mobileThreatDefenseConnectorId", input)
	}

	return nil
}

// ValidateDeviceManagementMobileThreatDefenseConnectorID checks that 'input' can be parsed as a Device Management Mobile Threat Defense Connector ID
func ValidateDeviceManagementMobileThreatDefenseConnectorID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDeviceManagementMobileThreatDefenseConnectorID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Device Management Mobile Threat Defense Connector ID
func (id DeviceManagementMobileThreatDefenseConnectorId) ID() string {
	fmtString := "/deviceManagement/mobileThreatDefenseConnectors/%s"
	return fmt.Sprintf(fmtString, id.MobileThreatDefenseConnectorId)
}

// Segments returns a slice of Resource ID Segments which comprise this Device Management Mobile Threat Defense Connector ID
func (id DeviceManagementMobileThreatDefenseConnectorId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("deviceManagement", "deviceManagement", "deviceManagement"),
		resourceids.StaticSegment("mobileThreatDefenseConnectors", "mobileThreatDefenseConnectors", "mobileThreatDefenseConnectors"),
		resourceids.UserSpecifiedSegment("mobileThreatDefenseConnectorId", "mobileThreatDefenseConnectorId"),
	}
}

// String returns a human-readable description of this Device Management Mobile Threat Defense Connector ID
func (id DeviceManagementMobileThreatDefenseConnectorId) String() string {
	components := []string{
		fmt.Sprintf("Mobile Threat Defense Connector: %q", id.MobileThreatDefenseConnectorId),
	}
	return fmt.Sprintf("Device Management Mobile Threat Defense Connector (%s)", strings.Join(components, "\n"))
}
