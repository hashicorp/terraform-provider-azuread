package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementEmbeddedSIMActivationCodePoolIdDeviceStateId{}

// DeviceManagementEmbeddedSIMActivationCodePoolIdDeviceStateId is a struct representing the Resource ID for a Device Management Embedded SIM Activation Code Pool Id Device State
type DeviceManagementEmbeddedSIMActivationCodePoolIdDeviceStateId struct {
	EmbeddedSIMActivationCodePoolId string
	EmbeddedSIMDeviceStateId        string
}

// NewDeviceManagementEmbeddedSIMActivationCodePoolIdDeviceStateID returns a new DeviceManagementEmbeddedSIMActivationCodePoolIdDeviceStateId struct
func NewDeviceManagementEmbeddedSIMActivationCodePoolIdDeviceStateID(embeddedSIMActivationCodePoolId string, embeddedSIMDeviceStateId string) DeviceManagementEmbeddedSIMActivationCodePoolIdDeviceStateId {
	return DeviceManagementEmbeddedSIMActivationCodePoolIdDeviceStateId{
		EmbeddedSIMActivationCodePoolId: embeddedSIMActivationCodePoolId,
		EmbeddedSIMDeviceStateId:        embeddedSIMDeviceStateId,
	}
}

// ParseDeviceManagementEmbeddedSIMActivationCodePoolIdDeviceStateID parses 'input' into a DeviceManagementEmbeddedSIMActivationCodePoolIdDeviceStateId
func ParseDeviceManagementEmbeddedSIMActivationCodePoolIdDeviceStateID(input string) (*DeviceManagementEmbeddedSIMActivationCodePoolIdDeviceStateId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementEmbeddedSIMActivationCodePoolIdDeviceStateId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementEmbeddedSIMActivationCodePoolIdDeviceStateId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseDeviceManagementEmbeddedSIMActivationCodePoolIdDeviceStateIDInsensitively parses 'input' case-insensitively into a DeviceManagementEmbeddedSIMActivationCodePoolIdDeviceStateId
// note: this method should only be used for API response data and not user input
func ParseDeviceManagementEmbeddedSIMActivationCodePoolIdDeviceStateIDInsensitively(input string) (*DeviceManagementEmbeddedSIMActivationCodePoolIdDeviceStateId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementEmbeddedSIMActivationCodePoolIdDeviceStateId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementEmbeddedSIMActivationCodePoolIdDeviceStateId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *DeviceManagementEmbeddedSIMActivationCodePoolIdDeviceStateId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.EmbeddedSIMActivationCodePoolId, ok = input.Parsed["embeddedSIMActivationCodePoolId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "embeddedSIMActivationCodePoolId", input)
	}

	if id.EmbeddedSIMDeviceStateId, ok = input.Parsed["embeddedSIMDeviceStateId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "embeddedSIMDeviceStateId", input)
	}

	return nil
}

// ValidateDeviceManagementEmbeddedSIMActivationCodePoolIdDeviceStateID checks that 'input' can be parsed as a Device Management Embedded SIM Activation Code Pool Id Device State ID
func ValidateDeviceManagementEmbeddedSIMActivationCodePoolIdDeviceStateID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDeviceManagementEmbeddedSIMActivationCodePoolIdDeviceStateID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Device Management Embedded SIM Activation Code Pool Id Device State ID
func (id DeviceManagementEmbeddedSIMActivationCodePoolIdDeviceStateId) ID() string {
	fmtString := "/deviceManagement/embeddedSIMActivationCodePools/%s/deviceStates/%s"
	return fmt.Sprintf(fmtString, id.EmbeddedSIMActivationCodePoolId, id.EmbeddedSIMDeviceStateId)
}

// Segments returns a slice of Resource ID Segments which comprise this Device Management Embedded SIM Activation Code Pool Id Device State ID
func (id DeviceManagementEmbeddedSIMActivationCodePoolIdDeviceStateId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("deviceManagement", "deviceManagement", "deviceManagement"),
		resourceids.StaticSegment("embeddedSIMActivationCodePools", "embeddedSIMActivationCodePools", "embeddedSIMActivationCodePools"),
		resourceids.UserSpecifiedSegment("embeddedSIMActivationCodePoolId", "embeddedSIMActivationCodePoolId"),
		resourceids.StaticSegment("deviceStates", "deviceStates", "deviceStates"),
		resourceids.UserSpecifiedSegment("embeddedSIMDeviceStateId", "embeddedSIMDeviceStateId"),
	}
}

// String returns a human-readable description of this Device Management Embedded SIM Activation Code Pool Id Device State ID
func (id DeviceManagementEmbeddedSIMActivationCodePoolIdDeviceStateId) String() string {
	components := []string{
		fmt.Sprintf("Embedded SIM Activation Code Pool: %q", id.EmbeddedSIMActivationCodePoolId),
		fmt.Sprintf("Embedded SIM Device State: %q", id.EmbeddedSIMDeviceStateId),
	}
	return fmt.Sprintf("Device Management Embedded SIM Activation Code Pool Id Device State (%s)", strings.Join(components, "\n"))
}
