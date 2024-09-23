package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementNdesConnectorId{}

// DeviceManagementNdesConnectorId is a struct representing the Resource ID for a Device Management Ndes Connector
type DeviceManagementNdesConnectorId struct {
	NdesConnectorId string
}

// NewDeviceManagementNdesConnectorID returns a new DeviceManagementNdesConnectorId struct
func NewDeviceManagementNdesConnectorID(ndesConnectorId string) DeviceManagementNdesConnectorId {
	return DeviceManagementNdesConnectorId{
		NdesConnectorId: ndesConnectorId,
	}
}

// ParseDeviceManagementNdesConnectorID parses 'input' into a DeviceManagementNdesConnectorId
func ParseDeviceManagementNdesConnectorID(input string) (*DeviceManagementNdesConnectorId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementNdesConnectorId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementNdesConnectorId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseDeviceManagementNdesConnectorIDInsensitively parses 'input' case-insensitively into a DeviceManagementNdesConnectorId
// note: this method should only be used for API response data and not user input
func ParseDeviceManagementNdesConnectorIDInsensitively(input string) (*DeviceManagementNdesConnectorId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementNdesConnectorId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementNdesConnectorId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *DeviceManagementNdesConnectorId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.NdesConnectorId, ok = input.Parsed["ndesConnectorId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "ndesConnectorId", input)
	}

	return nil
}

// ValidateDeviceManagementNdesConnectorID checks that 'input' can be parsed as a Device Management Ndes Connector ID
func ValidateDeviceManagementNdesConnectorID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDeviceManagementNdesConnectorID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Device Management Ndes Connector ID
func (id DeviceManagementNdesConnectorId) ID() string {
	fmtString := "/deviceManagement/ndesConnectors/%s"
	return fmt.Sprintf(fmtString, id.NdesConnectorId)
}

// Segments returns a slice of Resource ID Segments which comprise this Device Management Ndes Connector ID
func (id DeviceManagementNdesConnectorId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("deviceManagement", "deviceManagement", "deviceManagement"),
		resourceids.StaticSegment("ndesConnectors", "ndesConnectors", "ndesConnectors"),
		resourceids.UserSpecifiedSegment("ndesConnectorId", "ndesConnectorId"),
	}
}

// String returns a human-readable description of this Device Management Ndes Connector ID
func (id DeviceManagementNdesConnectorId) String() string {
	components := []string{
		fmt.Sprintf("Ndes Connector: %q", id.NdesConnectorId),
	}
	return fmt.Sprintf("Device Management Ndes Connector (%s)", strings.Join(components, "\n"))
}
