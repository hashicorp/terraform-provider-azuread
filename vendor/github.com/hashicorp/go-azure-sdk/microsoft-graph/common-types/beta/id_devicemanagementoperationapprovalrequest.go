package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementOperationApprovalRequestId{}

// DeviceManagementOperationApprovalRequestId is a struct representing the Resource ID for a Device Management Operation Approval Request
type DeviceManagementOperationApprovalRequestId struct {
	OperationApprovalRequestId string
}

// NewDeviceManagementOperationApprovalRequestID returns a new DeviceManagementOperationApprovalRequestId struct
func NewDeviceManagementOperationApprovalRequestID(operationApprovalRequestId string) DeviceManagementOperationApprovalRequestId {
	return DeviceManagementOperationApprovalRequestId{
		OperationApprovalRequestId: operationApprovalRequestId,
	}
}

// ParseDeviceManagementOperationApprovalRequestID parses 'input' into a DeviceManagementOperationApprovalRequestId
func ParseDeviceManagementOperationApprovalRequestID(input string) (*DeviceManagementOperationApprovalRequestId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementOperationApprovalRequestId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementOperationApprovalRequestId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseDeviceManagementOperationApprovalRequestIDInsensitively parses 'input' case-insensitively into a DeviceManagementOperationApprovalRequestId
// note: this method should only be used for API response data and not user input
func ParseDeviceManagementOperationApprovalRequestIDInsensitively(input string) (*DeviceManagementOperationApprovalRequestId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementOperationApprovalRequestId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementOperationApprovalRequestId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *DeviceManagementOperationApprovalRequestId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.OperationApprovalRequestId, ok = input.Parsed["operationApprovalRequestId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "operationApprovalRequestId", input)
	}

	return nil
}

// ValidateDeviceManagementOperationApprovalRequestID checks that 'input' can be parsed as a Device Management Operation Approval Request ID
func ValidateDeviceManagementOperationApprovalRequestID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDeviceManagementOperationApprovalRequestID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Device Management Operation Approval Request ID
func (id DeviceManagementOperationApprovalRequestId) ID() string {
	fmtString := "/deviceManagement/operationApprovalRequests/%s"
	return fmt.Sprintf(fmtString, id.OperationApprovalRequestId)
}

// Segments returns a slice of Resource ID Segments which comprise this Device Management Operation Approval Request ID
func (id DeviceManagementOperationApprovalRequestId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("deviceManagement", "deviceManagement", "deviceManagement"),
		resourceids.StaticSegment("operationApprovalRequests", "operationApprovalRequests", "operationApprovalRequests"),
		resourceids.UserSpecifiedSegment("operationApprovalRequestId", "operationApprovalRequestId"),
	}
}

// String returns a human-readable description of this Device Management Operation Approval Request ID
func (id DeviceManagementOperationApprovalRequestId) String() string {
	components := []string{
		fmt.Sprintf("Operation Approval Request: %q", id.OperationApprovalRequestId),
	}
	return fmt.Sprintf("Device Management Operation Approval Request (%s)", strings.Join(components, "\n"))
}
