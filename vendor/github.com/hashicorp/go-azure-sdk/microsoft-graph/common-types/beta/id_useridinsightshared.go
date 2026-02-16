package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdInsightSharedId{}

// UserIdInsightSharedId is a struct representing the Resource ID for a User Id Insight Shared
type UserIdInsightSharedId struct {
	UserId          string
	SharedInsightId string
}

// NewUserIdInsightSharedID returns a new UserIdInsightSharedId struct
func NewUserIdInsightSharedID(userId string, sharedInsightId string) UserIdInsightSharedId {
	return UserIdInsightSharedId{
		UserId:          userId,
		SharedInsightId: sharedInsightId,
	}
}

// ParseUserIdInsightSharedID parses 'input' into a UserIdInsightSharedId
func ParseUserIdInsightSharedID(input string) (*UserIdInsightSharedId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdInsightSharedId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdInsightSharedId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseUserIdInsightSharedIDInsensitively parses 'input' case-insensitively into a UserIdInsightSharedId
// note: this method should only be used for API response data and not user input
func ParseUserIdInsightSharedIDInsensitively(input string) (*UserIdInsightSharedId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdInsightSharedId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdInsightSharedId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *UserIdInsightSharedId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.UserId, ok = input.Parsed["userId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "userId", input)
	}

	if id.SharedInsightId, ok = input.Parsed["sharedInsightId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "sharedInsightId", input)
	}

	return nil
}

// ValidateUserIdInsightSharedID checks that 'input' can be parsed as a User Id Insight Shared ID
func ValidateUserIdInsightSharedID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserIdInsightSharedID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Id Insight Shared ID
func (id UserIdInsightSharedId) ID() string {
	fmtString := "/users/%s/insights/shared/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.SharedInsightId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Id Insight Shared ID
func (id UserIdInsightSharedId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("users", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userId"),
		resourceids.StaticSegment("insights", "insights", "insights"),
		resourceids.StaticSegment("shared", "shared", "shared"),
		resourceids.UserSpecifiedSegment("sharedInsightId", "sharedInsightId"),
	}
}

// String returns a human-readable description of this User Id Insight Shared ID
func (id UserIdInsightSharedId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Shared Insight: %q", id.SharedInsightId),
	}
	return fmt.Sprintf("User Id Insight Shared (%s)", strings.Join(components, "\n"))
}
