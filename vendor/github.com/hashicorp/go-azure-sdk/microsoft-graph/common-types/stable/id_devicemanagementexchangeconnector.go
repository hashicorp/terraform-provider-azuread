package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementExchangeConnectorId{}

// DeviceManagementExchangeConnectorId is a struct representing the Resource ID for a Device Management Exchange Connector
type DeviceManagementExchangeConnectorId struct {
	DeviceManagementExchangeConnectorId string
}

// NewDeviceManagementExchangeConnectorID returns a new DeviceManagementExchangeConnectorId struct
func NewDeviceManagementExchangeConnectorID(deviceManagementExchangeConnectorId string) DeviceManagementExchangeConnectorId {
	return DeviceManagementExchangeConnectorId{
		DeviceManagementExchangeConnectorId: deviceManagementExchangeConnectorId,
	}
}

// ParseDeviceManagementExchangeConnectorID parses 'input' into a DeviceManagementExchangeConnectorId
func ParseDeviceManagementExchangeConnectorID(input string) (*DeviceManagementExchangeConnectorId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementExchangeConnectorId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementExchangeConnectorId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseDeviceManagementExchangeConnectorIDInsensitively parses 'input' case-insensitively into a DeviceManagementExchangeConnectorId
// note: this method should only be used for API response data and not user input
func ParseDeviceManagementExchangeConnectorIDInsensitively(input string) (*DeviceManagementExchangeConnectorId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementExchangeConnectorId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementExchangeConnectorId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *DeviceManagementExchangeConnectorId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.DeviceManagementExchangeConnectorId, ok = input.Parsed["deviceManagementExchangeConnectorId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "deviceManagementExchangeConnectorId", input)
	}

	return nil
}

// ValidateDeviceManagementExchangeConnectorID checks that 'input' can be parsed as a Device Management Exchange Connector ID
func ValidateDeviceManagementExchangeConnectorID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDeviceManagementExchangeConnectorID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Device Management Exchange Connector ID
func (id DeviceManagementExchangeConnectorId) ID() string {
	fmtString := "/deviceManagement/exchangeConnectors/%s"
	return fmt.Sprintf(fmtString, id.DeviceManagementExchangeConnectorId)
}

// Segments returns a slice of Resource ID Segments which comprise this Device Management Exchange Connector ID
func (id DeviceManagementExchangeConnectorId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("deviceManagement", "deviceManagement", "deviceManagement"),
		resourceids.StaticSegment("exchangeConnectors", "exchangeConnectors", "exchangeConnectors"),
		resourceids.UserSpecifiedSegment("deviceManagementExchangeConnectorId", "deviceManagementExchangeConnectorId"),
	}
}

// String returns a human-readable description of this Device Management Exchange Connector ID
func (id DeviceManagementExchangeConnectorId) String() string {
	components := []string{
		fmt.Sprintf("Device Management Exchange Connector: %q", id.DeviceManagementExchangeConnectorId),
	}
	return fmt.Sprintf("Device Management Exchange Connector (%s)", strings.Join(components, "\n"))
}
