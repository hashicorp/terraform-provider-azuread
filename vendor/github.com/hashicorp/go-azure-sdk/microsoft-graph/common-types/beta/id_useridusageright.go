package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdUsageRightId{}

// UserIdUsageRightId is a struct representing the Resource ID for a User Id Usage Right
type UserIdUsageRightId struct {
	UserId       string
	UsageRightId string
}

// NewUserIdUsageRightID returns a new UserIdUsageRightId struct
func NewUserIdUsageRightID(userId string, usageRightId string) UserIdUsageRightId {
	return UserIdUsageRightId{
		UserId:       userId,
		UsageRightId: usageRightId,
	}
}

// ParseUserIdUsageRightID parses 'input' into a UserIdUsageRightId
func ParseUserIdUsageRightID(input string) (*UserIdUsageRightId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdUsageRightId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdUsageRightId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseUserIdUsageRightIDInsensitively parses 'input' case-insensitively into a UserIdUsageRightId
// note: this method should only be used for API response data and not user input
func ParseUserIdUsageRightIDInsensitively(input string) (*UserIdUsageRightId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdUsageRightId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdUsageRightId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *UserIdUsageRightId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.UserId, ok = input.Parsed["userId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "userId", input)
	}

	if id.UsageRightId, ok = input.Parsed["usageRightId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "usageRightId", input)
	}

	return nil
}

// ValidateUserIdUsageRightID checks that 'input' can be parsed as a User Id Usage Right ID
func ValidateUserIdUsageRightID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserIdUsageRightID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Id Usage Right ID
func (id UserIdUsageRightId) ID() string {
	fmtString := "/users/%s/usageRights/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.UsageRightId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Id Usage Right ID
func (id UserIdUsageRightId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("users", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userId"),
		resourceids.StaticSegment("usageRights", "usageRights", "usageRights"),
		resourceids.UserSpecifiedSegment("usageRightId", "usageRightId"),
	}
}

// String returns a human-readable description of this User Id Usage Right ID
func (id UserIdUsageRightId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Usage Right: %q", id.UsageRightId),
	}
	return fmt.Sprintf("User Id Usage Right (%s)", strings.Join(components, "\n"))
}
