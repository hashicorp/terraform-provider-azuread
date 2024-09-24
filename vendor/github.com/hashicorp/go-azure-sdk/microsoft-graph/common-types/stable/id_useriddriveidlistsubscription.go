package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdDriveIdListSubscriptionId{}

// UserIdDriveIdListSubscriptionId is a struct representing the Resource ID for a User Id Drive Id List Subscription
type UserIdDriveIdListSubscriptionId struct {
	UserId         string
	DriveId        string
	SubscriptionId string
}

// NewUserIdDriveIdListSubscriptionID returns a new UserIdDriveIdListSubscriptionId struct
func NewUserIdDriveIdListSubscriptionID(userId string, driveId string, subscriptionId string) UserIdDriveIdListSubscriptionId {
	return UserIdDriveIdListSubscriptionId{
		UserId:         userId,
		DriveId:        driveId,
		SubscriptionId: subscriptionId,
	}
}

// ParseUserIdDriveIdListSubscriptionID parses 'input' into a UserIdDriveIdListSubscriptionId
func ParseUserIdDriveIdListSubscriptionID(input string) (*UserIdDriveIdListSubscriptionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdDriveIdListSubscriptionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdDriveIdListSubscriptionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseUserIdDriveIdListSubscriptionIDInsensitively parses 'input' case-insensitively into a UserIdDriveIdListSubscriptionId
// note: this method should only be used for API response data and not user input
func ParseUserIdDriveIdListSubscriptionIDInsensitively(input string) (*UserIdDriveIdListSubscriptionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdDriveIdListSubscriptionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdDriveIdListSubscriptionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *UserIdDriveIdListSubscriptionId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.UserId, ok = input.Parsed["userId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "userId", input)
	}

	if id.DriveId, ok = input.Parsed["driveId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "driveId", input)
	}

	if id.SubscriptionId, ok = input.Parsed["subscriptionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", input)
	}

	return nil
}

// ValidateUserIdDriveIdListSubscriptionID checks that 'input' can be parsed as a User Id Drive Id List Subscription ID
func ValidateUserIdDriveIdListSubscriptionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserIdDriveIdListSubscriptionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Id Drive Id List Subscription ID
func (id UserIdDriveIdListSubscriptionId) ID() string {
	fmtString := "/users/%s/drives/%s/list/subscriptions/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.DriveId, id.SubscriptionId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Id Drive Id List Subscription ID
func (id UserIdDriveIdListSubscriptionId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("users", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userId"),
		resourceids.StaticSegment("drives", "drives", "drives"),
		resourceids.UserSpecifiedSegment("driveId", "driveId"),
		resourceids.StaticSegment("list", "list", "list"),
		resourceids.StaticSegment("subscriptions", "subscriptions", "subscriptions"),
		resourceids.UserSpecifiedSegment("subscriptionId", "subscriptionId"),
	}
}

// String returns a human-readable description of this User Id Drive Id List Subscription ID
func (id UserIdDriveIdListSubscriptionId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Drive: %q", id.DriveId),
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
	}
	return fmt.Sprintf("User Id Drive Id List Subscription (%s)", strings.Join(components, "\n"))
}
