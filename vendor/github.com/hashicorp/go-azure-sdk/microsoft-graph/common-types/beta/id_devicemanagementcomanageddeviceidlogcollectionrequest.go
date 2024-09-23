package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementComanagedDeviceIdLogCollectionRequestId{}

// DeviceManagementComanagedDeviceIdLogCollectionRequestId is a struct representing the Resource ID for a Device Management Comanaged Device Id Log Collection Request
type DeviceManagementComanagedDeviceIdLogCollectionRequestId struct {
	ManagedDeviceId               string
	DeviceLogCollectionResponseId string
}

// NewDeviceManagementComanagedDeviceIdLogCollectionRequestID returns a new DeviceManagementComanagedDeviceIdLogCollectionRequestId struct
func NewDeviceManagementComanagedDeviceIdLogCollectionRequestID(managedDeviceId string, deviceLogCollectionResponseId string) DeviceManagementComanagedDeviceIdLogCollectionRequestId {
	return DeviceManagementComanagedDeviceIdLogCollectionRequestId{
		ManagedDeviceId:               managedDeviceId,
		DeviceLogCollectionResponseId: deviceLogCollectionResponseId,
	}
}

// ParseDeviceManagementComanagedDeviceIdLogCollectionRequestID parses 'input' into a DeviceManagementComanagedDeviceIdLogCollectionRequestId
func ParseDeviceManagementComanagedDeviceIdLogCollectionRequestID(input string) (*DeviceManagementComanagedDeviceIdLogCollectionRequestId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementComanagedDeviceIdLogCollectionRequestId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementComanagedDeviceIdLogCollectionRequestId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseDeviceManagementComanagedDeviceIdLogCollectionRequestIDInsensitively parses 'input' case-insensitively into a DeviceManagementComanagedDeviceIdLogCollectionRequestId
// note: this method should only be used for API response data and not user input
func ParseDeviceManagementComanagedDeviceIdLogCollectionRequestIDInsensitively(input string) (*DeviceManagementComanagedDeviceIdLogCollectionRequestId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementComanagedDeviceIdLogCollectionRequestId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementComanagedDeviceIdLogCollectionRequestId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *DeviceManagementComanagedDeviceIdLogCollectionRequestId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.ManagedDeviceId, ok = input.Parsed["managedDeviceId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "managedDeviceId", input)
	}

	if id.DeviceLogCollectionResponseId, ok = input.Parsed["deviceLogCollectionResponseId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "deviceLogCollectionResponseId", input)
	}

	return nil
}

// ValidateDeviceManagementComanagedDeviceIdLogCollectionRequestID checks that 'input' can be parsed as a Device Management Comanaged Device Id Log Collection Request ID
func ValidateDeviceManagementComanagedDeviceIdLogCollectionRequestID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDeviceManagementComanagedDeviceIdLogCollectionRequestID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Device Management Comanaged Device Id Log Collection Request ID
func (id DeviceManagementComanagedDeviceIdLogCollectionRequestId) ID() string {
	fmtString := "/deviceManagement/comanagedDevices/%s/logCollectionRequests/%s"
	return fmt.Sprintf(fmtString, id.ManagedDeviceId, id.DeviceLogCollectionResponseId)
}

// Segments returns a slice of Resource ID Segments which comprise this Device Management Comanaged Device Id Log Collection Request ID
func (id DeviceManagementComanagedDeviceIdLogCollectionRequestId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("deviceManagement", "deviceManagement", "deviceManagement"),
		resourceids.StaticSegment("comanagedDevices", "comanagedDevices", "comanagedDevices"),
		resourceids.UserSpecifiedSegment("managedDeviceId", "managedDeviceId"),
		resourceids.StaticSegment("logCollectionRequests", "logCollectionRequests", "logCollectionRequests"),
		resourceids.UserSpecifiedSegment("deviceLogCollectionResponseId", "deviceLogCollectionResponseId"),
	}
}

// String returns a human-readable description of this Device Management Comanaged Device Id Log Collection Request ID
func (id DeviceManagementComanagedDeviceIdLogCollectionRequestId) String() string {
	components := []string{
		fmt.Sprintf("Managed Device: %q", id.ManagedDeviceId),
		fmt.Sprintf("Device Log Collection Response: %q", id.DeviceLogCollectionResponseId),
	}
	return fmt.Sprintf("Device Management Comanaged Device Id Log Collection Request (%s)", strings.Join(components, "\n"))
}
