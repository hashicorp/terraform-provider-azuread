package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeManagedDeviceIdLogCollectionRequestId{}

// MeManagedDeviceIdLogCollectionRequestId is a struct representing the Resource ID for a Me Managed Device Id Log Collection Request
type MeManagedDeviceIdLogCollectionRequestId struct {
	ManagedDeviceId               string
	DeviceLogCollectionResponseId string
}

// NewMeManagedDeviceIdLogCollectionRequestID returns a new MeManagedDeviceIdLogCollectionRequestId struct
func NewMeManagedDeviceIdLogCollectionRequestID(managedDeviceId string, deviceLogCollectionResponseId string) MeManagedDeviceIdLogCollectionRequestId {
	return MeManagedDeviceIdLogCollectionRequestId{
		ManagedDeviceId:               managedDeviceId,
		DeviceLogCollectionResponseId: deviceLogCollectionResponseId,
	}
}

// ParseMeManagedDeviceIdLogCollectionRequestID parses 'input' into a MeManagedDeviceIdLogCollectionRequestId
func ParseMeManagedDeviceIdLogCollectionRequestID(input string) (*MeManagedDeviceIdLogCollectionRequestId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeManagedDeviceIdLogCollectionRequestId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeManagedDeviceIdLogCollectionRequestId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseMeManagedDeviceIdLogCollectionRequestIDInsensitively parses 'input' case-insensitively into a MeManagedDeviceIdLogCollectionRequestId
// note: this method should only be used for API response data and not user input
func ParseMeManagedDeviceIdLogCollectionRequestIDInsensitively(input string) (*MeManagedDeviceIdLogCollectionRequestId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeManagedDeviceIdLogCollectionRequestId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeManagedDeviceIdLogCollectionRequestId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *MeManagedDeviceIdLogCollectionRequestId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.ManagedDeviceId, ok = input.Parsed["managedDeviceId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "managedDeviceId", input)
	}

	if id.DeviceLogCollectionResponseId, ok = input.Parsed["deviceLogCollectionResponseId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "deviceLogCollectionResponseId", input)
	}

	return nil
}

// ValidateMeManagedDeviceIdLogCollectionRequestID checks that 'input' can be parsed as a Me Managed Device Id Log Collection Request ID
func ValidateMeManagedDeviceIdLogCollectionRequestID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeManagedDeviceIdLogCollectionRequestID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Managed Device Id Log Collection Request ID
func (id MeManagedDeviceIdLogCollectionRequestId) ID() string {
	fmtString := "/me/managedDevices/%s/logCollectionRequests/%s"
	return fmt.Sprintf(fmtString, id.ManagedDeviceId, id.DeviceLogCollectionResponseId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Managed Device Id Log Collection Request ID
func (id MeManagedDeviceIdLogCollectionRequestId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("me", "me", "me"),
		resourceids.StaticSegment("managedDevices", "managedDevices", "managedDevices"),
		resourceids.UserSpecifiedSegment("managedDeviceId", "managedDeviceId"),
		resourceids.StaticSegment("logCollectionRequests", "logCollectionRequests", "logCollectionRequests"),
		resourceids.UserSpecifiedSegment("deviceLogCollectionResponseId", "deviceLogCollectionResponseId"),
	}
}

// String returns a human-readable description of this Me Managed Device Id Log Collection Request ID
func (id MeManagedDeviceIdLogCollectionRequestId) String() string {
	components := []string{
		fmt.Sprintf("Managed Device: %q", id.ManagedDeviceId),
		fmt.Sprintf("Device Log Collection Response: %q", id.DeviceLogCollectionResponseId),
	}
	return fmt.Sprintf("Me Managed Device Id Log Collection Request (%s)", strings.Join(components, "\n"))
}
