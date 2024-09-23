package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementManagedDeviceIdLogCollectionRequestId{}

// DeviceManagementManagedDeviceIdLogCollectionRequestId is a struct representing the Resource ID for a Device Management Managed Device Id Log Collection Request
type DeviceManagementManagedDeviceIdLogCollectionRequestId struct {
	ManagedDeviceId               string
	DeviceLogCollectionResponseId string
}

// NewDeviceManagementManagedDeviceIdLogCollectionRequestID returns a new DeviceManagementManagedDeviceIdLogCollectionRequestId struct
func NewDeviceManagementManagedDeviceIdLogCollectionRequestID(managedDeviceId string, deviceLogCollectionResponseId string) DeviceManagementManagedDeviceIdLogCollectionRequestId {
	return DeviceManagementManagedDeviceIdLogCollectionRequestId{
		ManagedDeviceId:               managedDeviceId,
		DeviceLogCollectionResponseId: deviceLogCollectionResponseId,
	}
}

// ParseDeviceManagementManagedDeviceIdLogCollectionRequestID parses 'input' into a DeviceManagementManagedDeviceIdLogCollectionRequestId
func ParseDeviceManagementManagedDeviceIdLogCollectionRequestID(input string) (*DeviceManagementManagedDeviceIdLogCollectionRequestId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementManagedDeviceIdLogCollectionRequestId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementManagedDeviceIdLogCollectionRequestId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseDeviceManagementManagedDeviceIdLogCollectionRequestIDInsensitively parses 'input' case-insensitively into a DeviceManagementManagedDeviceIdLogCollectionRequestId
// note: this method should only be used for API response data and not user input
func ParseDeviceManagementManagedDeviceIdLogCollectionRequestIDInsensitively(input string) (*DeviceManagementManagedDeviceIdLogCollectionRequestId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementManagedDeviceIdLogCollectionRequestId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementManagedDeviceIdLogCollectionRequestId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *DeviceManagementManagedDeviceIdLogCollectionRequestId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.ManagedDeviceId, ok = input.Parsed["managedDeviceId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "managedDeviceId", input)
	}

	if id.DeviceLogCollectionResponseId, ok = input.Parsed["deviceLogCollectionResponseId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "deviceLogCollectionResponseId", input)
	}

	return nil
}

// ValidateDeviceManagementManagedDeviceIdLogCollectionRequestID checks that 'input' can be parsed as a Device Management Managed Device Id Log Collection Request ID
func ValidateDeviceManagementManagedDeviceIdLogCollectionRequestID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDeviceManagementManagedDeviceIdLogCollectionRequestID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Device Management Managed Device Id Log Collection Request ID
func (id DeviceManagementManagedDeviceIdLogCollectionRequestId) ID() string {
	fmtString := "/deviceManagement/managedDevices/%s/logCollectionRequests/%s"
	return fmt.Sprintf(fmtString, id.ManagedDeviceId, id.DeviceLogCollectionResponseId)
}

// Segments returns a slice of Resource ID Segments which comprise this Device Management Managed Device Id Log Collection Request ID
func (id DeviceManagementManagedDeviceIdLogCollectionRequestId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("deviceManagement", "deviceManagement", "deviceManagement"),
		resourceids.StaticSegment("managedDevices", "managedDevices", "managedDevices"),
		resourceids.UserSpecifiedSegment("managedDeviceId", "managedDeviceId"),
		resourceids.StaticSegment("logCollectionRequests", "logCollectionRequests", "logCollectionRequests"),
		resourceids.UserSpecifiedSegment("deviceLogCollectionResponseId", "deviceLogCollectionResponseId"),
	}
}

// String returns a human-readable description of this Device Management Managed Device Id Log Collection Request ID
func (id DeviceManagementManagedDeviceIdLogCollectionRequestId) String() string {
	components := []string{
		fmt.Sprintf("Managed Device: %q", id.ManagedDeviceId),
		fmt.Sprintf("Device Log Collection Response: %q", id.DeviceLogCollectionResponseId),
	}
	return fmt.Sprintf("Device Management Managed Device Id Log Collection Request (%s)", strings.Join(components, "\n"))
}
