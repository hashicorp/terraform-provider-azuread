package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdInsightUsedId{}

// UserIdInsightUsedId is a struct representing the Resource ID for a User Id Insight Used
type UserIdInsightUsedId struct {
	UserId        string
	UsedInsightId string
}

// NewUserIdInsightUsedID returns a new UserIdInsightUsedId struct
func NewUserIdInsightUsedID(userId string, usedInsightId string) UserIdInsightUsedId {
	return UserIdInsightUsedId{
		UserId:        userId,
		UsedInsightId: usedInsightId,
	}
}

// ParseUserIdInsightUsedID parses 'input' into a UserIdInsightUsedId
func ParseUserIdInsightUsedID(input string) (*UserIdInsightUsedId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdInsightUsedId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdInsightUsedId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseUserIdInsightUsedIDInsensitively parses 'input' case-insensitively into a UserIdInsightUsedId
// note: this method should only be used for API response data and not user input
func ParseUserIdInsightUsedIDInsensitively(input string) (*UserIdInsightUsedId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdInsightUsedId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdInsightUsedId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *UserIdInsightUsedId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.UserId, ok = input.Parsed["userId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "userId", input)
	}

	if id.UsedInsightId, ok = input.Parsed["usedInsightId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "usedInsightId", input)
	}

	return nil
}

// ValidateUserIdInsightUsedID checks that 'input' can be parsed as a User Id Insight Used ID
func ValidateUserIdInsightUsedID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserIdInsightUsedID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Id Insight Used ID
func (id UserIdInsightUsedId) ID() string {
	fmtString := "/users/%s/insights/used/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.UsedInsightId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Id Insight Used ID
func (id UserIdInsightUsedId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("users", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userId"),
		resourceids.StaticSegment("insights", "insights", "insights"),
		resourceids.StaticSegment("used", "used", "used"),
		resourceids.UserSpecifiedSegment("usedInsightId", "usedInsightId"),
	}
}

// String returns a human-readable description of this User Id Insight Used ID
func (id UserIdInsightUsedId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Used Insight: %q", id.UsedInsightId),
	}
	return fmt.Sprintf("User Id Insight Used (%s)", strings.Join(components, "\n"))
}
