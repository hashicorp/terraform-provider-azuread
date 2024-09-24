package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdNotificationId{}

// UserIdNotificationId is a struct representing the Resource ID for a User Id Notification
type UserIdNotificationId struct {
	UserId         string
	NotificationId string
}

// NewUserIdNotificationID returns a new UserIdNotificationId struct
func NewUserIdNotificationID(userId string, notificationId string) UserIdNotificationId {
	return UserIdNotificationId{
		UserId:         userId,
		NotificationId: notificationId,
	}
}

// ParseUserIdNotificationID parses 'input' into a UserIdNotificationId
func ParseUserIdNotificationID(input string) (*UserIdNotificationId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdNotificationId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdNotificationId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseUserIdNotificationIDInsensitively parses 'input' case-insensitively into a UserIdNotificationId
// note: this method should only be used for API response data and not user input
func ParseUserIdNotificationIDInsensitively(input string) (*UserIdNotificationId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdNotificationId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdNotificationId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *UserIdNotificationId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.UserId, ok = input.Parsed["userId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "userId", input)
	}

	if id.NotificationId, ok = input.Parsed["notificationId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "notificationId", input)
	}

	return nil
}

// ValidateUserIdNotificationID checks that 'input' can be parsed as a User Id Notification ID
func ValidateUserIdNotificationID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserIdNotificationID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Id Notification ID
func (id UserIdNotificationId) ID() string {
	fmtString := "/users/%s/notifications/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.NotificationId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Id Notification ID
func (id UserIdNotificationId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("users", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userId"),
		resourceids.StaticSegment("notifications", "notifications", "notifications"),
		resourceids.UserSpecifiedSegment("notificationId", "notificationId"),
	}
}

// String returns a human-readable description of this User Id Notification ID
func (id UserIdNotificationId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Notification: %q", id.NotificationId),
	}
	return fmt.Sprintf("User Id Notification (%s)", strings.Join(components, "\n"))
}
