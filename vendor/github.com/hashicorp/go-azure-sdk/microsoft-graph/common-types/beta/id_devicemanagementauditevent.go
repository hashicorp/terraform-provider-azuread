package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementAuditEventId{}

// DeviceManagementAuditEventId is a struct representing the Resource ID for a Device Management Audit Event
type DeviceManagementAuditEventId struct {
	AuditEventId string
}

// NewDeviceManagementAuditEventID returns a new DeviceManagementAuditEventId struct
func NewDeviceManagementAuditEventID(auditEventId string) DeviceManagementAuditEventId {
	return DeviceManagementAuditEventId{
		AuditEventId: auditEventId,
	}
}

// ParseDeviceManagementAuditEventID parses 'input' into a DeviceManagementAuditEventId
func ParseDeviceManagementAuditEventID(input string) (*DeviceManagementAuditEventId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementAuditEventId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementAuditEventId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseDeviceManagementAuditEventIDInsensitively parses 'input' case-insensitively into a DeviceManagementAuditEventId
// note: this method should only be used for API response data and not user input
func ParseDeviceManagementAuditEventIDInsensitively(input string) (*DeviceManagementAuditEventId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementAuditEventId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementAuditEventId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *DeviceManagementAuditEventId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.AuditEventId, ok = input.Parsed["auditEventId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "auditEventId", input)
	}

	return nil
}

// ValidateDeviceManagementAuditEventID checks that 'input' can be parsed as a Device Management Audit Event ID
func ValidateDeviceManagementAuditEventID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDeviceManagementAuditEventID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Device Management Audit Event ID
func (id DeviceManagementAuditEventId) ID() string {
	fmtString := "/deviceManagement/auditEvents/%s"
	return fmt.Sprintf(fmtString, id.AuditEventId)
}

// Segments returns a slice of Resource ID Segments which comprise this Device Management Audit Event ID
func (id DeviceManagementAuditEventId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("deviceManagement", "deviceManagement", "deviceManagement"),
		resourceids.StaticSegment("auditEvents", "auditEvents", "auditEvents"),
		resourceids.UserSpecifiedSegment("auditEventId", "auditEventId"),
	}
}

// String returns a human-readable description of this Device Management Audit Event ID
func (id DeviceManagementAuditEventId) String() string {
	components := []string{
		fmt.Sprintf("Audit Event: %q", id.AuditEventId),
	}
	return fmt.Sprintf("Device Management Audit Event (%s)", strings.Join(components, "\n"))
}
