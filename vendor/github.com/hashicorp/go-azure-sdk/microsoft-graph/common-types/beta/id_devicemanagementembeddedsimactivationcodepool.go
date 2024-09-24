package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementEmbeddedSIMActivationCodePoolId{}

// DeviceManagementEmbeddedSIMActivationCodePoolId is a struct representing the Resource ID for a Device Management Embedded S I M Activation Code Pool
type DeviceManagementEmbeddedSIMActivationCodePoolId struct {
	EmbeddedSIMActivationCodePoolId string
}

// NewDeviceManagementEmbeddedSIMActivationCodePoolID returns a new DeviceManagementEmbeddedSIMActivationCodePoolId struct
func NewDeviceManagementEmbeddedSIMActivationCodePoolID(embeddedSIMActivationCodePoolId string) DeviceManagementEmbeddedSIMActivationCodePoolId {
	return DeviceManagementEmbeddedSIMActivationCodePoolId{
		EmbeddedSIMActivationCodePoolId: embeddedSIMActivationCodePoolId,
	}
}

// ParseDeviceManagementEmbeddedSIMActivationCodePoolID parses 'input' into a DeviceManagementEmbeddedSIMActivationCodePoolId
func ParseDeviceManagementEmbeddedSIMActivationCodePoolID(input string) (*DeviceManagementEmbeddedSIMActivationCodePoolId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementEmbeddedSIMActivationCodePoolId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementEmbeddedSIMActivationCodePoolId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseDeviceManagementEmbeddedSIMActivationCodePoolIDInsensitively parses 'input' case-insensitively into a DeviceManagementEmbeddedSIMActivationCodePoolId
// note: this method should only be used for API response data and not user input
func ParseDeviceManagementEmbeddedSIMActivationCodePoolIDInsensitively(input string) (*DeviceManagementEmbeddedSIMActivationCodePoolId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementEmbeddedSIMActivationCodePoolId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementEmbeddedSIMActivationCodePoolId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *DeviceManagementEmbeddedSIMActivationCodePoolId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.EmbeddedSIMActivationCodePoolId, ok = input.Parsed["embeddedSIMActivationCodePoolId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "embeddedSIMActivationCodePoolId", input)
	}

	return nil
}

// ValidateDeviceManagementEmbeddedSIMActivationCodePoolID checks that 'input' can be parsed as a Device Management Embedded S I M Activation Code Pool ID
func ValidateDeviceManagementEmbeddedSIMActivationCodePoolID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDeviceManagementEmbeddedSIMActivationCodePoolID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Device Management Embedded S I M Activation Code Pool ID
func (id DeviceManagementEmbeddedSIMActivationCodePoolId) ID() string {
	fmtString := "/deviceManagement/embeddedSIMActivationCodePools/%s"
	return fmt.Sprintf(fmtString, id.EmbeddedSIMActivationCodePoolId)
}

// Segments returns a slice of Resource ID Segments which comprise this Device Management Embedded S I M Activation Code Pool ID
func (id DeviceManagementEmbeddedSIMActivationCodePoolId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("deviceManagement", "deviceManagement", "deviceManagement"),
		resourceids.StaticSegment("embeddedSIMActivationCodePools", "embeddedSIMActivationCodePools", "embeddedSIMActivationCodePools"),
		resourceids.UserSpecifiedSegment("embeddedSIMActivationCodePoolId", "embeddedSIMActivationCodePoolId"),
	}
}

// String returns a human-readable description of this Device Management Embedded S I M Activation Code Pool ID
func (id DeviceManagementEmbeddedSIMActivationCodePoolId) String() string {
	components := []string{
		fmt.Sprintf("Embedded S I M Activation Code Pool: %q", id.EmbeddedSIMActivationCodePoolId),
	}
	return fmt.Sprintf("Device Management Embedded S I M Activation Code Pool (%s)", strings.Join(components, "\n"))
}
