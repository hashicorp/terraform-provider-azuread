package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementNotificationMessageTemplateId{}

// DeviceManagementNotificationMessageTemplateId is a struct representing the Resource ID for a Device Management Notification Message Template
type DeviceManagementNotificationMessageTemplateId struct {
	NotificationMessageTemplateId string
}

// NewDeviceManagementNotificationMessageTemplateID returns a new DeviceManagementNotificationMessageTemplateId struct
func NewDeviceManagementNotificationMessageTemplateID(notificationMessageTemplateId string) DeviceManagementNotificationMessageTemplateId {
	return DeviceManagementNotificationMessageTemplateId{
		NotificationMessageTemplateId: notificationMessageTemplateId,
	}
}

// ParseDeviceManagementNotificationMessageTemplateID parses 'input' into a DeviceManagementNotificationMessageTemplateId
func ParseDeviceManagementNotificationMessageTemplateID(input string) (*DeviceManagementNotificationMessageTemplateId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementNotificationMessageTemplateId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementNotificationMessageTemplateId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseDeviceManagementNotificationMessageTemplateIDInsensitively parses 'input' case-insensitively into a DeviceManagementNotificationMessageTemplateId
// note: this method should only be used for API response data and not user input
func ParseDeviceManagementNotificationMessageTemplateIDInsensitively(input string) (*DeviceManagementNotificationMessageTemplateId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementNotificationMessageTemplateId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementNotificationMessageTemplateId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *DeviceManagementNotificationMessageTemplateId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.NotificationMessageTemplateId, ok = input.Parsed["notificationMessageTemplateId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "notificationMessageTemplateId", input)
	}

	return nil
}

// ValidateDeviceManagementNotificationMessageTemplateID checks that 'input' can be parsed as a Device Management Notification Message Template ID
func ValidateDeviceManagementNotificationMessageTemplateID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDeviceManagementNotificationMessageTemplateID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Device Management Notification Message Template ID
func (id DeviceManagementNotificationMessageTemplateId) ID() string {
	fmtString := "/deviceManagement/notificationMessageTemplates/%s"
	return fmt.Sprintf(fmtString, id.NotificationMessageTemplateId)
}

// Segments returns a slice of Resource ID Segments which comprise this Device Management Notification Message Template ID
func (id DeviceManagementNotificationMessageTemplateId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("deviceManagement", "deviceManagement", "deviceManagement"),
		resourceids.StaticSegment("notificationMessageTemplates", "notificationMessageTemplates", "notificationMessageTemplates"),
		resourceids.UserSpecifiedSegment("notificationMessageTemplateId", "notificationMessageTemplateId"),
	}
}

// String returns a human-readable description of this Device Management Notification Message Template ID
func (id DeviceManagementNotificationMessageTemplateId) String() string {
	components := []string{
		fmt.Sprintf("Notification Message Template: %q", id.NotificationMessageTemplateId),
	}
	return fmt.Sprintf("Device Management Notification Message Template (%s)", strings.Join(components, "\n"))
}
