package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdDriveIdItemIdSubscriptionId{}

// UserIdDriveIdItemIdSubscriptionId is a struct representing the Resource ID for a User Id Drive Id Item Id Subscription
type UserIdDriveIdItemIdSubscriptionId struct {
	UserId         string
	DriveId        string
	DriveItemId    string
	SubscriptionId string
}

// NewUserIdDriveIdItemIdSubscriptionID returns a new UserIdDriveIdItemIdSubscriptionId struct
func NewUserIdDriveIdItemIdSubscriptionID(userId string, driveId string, driveItemId string, subscriptionId string) UserIdDriveIdItemIdSubscriptionId {
	return UserIdDriveIdItemIdSubscriptionId{
		UserId:         userId,
		DriveId:        driveId,
		DriveItemId:    driveItemId,
		SubscriptionId: subscriptionId,
	}
}

// ParseUserIdDriveIdItemIdSubscriptionID parses 'input' into a UserIdDriveIdItemIdSubscriptionId
func ParseUserIdDriveIdItemIdSubscriptionID(input string) (*UserIdDriveIdItemIdSubscriptionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdDriveIdItemIdSubscriptionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdDriveIdItemIdSubscriptionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseUserIdDriveIdItemIdSubscriptionIDInsensitively parses 'input' case-insensitively into a UserIdDriveIdItemIdSubscriptionId
// note: this method should only be used for API response data and not user input
func ParseUserIdDriveIdItemIdSubscriptionIDInsensitively(input string) (*UserIdDriveIdItemIdSubscriptionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdDriveIdItemIdSubscriptionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdDriveIdItemIdSubscriptionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *UserIdDriveIdItemIdSubscriptionId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.UserId, ok = input.Parsed["userId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "userId", input)
	}

	if id.DriveId, ok = input.Parsed["driveId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "driveId", input)
	}

	if id.DriveItemId, ok = input.Parsed["driveItemId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "driveItemId", input)
	}

	if id.SubscriptionId, ok = input.Parsed["subscriptionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", input)
	}

	return nil
}

// ValidateUserIdDriveIdItemIdSubscriptionID checks that 'input' can be parsed as a User Id Drive Id Item Id Subscription ID
func ValidateUserIdDriveIdItemIdSubscriptionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserIdDriveIdItemIdSubscriptionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Id Drive Id Item Id Subscription ID
func (id UserIdDriveIdItemIdSubscriptionId) ID() string {
	fmtString := "/users/%s/drives/%s/items/%s/subscriptions/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.DriveId, id.DriveItemId, id.SubscriptionId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Id Drive Id Item Id Subscription ID
func (id UserIdDriveIdItemIdSubscriptionId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("users", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userId"),
		resourceids.StaticSegment("drives", "drives", "drives"),
		resourceids.UserSpecifiedSegment("driveId", "driveId"),
		resourceids.StaticSegment("items", "items", "items"),
		resourceids.UserSpecifiedSegment("driveItemId", "driveItemId"),
		resourceids.StaticSegment("subscriptions", "subscriptions", "subscriptions"),
		resourceids.UserSpecifiedSegment("subscriptionId", "subscriptionId"),
	}
}

// String returns a human-readable description of this User Id Drive Id Item Id Subscription ID
func (id UserIdDriveIdItemIdSubscriptionId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Drive: %q", id.DriveId),
		fmt.Sprintf("Drive Item: %q", id.DriveItemId),
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
	}
	return fmt.Sprintf("User Id Drive Id Item Id Subscription (%s)", strings.Join(components, "\n"))
}
