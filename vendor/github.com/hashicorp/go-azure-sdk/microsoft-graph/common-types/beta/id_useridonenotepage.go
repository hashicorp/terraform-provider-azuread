package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdOnenotePageId{}

// UserIdOnenotePageId is a struct representing the Resource ID for a User Id Onenote Page
type UserIdOnenotePageId struct {
	UserId        string
	OnenotePageId string
}

// NewUserIdOnenotePageID returns a new UserIdOnenotePageId struct
func NewUserIdOnenotePageID(userId string, onenotePageId string) UserIdOnenotePageId {
	return UserIdOnenotePageId{
		UserId:        userId,
		OnenotePageId: onenotePageId,
	}
}

// ParseUserIdOnenotePageID parses 'input' into a UserIdOnenotePageId
func ParseUserIdOnenotePageID(input string) (*UserIdOnenotePageId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdOnenotePageId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdOnenotePageId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseUserIdOnenotePageIDInsensitively parses 'input' case-insensitively into a UserIdOnenotePageId
// note: this method should only be used for API response data and not user input
func ParseUserIdOnenotePageIDInsensitively(input string) (*UserIdOnenotePageId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdOnenotePageId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdOnenotePageId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *UserIdOnenotePageId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.UserId, ok = input.Parsed["userId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "userId", input)
	}

	if id.OnenotePageId, ok = input.Parsed["onenotePageId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "onenotePageId", input)
	}

	return nil
}

// ValidateUserIdOnenotePageID checks that 'input' can be parsed as a User Id Onenote Page ID
func ValidateUserIdOnenotePageID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserIdOnenotePageID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Id Onenote Page ID
func (id UserIdOnenotePageId) ID() string {
	fmtString := "/users/%s/onenote/pages/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.OnenotePageId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Id Onenote Page ID
func (id UserIdOnenotePageId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("users", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userId"),
		resourceids.StaticSegment("onenote", "onenote", "onenote"),
		resourceids.StaticSegment("pages", "pages", "pages"),
		resourceids.UserSpecifiedSegment("onenotePageId", "onenotePageId"),
	}
}

// String returns a human-readable description of this User Id Onenote Page ID
func (id UserIdOnenotePageId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Onenote Page: %q", id.OnenotePageId),
	}
	return fmt.Sprintf("User Id Onenote Page (%s)", strings.Join(components, "\n"))
}
