package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementRemoteActionAuditId{}

// DeviceManagementRemoteActionAuditId is a struct representing the Resource ID for a Device Management Remote Action Audit
type DeviceManagementRemoteActionAuditId struct {
	RemoteActionAuditId string
}

// NewDeviceManagementRemoteActionAuditID returns a new DeviceManagementRemoteActionAuditId struct
func NewDeviceManagementRemoteActionAuditID(remoteActionAuditId string) DeviceManagementRemoteActionAuditId {
	return DeviceManagementRemoteActionAuditId{
		RemoteActionAuditId: remoteActionAuditId,
	}
}

// ParseDeviceManagementRemoteActionAuditID parses 'input' into a DeviceManagementRemoteActionAuditId
func ParseDeviceManagementRemoteActionAuditID(input string) (*DeviceManagementRemoteActionAuditId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementRemoteActionAuditId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementRemoteActionAuditId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseDeviceManagementRemoteActionAuditIDInsensitively parses 'input' case-insensitively into a DeviceManagementRemoteActionAuditId
// note: this method should only be used for API response data and not user input
func ParseDeviceManagementRemoteActionAuditIDInsensitively(input string) (*DeviceManagementRemoteActionAuditId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementRemoteActionAuditId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementRemoteActionAuditId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *DeviceManagementRemoteActionAuditId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.RemoteActionAuditId, ok = input.Parsed["remoteActionAuditId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "remoteActionAuditId", input)
	}

	return nil
}

// ValidateDeviceManagementRemoteActionAuditID checks that 'input' can be parsed as a Device Management Remote Action Audit ID
func ValidateDeviceManagementRemoteActionAuditID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDeviceManagementRemoteActionAuditID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Device Management Remote Action Audit ID
func (id DeviceManagementRemoteActionAuditId) ID() string {
	fmtString := "/deviceManagement/remoteActionAudits/%s"
	return fmt.Sprintf(fmtString, id.RemoteActionAuditId)
}

// Segments returns a slice of Resource ID Segments which comprise this Device Management Remote Action Audit ID
func (id DeviceManagementRemoteActionAuditId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("deviceManagement", "deviceManagement", "deviceManagement"),
		resourceids.StaticSegment("remoteActionAudits", "remoteActionAudits", "remoteActionAudits"),
		resourceids.UserSpecifiedSegment("remoteActionAuditId", "remoteActionAuditId"),
	}
}

// String returns a human-readable description of this Device Management Remote Action Audit ID
func (id DeviceManagementRemoteActionAuditId) String() string {
	components := []string{
		fmt.Sprintf("Remote Action Audit: %q", id.RemoteActionAuditId),
	}
	return fmt.Sprintf("Device Management Remote Action Audit (%s)", strings.Join(components, "\n"))
}
