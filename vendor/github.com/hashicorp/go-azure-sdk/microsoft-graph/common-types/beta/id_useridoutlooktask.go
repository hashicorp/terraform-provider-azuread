package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdOutlookTaskId{}

// UserIdOutlookTaskId is a struct representing the Resource ID for a User Id Outlook Task
type UserIdOutlookTaskId struct {
	UserId        string
	OutlookTaskId string
}

// NewUserIdOutlookTaskID returns a new UserIdOutlookTaskId struct
func NewUserIdOutlookTaskID(userId string, outlookTaskId string) UserIdOutlookTaskId {
	return UserIdOutlookTaskId{
		UserId:        userId,
		OutlookTaskId: outlookTaskId,
	}
}

// ParseUserIdOutlookTaskID parses 'input' into a UserIdOutlookTaskId
func ParseUserIdOutlookTaskID(input string) (*UserIdOutlookTaskId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdOutlookTaskId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdOutlookTaskId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseUserIdOutlookTaskIDInsensitively parses 'input' case-insensitively into a UserIdOutlookTaskId
// note: this method should only be used for API response data and not user input
func ParseUserIdOutlookTaskIDInsensitively(input string) (*UserIdOutlookTaskId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdOutlookTaskId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdOutlookTaskId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *UserIdOutlookTaskId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.UserId, ok = input.Parsed["userId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "userId", input)
	}

	if id.OutlookTaskId, ok = input.Parsed["outlookTaskId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "outlookTaskId", input)
	}

	return nil
}

// ValidateUserIdOutlookTaskID checks that 'input' can be parsed as a User Id Outlook Task ID
func ValidateUserIdOutlookTaskID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserIdOutlookTaskID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Id Outlook Task ID
func (id UserIdOutlookTaskId) ID() string {
	fmtString := "/users/%s/outlook/tasks/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.OutlookTaskId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Id Outlook Task ID
func (id UserIdOutlookTaskId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("users", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userId"),
		resourceids.StaticSegment("outlook", "outlook", "outlook"),
		resourceids.StaticSegment("tasks", "tasks", "tasks"),
		resourceids.UserSpecifiedSegment("outlookTaskId", "outlookTaskId"),
	}
}

// String returns a human-readable description of this User Id Outlook Task ID
func (id UserIdOutlookTaskId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Outlook Task: %q", id.OutlookTaskId),
	}
	return fmt.Sprintf("User Id Outlook Task (%s)", strings.Join(components, "\n"))
}
