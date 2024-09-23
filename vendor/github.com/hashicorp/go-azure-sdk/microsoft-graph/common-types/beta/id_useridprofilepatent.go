package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdProfilePatentId{}

// UserIdProfilePatentId is a struct representing the Resource ID for a User Id Profile Patent
type UserIdProfilePatentId struct {
	UserId       string
	ItemPatentId string
}

// NewUserIdProfilePatentID returns a new UserIdProfilePatentId struct
func NewUserIdProfilePatentID(userId string, itemPatentId string) UserIdProfilePatentId {
	return UserIdProfilePatentId{
		UserId:       userId,
		ItemPatentId: itemPatentId,
	}
}

// ParseUserIdProfilePatentID parses 'input' into a UserIdProfilePatentId
func ParseUserIdProfilePatentID(input string) (*UserIdProfilePatentId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdProfilePatentId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdProfilePatentId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseUserIdProfilePatentIDInsensitively parses 'input' case-insensitively into a UserIdProfilePatentId
// note: this method should only be used for API response data and not user input
func ParseUserIdProfilePatentIDInsensitively(input string) (*UserIdProfilePatentId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdProfilePatentId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdProfilePatentId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *UserIdProfilePatentId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.UserId, ok = input.Parsed["userId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "userId", input)
	}

	if id.ItemPatentId, ok = input.Parsed["itemPatentId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "itemPatentId", input)
	}

	return nil
}

// ValidateUserIdProfilePatentID checks that 'input' can be parsed as a User Id Profile Patent ID
func ValidateUserIdProfilePatentID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserIdProfilePatentID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Id Profile Patent ID
func (id UserIdProfilePatentId) ID() string {
	fmtString := "/users/%s/profile/patents/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.ItemPatentId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Id Profile Patent ID
func (id UserIdProfilePatentId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("users", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userId"),
		resourceids.StaticSegment("profile", "profile", "profile"),
		resourceids.StaticSegment("patents", "patents", "patents"),
		resourceids.UserSpecifiedSegment("itemPatentId", "itemPatentId"),
	}
}

// String returns a human-readable description of this User Id Profile Patent ID
func (id UserIdProfilePatentId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Item Patent: %q", id.ItemPatentId),
	}
	return fmt.Sprintf("User Id Profile Patent (%s)", strings.Join(components, "\n"))
}
