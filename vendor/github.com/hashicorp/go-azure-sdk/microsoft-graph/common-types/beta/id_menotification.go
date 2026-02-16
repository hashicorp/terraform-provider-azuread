package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeNotificationId{}

// MeNotificationId is a struct representing the Resource ID for a Me Notification
type MeNotificationId struct {
	NotificationId string
}

// NewMeNotificationID returns a new MeNotificationId struct
func NewMeNotificationID(notificationId string) MeNotificationId {
	return MeNotificationId{
		NotificationId: notificationId,
	}
}

// ParseMeNotificationID parses 'input' into a MeNotificationId
func ParseMeNotificationID(input string) (*MeNotificationId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeNotificationId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeNotificationId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseMeNotificationIDInsensitively parses 'input' case-insensitively into a MeNotificationId
// note: this method should only be used for API response data and not user input
func ParseMeNotificationIDInsensitively(input string) (*MeNotificationId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeNotificationId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeNotificationId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *MeNotificationId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.NotificationId, ok = input.Parsed["notificationId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "notificationId", input)
	}

	return nil
}

// ValidateMeNotificationID checks that 'input' can be parsed as a Me Notification ID
func ValidateMeNotificationID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeNotificationID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Notification ID
func (id MeNotificationId) ID() string {
	fmtString := "/me/notifications/%s"
	return fmt.Sprintf(fmtString, id.NotificationId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Notification ID
func (id MeNotificationId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("me", "me", "me"),
		resourceids.StaticSegment("notifications", "notifications", "notifications"),
		resourceids.UserSpecifiedSegment("notificationId", "notificationId"),
	}
}

// String returns a human-readable description of this Me Notification ID
func (id MeNotificationId) String() string {
	components := []string{
		fmt.Sprintf("Notification: %q", id.NotificationId),
	}
	return fmt.Sprintf("Me Notification (%s)", strings.Join(components, "\n"))
}
