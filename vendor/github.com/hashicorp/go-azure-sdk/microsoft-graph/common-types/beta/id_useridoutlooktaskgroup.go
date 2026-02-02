package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdOutlookTaskGroupId{}

// UserIdOutlookTaskGroupId is a struct representing the Resource ID for a User Id Outlook Task Group
type UserIdOutlookTaskGroupId struct {
	UserId             string
	OutlookTaskGroupId string
}

// NewUserIdOutlookTaskGroupID returns a new UserIdOutlookTaskGroupId struct
func NewUserIdOutlookTaskGroupID(userId string, outlookTaskGroupId string) UserIdOutlookTaskGroupId {
	return UserIdOutlookTaskGroupId{
		UserId:             userId,
		OutlookTaskGroupId: outlookTaskGroupId,
	}
}

// ParseUserIdOutlookTaskGroupID parses 'input' into a UserIdOutlookTaskGroupId
func ParseUserIdOutlookTaskGroupID(input string) (*UserIdOutlookTaskGroupId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdOutlookTaskGroupId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdOutlookTaskGroupId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseUserIdOutlookTaskGroupIDInsensitively parses 'input' case-insensitively into a UserIdOutlookTaskGroupId
// note: this method should only be used for API response data and not user input
func ParseUserIdOutlookTaskGroupIDInsensitively(input string) (*UserIdOutlookTaskGroupId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdOutlookTaskGroupId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdOutlookTaskGroupId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *UserIdOutlookTaskGroupId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.UserId, ok = input.Parsed["userId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "userId", input)
	}

	if id.OutlookTaskGroupId, ok = input.Parsed["outlookTaskGroupId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "outlookTaskGroupId", input)
	}

	return nil
}

// ValidateUserIdOutlookTaskGroupID checks that 'input' can be parsed as a User Id Outlook Task Group ID
func ValidateUserIdOutlookTaskGroupID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserIdOutlookTaskGroupID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Id Outlook Task Group ID
func (id UserIdOutlookTaskGroupId) ID() string {
	fmtString := "/users/%s/outlook/taskGroups/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.OutlookTaskGroupId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Id Outlook Task Group ID
func (id UserIdOutlookTaskGroupId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("users", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userId"),
		resourceids.StaticSegment("outlook", "outlook", "outlook"),
		resourceids.StaticSegment("taskGroups", "taskGroups", "taskGroups"),
		resourceids.UserSpecifiedSegment("outlookTaskGroupId", "outlookTaskGroupId"),
	}
}

// String returns a human-readable description of this User Id Outlook Task Group ID
func (id UserIdOutlookTaskGroupId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Outlook Task Group: %q", id.OutlookTaskGroupId),
	}
	return fmt.Sprintf("User Id Outlook Task Group (%s)", strings.Join(components, "\n"))
}
