package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdOnenoteResourceId{}

// UserIdOnenoteResourceId is a struct representing the Resource ID for a User Id Onenote Resource
type UserIdOnenoteResourceId struct {
	UserId            string
	OnenoteResourceId string
}

// NewUserIdOnenoteResourceID returns a new UserIdOnenoteResourceId struct
func NewUserIdOnenoteResourceID(userId string, onenoteResourceId string) UserIdOnenoteResourceId {
	return UserIdOnenoteResourceId{
		UserId:            userId,
		OnenoteResourceId: onenoteResourceId,
	}
}

// ParseUserIdOnenoteResourceID parses 'input' into a UserIdOnenoteResourceId
func ParseUserIdOnenoteResourceID(input string) (*UserIdOnenoteResourceId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdOnenoteResourceId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdOnenoteResourceId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseUserIdOnenoteResourceIDInsensitively parses 'input' case-insensitively into a UserIdOnenoteResourceId
// note: this method should only be used for API response data and not user input
func ParseUserIdOnenoteResourceIDInsensitively(input string) (*UserIdOnenoteResourceId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdOnenoteResourceId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdOnenoteResourceId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *UserIdOnenoteResourceId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.UserId, ok = input.Parsed["userId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "userId", input)
	}

	if id.OnenoteResourceId, ok = input.Parsed["onenoteResourceId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "onenoteResourceId", input)
	}

	return nil
}

// ValidateUserIdOnenoteResourceID checks that 'input' can be parsed as a User Id Onenote Resource ID
func ValidateUserIdOnenoteResourceID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserIdOnenoteResourceID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Id Onenote Resource ID
func (id UserIdOnenoteResourceId) ID() string {
	fmtString := "/users/%s/onenote/resources/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.OnenoteResourceId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Id Onenote Resource ID
func (id UserIdOnenoteResourceId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("users", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userId"),
		resourceids.StaticSegment("onenote", "onenote", "onenote"),
		resourceids.StaticSegment("resources", "resources", "resources"),
		resourceids.UserSpecifiedSegment("onenoteResourceId", "onenoteResourceId"),
	}
}

// String returns a human-readable description of this User Id Onenote Resource ID
func (id UserIdOnenoteResourceId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Onenote Resource: %q", id.OnenoteResourceId),
	}
	return fmt.Sprintf("User Id Onenote Resource (%s)", strings.Join(components, "\n"))
}
