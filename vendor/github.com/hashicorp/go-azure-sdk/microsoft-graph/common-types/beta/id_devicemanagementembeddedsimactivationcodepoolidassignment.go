package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementEmbeddedSIMActivationCodePoolIdAssignmentId{}

// DeviceManagementEmbeddedSIMActivationCodePoolIdAssignmentId is a struct representing the Resource ID for a Device Management Embedded SIM Activation Code Pool Id Assignment
type DeviceManagementEmbeddedSIMActivationCodePoolIdAssignmentId struct {
	EmbeddedSIMActivationCodePoolId           string
	EmbeddedSIMActivationCodePoolAssignmentId string
}

// NewDeviceManagementEmbeddedSIMActivationCodePoolIdAssignmentID returns a new DeviceManagementEmbeddedSIMActivationCodePoolIdAssignmentId struct
func NewDeviceManagementEmbeddedSIMActivationCodePoolIdAssignmentID(embeddedSIMActivationCodePoolId string, embeddedSIMActivationCodePoolAssignmentId string) DeviceManagementEmbeddedSIMActivationCodePoolIdAssignmentId {
	return DeviceManagementEmbeddedSIMActivationCodePoolIdAssignmentId{
		EmbeddedSIMActivationCodePoolId:           embeddedSIMActivationCodePoolId,
		EmbeddedSIMActivationCodePoolAssignmentId: embeddedSIMActivationCodePoolAssignmentId,
	}
}

// ParseDeviceManagementEmbeddedSIMActivationCodePoolIdAssignmentID parses 'input' into a DeviceManagementEmbeddedSIMActivationCodePoolIdAssignmentId
func ParseDeviceManagementEmbeddedSIMActivationCodePoolIdAssignmentID(input string) (*DeviceManagementEmbeddedSIMActivationCodePoolIdAssignmentId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementEmbeddedSIMActivationCodePoolIdAssignmentId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementEmbeddedSIMActivationCodePoolIdAssignmentId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseDeviceManagementEmbeddedSIMActivationCodePoolIdAssignmentIDInsensitively parses 'input' case-insensitively into a DeviceManagementEmbeddedSIMActivationCodePoolIdAssignmentId
// note: this method should only be used for API response data and not user input
func ParseDeviceManagementEmbeddedSIMActivationCodePoolIdAssignmentIDInsensitively(input string) (*DeviceManagementEmbeddedSIMActivationCodePoolIdAssignmentId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementEmbeddedSIMActivationCodePoolIdAssignmentId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementEmbeddedSIMActivationCodePoolIdAssignmentId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *DeviceManagementEmbeddedSIMActivationCodePoolIdAssignmentId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.EmbeddedSIMActivationCodePoolId, ok = input.Parsed["embeddedSIMActivationCodePoolId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "embeddedSIMActivationCodePoolId", input)
	}

	if id.EmbeddedSIMActivationCodePoolAssignmentId, ok = input.Parsed["embeddedSIMActivationCodePoolAssignmentId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "embeddedSIMActivationCodePoolAssignmentId", input)
	}

	return nil
}

// ValidateDeviceManagementEmbeddedSIMActivationCodePoolIdAssignmentID checks that 'input' can be parsed as a Device Management Embedded SIM Activation Code Pool Id Assignment ID
func ValidateDeviceManagementEmbeddedSIMActivationCodePoolIdAssignmentID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDeviceManagementEmbeddedSIMActivationCodePoolIdAssignmentID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Device Management Embedded SIM Activation Code Pool Id Assignment ID
func (id DeviceManagementEmbeddedSIMActivationCodePoolIdAssignmentId) ID() string {
	fmtString := "/deviceManagement/embeddedSIMActivationCodePools/%s/assignments/%s"
	return fmt.Sprintf(fmtString, id.EmbeddedSIMActivationCodePoolId, id.EmbeddedSIMActivationCodePoolAssignmentId)
}

// Segments returns a slice of Resource ID Segments which comprise this Device Management Embedded SIM Activation Code Pool Id Assignment ID
func (id DeviceManagementEmbeddedSIMActivationCodePoolIdAssignmentId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("deviceManagement", "deviceManagement", "deviceManagement"),
		resourceids.StaticSegment("embeddedSIMActivationCodePools", "embeddedSIMActivationCodePools", "embeddedSIMActivationCodePools"),
		resourceids.UserSpecifiedSegment("embeddedSIMActivationCodePoolId", "embeddedSIMActivationCodePoolId"),
		resourceids.StaticSegment("assignments", "assignments", "assignments"),
		resourceids.UserSpecifiedSegment("embeddedSIMActivationCodePoolAssignmentId", "embeddedSIMActivationCodePoolAssignmentId"),
	}
}

// String returns a human-readable description of this Device Management Embedded SIM Activation Code Pool Id Assignment ID
func (id DeviceManagementEmbeddedSIMActivationCodePoolIdAssignmentId) String() string {
	components := []string{
		fmt.Sprintf("Embedded SIM Activation Code Pool: %q", id.EmbeddedSIMActivationCodePoolId),
		fmt.Sprintf("Embedded SIM Activation Code Pool Assignment: %q", id.EmbeddedSIMActivationCodePoolAssignmentId),
	}
	return fmt.Sprintf("Device Management Embedded SIM Activation Code Pool Id Assignment (%s)", strings.Join(components, "\n"))
}
