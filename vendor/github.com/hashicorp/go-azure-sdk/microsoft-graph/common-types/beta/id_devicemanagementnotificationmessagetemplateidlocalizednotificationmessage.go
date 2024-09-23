package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementNotificationMessageTemplateIdLocalizedNotificationMessageId{}

// DeviceManagementNotificationMessageTemplateIdLocalizedNotificationMessageId is a struct representing the Resource ID for a Device Management Notification Message Template Id Localized Notification Message
type DeviceManagementNotificationMessageTemplateIdLocalizedNotificationMessageId struct {
	NotificationMessageTemplateId  string
	LocalizedNotificationMessageId string
}

// NewDeviceManagementNotificationMessageTemplateIdLocalizedNotificationMessageID returns a new DeviceManagementNotificationMessageTemplateIdLocalizedNotificationMessageId struct
func NewDeviceManagementNotificationMessageTemplateIdLocalizedNotificationMessageID(notificationMessageTemplateId string, localizedNotificationMessageId string) DeviceManagementNotificationMessageTemplateIdLocalizedNotificationMessageId {
	return DeviceManagementNotificationMessageTemplateIdLocalizedNotificationMessageId{
		NotificationMessageTemplateId:  notificationMessageTemplateId,
		LocalizedNotificationMessageId: localizedNotificationMessageId,
	}
}

// ParseDeviceManagementNotificationMessageTemplateIdLocalizedNotificationMessageID parses 'input' into a DeviceManagementNotificationMessageTemplateIdLocalizedNotificationMessageId
func ParseDeviceManagementNotificationMessageTemplateIdLocalizedNotificationMessageID(input string) (*DeviceManagementNotificationMessageTemplateIdLocalizedNotificationMessageId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementNotificationMessageTemplateIdLocalizedNotificationMessageId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementNotificationMessageTemplateIdLocalizedNotificationMessageId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseDeviceManagementNotificationMessageTemplateIdLocalizedNotificationMessageIDInsensitively parses 'input' case-insensitively into a DeviceManagementNotificationMessageTemplateIdLocalizedNotificationMessageId
// note: this method should only be used for API response data and not user input
func ParseDeviceManagementNotificationMessageTemplateIdLocalizedNotificationMessageIDInsensitively(input string) (*DeviceManagementNotificationMessageTemplateIdLocalizedNotificationMessageId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementNotificationMessageTemplateIdLocalizedNotificationMessageId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementNotificationMessageTemplateIdLocalizedNotificationMessageId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *DeviceManagementNotificationMessageTemplateIdLocalizedNotificationMessageId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.NotificationMessageTemplateId, ok = input.Parsed["notificationMessageTemplateId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "notificationMessageTemplateId", input)
	}

	if id.LocalizedNotificationMessageId, ok = input.Parsed["localizedNotificationMessageId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "localizedNotificationMessageId", input)
	}

	return nil
}

// ValidateDeviceManagementNotificationMessageTemplateIdLocalizedNotificationMessageID checks that 'input' can be parsed as a Device Management Notification Message Template Id Localized Notification Message ID
func ValidateDeviceManagementNotificationMessageTemplateIdLocalizedNotificationMessageID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDeviceManagementNotificationMessageTemplateIdLocalizedNotificationMessageID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Device Management Notification Message Template Id Localized Notification Message ID
func (id DeviceManagementNotificationMessageTemplateIdLocalizedNotificationMessageId) ID() string {
	fmtString := "/deviceManagement/notificationMessageTemplates/%s/localizedNotificationMessages/%s"
	return fmt.Sprintf(fmtString, id.NotificationMessageTemplateId, id.LocalizedNotificationMessageId)
}

// Segments returns a slice of Resource ID Segments which comprise this Device Management Notification Message Template Id Localized Notification Message ID
func (id DeviceManagementNotificationMessageTemplateIdLocalizedNotificationMessageId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("deviceManagement", "deviceManagement", "deviceManagement"),
		resourceids.StaticSegment("notificationMessageTemplates", "notificationMessageTemplates", "notificationMessageTemplates"),
		resourceids.UserSpecifiedSegment("notificationMessageTemplateId", "notificationMessageTemplateId"),
		resourceids.StaticSegment("localizedNotificationMessages", "localizedNotificationMessages", "localizedNotificationMessages"),
		resourceids.UserSpecifiedSegment("localizedNotificationMessageId", "localizedNotificationMessageId"),
	}
}

// String returns a human-readable description of this Device Management Notification Message Template Id Localized Notification Message ID
func (id DeviceManagementNotificationMessageTemplateIdLocalizedNotificationMessageId) String() string {
	components := []string{
		fmt.Sprintf("Notification Message Template: %q", id.NotificationMessageTemplateId),
		fmt.Sprintf("Localized Notification Message: %q", id.LocalizedNotificationMessageId),
	}
	return fmt.Sprintf("Device Management Notification Message Template Id Localized Notification Message (%s)", strings.Join(components, "\n"))
}
