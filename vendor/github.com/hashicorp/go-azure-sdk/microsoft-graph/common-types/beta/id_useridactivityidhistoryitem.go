package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdActivityIdHistoryItemId{}

// UserIdActivityIdHistoryItemId is a struct representing the Resource ID for a User Id Activity Id History Item
type UserIdActivityIdHistoryItemId struct {
	UserId                string
	UserActivityId        string
	ActivityHistoryItemId string
}

// NewUserIdActivityIdHistoryItemID returns a new UserIdActivityIdHistoryItemId struct
func NewUserIdActivityIdHistoryItemID(userId string, userActivityId string, activityHistoryItemId string) UserIdActivityIdHistoryItemId {
	return UserIdActivityIdHistoryItemId{
		UserId:                userId,
		UserActivityId:        userActivityId,
		ActivityHistoryItemId: activityHistoryItemId,
	}
}

// ParseUserIdActivityIdHistoryItemID parses 'input' into a UserIdActivityIdHistoryItemId
func ParseUserIdActivityIdHistoryItemID(input string) (*UserIdActivityIdHistoryItemId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdActivityIdHistoryItemId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdActivityIdHistoryItemId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseUserIdActivityIdHistoryItemIDInsensitively parses 'input' case-insensitively into a UserIdActivityIdHistoryItemId
// note: this method should only be used for API response data and not user input
func ParseUserIdActivityIdHistoryItemIDInsensitively(input string) (*UserIdActivityIdHistoryItemId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdActivityIdHistoryItemId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdActivityIdHistoryItemId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *UserIdActivityIdHistoryItemId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.UserId, ok = input.Parsed["userId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "userId", input)
	}

	if id.UserActivityId, ok = input.Parsed["userActivityId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "userActivityId", input)
	}

	if id.ActivityHistoryItemId, ok = input.Parsed["activityHistoryItemId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "activityHistoryItemId", input)
	}

	return nil
}

// ValidateUserIdActivityIdHistoryItemID checks that 'input' can be parsed as a User Id Activity Id History Item ID
func ValidateUserIdActivityIdHistoryItemID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserIdActivityIdHistoryItemID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Id Activity Id History Item ID
func (id UserIdActivityIdHistoryItemId) ID() string {
	fmtString := "/users/%s/activities/%s/historyItems/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.UserActivityId, id.ActivityHistoryItemId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Id Activity Id History Item ID
func (id UserIdActivityIdHistoryItemId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("users", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userId"),
		resourceids.StaticSegment("activities", "activities", "activities"),
		resourceids.UserSpecifiedSegment("userActivityId", "userActivityId"),
		resourceids.StaticSegment("historyItems", "historyItems", "historyItems"),
		resourceids.UserSpecifiedSegment("activityHistoryItemId", "activityHistoryItemId"),
	}
}

// String returns a human-readable description of this User Id Activity Id History Item ID
func (id UserIdActivityIdHistoryItemId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("User Activity: %q", id.UserActivityId),
		fmt.Sprintf("Activity History Item: %q", id.ActivityHistoryItemId),
	}
	return fmt.Sprintf("User Id Activity Id History Item (%s)", strings.Join(components, "\n"))
}
