package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdDriveIdRootSubscriptionId{}

// UserIdDriveIdRootSubscriptionId is a struct representing the Resource ID for a User Id Drive Id Root Subscription
type UserIdDriveIdRootSubscriptionId struct {
	UserId         string
	DriveId        string
	SubscriptionId string
}

// NewUserIdDriveIdRootSubscriptionID returns a new UserIdDriveIdRootSubscriptionId struct
func NewUserIdDriveIdRootSubscriptionID(userId string, driveId string, subscriptionId string) UserIdDriveIdRootSubscriptionId {
	return UserIdDriveIdRootSubscriptionId{
		UserId:         userId,
		DriveId:        driveId,
		SubscriptionId: subscriptionId,
	}
}

// ParseUserIdDriveIdRootSubscriptionID parses 'input' into a UserIdDriveIdRootSubscriptionId
func ParseUserIdDriveIdRootSubscriptionID(input string) (*UserIdDriveIdRootSubscriptionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdDriveIdRootSubscriptionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdDriveIdRootSubscriptionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseUserIdDriveIdRootSubscriptionIDInsensitively parses 'input' case-insensitively into a UserIdDriveIdRootSubscriptionId
// note: this method should only be used for API response data and not user input
func ParseUserIdDriveIdRootSubscriptionIDInsensitively(input string) (*UserIdDriveIdRootSubscriptionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdDriveIdRootSubscriptionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdDriveIdRootSubscriptionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *UserIdDriveIdRootSubscriptionId) FromParseResult(input resourceids.ParseResult) error {
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

// ValidateUserIdDriveIdRootSubscriptionID checks that 'input' can be parsed as a User Id Drive Id Root Subscription ID
func ValidateUserIdDriveIdRootSubscriptionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserIdDriveIdRootSubscriptionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Id Drive Id Root Subscription ID
func (id UserIdDriveIdRootSubscriptionId) ID() string {
	fmtString := "/users/%s/drives/%s/root/subscriptions/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.DriveId, id.SubscriptionId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Id Drive Id Root Subscription ID
func (id UserIdDriveIdRootSubscriptionId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("users", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userId"),
		resourceids.StaticSegment("drives", "drives", "drives"),
		resourceids.UserSpecifiedSegment("driveId", "driveId"),
		resourceids.StaticSegment("root", "root", "root"),
		resourceids.StaticSegment("subscriptions", "subscriptions", "subscriptions"),
		resourceids.UserSpecifiedSegment("subscriptionId", "subscriptionId"),
	}
}

// String returns a human-readable description of this User Id Drive Id Root Subscription ID
func (id UserIdDriveIdRootSubscriptionId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Drive: %q", id.DriveId),
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
	}
	return fmt.Sprintf("User Id Drive Id Root Subscription (%s)", strings.Join(components, "\n"))
}
